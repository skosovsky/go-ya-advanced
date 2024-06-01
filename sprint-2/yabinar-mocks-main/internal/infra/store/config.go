package store

import (
	"mocks/internal/infra/store/memory"
	"mocks/internal/infra/store/redis"
)

type Config struct {
	Memory *memory.Config
	Redis  *redis.Config
}
