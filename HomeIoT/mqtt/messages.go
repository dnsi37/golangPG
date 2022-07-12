package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
)

type LedOrder struct {
	Order bool
	json  *json.RawMessage
}

func (l *LedOrder) On() *LedOrder {
	jsonMsg := json.RawMessage(`{"order" : "on" }`)
	newOrder := &LedOrder{true, &jsonMsg}
	return newOrder
}

func (l *LedOrder) Off() *LedOrder {
	jsonMsg := json.RawMessage(`{"order" : "off" }`)
	newOrder := &LedOrder{false, &jsonMsg}
	return newOrder
}

func JsonTest () {

	jsonMsg := json.RawMessage(`{"success":true,"validationErrors":null,"message":null,"code":0}`)
	stringMsg :=`{"success":true,"validationErrors":null,"message":null,"code":0}`
	var buffjson struct {
		Success bool
	}
	err := json.Unmarshal(jsonMsg,&buffjson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonMsg)
	fmt.Println([]byte(stringMsg))
	

}