package officerepository

import (
	"context"
	"time"
)

type Office struct {
	ID        int64     `db:"id"        gorm:"id;primaryKey"`
	Name      string    `db:"name"      gorm:"name"`
	Address   string    `db:"address"   gorm:"address"`
	CreatedAt time.Time `db:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `db:"updated_at" gorm:"updated_at"`
}

type OfficeRepository interface {
	Create(ctx context.Context, e *Office) error
	Get(ctx context.Context, id int64) (*Office, error)
	List(ctx context.Context) ([]*Office, error)
	Update(ctx context.Context, e *Office) error
	Delete(ctx context.Context, id int64) error
}
