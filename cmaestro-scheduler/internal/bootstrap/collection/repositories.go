package collection

import (
	"context"

	"github.com/cactus-platform/cmaestro-scheduler/internal/config"
)

type Repositories struct {
}

func NewRepositories(ctx context.Context, runtimeConfig *config.RuntimeConfig, conns *Connections) (*Repositories, error) {
	return &Repositories{}, nil
}
