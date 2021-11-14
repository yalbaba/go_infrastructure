package liveearth_primary

type LiveWeather struct {
	AreaId      int    `json:"area_id" xorm:"not null pk comment('地区id') INT(11)"`
	WeatherType string `json:"weather_type" xorm:"not null comment('天气类型') VARCHAR(8)"`
	WindDir     string `json:"wind_dir" xorm:"not null comment('风向') VARCHAR(8)"`
	WindScale   string `json:"wind_scale" xorm:"not null comment('风级') VARCHAR(8)"`
	Temp        string `json:"temp" xorm:"not null comment('温度') VARCHAR(8)"`
}

func (*LiveWeather) TableName() string {
	return "liveearth_primary.live_weather"
}
