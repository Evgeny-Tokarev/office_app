// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package employee_repository

import (
	"database/sql"
	"time"
)

type Employee struct {
	ID        int64          `db:"id"`
	Name      string         `db:"name"`
	Age       int32          `db:"age"`
	OfficeID  int64          `db:"office_id"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	ImgFile   sql.NullString `db:"img_file"`
}
