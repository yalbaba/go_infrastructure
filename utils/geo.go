/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/12/9
   Description :
-------------------------------------------------
*/

package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/kayon/iploc"
	"github.com/valyala/fasthttp"
	"xorm.io/xorm"

	jsoniter "github.com/json-iterator/go"

	"liveearth/infrastructure/consts"
)

var Geo = new(geoUtil)

type geoUtil struct{}

const MaxGeoDistance = 1e9 // 百万公里

// 计算两个坐标的距离, 输出单位:米
func (*geoUtil) Distances(lon1, lat1, lon2, lat2 float64) float64 {
	if lon1 == 0 && lat1 == 0 {
		return 0
	}
	if lon2 == 0 && lat2 == 0 {
		return 0
	}

	radians := func(d float64) float64 {
		r := d * math.Pi / 180.0
		if d < 0 {
			r = -math.Abs(r)
		}
		return r
	}
	lon1, lat1, lon2, lat2 = radians(lon1), radians(lat1), radians(lon2), radians(lat2)
	dlon, dlat := lon2-lon1, lat2-lat1

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	return 2 * math.Asin(math.Sqrt(a)) * 6370996.81
}

// 地理中心点, 将每个经纬度转化成x,y,z的坐标值。然后根据根据x,y,z的值，寻找3D坐标系中的中心点
// GeoMidPoint({lon,lat},{lon,lat}) lon,lat
func (*geoUtil) MidPoint(coords ...[]float64) (float64, float64) {
	var x, y, z float64
	for _, coord := range coords {
		lon := coord[0] * math.Pi / 180
		lat := coord[1] * math.Pi / 180

		a := math.Cos(lat) * math.Cos(lon)
		b := math.Cos(lat) * math.Sin(lon)
		c := math.Sin(lat)

		x += a
		y += b
		z += c
	}

	coord_num := float64(len(coords))
	x /= coord_num
	y /= coord_num
	z /= coord_num

	lon := math.Atan2(y, x)
	hyp := math.Sqrt(x*x + y*y)
	lat := math.Atan2(z, hyp)

	lon *= 180 / math.Pi
	lat *= 180 / math.Pi

	return lon, lat
}

// 平均经纬度, 将经纬度坐标看成是平面坐标，直接计算经度和纬度的平均值(该方法只是大致的估算方法，仅适合距离在400KM以内的点)
// GeoAveragePoint({lon,lat},{lon,lat}) lon,lat
func (*geoUtil) AveragePoint(coords ...[]float64) (float64, float64) {
	var x, y float64
	for _, coord := range coords {
		lon := coord[0]
		lat := coord[1]
		x += lon
		y += lat
	}

	coord_num := float64(len(coords))
	x /= coord_num
	y /= coord_num
	return x, y
}

// 根据距离输出展示文本
func (*geoUtil) MakeDistanceText(distance float64) string {
	if distance == MaxGeoDistance {
		return ""
	}
	if distance < 1000 {
		return fmt.Sprintf("%dm", int(distance))
	}
	return fmt.Sprintf("%dkm", int(distance/1000))
}

var Municipalities = map[string]struct{}{
	"北京": {}, "上海": {}, "天津": {}, "重庆": {},
	"北京市": {}, "上海市": {}, "天津市": {}, "重庆市": {},
	"北京城区": {}, "上海城区": {}, "天津城区": {}, "重庆城区": {},
	"11": {}, "31": {}, "12": {}, "50": {},
	"911": {}, "931": {}, "912": {}, "955": {},
	"BJ": {}, "SH": {}, "TJ": {}, "CQ": {}, "bj": {}, "sh": {}, "tj": {}, "cq": {},
}

// 是否为直辖市
func (*geoUtil) IsMunicipalities(name string) bool {
	_, ok := Municipalities[name]
	return ok
}

// 获取ip所在区域id
func (*geoUtil) GetDistrictIdsOfIp(session *xorm.Session, ipLoc *iploc.Locator, ipStr string) ([]int, error) {
	ipDetail := ipLoc.Find(ipStr)
	if ipDetail == nil {
		return []int{consts.DistrictIdAll}, nil
	}

	var districtName string
	if len(ipDetail.City) != 0 {
		districtName = ipDetail.City
	} else if len(ipDetail.Province) != 0 && len(ipDetail.City) == 0 {
		districtName = ipDetail.Province
	} else {
		districtName = ipDetail.Country
	}

	if len(districtName) == 0 {
		return []int{consts.DistrictIdAll}, nil
	}

	// 获取用户区域
	var parentIds string
	_, err := session.SQL("select parent_ids from liveearth_primary.live_district where name like ? limit 1", districtName+"%").Get(&parentIds)
	if err != nil {
		return []int{consts.DistrictIdAll}, fmt.Errorf("utils.Geo.GetDistrictIdsOfIp err:%v", err)
	}

	var dist []int
	_ = jsoniter.UnmarshalFromString(parentIds, &dist)
	if len(dist) == 0 {
		return []int{consts.DistrictIdAll}, nil
	}

	return dist, nil
}

// 根据经纬度获取所在区域id
func (u *geoUtil) GetDistrictIdsOfGeo(se *xorm.Session, lon, lat float64) ([]int, error) {
	detail, err := u.GaodeReGeo(lon, lat)
	if err != nil {
		return []int{consts.DistrictIdAll}, nil
	}

	var districtName string
	if len(detail.City) != 0 {
		districtName = detail.City
	} else if len(detail.Province) != 0 && len(detail.City) == 0 {
		districtName = detail.Province
	} else {
		districtName = detail.District
	}

	if len(districtName) == 0 {
		return []int{consts.DistrictIdAll}, nil
	}

	// 获取用户区域
	var parentIds string
	_, err = se.SQL("select parent_ids from liveearth_primary.live_district where name like ? limit 1", districtName+"%").Get(&parentIds)
	if err != nil {
		return []int{consts.DistrictIdAll}, fmt.Errorf("utils.Geo.GetDistrictIdsOfIp err:%v", err)
	}

	var dist []int
	_ = jsoniter.UnmarshalFromString(parentIds, &dist)
	if len(dist) == 0 {
		return []int{consts.DistrictIdAll}, nil
	}

	return dist, nil
}

var Municipality = map[string]int{"上海市": 0, "重庆市": 0, "北京市": 0, "天津市": 0, "香港特别行政区": 0, "澳门特别行政区": 0}

func MunicipalityCityIsNull(province, city string) string {
	if _, ok := Municipality[province]; ok {
		city = ""
	}
	return city
}

// 数量解析，传入参数
type CountFootprintInput struct {
	Nation    string  `json:"nation"`     // 国家
	Province  string  `json:"province"`   // 省
	City      string  `json:"city"`       // 城市
	Lat       float64 `json:"lat"`        // 纬度
	Lon       float64 `json:"lon"`        // 经度
	MediaType int     `json:"media_type"` // 媒体类型
}

// 数量解析。传出参数
type CountFootprintOut struct {
	NationCount     int    `json:"nation_count"`
	ProvinceCount   int    `json:"province_count"`
	CityCount       int    `json:"city_count"`
	Locations       string `json:"locations"`
	PictureCount    int    `json:"picture_count"`
	VideoCount      int    `json:"video_count"`
	CFlytoRectangle string `json:"c_flyto_rectangle"`
}

type ChangeLocationsModel struct {
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
}

type LatLon struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type ChangeLocationsJsonModel struct {
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Location LatLon `json:"location"`
}

func ParseCountFootprint(inputs []*CountFootprintInput) *CountFootprintOut {
	nationSet := map[string]int{}
	citySet := map[string]int{}
	provinceSet := map[string]int{}
	locations := map[ChangeLocationsModel]*LatLon{}
	var pictureCount, videoCount int

	for _, image := range inputs {
		if image.Nation != "" {
			nationSet[image.Nation] = 1
		}
		if image.City == "" || MunicipalityCityIsNull(image.Province, image.City) == "" {
			// 是直辖市或区或者国外的城市
			citySet[image.Province] = 1
		} else {
			citySet[image.City] = 1
		}
		if image.Province != "" {
			provinceSet[image.Province] = 1
		}
		if image.Nation != "" || image.Province != "" {
			l := ChangeLocationsModel{
				Nation:   image.Nation,
				Province: image.Province,
				City:     image.City}
			locations[l] = &LatLon{Lon: image.Lon,
				Lat: image.Lat}
		}
		switch image.MediaType {
		case 10:
			pictureCount++
		case 20:
			videoCount++
		}
	}
	var locs []*ChangeLocationsJsonModel
	for loc, latLon := range locations {
		var ll = new(ChangeLocationsJsonModel)
		ll.City = loc.City
		ll.Nation = loc.Nation
		ll.Province = loc.Province
		ll.Location = *latLon
		locs = append(locs, ll)
	}

	var jsonter = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := jsonter.Marshal(locs)
	var locJson string
	if err != nil {
		locJson = ""
	} else {
		locJson = string(b)
	}
	rectangle := Retangle(inputs)
	return &CountFootprintOut{
		Locations:       locJson,
		NationCount:     len(nationSet),
		ProvinceCount:   len(provinceSet),
		CityCount:       len(citySet),
		PictureCount:    pictureCount,
		VideoCount:      videoCount,
		CFlytoRectangle: rectangle,
	}
}

/*
   可视故事的可视范围
*/
func Retangle(inputs []*CountFootprintInput) string {

	earthC := 40076000.00 / 2
	maxLat := -90.00
	maxLon := -180.00
	minLat := 90.00
	minLon := 180.00
	for _, input := range inputs {
		if input.Lon > maxLon {
			maxLon = input.Lon
		}
		if input.Lon < minLon {
			minLon = input.Lon
		}
		if input.Lat > maxLat {
			maxLat = input.Lat
		}
		if input.Lat < minLat {
			minLat = input.Lat
		}
	}
	l := GeoDistances(maxLon, maxLat, minLon, minLat)
	if l >= earthC {
		return "0,0;0,0"
	}
	return fmt.Sprintf("%G,%G;%G,%G", minLon, minLat, maxLon, maxLat)
}

func GeoDistances(lon1, lat1, lon2, lat2 float64) float64 {
	if lon1 == 0 && lat1 == 0 {
		return 0
	}
	if lon2 == 0 && lat2 == 0 {
		return 0
	}

	radians := func(d float64) float64 {
		r := d * math.Pi / 180.0
		if d < 0 {
			r = -math.Abs(r)
		}
		return r
	}
	lon1, lat1, lon2, lat2 = radians(lon1), radians(lat1), radians(lon2), radians(lat2)
	dlon, dlat := lon2-lon1, lat2-lat1

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	return 2 * math.Asin(math.Sqrt(a)) * 6370996.81
}

func Bl2height(lat, lon float64) (int, error) {
	return 0, nil
	statusCode, body, err := fasthttp.Get(nil, fmt.Sprintf("http://earth-backup.national-space.com/getheight?l=%f&b=%f", lon, lat))
	if err != nil {
		return 0, err
	}
	if statusCode != fasthttp.StatusOK {
		return 0, errors.New("statusCode is not 200")
	}
	f, err := strconv.ParseFloat(string(body), 64)
	if err != nil {
		return 0, errors.New("格式转换错误")
	}
	return int(f), nil
}

// 根据经度计算时差
func GetDiffTime(lon, lat float64) int {
	const beijingZone = 116.383
	if lon > 73.66 && lon < 135.05 && lat > 3.86 && lat < 53.55 { // 中国区域
		return 0
	} else {
		zone := int(math.Floor((lon-beijingZone)/15+0.5)) + 1 // 加0.5然后向下取整，类似4舍5入
		return zone
	}
}

// 判断是否在中国
func IsInsideChina(lon, lat float64) bool {
	InSideRectangl := [][]float64{
		// [Rectangle rectangleWithLa1:49.220400 lo1:079.446200 la2:42.889900 lo2:096.330000],
		// 左上，右下
		{49.220400, 79.446200, 42.889900, 96.330000},
		{54.141500, 109.687200, 39.374200, 135.000200},
		{42.889900, 73.124600, 29.529700, 124.143255},
		{29.529700, 82.968400, 26.718600, 97.035200},
		{29.529700, 97.025300, 20.414096, 124.367395},
		{20.414096, 107.975793, 17.871542, 111.744104},
	}
	OutSideRectangl := [][]float64{
		{25.398623, 119.921265, 21.785006, 122.497559},
		{22.284000, 101.865200, 20.098800, 106.665000},
		{21.542200, 106.452500, 20.487800, 108.051000},
		{55.817500, 109.032300, 50.325700, 119.127000},
		{55.817500, 127.456800, 49.557400, 137.022700},
		{44.892200, 131.266200, 42.569200, 137.022700},
	}
	point := []float64{lat, lon}
	for _, inRec := range InSideRectangl {
		isIn, _ := inRectangle(point, inRec)
		if isIn {
			for _, outRec := range OutSideRectangl {
				isOut, _ := inRectangle(point, outRec)
				if isOut {
					return false
				}
			}
			return true
		}
	}
	return false
}

func inRectangle(point []float64, rectangle []float64) (bool, error) {
	/*
		point [纬度，经度]，rectangle[左上纬度，左上经度，右下维度，右下经度]
	*/
	if len(point) != 2 || len(rectangle) != 4 {
		return false, errors.New("输入经纬度长度不合法")
	}
	if math.Min(rectangle[0], rectangle[2]) <= point[0] && point[0] <= math.Max(rectangle[0], rectangle[2]) && math.Min(rectangle[1], rectangle[3]) <= point[1] && point[1] <= math.Max(rectangle[1], rectangle[3]) {
		return true, nil
	}
	return false, nil
}

const (
	Pi  = 3.1415926535897932384626 // π
	Ee  = 0.00669342162296594323   // 偏心率平方
	R   = 6378245.0                // 长半轴
	XPi = 3.14159265358979324 * 3000.0 / 180.0
)

// wgs84转高德坐标系
func Wgs84ToGcj02(lon, lat float64) (float64, float64) {
	if !IsInsideChina(lon, lat) {
		return lon, lat
	}
	dlat := transformlat(lon-105.0, lat-35.0)
	dlon := transformlon(lon-105.0, lat-35.0)
	radlat := lat / 180.0 * Pi
	magic := math.Sin(radlat)
	magic = 1 - Ee*magic*magic
	sqrtmagic := math.Sqrt(magic)
	dlat = (dlat * 180.0) / ((R * (1 - Ee)) / (magic * sqrtmagic) * Pi)
	dlng := (dlon * 180.0) / (R / sqrtmagic * math.Cos(radlat) * Pi)
	mglat := lat + dlat
	mglng := lon + dlng
	return mglng, mglat
}

// 高德转wgs84
func Gcj02ToWgs84(lon, lat float64) (float64, float64) {
	if !IsInsideChina(lon, lat) {
		return lon, lat
	}
	dlat := transformlat(lon-105.0, lat-35.0)
	dlng := transformlon(lon-105.0, lat-35.0)
	radlat := lat / 180.0 * Pi
	magic := math.Sin(radlat)
	magic = 1 - Ee*magic*magic
	sqrtmagic := math.Sqrt(magic)
	dlat = (dlat * 180.0) / ((R * (1 - Ee)) / (magic * sqrtmagic) * Pi)
	dlng = (dlng * 180.0) / (R / sqrtmagic * math.Cos(radlat) * Pi)
	mglat := lat + dlat
	mglng := lon + dlng
	return lon*2 - mglng, lat*2 - mglat
}

// wgs84转百度
func Wgs84ToBd09(lon, lat float64) (float64, float64) {
	lon, lat = Wgs84ToGcj02(lon, lat)
	return gcj02ToBd09(lon, lat)
}

// 百度转wgs84
func Bd09ToWgs84(lon, lat float64) (float64, float64) {
	lon, lat = bd09Togcj02(lon, lat)
	return Gcj02ToWgs84(lon, lat)
}

func transformlat(lon, lat float64) float64 {
	ret := -100.0 + 2.0*lon + 3.0*lat + 0.2*lat*lat + 0.1*lon*lat + 0.2*math.Sqrt(math.Abs(lon))
	ret += (20.0*math.Sin(6.0*lon*Pi) + 20.0*math.Sin(2.0*lon*Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lat*Pi) + 40.0*math.Sin(lat/3.0*Pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(lat/12.0*Pi) + 320*math.Sin(lat*Pi/30.0)) * 2.0 / 3.0
	return ret
}
func transformlon(lon, lat float64) float64 {
	ret := 300.0 + lon + 2.0*lat + 0.1*lon*lon + 0.1*lon*lat + 0.1*math.Sqrt(math.Abs(lon))
	ret += (20.0*math.Sin(6.0*lon*Pi) + 20.0*math.Sin(2.0*lon*Pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lon*Pi) + 40.0*math.Sin(lon/3.0*Pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(lon/12.0*Pi) + 300.0*math.Sin(lon/30.0*Pi)) * 2.0 / 3.0
	return ret
}

// 高德转百度
func gcj02ToBd09(lon, lat float64) (float64, float64) {
	z := math.Sqrt(lon*lon+lat*lat) + 0.000003*math.Cos(lon*XPi)
	theta := math.Atan2(lat, lon) + 0.000003*math.Cos(lon*XPi)
	bd_lng := z*math.Cos(theta) + 0.0065
	bd_lat := z*math.Sin(theta) + 0.006
	return bd_lng, bd_lat
}

// 百度转高德
func bd09Togcj02(lon, lat float64) (float64, float64) {
	x := lon - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*XPi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*XPi)
	gg_lng := z * math.Cos(theta)
	gg_lat := z * math.Sin(theta)
	return gg_lng, gg_lat
}

func GetPrecisionByOneMeter(numStr string) string {
	pointNum := strings.Index(numStr, ".")
	if pointNum > 0 && len(numStr) >= pointNum+4 {
		return numStr[:pointNum+4]
	}
	return numStr
}
