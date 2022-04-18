package kelas

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type KelasRepositoryImpl struct {
	DB *sql.DB
}

func NewKelasRepository(db *sql.DB) KelasRepository {
	return &KelasRepositoryImpl{db}
}
func (repository *KelasRepositoryImpl) Insert(ctx context.Context, kelas entity.Kelas) (entity.Kelas, error) {
	script := "INSERT INTO Kelas(jumlah_murid, nama_sekolah) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, kelas.Jumlah_murid, kelas.Nama_sekolah)
	if err != nil {
		return kelas, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return kelas, err
	}
	kelas.Id = int32(id)
	return kelas, nil
}
func (repository *KelasRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Kelas, error) {
	script := "SELECt id, merk, harga FROM handphone WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	Kelas := entity.Kelas{}
	if err != nil {
		return Kelas, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&Kelas.Id, &Kelas.Jumlah_murid, &Kelas.Nama_sekolah)
		return Kelas, nil
	} else {
		//tidak ada
		return Kelas, errors.New("Id" + strconv.Itoa(int(id)) + " Nor Found")
	}
}

func (repository *KelasRepositoryImpl) FindAll(ctx context.Context) ([]entity.Kelas, error) {
	script := "SELECt id, jumlah_murid, nama_sekolah FROM kelas "
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Kilas []entity.Kelas
	for rows.Next() {
		Kelas := entity.Kelas{}
		rows.Scan(&Kelas.Id, &Kelas.Jumlah_murid, &Kelas.Nama_sekolah)
		Kilas = append(Kilas, Kelas)
	}
	return Kilas, nil
}

func (repository *KelasRepositoryImpl) Update(ctx context.Context, id int32, Kelas entity.Kelas) (entity.Kelas, error) {
	//TODO implement me
	script := "SELECT id, merk, harga FROM handphone WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return Kelas, err
	}
	if rows.Next() {
		// yes
		script := "UPDATE Kelas SET jumlah_murid = ?, nama_sekolah = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, Kelas.Jumlah_murid, Kelas.Nama_sekolah, id)
		if err != nil {
			return Kelas, err
		}
		Kelas.Id = id
		return Kelas, nil
	} else {
		// no
		return Kelas, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}

func (repository *KelasRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, merk, harga FROM handphone WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {

		script := "DELETE FROM handphone WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {

		return "Gagal", errors.New(("Id" + strconv.Itoa(int(id)) + "tidak ada"))
	}
}
