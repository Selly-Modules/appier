package appier

// Pull ...
type Pull struct{}

// ProductUpsert ...
func (Pull) ProductUpsert(payload Payload) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPullProductUpsert, payload.Data)
}

// PingService ...
func (Pull) PingService(payload Payload) (bool, error) {
	return publishWithJetStream(JetStreamAppierService, SubjectPingPullService, payload.Data)
}
