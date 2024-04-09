// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package office_repository

import (
	"context"
)

type Querier interface {
	AttachePhoto(ctx context.Context, arg AttachePhotoParams) error
	CreateOffice(ctx context.Context, arg CreateOfficeParams) (Office, error)
	DeleteOffice(ctx context.Context, id int64) error
	GetImagePath(ctx context.Context, id int64) (string, error)
	// @sql postgresql
	GetOffice(ctx context.Context, id int64) (Office, error)
	ListOffices(ctx context.Context) ([]Office, error)
	UpdateOffice(ctx context.Context, arg UpdateOfficeParams) error
}

var _ Querier = (*Queries)(nil)
