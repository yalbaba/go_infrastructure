package liveearth_primary

type DataWeekUserDistribution struct {
	Id                 int  `json:"id" xorm:"not null pk autoincr INT(10)"`
	NewUser            int  `json:"new_user" gorm:"column:new_user"`
	BackUser           int  `json:"back_user" gorm:"column:back_user"`
	OneWeekActiveUser  int  `json:"one_week_active_user" gorm:"column:one_week_active_user"`
	TwoWeekActiveUser  int  `json:"two_week_active_user" gorm:"column:two_week_active_user"`
	FourWeekActiveUser int  `json:"four_week_active_user" gorm:"column:four_week_active_user"`
	LoyalUser          int  `json:"loyal_user" gorm:"column:loyal_user"`
	Sort               int  `json:"sort" gorm:"column:sort"`
	AppType            int8 `json:"app_type" gorm:"column:app_type"` // 20：安卓，30：ios
}

func (m *DataWeekUserDistribution) TableName() string {
	return "liveearth_primary.data_week_user_distribution"
}
