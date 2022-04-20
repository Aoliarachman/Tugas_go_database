package rumah

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type RumahRepositoryImpl struct {
	DB *sql.DB
}

func NewRumahRepository(db *sql.DB) *RumahRepositoryImpl {
	return &RumahRepositoryImpl{db}
}
func (repository *RumahRepositoryImpl) Insert(ctx context.Context, rumah entity.Rumah) (entity.Rumah, error) {
	script := "INSERT INTO Rumah(harga, luas) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, rumah.Harga, rumah.Luas)
	if err != nil {
		return rumah, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return rumah, err
	}
	rumah.Id = int32(id)
	return rumah, nil
}
func (repository *RumahRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Rumah, error) {
	script := "SELECt id, harga, luas FROM rumah WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	rumah := entity.Rumah{}
	if err != nil {
		return rumah, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&rumah.Id, &rumah.Harga, &rumah.Luas)
		return rumah, nil
	} else {
		//tidak ada
		return rumah, errors.New("Id" + strconv.Itoa(int(id)) + " Nor Found")
	}
}

func (repository *RumahRepositoryImpl) FindAll(ctx context.Context) ([]entity.Rumah, error) {
	script := "SELECt id, Harga, luas FROM Rumah "
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var home []entity.Rumah
	for rows.Next() {
		Rumah := entity.Rumah{}
		rows.Scan(&Rumah.Id, &Rumah.Harga, &Rumah.Luas)
		home = append(home, Rumah)
	}
	return home, nil
}

func (repository *RumahRepositoryImpl) Update(ctx context.Context, rumah entity.Rumah) (entity.Rumah, error) {
	script := "UPDATE Rumah SET Harga = ? WHERE  luas = ?"
	result, err := repository.DB.ExecContext(ctx, script, rumah.Harga, rumah.Luas)
	if err != nil {
		return rumah, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return rumah, err
	}
	if rowCnt == 0 {
		return rumah, err
	}
	return rumah, nil
}

func (repository *RumahRepositoryImpl) Delete(ctx context.Context, rumah entity.Rumah) (entity.Rumah, error) {
	script := "DELETE FROM rumah WHERE Merk = ?"
	result, err := repository.DB.ExecContext(ctx, script, rumah.Luas)
	if err != nil {
		return rumah, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return rumah, err
	}
	if rowCnt == 0 {
		return rumah, err
	}
	return rumah, nil
}
