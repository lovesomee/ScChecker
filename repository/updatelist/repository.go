package updatelist

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

type IRepository interface {
	SelectUpdateList(ctx context.Context) ([]string, error)
	InsertUpdateList(ctx context.Context, updateList []string) error
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

func (r *Repository) SelectUpdateList(ctx context.Context) ([]string, error) {
	var updateList []string
	if err := r.db.SelectContext(ctx, &updateList, selectUpdateListSql); err != nil {
		return nil, fmt.Errorf("get query error: %w", err)
	}

	return updateList, nil
}

func (r *Repository) InsertUpdateList(ctx context.Context, updateList []string) error {
	values := make([]string, len(updateList))
	for i := range updateList {
		values[i] = fmt.Sprintf("($%d)", i+1)
	}
	query := fmt.Sprintf("INSERT INTO update_list (item_id) VALUES %s", strings.Join(values, ","))

	// Подготовка запроса
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение запроса
	_, err = stmt.Exec(toInterfaceSlice(updateList)...)
	return err
}

func toInterfaceSlice(updateList []string) []interface{} {
	result := make([]interface{}, len(updateList))
	for i, v := range updateList {
		result[i] = v
	}
	return result
}
