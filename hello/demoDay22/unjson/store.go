package unjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {

	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("error=", err)
		return false
	}
	filePath := "f:/hello.txt"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("error=", err)
		return false
	}
	return true
}

func (this *Monster) Restore() bool {
	filePath := "f:/hello.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("error=", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("error", err)
		return false
	}
	return true
}
