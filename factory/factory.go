package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	udata "expenstracker/features/users/data"
	udeliv "expenstracker/features/users/delivery"
	ucase "expenstracker/features/users/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	validator := validator.New()

	userData := udata.New(db)
	userCase := ucase.New(userData, validator)
	userHandler := udeliv.New(userCase, userData)
	udeliv.RouteUser(e, userHandler)

}
