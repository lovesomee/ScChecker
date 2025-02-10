package updatelist

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	SelectUpdateList() ([]DbUpdateList, error)
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

//go:embed sql/select_update_list.sql
var selectUpdateListSql string

func (r *Repository) SelectUpdateList() ([]DbUpdateList, error) {
	var updateList []DbUpdateList
	if err := r.db.Get(&updateList, selectUpdateListSql); err != nil {
		return nil, err
	}

	return updateList, nil
}
