package guide

import "time"

type LandscapeGoodsRelationship struct {
	Id              int       `json:"id" gorm:"column:id" xorm:"not null pk autoincr INT(10)"`
	LandscapeViewId int64     `json:"landscape_view_id" gorm:"column:landscape_view_id"` // 景区视图id
	GoodsId         int64     `json:"goods_id" gorm:"column:goods_id"`                   // 商品id
	GoodsTitle      string    `json:"goods_title" gorm:"column:goods_title"`             // 商品名称
	CoverUrl        string    `json:"cover_url" gorm:"column:cover_url"`                 // 商品封面图片
	ItemImages      string    `json:"item_images" gorm:"column:item_images"`             // 图片详情
	DetailUrl       string    `json:"detail_url" gorm:"column:detail_url"`               // 商品链接
	Norm            string    `json:"norm" gorm:"column:norm"`                           // 规格
	GoodsPrice      int       `json:"goods_price" gorm:"column:goods_price"`             // 商品价格(单位:分)
	OriginPrice     string    `json:"origin_price" gorm:"column:origin_price"`           // 划线价格(单位:元)
	SortId          int       `json:"sort_id" gorm:"column:sort_id"`                     // 排序
	UpdateTime      time.Time `json:"update_time" gorm:"column:update_time"`
	GoodsType       int8      `json:"goods_type" gorm:"column:goods_type"`         // 商品类别（10：景区商品，20：公共商品）(暂时保留)
	GoodsPlatform   int8      `json:"goods_platform" gorm:"column:goods_platform"` // 商品所属平台(10：有赞，20：淘宝，30：拼多多)
	Recommend       int8      `json:"recommend" gorm:"column:recommend"`           // 是否推荐(20：批量推荐，19：直播间自己推荐，10取消推荐) 暂时保留
	Start           time.Time `json:"start" gorm:"column:start"`                   // 推荐开始时间：暂时保留
	End             time.Time `json:"end" gorm:"column:end"`                       // 推荐结束时间：暂时保留
	CreateTime      time.Time `json:"create_time" gorm:"column:create_time"`       // 创建时间
	State           int8      `json:"state" gorm:"column:state"`                   // 状态：10：销售中，20：已删除
	Operator        string    `json:"operator" gorm:"column:operator"`             // 操作人名称
}

func (m *LandscapeGoodsRelationship) TableName() string {
	return "guide.landscape_goods_relationship"
}
