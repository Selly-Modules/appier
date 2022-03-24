package appier

import "github.com/Selly-Modules/natsio"

// Config ...
type Config struct {
	Redis RedisCfg
	Nats  natsio.Config
}
