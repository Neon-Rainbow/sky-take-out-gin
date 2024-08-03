package setmeal

import (
	"context"
	SQLmodel "sky-take-out-gin/model/sql"
)

type SetmealDAOInterface interface {
	CreateSetmeal(ctx context.Context, setmeal *SQLmodel.Setmeal) error
	GetSetmealByID(ctx context.Context, id int64) (*SQLmodel.Setmeal, error)
	GetSetmealPage(ctx context.Context, CategoryID int64, page, pageSize int) ([]SQLmodel.Setmeal, error)
	SearchSetmeals(ctx context.Context, condition string, args ...interface{}) ([]SQLmodel.Setmeal, error)
	UpdateSetmeal(ctx context.Context, setmeal *SQLmodel.Setmeal) error
	UpdateSetmealStatus(ctx context.Context, id int64, status int) error
	DeleteSetmeals(ctx context.Context, ids []int64) error
	GetSetmeals(ctx context.Context, categoryID int64, name string, status int, offset int, limit int) ([]SQLmodel.Setmeal, int64, error)
	SetmealRawSQL(ctx context.Context, sql string, values ...interface{}) ([]SQLmodel.Setmeal, error)
}
