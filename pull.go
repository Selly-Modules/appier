package appier

// Pull ...
type Pull struct{}

// ProductUpsert ...
func (Pull) ProductUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullProductUpsert, payload)
}

// BrandUpsert ...
func (Pull) BrandUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullBrandUpsert, payload)
}

// CategoryUpsert ...
func (Pull) CategoryUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullCategoryUpsert, payload)
}

// SubCategoryUpsert ...
func (Pull) SubCategoryUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullSubCategoryUpsert, payload)
}

// SupplierUpsert ...
func (Pull) SupplierUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullSupplierUpsert, payload)
}

// InventoryUpsert ...
func (Pull) InventoryUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullInventoryUpsert, payload)
}

// SKUUpsert ...
func (Pull) SKUUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullSKUUpsert, payload)
}

// PropertyUpsert ...
func (Pull) PropertyUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullPropertyUpsert, payload)
}

// PropertyValueUpsert ...
func (Pull) PropertyValueUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullPropertyValueUpsert, payload)
}
