package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"` //反射机制
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func main() {
	// 结构体 json
	var monster Monster = Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-01",
		Sal:      6000.1,
		Skill:    "牛魔拳",
	}
	monsterJson, err := json.Marshal(monster)
	if err == nil {
		fmt.Printf("json encode: %s \n", monsterJson)
	}
	// map json
	var mapData map[string]interface{} = make(map[string]interface{})
	mapData["name"] = "红孩儿"
	mapData["age"] = 123
	mapJson, err1 := json.Marshal(mapData)
	if err1 == nil {
		fmt.Printf("mapjson: %s \n", mapJson)
	}
	// slice json

	var sliceData []map[string]interface{}
	var sm1 = make(map[string]interface{})
	sm1["name"] = "西施"
	sm1["age"] = 20
	sm1["monster"] = monster
	var sm2 = make(map[string]interface{})
	sm2["name"] = "西施"
	sm2["age"] = 20
	sm2["monster"] = monster
	sliceData = append(sliceData, sm1, sm2)
	sliceJson, err3 := json.Marshal(sliceData)
	if err3 == nil {
		fmt.Printf("sliceJson: %s \n", sliceJson)
	}

	// json反序列化
	err4 := json.Unmarshal(sliceJson, &sliceData)
	if err4 == nil {
		fmt.Printf("json unserialize: %v \n", sliceData)
	} else {
		fmt.Printf("unmarshall err:%v \n", err4)
	}

}
