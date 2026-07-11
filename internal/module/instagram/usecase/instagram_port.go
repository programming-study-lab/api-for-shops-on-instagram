package usecase

import (
	"context"
)

type InstagramUsecase interface {
	InstagramGetInfo(ctx context.Context) (*[]map[string]any, error)
}
