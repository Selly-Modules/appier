package appier

import (
	"encoding/json"
	"fmt"
)

// toBytes ...
func toBytes(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}

// getRedisKey ...
func getRedisKey(prefix string, targetID string) string {
	return fmt.Sprintf("%s%s_%s", RedisSyncAppierPrefix, prefix, targetID)
}
