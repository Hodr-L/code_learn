package sum

import (
	"strings"
)

func Makesuffix(suffix string) func(string) string {

	return func(name string) string {
		if !strings.HasSuffix(name, suffix) { //如果name没有指定后缀，则加上，否则返回原来的名字
			return name + suffix
		}
		return name
	}
}
