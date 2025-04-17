package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/marifat_ac_backend/storage/postgres"
	"github.com/jasurxaydarov/marifat_ac_backend/storage/repoi"
)

type StorageI interface {
	UserRepo() repoi.UserRepoI
	TeacherRepo() repoi.TeacherRepoI
}

type storage struct {
	userRepo    repoi.UserRepoI
	teacherRepo repoi.TeacherRepoI
}

func NewStorage(db *pgx.Conn) StorageI {

	return &storage{
		userRepo: postgres.NewUserREpo(db),
	}
}

func (s *storage) UserRepo() repoi.UserRepoI {

	return s.userRepo
}

func (s *storage) TeacherRepo() repoi.TeacherRepoI {
	return s.teacherRepo
}
