package tas

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestTasInsert(t *testing.T) {
	TasRepository := NewTasRepository(belajar_db.GetConnection())

	ctx := context.Background()
	Tas := entity.Tas{
		Harga: 70000,
		Merk:  "Boogie",
	}

	result, err := TasRepository.Insert(ctx, Tas)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	TasRepository := NewTasRepository(belajar_db.GetConnection())

	tas, err := TasRepository.FindById(context.Background(), 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(tas)
}

func TestFindAll(t *testing.T) {
	TasRepository := NewTasRepository(belajar_db.GetConnection())

	Tas, err := TasRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, Tas := range Tas {
		fmt.Println(Tas)
	}
}

func TestUpdate(t *testing.T) {
	tasRepository := NewTasRepository(belajar_db.GetConnection())

	ctx := context.Background()
	tas := entity.Tas{
		Harga: 500000,
		Merk:  "Lois",
	}
	result, err := tasRepository.Update(ctx, tas)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	rumahRepository := NewTasRepository(belajar_db.GetConnection())

	ctx := context.Background()
	tas := entity.Tas{
		Merk: "nike",
	}
	result, err := rumahRepository.Delete(ctx, tas)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
