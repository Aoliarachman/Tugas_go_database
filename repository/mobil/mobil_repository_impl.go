package mobil

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type MobilRepositoryImpl struct {
	DB *sql.DB
}

func NewMobilRepository(db *sql.DB) *MobilRepositoryImpl {
	return &MobilRepositoryImpl{db}
}
func (repository *MobilRepositoryImpl) Insert(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error) {
	script := "INSERT INTO Mobil(harga, Merk) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, mobil.Harga, mobil.Merk)
	if err != nil {
		return mobil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return mobil, err
	}
	mobil.Id = int32(id)
	return mobil, nil
}
func (repository *MobilRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Mobil, error) {
	script := "SELECt id, merk, harga FROM mobil WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	mobil := entity.Mobil{}
	if err != nil {
		return mobil, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&mobil.Id, &mobil.Harga, &mobil.Merk)
		return mobil, nil
	} else {
		//tidak ada
		return mobil, errors.New("Id" + strconv.Itoa(int(id)) + " Nor Found")
	}
}

func (repository *MobilRepositoryImpl) FindAll(ctx context.Context) ([]entity.Mobil, error) {
	script := "SELECt id, Harga, Merk FROM Mobil "
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Mobile []entity.Mobil
	for rows.Next() {
		Mobil := entity.Mobil{}
		rows.Scan(&Mobil.Id, &Mobil.Harga, &Mobil.Merk)
		Mobile = append(Mobile, Mobil)
	}
	return Mobile, nil
}

func (repository *MobilRepositoryImpl) Update(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error) {
	script := "UPDATE Mobil SET Harga = ? WHERE  Merk = ?"
	result, err := repository.DB.ExecContext(ctx, script, mobil.Harga, mobil.Merk)
	if err != nil {
		return mobil, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return mobil, err
	}
	if rowCnt == 0 {
		return mobil, err
	}
	return mobil, nil
}

func (repository *MobilRepositoryImpl) Delete(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error) {
	script := "DELETE FROM mobil WHERE Merk = ?"
	result, err := repository.DB.ExecContext(ctx, script, mobil.Merk)
	if err != nil {
		return mobil, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return mobil, err
	}
	if rowCnt == 0 {
		return mobil, err
	}
	return mobil, nil
}
