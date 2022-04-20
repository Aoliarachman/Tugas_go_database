package tas

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type TasRepositoryImpl struct {
	DB *sql.DB
}

func NewTasRepository(db *sql.DB) *TasRepositoryImpl {
	return &TasRepositoryImpl{db}
}
func (repository *TasRepositoryImpl) Insert(ctx context.Context, tas entity.Tas) (entity.Tas, error) {
	script := "INSERT INTO Tas(Harga, Merk) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, tas.Harga, tas.Merk)
	if err != nil {
		return tas, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return tas, err
	}
	tas.Id = int32(id)
	return tas, nil
}
func (repository *TasRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Tas, error) {
	script := "SELECt id, harga, Merk FROM tas WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	tas := entity.Tas{}
	if err != nil {
		return tas, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&tas.Id, &tas.Harga, &tas.Merk)
		return tas, nil
	} else {
		//tidak ada
		return tas, errors.New("Id" + strconv.Itoa(int(id)) + " Nor Found")
	}
}

func (repository *TasRepositoryImpl) FindAll(ctx context.Context) ([]entity.Tas, error) {
	script := "SELECt id, Harga, Merk FROM Tas "
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Bag []entity.Tas
	for rows.Next() {
		Tas := entity.Tas{}
		rows.Scan(&Tas.Id, &Tas.Harga, &Tas.Merk)
		Bag = append(Bag, Tas)
	}
	return Bag, nil
}

func (repository *TasRepositoryImpl) Update(ctx context.Context, tas entity.Tas) (entity.Tas, error) {
	script := "UPDATE Tas SET harga = ? WHERE merk = ?"
	result, err := repository.DB.ExecContext(ctx, script, tas.Harga, tas.Merk)
	if err != nil {
		return tas, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return tas, err
	}
	if rowCnt == 0 {
		return tas, err
	}
	return tas, nil
}

func (repository *TasRepositoryImpl) Delete(ctx context.Context, tas entity.Tas) (entity.Tas, error) {
	script := "DELETE FROM tas WHERE Merk = ?"
	result, err := repository.DB.ExecContext(ctx, script, tas.Merk)
	if err != nil {
		return tas, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return tas, err
	}
	if rowCnt == 0 {
		return tas, err
	}
	return tas, nil
}
