package liveearth_primary

type BirdsViewDevice struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	DeviceCode string `json:"device_code" gorm:"column:device_code"` // 设备id
	DeviceName string `json:"device_name" gorm:"column:device_name"` // 设备名称
	StreamName string `json:"stream_name"`
}

func (m *BirdsViewDevice) TableName() string {
	return "liveearth_primary.birds_view_device"
}
