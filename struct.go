package appier

// RequestBody ...
type RequestBody struct {
	Body []byte `json:"body"`
}

// Payload ...
type Payload struct {
	Data []byte
}

// RedisCfg ...
type RedisCfg struct {
	URI      string
	Password string
}
