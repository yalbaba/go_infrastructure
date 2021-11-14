package liveearth_primary

type AdminPushTemplate struct {
	Id int `json:"id" xorm:"not null pk autoincr comment('pk') INT(10)"`
}

func (*AdminPushTemplate) TableName() string {
	return "liveearth_primary.admin_push_template"
}
