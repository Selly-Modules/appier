package appier

import (
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
	for {
		keys, values := redisdb.GetWithPrefixPattern(pattern)
		if len(keys) == 0 {
			fmt.Println("No keys data")
			return
		}

		if len(values) == 0 {
			fmt.Println("No values DATA")
			return
		}

		processSyncToServiceAppier(typeData, values)

		//// Del keys
		//for _, key := range keys {
		//	redisdb.DelKey(context.Background(), key)
		//}
	}
}

// processSyncToServiceAppier ...
func processSyncToServiceAppier(typeData string, values []string) {
	fmt.Println("Sends")
	for {
		if len(values) <= 20 {
			syncDataByList(typeData, values)
			break
		}

		var list = values[0:19]
		syncDataByList(typeData, list)
		values = values[20:]
	}
	return
}

// syncDataByList ...
func syncDataByList(typeData string, data []string) {
	payload := toBytes(data)
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
}
