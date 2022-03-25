package appier

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/redisdb"
)

// SyncToService ...
type SyncToService struct{}

const limitGetKeyRedis int64 = 20

var listTypeSync = []string{
	RedisSyncProduct,
	RedisSyncSupplier,
	RedisSyncInventory,
	RedisSyncBrand,
	RedisSyncCategory,
	RedisSyncSubCategory,
	RedisSyncProperty,
	RedisSyncPropertyValue,
}

// syncData ...
func (s SyncToService) syncData() {
	for _, item := range listTypeSync {
		keyProductPattern := getRedisPrefixPattern(item)
		handleSyncDataToServiceAppier(keyProductPattern, item)
	}
	fmt.Println("Sync data to service done!")
}

// handleSyncDataToServiceAppier ...
func handleSyncDataToServiceAppier(pattern string, typeData string) {
	fmt.Println("keyProductPattern", pattern)

	for {
		keys, values := redisdb.GetWithPrefixPattern(pattern, limitGetKeyRedis)
		if len(keys) == 0 {
			return
		}

		fmt.Println("keys:", keys)
		fmt.Println("values", values)

		payload := toBytes(values)
		// Convert data
		switch typeData {
		case RedisSyncProduct:
			Pull{}.ProductUpsert(payload)
		case RedisSyncSKU:
			Pull{}.SKUUpsert(payload)
		case RedisSyncCategory:
			Pull{}.CategoryUpsert(payload)
		case RedisSyncSubCategory:
			Pull{}.SubCategoryUpsert(payload)
		case RedisSyncInventory:
			Pull{}.InventoryUpsert(payload)
		case RedisSyncBrand:
			Pull{}.BrandUpsert(payload)
		case RedisSyncSupplier:
			Pull{}.SupplierUpsert(payload)
		case RedisSyncProperty:
			Pull{}.PropertyUpsert(payload)
		case RedisSyncPropertyValue:
			Pull{}.PropertyValueUpsert(payload)
		}

		// Del keys
		for _, key := range keys {
			redisdb.DelKey(context.Background(), key)
		}
	}
}
