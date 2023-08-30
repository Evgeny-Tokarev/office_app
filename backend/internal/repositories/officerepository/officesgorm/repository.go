package officesgorm

import (
	"context"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/officerepository"
	"gorm.io/gorm"
)

type OfficeGorm struct {
	db *gorm.DB
}

func New(db *gorm.DB) *OfficeGorm {
	return &OfficeGorm{
		db: db,
	}
}

func (r *OfficeGorm) Create(ctx context.Context, e *officerepository.Office) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *OfficeGorm) Get(ctx context.Context, id int64) (*officerepository.Office, error) {
	e := new(officerepository.Office)
	err := r.db.WithContext(ctx).First(e, id).Error
	return e, err
}

func (r *OfficeGorm) List(ctx context.Context) ([]*officerepository.Office, error) {
	var list []*officerepository.Office
	err := r.db.WithContext(ctx).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, err
}

func (r *OfficeGorm) Update(ctx context.Context, e *officerepository.Office) error {
	return r.db.WithContext(ctx).Save(e).Error
}

func (r *OfficeGorm) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&officerepository.Office{ID: id}).Error
}
