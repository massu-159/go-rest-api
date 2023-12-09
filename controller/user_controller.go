package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
	CSRFToken(c echo.Context) error
}

// 構造体を作成
type userController struct {
	uu usecase.IUserUseCase
}

// コンストラクタを作成
// 依存関係の注入を行う(di)
// 構造体の実体を作成する
func NewUserController(uu usecase.IUserUseCase) IUserController {
	return &userController{uu}
}

// ユーザーを新規登録する
func (uc *userController) SignUp(c echo.Context) error {
	// リクエストボディを取得する
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// ユーザーを新規登録する
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)
}

// ユーザーをログインさせる
func (uc *userController) Login(c echo.Context) error {
	// リクエストボディを取得する
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// ユーザーをログインさせる
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// cookieにトークンをセットする
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour) // 有効期限は24時間
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = false // httpsで通信する場合はtrueに変える
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// ユーザーをログアウトさせる
func (uc *userController) LogOut(c echo.Context) error {
	// cookieを削除する
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = false // httpsで通信する場合はtrueに変える
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// CSRFトークンを取得する
func (uc *userController) CSRFToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrfToken": token,
	})
}
