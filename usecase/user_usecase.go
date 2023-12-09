package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

// 構造体を作成
type userUseCase struct {
	ur repository.IUserRepository
}

// コンストラクタを作成
// 依存関係の注入を行う(di)
// 構造体の実体を作成する
func NewUserUseCase(ur repository.IUserRepository) IUserUseCase {
	return &userUseCase{ur}
}

// ユーザーを新規登録する
func (uu *userUseCase) SignUp(user model.User) (model.UserResponse, error) {
	// パスワードをハッシュ化する
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	// ユーザーを作成する
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	// レスポンス用のユーザーに値を詰め替える
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

// ユーザーをログインさせる
func (uu *userUseCase) Login(user model.User) (string, error) {
	// ユーザーを取得する
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// パスワードを比較する
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// JWTを発行する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(), // 有効期限は12時間
	})
	// 環境変数からSECRETを取得する
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
