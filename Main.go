package main

import (
	"encoding/json"
	"syscall/js"
	"time"
	"webassembly/encrypt"
)
const key = "ZWdgbSRWh7YlXjZV96qOqyOjbYZiiHJJ"
func JsonEncode(data interface{}) string {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	return string(jsonByte)
}
func encryptFace() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		inputJSON := args[0].String()
		data:=map[string]interface{}{
			"json":inputJSON,
			"time":time.Now().Unix(),
		}
		aesEncrypt := encrypt.AesEncrypt{}
		encrypted,_:= aesEncrypt.Encrypt(JsonEncode(data), key)
		return encrypted
	})
	return jsonFunc
}

func main() {
	js.Global().Set("beforeSend", encryptFace())
	select {

	}
}