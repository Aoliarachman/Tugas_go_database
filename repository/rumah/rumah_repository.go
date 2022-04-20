package rumah

import (
	"context"
	"go_database/entity"
)

type RumahRepository interface {
	Insert(ctx context.Context, rumah entity.Rumah) (entity.Rumah, error)
	FindById(ctx context.Context, id int32) (entity.Rumah, error)
	FindAll(ctx context.Context) ([]entity.Rumah, error)
	Update(ctx context.Context, rumah entity.Rumah) (entity.Rumah, error)
	Delete(ctx context.Context) (string, error)
}
