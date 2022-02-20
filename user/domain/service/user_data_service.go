package service

import (
	"errors"
	"go-micro/user/domain/model"
	"go-micro/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

//业务逻辑代码
type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) (err error)
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isOk bool, err error)
}

func NewUserDataService(UserRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: UserRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

//加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

//验证用户密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

//插入用户
func (u *UserDataService) AddUser(user *model.User) (userID int64, err error) {
	pwdByte, err := GeneratePassword(user.HashedPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashedPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

//删除用户
func (u *UserDataService) DeleteUser(userId int64) error {
	return u.UserRepository.DeleteUserByID(userId)
}

//更新用户
func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	//判断是否更新了密码
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashedPassword)
		if err != nil {
			return err
		}
		user.HashedPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

//根据用户名查找用户
func (u *UserDataService) FindUserByName(userName string) (*model.User, error) {
	return u.UserRepository.FindUserByName(userName)
}

//检查密码是否正确
func (u *UserDataService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashedPassword)
}
