package jsondemo

import (
	"encoding/json"
	"fmt"
)

const jsonStr = `{"id":34,"username":"xxxxx username","password":"mypassword is password","listObj":["111","222","333","444"]}`

// JsonFromString 一个对象转为json string
func JsonFromString() TestJsonObj {
	jsonObj := TestJsonObj{}
	err := json.Unmarshal([]byte(jsonStr), &jsonObj)
	if err != nil {
		fmt.Println(err)
	}
	return jsonObj
}

// StringFromJson string 转为json对象
func StringFromJson(obj TestJsonObj) string {
	resp, _ := json.Marshal(obj)
	return string(resp)
}
