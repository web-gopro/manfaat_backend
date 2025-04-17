package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/marifat_ac_backend/models"
	"github.com/jasurxaydarov/marifat_ac_backend/pgx/helpers"
	"github.com/jasurxaydarov/marifat_ac_backend/storage/repoi"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserREpo(db *pgx.Conn) repoi.UserRepoI {

	return &UserRepo{db: db}
}

// Create users
func (u *UserRepo) CreateUser(ctx context.Context, req models.UserReq) (*models.User, error) {
	var id uuid.UUID

	query := `
		INSERT INTO users(
			user_id,
			username,
			usersurname,
			user_number,
			email,
			password,
			user_role
		)VALUES(
			$1, $2, $3, $4, $5, $6, $7
		)
	
		`

	id = uuid.New()

	_, err := u.db.Exec(ctx,
		query,
		id,
		req.User_name,
		req.User_surname,
		req.User_number,
		req.User_email,
		req.Password,
		req.User_role,
	)

	if err != nil {
		fmt.Println("err on create user", err)
		return nil, err
	}

	resp, err := u.GetUser(ctx, id.String())

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

// Get user By Id
func (u *UserRepo) GetUser(ctx context.Context, req string) (*models.User, error) {

	var resp models.User
	query := `
		SELECT 
			user_id,
			username,
			usersurname,
			user_number,
			email,
			password,
			user_role
			
		FROM
			users
		WHERE
			user_id = $1

	`

	row := u.db.QueryRow(ctx, query, req)

	err := row.Scan(
		&resp.User_id,
		&resp.User_name,
		&resp.User_surname,
		&resp.User_number,
		&resp.User_email,
		&resp.Password,
		&resp.User_role,
	)

	if err != nil {

		fmt.Println("err on getting user", err)
		return nil, err
	}

	return &resp, nil
}

// Get users by list
func (u *UserRepo) GetUsers(ctx context.Context, req models.GetList) (*models.UserList, error) {

	offset := (req.Page - 1) * req.Limit

	var resp models.User

	var res models.UserList

	query := `
		SELECT 
			user_id,
			username,
			usersurname,
			user_number,
			email,
			password,
			user_role
		FROM 
    		users
		LIMIT $1 OFFSET $2;
	`

	row, err := u.db.Query(
		ctx,
		query,
		req.Limit,
		offset,
	)

	if err != nil {
		fmt.Println("err on get users list Query ", err.Error())
		return nil, err
	}

	for row.Next() {

		row.Scan(
			&resp.User_id,
			&resp.User_name,
			&resp.User_surname,
			&resp.User_number,
			&resp.User_email,
			&resp.Password,
			&resp.User_number,
		)

		if err != nil {
			fmt.Println("err 0n get user list row.Scan ", err.Error())
			return nil, err
		}

		res.Count++

		res.User = append(res.User, resp)

	}
	return &res, nil
}

// Check user is exists or not
func (u *UserRepo) IsExists(ctx context.Context, req models.IsExists) (*models.IsExistsResp, error) {

	var isExists bool

	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s = '%s')", req.TableName, req.ClomunName, req.ExpValue)

	err := u.db.QueryRow(ctx, query).Scan(&isExists)

	if err != nil {

		fmt.Println("err on check exists ", err)
		return nil, err

	}

	return &models.IsExistsResp{IsExists: isExists}, nil
}

// User login
func (u *UserRepo) LogIn(ctx context.Context, req models.UserLogIn) (*models.Claim, error) {

	var userId, gmail, hashPassword, userRole string

	query := `
		SELECT
			user_id,
			email,
			password,
			user_role
		FROM
			users
		WHERE	
			email =$1
	`

	err := u.db.QueryRow(
		ctx,
		query,
		req.User_email,
	).Scan(
		&userId,
		&gmail,
		&hashPassword,
		&userRole,
	)

	if err != nil {

		fmt.Println("err on login QueryRow", err.Error())
		return nil, err
	}

	if !helpers.CompareHashPassword(hashPassword, req.Password) {

		return nil, errors.New("user password is incorrect")

	}

	return &models.Claim{UserId: userId, UserRole: userRole}, nil
}

func (u *UserRepo) ForReq(ctx context.Context, req models.For_req) (*models.For_req, error) {

	query := `
		INSERT INTO for_req(
			data_id,
			name,
			number,
			description
		)VALUES(
			$1, $2, $3, $4
		)
	`

	_, err := u.db.Exec(
		ctx, query,
		req.Data_id,
		req.Name,
		req.Number,
		req.Description,
	)

	if err != nil {

		fmt.Println("err on create  ForReq", err.Error())
		return nil, err
	}

	resp, err := u.GetForReq(ctx, models.Id{Id: req.Data_id})

	if err != nil {

		fmt.Println("err on get  GetForReq", err.Error())
		return nil, err
	}

	return resp, nil
}

func (u *UserRepo) GetForReq(ctx context.Context, req models.Id) (*models.For_req, error) {

	var resp models.For_req

	query := `
		SELECT 
			data_id,
			name,
			number,
			description

		FROM
			for_req
		WHERE
			data_id =$1

	`

	err := u.db.QueryRow(
		ctx,
		query,
		req.Id,
	).Scan(
		&resp.Data_id,
		&resp.Name,
		&resp.Number,
		&resp.Description,
	)

	if err != nil {
		fmt.Println("err on db GetForReq", err.Error())

		return nil, err
	}

	return &resp, nil
}
