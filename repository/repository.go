package repository

import (
	"github.com/jmoiron/sqlx"

	"exportimportcsv/model"
)

type Repository interface {
	CreateUser(guid, email, password, roleName string) (data model.User, err error)
	GetUser() (data []model.ExportUser, err error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(guid, email, password, roleName string) (data model.User, err error) {
	var db = r.db

	const statement = `
		INSERT INTO users 
				(guid, email, password, role_guid, created_at)
		VALUES 
				($1, $2, $3, (SELECT guid FROM roles WHERE name = $4), (now() at time zone 'UTC')::TIMESTAMP)
		RETURNING
				id, guid, email, role_guid, created_at, updated_at
	`

	row := db.QueryRow(statement, guid, email, password, roleName)

	if err = row.Scan(
		&data.ID,
		&data.GUID,
		&data.Email,
		&data.RoleGUID,
		&data.CreatedAt,
		&data.UpdatedAt,
	); err != nil {
		return
	}

	return
}

func (r *repository) GetUser() (data []model.ExportUser, err error) {
	var (
		db     = r.db
		number int
	)

	const statement = `
		SELECT email, roles.name, created_at
		FROM users
		LEFT JOIN roles ON users.role_guid = roles.guid
	`

	rows, err := db.Query(statement)
	if err != nil {
		return
	}

	for rows.Next() {
		var user model.ExportUser

		number++
		user.NO = number

		if err = rows.Scan(
			&user.Email,
			&user.RoleName,
			&user.CreatedAt,
		); err != nil {
			return
		}

		data = append(data, user)
	}

	return
}
