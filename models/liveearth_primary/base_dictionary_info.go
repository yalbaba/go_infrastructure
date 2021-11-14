package liveearth_primary

// BaseDictionaryInfo ...
type BaseDictionaryInfo struct {
	Id uint32 `json:"id"`
	// Name 名称
	Name string `json:"name"`
	// Value 值
	Value string `json:"value"`
	// TypeCode 分类code
	TypeCode string `json:"type_code"`
	// TypeName 分类名称
	TypeName string `json:"type_name"`
	SortId   uint8  `json:"sort_id"`
	// State 状态: 10 = 禁用; 20 = 启用
	State uint8 `json:"state"`
}

func (*BaseDictionaryInfo) TableName() string {
	return "liveearth_primary.base_dictionary_info"
}
