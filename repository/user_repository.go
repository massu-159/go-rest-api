package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// 構造体を作成
type UserRepository struct {
	db *gorm.DB
}

// コンストラクタを作成
// 依存関係の注入を行う(di)
// 構造体の実体を作成する
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

// ユーザーを取得する
func (ur *UserRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザーを作成する
func (ur *UserRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
