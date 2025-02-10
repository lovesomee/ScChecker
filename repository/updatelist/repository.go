package updatelist

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	"sc-profile/models"
)

type IRepository interface {
	SelectUpdateList() ([]models.UpdateList, error)
	InsertUpdateList(updateList models.UpdateList) error
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

//go:embed sql/select_update_list.sql
var selectUpdateListSql string

func (r *Repository) SelectUpdateList() ([]models.UpdateList, error) {
	var updateList []DbUpdateList
	if err := r.db.Get(&updateList, selectUpdateListSql); err != nil {
		return nil, err
	}

	return DbUpdateListsToUpdateLists(updateList), nil
}

//go:embed sql/insert_update_list.sql
var insertUpdateListSql string

func (r *Repository) InsertUpdateList(updateList models.UpdateList) error {
	stmt, err := r.db.PrepareNamed(insertUpdateListSql)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(UpdateListToDb(updateList)); err != nil {
		return err
	}

	return nil
}
