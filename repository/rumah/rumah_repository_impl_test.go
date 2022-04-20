package rumah

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestRumahInsert(t *testing.T) {
	rumahRepository := NewRumahRepository(belajar_db.GetConnection())

	ctx := context.Background()
	Rumah := entity.Rumah{
		Harga: 3.000000,
		Luas:  "200 hektar",
	}

	result, err := rumahRepository.Insert(ctx, Rumah)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	RumahRepository := NewRumahRepository(belajar_db.GetConnection())

	rumah, err := RumahRepository.FindById(context.Background(), 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(rumah)
}

func TestFindAll(t *testing.T) {
	RumahRepository := NewRumahRepository(belajar_db.GetConnection())

	Rumah, err := RumahRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, Rumah := range Rumah {
		fmt.Println(Rumah)
	}
}

func TestUpdate(t *testing.T) {
	rumahRepository := NewRumahRepository(belajar_db.GetConnection())

	ctx := context.Background()
	rumah := entity.Rumah{
		Harga: 4.000000,
		Luas:  "200 meter",
	}
	result, err := rumahRepository.Update(ctx, rumah)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	rumahRepository := NewRumahRepository(belajar_db.GetConnection())

	ctx := context.Background()
	rumah := entity.Rumah{
		Luas: "200 meter",
	}
	result, err := rumahRepository.Delete(ctx, rumah)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
