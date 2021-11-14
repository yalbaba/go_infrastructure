package liveearth_primary

import "time"

type DataProvinceDistribution struct {
	Id      int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Ah      int       `json:"ah" gorm:"column:ah"`             // 安徽
	Bj      int       `json:"bj" gorm:"column:bj"`             // 北京
	Fj      int       `json:"fj" gorm:"column:fj"`             // 福建
	Gs      int       `json:"gs" gorm:"column:gs"`             // 甘肃
	Gd      int       `json:"gd" gorm:"column:gd"`             // 广东
	Gx      int       `json:"gx" gorm:"column:gx"`             // 广西
	Gz      int       `json:"gz" gorm:"column:gz"`             // 贵州
	Hi      int       `json:"hi" gorm:"column:hi"`             // 海南
	He      int       `json:"he" gorm:"column:he"`             // 河北
	Ha      int       `json:"ha" gorm:"column:ha"`             // 河南
	Hl      int       `json:"hl" gorm:"column:hl"`             // 黑龙江
	Hb      int       `json:"hb" gorm:"column:hb"`             // 湖北
	Hn      int       `json:"hn" gorm:"column:hn"`             // 湖南
	Jl      int       `json:"jl" gorm:"column:jl"`             // 吉林
	Js      int       `json:"js" gorm:"column:js"`             // 江苏
	Jx      int       `json:"jx" gorm:"column:jx"`             // 江西
	Ln      int       `json:"ln" gorm:"column:ln"`             // 辽宁
	Nm      int       `json:"nm" gorm:"column:nm"`             // 内蒙古
	Nx      int       `json:"nx" gorm:"column:nx"`             // 宁夏
	Qh      int       `json:"qh" gorm:"column:qh"`             // 青海
	Sd      int       `json:"sd" gorm:"column:sd"`             // 山东
	Sx      int       `json:"sx" gorm:"column:sx"`             // 山西
	Sn      int       `json:"sn" gorm:"column:sn"`             // 陕西
	Sh      int       `json:"sh" gorm:"column:sh"`             // 上海
	Sc      int       `json:"sc" gorm:"column:sc"`             // 四川
	Tj      int       `json:"tj" gorm:"column:tj"`             // 天津
	Xz      int       `json:"xz" gorm:"column:xz"`             // 西藏
	Xj      int       `json:"xj" gorm:"column:xj"`             // 新疆
	Yn      int       `json:"yn" gorm:"column:yn"`             // 云南
	Zj      int       `json:"zj" gorm:"column:zj"`             // 浙江
	Cq      int       `json:"cq" gorm:"column:cq"`             // 重庆
	Mo      int       `json:"mo" gorm:"column:mo"`             // 澳门
	Hk      int       `json:"hk" gorm:"column:hk"`             // 香港
	Tw      int       `json:"tw" gorm:"column:tw"`             // 台湾
	Date    time.Time `json:"date" gorm:"column:date"`         // 日期
	AppType int8      `json:"app_type" gorm:"column:app_type"` // app终端类型
}

func (m *DataProvinceDistribution) TableName() string {
	return "liveearth_primary.data_province_distribution"
}
