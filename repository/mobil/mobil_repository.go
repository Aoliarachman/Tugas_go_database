package mobil

import (
	"context"
	"go_database/entity"
)

type MobilRepository interface {
	Insert(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error)
	FindById(ctx context.Context, id int32) (entity.Mobil, error)
	FindAll(ctx context.Context) ([]entity.Mobil, error)
	Update(ctx context.Context, id int32, mobil entity.Mobil) (entity.Mobil, error)
	Delete(ctx context.Context, id int32) (string, error)
}
