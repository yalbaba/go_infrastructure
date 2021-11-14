package consts

import (
	"testing"
)

func TestPosition_MakeShowAddress(t *testing.T) {
	p := &Position{Foreign: false, Nation: "中国", Province: "台湾", City: "嘉怡"}

	p.MakeShowAddress()

	t.Log(p)
}
