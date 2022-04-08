package main

import (
	"testing"
)

func TestAddupper(t *testing.T) {
	res := Adduper(10)
	if res != 55 {
		t.Fatalf("执行错误！期望值：%v 实际值：%v", 55, res)
	}
	t.Logf("执行正确...")

}
