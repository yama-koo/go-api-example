package controllers

import (
	"fmt"
	"strconv"

	"github.com/yama-koo/go-api-example/domain"
	"github.com/yama-koo/go-api-example/interface/database"
	"github.com/yama-koo/go-api-example/usecase"
)

// UserController struct
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController func
func NewUserController(sqlHandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// Create func
func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, fmt.Errorf("%s", err.Error()))
		return
	}

	c.JSON(201, nil)
}

// Index func
func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, fmt.Errorf("%s", err))
		return
	}

	c.JSON(200, users)
}

// Show func
func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.FindByID(id)
	if err != nil {
		c.JSON(500, fmt.Errorf("%s", err))
		return
	}

	c.JSON(200, user)
}
