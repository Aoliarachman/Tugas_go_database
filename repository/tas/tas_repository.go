package tas

import (
	"context"
	"go_database/entity"
)

type TasRepository interface {
	Insert(ctx context.Context, tas entity.Rumah) (entity.Tas, error)
	FindById(ctx context.Context, id int32) (entity.Tas, error)
	FindAll(ctx context.Context) ([]entity.Tas, error)
	Update(ctx context.Context, tas entity.Tas) (entity.Tas, error)
	Delete(ctx context.Context) (string, error)
}
