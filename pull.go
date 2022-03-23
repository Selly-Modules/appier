package appier

// Pull ...
type Pull struct{}

// ProductUpsert ...
func (Pull) ProductUpsert(payload []byte) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullProductUpsert, payload)
}
