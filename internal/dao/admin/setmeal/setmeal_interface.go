package setmeal

import (
	"context"
	model "sky-take-out-gin/model/sql"
)

type SetmealDAOInterface interface {
	CreateSetmeal(ctx context.Context, setmeal *model.Setmeal) error
	GetSetmealByID(ctx context.Context, id int64) (*model.Setmeal, error)
	GetSetmealPage(ctx context.Context, page, pageSize int) ([]model.Setmeal, error)
	SearchSetmeals(ctx context.Context, condition string, args ...interface{}) ([]model.Setmeal, error)
	UpdateSetmeal(ctx context.Context, setmeal *model.Setmeal) error
	UpdateSetmealStatus(ctx context.Context, id int64, status int) error
	DeleteSetmeals(ctx context.Context, ids []int64) error
	GetSetmeals(ctx context.Context, categoryID int64, name string, status int, offset int, limit int) ([]model.Setmeal, int64, error)
	SetmealRawSQL(ctx context.Context, sql string, values ...interface{}) ([]model.Setmeal, error)
}
