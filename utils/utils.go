package utils

import "encoding/json"

func JsonMarshalSilence(object interface{}) string {
	res, _ := json.Marshal(object)
	return string(res)
}
