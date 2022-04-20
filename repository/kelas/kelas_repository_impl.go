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

func NewKelasRepository(db *sql.DB) *KelasRepositoryImpl {
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
	script := "SELECt id, Jumlah_murid, Nama_sekolah FROM Kelas WHERE id = ? LIMIT 1"
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

func (repository *KelasRepositoryImpl) Update(ctx context.Context, kelas entity.Kelas) (entity.Kelas, error) {
	script := "UPDATE kelas SET kelas = ? WHERE  Nama_sekolah = ?"
	result, err := repository.DB.ExecContext(ctx, script, kelas.Jumlah_murid, kelas.Nama_sekolah)
	if err != nil {
		return kelas, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return kelas, err
	}
	if rowCnt == 0 {
		return kelas, err
	}
	return kelas, nil
}

func (repository *KelasRepositoryImpl) Delete(ctx context.Context, kelas entity.Kelas) (entity.Kelas, error) {
	script := "DELETE FROM kelas WHERE Jumlah_murid = ?"
	result, err := repository.DB.ExecContext(ctx, script, kelas.Jumlah_murid)
	if err != nil {
		return kelas, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return kelas, err
	}
	if rowCnt == 0 {
		return kelas, err
	}
	return kelas, nil
}
