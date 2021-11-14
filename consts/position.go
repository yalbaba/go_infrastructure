package consts

type PositionType int8

const (
	Area  PositionType = 10 * (iota + 1) // 区域
	Point                                // 点位
)

type Position struct {
	Foreign      bool         `json:"foreign"`       // true:国外
	PositionType PositionType `json:"position_type"` // 地址类型
	Lon          string       `json:"lon"`
	Lat          string       `json:"lat"`
	Nation       string       `json:"nation"`
	Province     string       `json:"province"`
	City         string       `json:"city"`
	County       string       `json:"county"`
	Town         string       `json:"town"`
	Address      string       `json:"address"`       // 城市下一级的详细地址
	FlyToHeight  int          `json:"fly_to_height"` // fly_to 高度, [发现]需要
	Distance     int          `json:"distance"`      // 距离
	ShowAddress  string       `json:"show_address"`  // 计算出用于展示的地址信息
	Continent    string       `json:"continent"`
}

func (p *Position) MakeShowAddress() *Position {

	if p.Foreign {
		if p.Nation != "" && p.City != "" {
			p.ShowAddress = p.Nation + "." + p.City
			return p
		}

		if p.Nation != "" {
			if p.Province != "" && p.City == "" {
				p.ShowAddress = p.Nation + "." + p.Province
				return p
			}

			if p.Province == "" && p.City == "" {
				p.ShowAddress = p.Nation
				return p
			}
		}
		return p
	}

	if p.Province != "" && p.City != "" {
		if p.Province == p.City {
			p.ShowAddress = p.Province
			return p
		}
		p.ShowAddress = p.Province + "." + p.City
		return p
	}

	if p.Province != "" && p.City == "" {
		p.ShowAddress = p.Province
		return p
	}

	if p.Province == "" && p.City != "" {
		p.ShowAddress = p.City
		return p
	}

	if p.Province == "" && p.City == "" {
		p.ShowAddress = p.Nation
		return p
	}

	return p
}
