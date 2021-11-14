package utils

import jsoniter "github.com/json-iterator/go"

func StructToMap(o interface{}) (map[string]interface{}, error) {

	mp := make(map[string]interface{})

	bt, err := jsoniter.Marshal(o)
	if err != nil {
		return nil, err
	}

	if err := jsoniter.Unmarshal(bt, &mp); err != nil {
		return nil, err
	}

	return mp, nil
}
