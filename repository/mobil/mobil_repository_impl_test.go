package mobil

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestMobilInsert(t *testing.T) {
	MobilRepository := NewMobilRepository(belajar_db.GetConnection())

	ctx := context.Background()
	Mobil := entity.Mobil{
		Harga: 4000,
		Merk:  "Honda",
	}

	result, err := MobilRepository.Insert(ctx, Mobil)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	MobilRepository := NewMobilRepository(belajar_db.GetConnection())

	mobil, err := MobilRepository.FindById(context.Background(), 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(mobil)
}

func TestFindAll(t *testing.T) {
	MobilRepository := NewMobilRepository(belajar_db.GetConnection())

	Mobil, err := MobilRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, Mobil := range Mobil {
		fmt.Println(Mobil)
	}
}

func TestUpdate(t *testing.T) {
	mobilRepository := NewMobilRepository(belajar_db.GetConnection())

	ctx := context.Background()
	mobil := entity.Mobil{
		Harga: 7000,
		Merk:  "Honda",
	}
	result, err := mobilRepository.Update(ctx, mobil)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	mobilRepository := NewMobilRepository(belajar_db.GetConnection())

	ctx := context.Background()
	mobil := entity.Mobil{
		Merk: "Honda",
	}
	result, err := mobilRepository.Delete(ctx, mobil)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
