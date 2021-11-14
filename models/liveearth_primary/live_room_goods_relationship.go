package liveearth_primary

import "time"

// LiveRoomGoodsRelationship ...
type LiveRoomGoodsRelationship struct {
	Id int64 `json:"id" xorm:"not null pk autoincr BIGINT(20)"`
	// LiveRoomID 直播间id
	LiveRoomId uint64 `json:"live_room_id"`
	// GoodsID 商品id
	GoodsId uint64 `json:"goods_id"`
	// GoodsTitle 商品名称
	GoodsTitle string `json:"goods_title"`
	// CoverURL 商品封面图片
	CoverUrl string `json:"cover_url"`
	// ItemImages 图片详情
	ItemImages string `json:"item_images"`
	// DetailURL 有赞详情
	DetailUrl string `json:"detail_url"`
	// Norm 规格
	Norm string `json:"norm"`
	// GoodsPrice 商品价格(单位:分)
	GoodsPrice int32 `json:"goods_price"`
	// OriginPrice 划线价(单位:元)
	OriginPrice string `json:"origin_price"`
	// SortID 排序
	SortId     int8      `json:"sort_id"`
	UpdateTime time.Time `json:"update_time"`
	// GoodsType 商品类别（10：直播间商品，20：公共商品）
	GoodsType int8 `json:"goods_type"`
	// GoodsPlatform 商品所属平台(10：有赞，20：淘宝，30：拼多多)
	GoodsPlatform int8 `json:"goods_platform"`
	// Recommend 是否推荐(0：是，1：否)
	Recommend int8 `json:"recommend"`
	// 用户终端：10：h5，20：安卓，30：ios
	TerminalType string `json:"terminal_type"`
	// Start 推荐开始时间
	Start time.Time `json:"start"`
	// End 推荐结束时间
	End time.Time `json:"end"`
}

func (*LiveRoomGoodsRelationship) TableName() string {
	return "liveearth_primary.live_room_goods_relationship"
}
