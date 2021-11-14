package liveearth_primary

import "time"

type DataPortraitDistribution struct {
	Id                   int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	EighteenTwentyfour   int       `json:"eighteen_twentyfour" gorm:"column:eighteen_twentyfour"`     // 18到24的比例
	TwentyfiveTwentynine int       `json:"twentyfive_twentynine" gorm:"column:twentyfive_twentynine"` // 25到29的比例
	ThirtyThirtyfour     int       `json:"thirty_thirtyfour" gorm:"column:thirty_thirtyfour"`         // 30到34的比例
	ThirtyfiveThirtynine int       `json:"thirtyfive_thirtynine" gorm:"column:thirtyfive_thirtynine"` // 35到39的比例
	FortyFortynine       int       `json:"forty_fortynine" gorm:"column:forty_fortynine"`             // 40到49的比例
	OverFifty            int       `json:"over_fifty" gorm:"column:over_fifty"`                       // 50以上的比例
	OneLine              int       `json:"one_line" gorm:"column:one_line"`                           // 一线城市比例
	TwoLine              int       `json:"two_line" gorm:"column:two_line"`                           // 二线
	ThreeLine            int       `json:"three_line" gorm:"column:three_line"`                       // 三线
	FourLine             int       `json:"four_line" gorm:"column:four_line"`                         // 四线
	FiveLine             int       `json:"five_line" gorm:"column:five_line"`                         // 五线
	SixLine              int       `json:"six_line" gorm:"column:six_line"`                           // 六线
	Males                int       `json:"males"`
	Females              int       `json:"females"`
	AppType              int8      `json:"app_type" gorm:"column:app_type"` // app类型：20：安卓，30：ios   10：全部
	Date                 time.Time `json:"date" gorm:"column:date"`         // 日期
}

func (m *DataPortraitDistribution) TableName() string {
	return "liveearth_primary.data_portrait_distribution"
}
