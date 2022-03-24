package appier

const (
	SubjectPullProductUpsert       = "selly.pull.product.upsert"
	SubjectPullBrandUpsert         = "selly.pull.brand.upsert"
	SubjectPullCategoryUpsert      = "selly.pull.category.upsert"
	SubjectPullSubCategoryUpsert   = "selly.pull.subcategory.upsert"
	SubjectPullInventoryUpsert     = "selly.pull.inventory.upsert"
	SubjectPullSupplierUpsert      = "selly.pull.supplier.upsert"
	SubjectPullPropertyUpsert      = "selly.pull.property.upsert"
	SubjectPullPropertyValueUpsert = "selly.pull.property_value.upsert"
	SubjectPullSKUUpsert           = "selly.pull.sku.upsert"
)

const (
	JetStreamAppierService = "Service_Appier"
)

const (
	RedisSyncAppierPrefix  = "appier_sync_"
	RedisSyncProduct       = "product"
	RedisSyncBrand         = "brand"
	RedisSyncCategory      = "category"
	RedisSyncSubCategory   = "subcategory"
	RedisSyncInventory     = "inventory"
	RedisSyncSupplier      = "supplier"
	RedisSyncProperty      = "property"
	RedisSyncPropertyValue = "property_value"
	RedisSyncSKU           = "sku"
)
