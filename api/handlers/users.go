package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jasurxaydarov/marifat_ac_backend/mail"
	"github.com/jasurxaydarov/marifat_ac_backend/models"
	"github.com/jasurxaydarov/marifat_ac_backend/pgx/helpers"
	"github.com/jasurxaydarov/marifat_ac_backend/token"
)



// UserCreate godoc
// @Summary Create a new user
// @Router /user [post]
// @Description Create a new user with the provided details
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserReq true "User object to create"
// @Success 201 {object} models.User "Created"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
func (h *Handlers) UserCreate(ctx *gin.Context) {

	var req models.UserReq

	ctx.BindJSON(&req)

	resp, err := h.storage.UserRepo().CreateUser(context.Background(), req)

	if err != nil {

		fmt.Println(err)
		ctx.JSON(500, err)
	}

	/*data := fmt.Sprintf(
	`Username:%s   :
	User Phone:%s  :
	User Email:%s  : `,
	resp.User_name,
	resp.User_number,
	resp.User_email)*/

	///mail.SendAdminMail(data)

	//fmt.Println(data)

	ctx.JSON(200, resp)

}

func (h *Handlers) CheckUser(ctx *gin.Context) {

	var reqBody models.UserIsExists

	err := ctx.BindJSON(&reqBody)

	if err != nil {

		fmt.Println("err on BindJSON", err)
		ctx.JSON(400, err)
		return
	}

	isExists, err := h.storage.UserRepo().IsExists(context.Background(), models.IsExists{
		TableName:  "users",
		ClomunName: "email",
		ExpValue:   reqBody.UserEmail,
	})

	if err != nil {

		fmt.Println("err on storage from users is exists ", err)
		ctx.JSON(500, err)

		return
	}

	if isExists.IsExists {
		ctx.JSON(201, models.IsExistsResp{
			IsExists: isExists.IsExists,
			Status:   "sign-in",
		})

		return
	}

	otp := models.UserOtpData{
		Otp:   mail.GenerateOtp(6),
		Email: reqBody.UserEmail,
	}

	otpDataB, err := json.Marshal(&otp)

	if err != nil {

		fmt.Println("err on json.Marshal", err)
		ctx.JSON(500, err)
		return
	}

	err = h.cache.Set(ctx, reqBody.UserEmail, string(otpDataB), 120)

	if err != nil {

		fmt.Println("err on redis cache set ", err)
		ctx.JSON(500, err)
		return
	}

	err = mail.SendMail([]string{reqBody.UserEmail}, otp.Otp)

	if err != nil {

		fmt.Println("err on sent otp to user email", err)
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(201, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "registr",
		Message: "we sent otp your email",
	})

}

func (h *Handlers) SignUp(ctx *gin.Context) {

	var otpData models.UserOtpData

	var reqBody models.UserReq

	err := ctx.BindJSON(&reqBody)
	if err != nil {

		fmt.Println("err on BindJSON", err)
		ctx.JSON(401, err.Error())
		return
	}

	otpSData, err := h.cache.GetDell(ctx, reqBody.User_email)

	fmt.Println(otpSData)

	if err != nil {
		fmt.Println("h.cache.GetDell", err)
		ctx.JSON(500, err.Error())
		return
	}

	if otpSData == "" {

		ctx.JSON(201, "otp code is expired")
		return
	}

	err = json.Unmarshal([]byte(otpSData), &otpData)

	if otpData.Otp != reqBody.Otp {

		ctx.JSON(405, "incorrect otp code")
		return
	}

	reqBody.Password, err = helpers.HashPassword(reqBody.Password)

	claim, err := h.storage.UserRepo().CreateUser(context.Background(), reqBody)

	accessToken, err := token.GenerateJWT(*&models.Claim{UserId: claim.User_id, UserRole: claim.User_role})

	if err != nil {

		ctx.JSON(500, err.Error())

		return
	}

	ctx.JSON(201, accessToken)
}

func (h *Handlers) LogIn(ctx *gin.Context) {

	var reqBody models.UserLogIn

	err := ctx.BindJSON(&reqBody)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}
	claim, err := h.storage.UserRepo().LogIn(context.Background(), reqBody)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	accessToken, err := token.GenerateJWT(*claim)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, accessToken)

}

func (h *Handlers) GetUserById(ctx *gin.Context) {

	var req models.GetById

	req.Id = ctx.Param("id")

	resp, err := h.storage.UserRepo().GetUser(context.Background(), req.Id)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)

}

func (h *Handlers) GetUsers(ctx *gin.Context) {
	var req models.GetList

	err := ctx.BindJSON(&req)

	if err != nil {

		fmt.Println("err on BindJSON", err.Error())
		ctx.JSON(401, err.Error())
		return
	}

	resp, err := h.storage.UserRepo().GetUsers(context.Background(), req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *Handlers) CreateForReq(ctx *gin.Context) {

	var req models.For_req

	err := ctx.BindJSON(&req)

	if err != nil {

		fmt.Println("err on CreateForReq in BindJSON", err)
		ctx.JSON(401, err.Error())
		return
	}

	req.Data_id = uuid.New().String()

	resp, err := h.storage.UserRepo().ForReq(context.Background(), req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	emaildata := fmt.Sprintf(
		`Assalomu allaykum admin bizda yangi user request bor:

		User Name : " %s ",
		User Number : " %s ",
		User Commet : " %s ",
		User ID : " %s ",
	
		`,
		resp.Name,
		resp.Number,
		resp.Description,
		resp.Data_id,
	)

	err = mail.SendAdminMail(emaildata)

	if err != nil {

		fmt.Println("err on err=mail.SendAdminMail(emaildata) in For req", err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)
}
