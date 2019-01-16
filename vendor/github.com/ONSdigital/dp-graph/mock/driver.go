package mock

import (
	"context"
)

type Mock struct {
	IsBackendUnreachable bool
	IsQueryInvalid       bool
	IsContentNotFound    bool
}

func (m *Mock) Close(ctx context.Context) error {
	return nil
}
