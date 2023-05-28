package ujson

import "encoding/json"

func JsonString(val interface{}) string {
	bs, _ := json.Marshal(val)
	return string(bs)
}
