package kelas

import (
	"context"
	"go_database/entity"
)

type KelasRepository interface {
	Insert(ctx context.Context, Kelas entity.Kelas) (entity.Kelas, error)
	FindById(ctx context.Context, id int32) (entity.Kelas, error)
	FindAll(ctx context.Context) ([]entity.Kelas, error)
	Update(ctx context.Context, id int32, kelas entity.Kelas) (entity.Kelas, error)
	Delete(ctx context.Context, id int32) (string, error)
}
