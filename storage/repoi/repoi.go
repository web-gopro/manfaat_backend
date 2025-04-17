package repoi

import (
	"context"

	"github.com/jasurxaydarov/marifat_ac_backend/models"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, req models.UserReq) (*models.User, error)
	GetUser(ctx context.Context, req string) (*models.User, error)
	GetUsers(ctx context.Context, req models.GetList) (*models.UserList, error)
	IsExists(ctx context.Context, req models.IsExists) (*models.IsExistsResp, error)
	LogIn(ctx context.Context, req models.UserLogIn) (*models.Claim, error)
	ForReq(ctx context.Context, req models.For_req) (*models.For_req, error)
	GetForReq(ctx context.Context, req models.Id) (*models.For_req, error)
}

type TeacherRepoI interface {
	CreateTeacher(ctx context.Context, req models.TeacherReq) (*models.Teacher, error)
	GetTeacher(ctx context.Context, req models.Id) (*models.Teacher, error)
}
