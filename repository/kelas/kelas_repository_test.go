package kelas

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestKelasInsert(t *testing.T) {
	KelasRepository := NewKelasRepository(belajar_db.GetConnection())

	ctx := context.Background()
	Kelas := entity.Kelas{
		Jumlah_murid: 3000,
		Nama_sekolah: "SDN CIJANTUNG 04",
	}

	result, err := KelasRepository.Insert(ctx, Kelas)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	KelasRepository := NewKelasRepository(belajar_db.GetConnection())

	comment, err := KelasRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	KelasRepository := NewKelasRepository(belajar_db.GetConnection())

	Kelas, err := KelasRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, Kelas := range Kelas {
		fmt.Println(Kelas)
	}
}
