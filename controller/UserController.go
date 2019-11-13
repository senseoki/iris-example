package controller

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/senseoki/iris_ex/entity"
	"github.com/senseoki/iris_ex/service"
	"github.com/senseoki/iris_ex/vo"
)

// UserController is ...
type UserController struct {
	Ctx iris.Context

	UserService service.UserService
}

// BeforeActivation is ...
func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Info("Inside /user")
		ctx.Next()
	}

	b.Handle("GET", "/user", "GetUser", anyMiddlewareHere)
	b.Handle("POST", "/user", "CreateUser", anyMiddlewareHere)
}

// GetUser is ...
func (c *UserController) GetUser() {
	userVO := &vo.User{
		RDBTX: c.Ctx.Values().Get("RDBTX").(*gorm.DB),
	}
	users := c.UserService.GetAll(userVO)
	// c.Ctx.JSON(map[string]interface{}{
	// 	"notice": "test",
	// })

	c.Ctx.JSON(users)
}

// CreateUser is ...
func (c *UserController) CreateUser() {
	timeNow := time.Now().UTC()

	userVO := &vo.User{
		RDBTX: c.Ctx.Values().Get("RDBTX").(*gorm.DB),
		User: &entity.User{
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
	}

	if err := c.Ctx.ReadJSON(userVO.User); err != nil {
		log.Println("Fail Verify Parameter ReadJSON")
		return
	}

	err := c.UserService.Create(userVO)
	if err != nil {
		c.Ctx.StatusCode(iris.StatusInternalServerError)
	}
}
