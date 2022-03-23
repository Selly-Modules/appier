package appier

import "encoding/json"

// toBytes ...
func toBytes(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}
