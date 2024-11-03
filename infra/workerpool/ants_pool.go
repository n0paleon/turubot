package workerpool

import (
	"github.com/panjf2000/ants/v2"
	"turubot/internal/domain"
)

var Pool *ants.Pool

func InitializePool(size int, logger domain.Logger) error {
	if size <= 0 {
		size = 10
		logger.Warn("invalid pool size, using 10 as default workerpool size")
	}

	pool, err := ants.NewPool(size, ants.WithPreAlloc(true))
	if err != nil {
		logger.Errorw("failed to initialize pool", "error", err)
		return err
	}

	Pool = pool
	logger.Infow("pool initialized", "size", size)

	return nil
}

func ClosePool() {
	Pool.Release()
}
