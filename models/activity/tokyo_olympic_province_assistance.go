package activity

type TokyoOlympicProvinceAssistance struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Province   string `json:"province" gorm:"column:province"`     // 省份
	Assistance int    `json:"assistance" gorm:"column:assistance"` // 助力数
}

func (m *TokyoOlympicProvinceAssistance) TableName() string {
	return "activity.tokyo_olympic_province_assistance"
}
