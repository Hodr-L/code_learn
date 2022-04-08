package unjson

import (
	"testing"
)

//func TestStore(t *testing.T) {
//	//先创建一个Monster实例
//	monster := Monster{
//		Name:  "红孩儿",
//		Age:   10,
//		Skill: "吐火",
//	}
//	res := monster.Store()
//	if !res {
//		t.Fatalf("错误，期望值：%v 实际值：%v", true, res)
//	}
//	t.Logf("测试成功！")
//}

func TestRestore(t *testing.T) {
	var monster Monster
	res := monster.Restore()
	if !res {
		t.Fatalf("错误，期望值：%v 实际值：%v", true, res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("错误，期望值：%v 实际值：%v", "红孩儿", monster.Name)
	}
	t.Logf("测试成功！")
}
