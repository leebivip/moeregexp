package moeregexp

import (
	"log"
	"testing"
)

func Test_IsMatch(t *testing.T) {
	path := "/mongo/wfw.eff/sdf李鹏"
	regstr := "^/mongo/[^:?&*/]+/[^:?&*/]+$"
	if IsMatch(regstr, path) == false {
		panic("路径格式不正确,应该为:/DB/C")
	}
	log.Print(IsMatch(regstr, path))
}
