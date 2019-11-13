package service

import (
	"github.com/senseoki/iris_ex/entity"
	"github.com/senseoki/iris_ex/vo"
)

// UserService is ...
type UserService interface {
	GetAll(*vo.User) []*entity.User
	Create(*vo.User)
}

// NewUserService is ...
func NewUserService() UserService {
	return &userService{}
}

type userService struct{}

func (s *userService) GetAll(userVO *vo.User) []*entity.User {
	users := []*entity.User{}
	userVO.RDBTX.Find(&users)
	return users
}

func (s *userService) Create(userVO *vo.User) {
	userVO.RDBTX.Create(userVO.User)
}
