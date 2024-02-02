// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package office_repository

import (
	"database/sql"
	"time"
)

type Office struct {
	ID        int64          `db:"id"`
	Name      string         `db:"name"`
	Address   string         `db:"address"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	ImgFile   sql.NullString `db:"img_file"`
}
