package appier

import (
	"errors"
	"fmt"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/redisdb"
)

// Client ...
type Client struct {
	Config        Config
	natsServer    natsio.Server
	natsJetStream natsio.JetStream
	Sync          Sync
	Pull          Pull
}

var client *Client

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

	// Connect redis
	if err := redisdb.Connect(config.Redis.URI, config.Redis.Password); err != nil {
		return nil, fmt.Errorf("redis connect failed: %v", err)
	}

	// Init schedule
	initSchedule()
	return client, nil
}

// publishWithJetStream ...
func publishWithJetStream(streamName, subject string, data []byte) (isPublished bool, err error) {

	// Publish jet stream
	if err = client.natsJetStream.Publish(streamName, subject, data); err != nil {
		return
	}

	isPublished = true
	return
}
