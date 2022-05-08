package utils

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/zlyuancn/zjve2"
	"github.com/zlyuancn/zstr"
)

const GaodeApiKey = "b5551625b69ac84fad89a71f06df7666"

// https://developer.amap.com/api/webservice/guide/api/georegeo#regeo
const GaodeReGeoApiUrl = "https://restapi.amap.com/v3/geocode/regeo"

// https://lbs.amap.com/api/webservice/guide/api/ipconfig
const GaodeIPPositionApiUrl = "https://restapi.amap.com/v3/ip"

var DefaultHttpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

type GaodeReGeoResult struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Regeocode struct {
		AddressComponent struct {
			City     string `json:"-"`
			Province string `json:"-"`
			Adcode   string `json:"adcode"`
			District string `json:"-"`
			Country  string `json:"country"`
		} `json:"addressComponent"`
	} `json:"regeocode"`
}

type ReGeoOut struct {
	Adcode   int    `json:"adcode"`
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
}

func (m *geoUtil) GaodeReGeo(lon, lat float64) (*ReGeoOut, error) {
	req, _ := http.NewRequest("GET", GaodeReGeoApiUrl, nil)
	q := req.URL.Query()
	q.Add("key", GaodeApiKey)
	q.Add("location", fmt.Sprintf("%.5f,%.5f", lon, lat))
	q.Add("extensions", "base")
	q.Add("output", "json")
	req.URL.RawQuery = q.Encode()

	resp, err := DefaultHttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %s", err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	result := new(GaodeReGeoResult)
	if err := jsoniter.Unmarshal(bs, &result); err != nil {
		return nil, fmt.Errorf("解码失败: %s", err)
	}

	if result.Status != "1" {
		return nil, fmt.Errorf("非预期的结果: %s", result.Info)
	}

	jve := zjve2.Load(bs)
	addressComponent := jve.Get("regeocode.addressComponent")
	if addressComponent.Type() != zjve2.Object {
		return nil, fmt.Errorf("非预期的结果: %s, %s", addressComponent.Path(), addressComponent.Type())
	}
	province := addressComponent.Get("province")
	if province.Type() == zjve2.String {
		result.Regeocode.AddressComponent.Province = province.MustStr().Val()
	}
	if m.IsMunicipalities(result.Regeocode.AddressComponent.Province) {
		result.Regeocode.AddressComponent.City = result.Regeocode.AddressComponent.Province
	} else {
		city := addressComponent.Get("city")
		if city.Type() == zjve2.String {
			result.Regeocode.AddressComponent.City = city.MustStr().Val()
		}
	}
	district := addressComponent.Get("district")
	if district.Type() == zjve2.String {
		result.Regeocode.AddressComponent.District = district.MustStr().Val()
	}
	return &ReGeoOut{
		Adcode:   zstr.ToIntDefault(result.Regeocode.AddressComponent.Adcode),
		Nation:   result.Regeocode.AddressComponent.Country,
		Province: result.Regeocode.AddressComponent.Province,
		City:     result.Regeocode.AddressComponent.City,
		District: result.Regeocode.AddressComponent.District,
	}, nil
}

type IPPositionOut struct {
	Province string `json:"province"`
	City     string `json:"city"`
}

func (m *geoUtil) GaodeIPPosition(ip string) (*IPPositionOut, error) {
	req, _ := http.NewRequest("GET", GaodeIPPositionApiUrl, nil)
	q := req.URL.Query()
	q.Add("key", GaodeApiKey)
	q.Add("output", "json")
	q.Add("ip", ip)
	req.URL.RawQuery = q.Encode()

	resp, err := DefaultHttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %s", err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	jve := zjve2.Load(bs)
	status := jve.Get("status")
	if status.Type() != zjve2.String {
		return nil, fmt.Errorf("解码失败, status类型为 %s, 值为 %s", status.Type(), status.String())
	}
	if status.MustStr().Val() != "1" {
		return nil, fmt.Errorf("非预期的结果: %v", jve.Get("info"))
	}

	var province string
	{
		a := jve.Get("province")
		if a.Type() != zjve2.String {
			return nil, fmt.Errorf("解码失败, province的类型为 %s, 值为 %s", a.Type(), a.String())
		}
		province = a.MustStr().Val()
	}
	city := province
	if !m.IsMunicipalities(province) {
		a := jve.Get("city")
		if a.Type() != zjve2.String {
			return nil, fmt.Errorf("解码失败, city的类型为 %s, 值为 %s", a.Type(), a.String())
		}
		city = a.MustStr().Val()
	}
	return &IPPositionOut{Province: province, City: city}, nil
}
