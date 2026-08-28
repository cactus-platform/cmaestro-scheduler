package collection

import (
	"context"

	"github.com/cactus-platform/cmaestro-scheduler/internal/config"
)

type Services struct {
}

func NewServices(ctx context.Context, runtimeConfig *config.RuntimeConfig, conns *Connections, repos *Repositories) (*Services, error) {
	return &Services{}, nil
}
