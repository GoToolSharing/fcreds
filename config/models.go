package config

import (
	"go.uber.org/zap"
)

type Settings struct {
	Verbose int
	Logger  *zap.Logger
}
