package api

import (
	"travel/api/handler"

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
	// travel.Use(middlewere.Middleware())

	users := travel.Group("/users")
	users.GET("/:id/GetProfile", h.GetProfile)
	users.GET("/:id/validate", h.ValidateUser)
	users.POST("/:id/editprofile", h.EditProfile)
	users.GET("/", h.GetUsers)
	users.DELETE("/:id/delete", h.DeleteUser)
	users.PUT("/update", h.UpdatePassword)
	users.POST("/:id/follow", h.Follow)
	users.GET("/:id/followers", h.GetFollowers)

	stories := travel.Group("/stories")
	users.POST("/:id/stories/create", h.CreateStory)
	stories.PUT("/:id/edit", h.EditStory)
	stories.GET("/", h.GetStories)
	stories.GET("/:id/fullinfo", h.GetStoryFullInfo)
	stories.DELETE("/:id/delete", h.DeleteStory)

	interactions := travel.Group("/interaction")
	interactions.POST("/comment/create", h.CreateComment)
	stories.GET("/:id/comments", h.GetComments)
	stories.POST("/:id/like", h.LikeStory)

	itineraries := travel.Group("/itineraries")
	users.POST("/:id/itineraries/create", h.CreateItineraries)
	itineraries.PUT("/:id/edit", h.EditItineraries)
	itineraries.GET("/", h.GetAllItineraries)
	itineraries.GET("/:id", h.GetItineraryFullInfo)
	itineraries.POST("/:id/comment", h.WriteCommentToItinerary)

	destination := travel.Group("/destination")
	destination.POST("/create", h.CreateDestination)
	destination.GET("/top", h.GetDestinations)
	
	return r
}
