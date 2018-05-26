package util

import (
	"encoding/json"
	"fmt"
)

func JsonEncode(v interface{}) []byte {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		//TODO panic
		err = fmt.Errorf("JsonEncode to json failed, %v", err)
		//return nil
		panic(err)
	}
	return b
}

func JsonDecode(v interface{}, b []byte) error {
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	return nil
}
