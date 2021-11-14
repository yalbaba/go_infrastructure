package liveearth_primary

// LiveDistrict ...
type LiveDistrict struct {
	// ID 自增数据库id
	Id int64 `json:"id"`
	// Adcode 行政区划代码
	Adcode int32 `json:"adcode"`
	// AreaCode 区号
	AreaCode int32 `json:"area_code"`
	// PostCode 邮政编码
	PostCode int32 `json:"post_code"`
	// Name 行政区划名称
	Name string `json:"name"`
	// DistrictPy 行政区划名称拼音
	DistrictPy string `json:"district_py"`
	// Nation 行政区划层级-国家
	Nation string `json:"nation"`
	// Province 行政区划层级-省
	Province string `json:"province"`
	// City 行政区划层级-市
	City string `json:"city"`
	// County 行政区划层级-区
	County string `json:"county"`
	// Level 行政区划等级
	Level int8 `json:"level"`
	// Altitude flyto高度
	Altitude int32 `json:"altitude"`
	// Lat 纬度
	Lat float64 `json:"lat"`
	// Lon 经度
	Lon float64 `json:"lon"`
	// ResidentPopulation 常住人口（万人）
	ResidentPopulation int32 `json:"resident_population"`
	// Station 驻地所在
	Station string `json:"station"`
	// ResidentArea 区划面积（平方公里）
	ResidentArea int32 `json:"resident_area"`
	// ParentName 行政区划父级名称
	ParentName string `json:"parent_name"`
	// Code 完整行政区划代码
	Code int32 `json:"code"`
	// Address 地址
	Address  string `json:"address"`
	ParentId uint32 `json:"parent_id"`
	// ParentIds 上级区域id
	ParentIds string `json:"parent_ids"`
}

func (*LiveDistrict) TableName() string {
	return "liveearth_primary.live_district"
}
