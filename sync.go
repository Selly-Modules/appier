package appier

import (
	"context"

	"github.com/Selly-Modules/redisdb"
)

// Sync ...
type Sync struct{}

// SyncProduct ...
func (Sync) SyncProduct(productID string, productJSONString string) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncProduct, productID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, productJSONString)
	return
}

// SyncBrand ...
func (Sync) SyncBrand(brandID, brandJSONString string) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncBrand, brandID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, brandJSONString)
	return
}

// SyncSupplier ...
func (Sync) SyncSupplier(supplierID, supplierJSONString string) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncSupplier, supplierID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, supplierJSONString)
	return
}

// SyncInventory ...
func (Sync) SyncInventory(inventoryID, inventoryJSONString string) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncInventory, inventoryID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, inventoryJSONString)
	return
}

// SyncCategory ...
func (Sync) SyncCategory(categoryID, categoryJSONString string) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncCategory, categoryID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, categoryJSONString)
	return
}

// SyncSubCategory ...
func (Sync) SyncSubCategory(subcategoryID, subCategoryJSONString string) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncSubCategory, subcategoryID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, subCategoryJSONString)
	return
}

// SyncProperty ...
func (Sync) SyncProperty(propertyID, propertyJSONString string) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncProperty, propertyID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, propertyJSONString)
	return
}

// SyncPropertyValue ...
func (Sync) SyncPropertyValue(propertyValueID, propertyValueJSONString string) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncPropertyValue, propertyValueID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, propertyValueJSONString)
	return
}

// SyncSKU ...
func (Sync) SyncSKU(skuID, skuJSONString string) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncSKU, skuID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, skuJSONString)
	return
}
