package api

import (
	"travel/api/handler"
	"travel/api/middlewere"

	_ "travel/api/docs"

	"github.com/gin-gonic/gin"
	swagerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

// @title TravelTales App
// @version 1.0
// @description Sayohatchilar o'rtasida tajriba almashish, sayohat rejalarini tuzish va boshqalar bilan bo'lishish imkoniyatini yaratish orqali global sayohat hamjamiyatini shakllantirish.

// @contact.name Saidakbar
// @contact.url http://www.support_me_with_smile
// @contact.email "pardaboyevsaidakbar103@gmail.com"
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:4444
// @BasePath /travel
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("swagger/*any", swagger.WrapHandler(swagerFiles.Handler))

	h := handler.NewHandler()
	travel := r.Group("/travel")
	travel.Use(middlewere.Middleware())

	users := travel.Group("/users")
	users.GET("/:id/GetProfile", h.GetProfile)

	return r
}
