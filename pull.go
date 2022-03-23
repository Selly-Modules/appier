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
