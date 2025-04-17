package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/marifat_ac_backend/models"
	"github.com/jasurxaydarov/marifat_ac_backend/storage/repoi"
)

type TeacherRepo struct {
	db *pgx.Conn
}

func NewTeacherRepo(db *pgx.Conn) repoi.TeacherRepoI {

	return &TeacherRepo{db: db}
}

func (p *TeacherRepo) CreateTeacher(ctx context.Context, req models.TeacherReq) (*models.Teacher, error) {

	return nil, nil
}
func (p *TeacherRepo) GetTeacher(ctx context.Context, req models.Id) (*models.Teacher, error) {

	return nil, nil
}
