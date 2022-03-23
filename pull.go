package appier

// Pull ...
type Pull struct{}

// ProductUpsert ...
func (Pull) ProductUpsert(payload Payload) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectRequestProductUpsert, toBytes(payload))
}
