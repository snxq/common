package common

import "encoding/json"

// ConvertByJSON 通过中间Json编码将x赋值给y
func ConvertByJSON(x, y interface{}) error {
	b, err := json.Marshal(x)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, y); err != nil {
		return err
	}
	return nil
}
