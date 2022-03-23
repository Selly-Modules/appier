package appier

import (
	"errors"
	"fmt"

	"github.com/Selly-Modules/natsio"
)

// Client ...
type Client struct {
	Config        Config
	natsServer    natsio.Server
	natsJetStream natsio.JetStream
}

var (
	client *Client
)

// NewClient
// Init client ...
func NewClient(config Config) (*Client, error) {
	if config.Nats.URL == "" {
		return nil, errors.New("nats url is required")
	}

	if err := natsio.Connect(config.Nats); err != nil {
		return nil, fmt.Errorf("nats connect failed: %v", err)
	}

	client = &Client{
		Config:        config,
		natsServer:    natsio.GetServer(),
		natsJetStream: natsio.GetJetStream(),
	}

	return client, nil
}

// publishWithJetStream ...
func publishWithJetStream(streamName, subject string, data []byte) (isPublished bool, err error) {
	var req = RequestBody{Body: data}

	// Publish jet stream
	if err = client.natsJetStream.Publish(streamName, subject, toBytes(req)); err != nil {
		return
	}

	isPublished = true
	return
}
