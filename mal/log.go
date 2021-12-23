package mal

import "go.uber.org/zap"

var log *zap.Logger

func init() {
	log = zap.L()
}

func SetLogger(logger *zap.Logger) {
	log = logger
}
