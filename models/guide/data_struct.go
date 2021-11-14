/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 */

package guide

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/twpayne/go-geom"

	"github.com/twpayne/go-geom/encoding/wkb"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JSON json.RawMessage

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// 通过自定义类型创建记录
type Point struct {
	Lon string
	Lat string
}

func (loc *Point) GeomFromText() string {
	return fmt.Sprintf("GeomFromText('Point(%v %v)')", loc.Lon, loc.Lat)
}

func (loc *Point) Scan(v interface{}) error {

	if b, ok := v.([]byte); ok {
		p, err := wkb.Unmarshal(b[4:])
		if err != nil {
			return err
		}

		loc.Lon = strconv.FormatFloat(p.FlatCoords()[0], 'E', -1, 64)
		loc.Lat = strconv.FormatFloat(p.FlatCoords()[1], 'E', -1, 64)
	}

	return nil
}

func (loc Point) GormDataType() string {
	return "point"
}

func (loc Point) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%v %v)", loc.Lon, loc.Lat)},
	}
}

func (p *Point) FromDB(data []byte) error {

	if len(data) == 0 {
		return nil
	}

	if len(data) != 25 {
		return errors.New("point data is incomplete")
	}

	point, err := wkb.Unmarshal(data[4:])
	if err != nil {
		return fmt.Errorf("point data parser err: %s", err)
	}
	vs := point.FlatCoords()
	if len(vs) != 2 {
		return fmt.Errorf("point data parser ok, but point size need 2, got %d", len(point.FlatCoords()))
	}
	p.Lon = fmt.Sprint(vs[0])
	p.Lat = fmt.Sprint(vs[1])

	return nil
}

func (p *Point) ToDB() ([]byte, error) {

	lon, err := strconv.ParseFloat(p.Lon, 64)
	if err != nil {
		return nil, err
	}
	lat, err := strconv.ParseFloat(p.Lat, 64)
	if err != nil {
		return nil, err
	}

	point := geom.NewPointFlat(geom.XY, []float64{lon, lat})
	data, err := wkb.Marshal(point, wkb.NDR)
	fmt.Println("p.ToDB", data)
	data = append([]byte{0, 0, 0, 0}, data...)
	return data, err
}
