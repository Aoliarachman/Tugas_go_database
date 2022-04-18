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

func NewMobilRepository(db *sql.DB) MobilRepository {
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

func (repository *MobilRepositoryImpl) Update(ctx context.Context, id int32, Mobil entity.Mobil) (entity.Mobil, error) {
	//TODO implement me
	script := "SELECT id, Harga, Merk FROM Mobil WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()

	if err != nil {
		return Mobil, err
	}
	if rows.Next() {
		// yes
		script := "UPDATE Mobil SET Harga = ?, Merk = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, Mobil.Harga, Mobil.Merk, id)
		if err != nil {
			return Mobil, err
		}
		Mobil.Id = id
		return Mobil, nil
	} else {
		// no
		return Mobil, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}

func (repository *MobilRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, Harga, Merk FROM Mobil WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {

		script := "DELETE FROM Mobil WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {

		return "Gagal", errors.New(("Id" + strconv.Itoa(int(id)) + "tidak ada"))
	}
}
