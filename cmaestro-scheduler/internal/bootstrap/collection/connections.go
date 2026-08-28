package collection

import (
	"context"
	"errors"
	"log"

	"github.com/cactus-platform/scheduler/internal/config"
)

type Connections struct {
	closeFns []func() error
}

func NewCollections(ctx context.Context, runtimeConfig *config.RuntimeConfig) (*Connections, error) {
	conn := &Connections{}

	success := false
	defer func() {
		if !success {
			_ = conn.Close()
		}
	}()

	success = true
	log.Println("Connections created!")
	return conn, nil
}

func (conn *Connections) addCloser(eventName string, value any) {
	log.Printf("adding [%s] to pools.event.close", eventName)
	closer, ok := value.(interface {
		Close() error
	})
	if ok {
		conn.closeFns = append(conn.closeFns, closer.Close)
	}
}

func closeIfPossible(value any) {
	closer, ok := value.(interface {
		Close() error
	})
	if ok {
		_ = closer.Close()
	}
}

func (conn *Connections) Close() error {
	log.Println("closing connections...")
	if conn == nil {
		return nil
	}

	var errs []error

	for i := len(conn.closeFns) - 1; i >= 0; i-- {
		if err := conn.closeFns[i](); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
