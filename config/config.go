package config

import (
	"fmt"
)

type Config struct {
	GRPCPort int
}

func (c Config) GRPCAddress() string {
	return fmt.Sprintf(":%d", c.GRPCPort)
}
