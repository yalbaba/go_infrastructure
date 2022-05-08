package utils

import (
	"fmt"
	"strconv"
)

var Convert = new(convertUtil)

type convertUtil struct{}

func (*convertUtil) IntsToInt64s(v ...int) []int64 {
	out := make([]int64, len(v))
	for i, a := range v {
		out[i] = int64(a)
	}
	return out
}
func (*convertUtil) IntsToString(v []int) []string {
	out := make([]string, len(v))
	for i, a := range v {
		out[i] = strconv.FormatInt(int64(a), 10)
	}
	return out
}
func (*convertUtil) IntsToInterfaces(v ...int) []interface{} {
	out := make([]interface{}, len(v))
	for i, a := range v {
		out[i] = a
	}
	return out
}

func (*convertUtil) Ints64ToInts(v ...int64) []int {
	out := make([]int, len(v))
	for i, a := range v {
		out[i] = int(a)
	}
	return out
}
func (*convertUtil) Ints64ToStrings(v ...int64) []string {
	out := make([]string, len(v))
	for i, a := range v {
		out[i] = strconv.FormatInt(a, 10)
	}
	return out
}
func (*convertUtil) Int64sToInterface(v []int64) []interface{} {
	out := make([]interface{}, len(v))
	for i, a := range v {
		out[i] = a
	}
	return out
}

func (*convertUtil) StringsToInts(ss ...string) ([]int, error) {
	ints := make([]int, len(ss))
	for i, s := range ss {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("[%s]不能转为int", s)
		}
		ints[i] = int(v)
	}
	return ints, nil
}
func (*convertUtil) StringsToInts64(ss ...string) ([]int64, error) {
	ints := make([]int64, len(ss))
	for i, s := range ss {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("[%s]不能转为int", s)
		}
		ints[i] = v
	}
	return ints, nil
}
func (*convertUtil) StringsToInterfaces(v ...string) []interface{} {
	out := make([]interface{}, len(v))
	for i, a := range v {
		out[i] = a
	}
	return out
}
