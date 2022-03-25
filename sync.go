package appier

import (
	"context"

	"github.com/Selly-Modules/redisdb"
)

// Sync ...
type Sync struct{}

// SyncProduct ...
func (Sync) SyncProduct(productID string, product interface{}) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncProduct, productID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, product)
	return
}

// SyncBrand ...
func (Sync) SyncBrand(brandID string, brand interface{}) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncBrand, brandID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, brand)
	return
}

// SyncSupplier ...
func (Sync) SyncSupplier(supplierID string, supplier interface{}) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncSupplier, supplierID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, supplier)
	return
}

// SyncInventory ...
func (Sync) SyncInventory(inventoryID string, inventory interface{}) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncInventory, inventoryID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, inventory)
	return
}

// SyncCategory ...
func (Sync) SyncCategory(categoryID string, category interface{}) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncCategory, categoryID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, category)
	return
}

// SyncSubCategory ...
func (Sync) SyncSubCategory(subcategoryID string, value interface{}) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncSubCategory, subcategoryID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, value)
	return
}

// SyncProperty ...
func (Sync) SyncProperty(propertyID string, value interface{}) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncProperty, propertyID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, value)
	return
}

// SyncPropertyValue ...
func (Sync) SyncPropertyValue(propertyValueID string, value interface{}) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncPropertyValue, propertyValueID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, value)
	return
}

// SyncSKU ...
func (Sync) SyncSKU(skuID string, value interface{}) {
	ctx := context.Background()
	key := getRedisKey(RedisSyncSKU, skuID)

	// Set redis
	redisdb.SetKeyValue(ctx, key, value)
	return
}
