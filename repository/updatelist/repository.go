package updatelist

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"sc-profile/models"
)

type IRepository interface {
	SelectUpdateList(ctx context.Context) ([]models.UpdateList, error)
	InsertUpdateList(ctx context.Context, updateList models.UpdateList) error
}

type Repository struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewRepository(logger *zap.Logger, db *sqlx.DB) *Repository {
	return &Repository{logger: logger, db: db}
}

//go:embed sql/select_update_list.sql
var selectUpdateListSql string

func (r *Repository) SelectUpdateList(ctx context.Context) ([]models.UpdateList, error) {
	var updateList []DbUpdateList
	if err := r.db.GetContext(ctx, &updateList, selectUpdateListSql); err != nil {
		return nil, fmt.Errorf("get query error: %w", err)
	}

	return DbUpdateListsToUpdateLists(updateList), nil
}

//go:embed sql/insert_update_list.sql
var insertUpdateListSql string

func (r *Repository) InsertUpdateList(ctx context.Context, updateList models.UpdateList) error {
	stmt, err := r.db.PrepareNamedContext(ctx, insertUpdateListSql)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, UpdateListToDb(updateList)); err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}
