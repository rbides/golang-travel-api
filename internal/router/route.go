package router

import (
	travelHandler "golang-travel-api/internal/handler/http"
	travelRepo "golang-travel-api/internal/repository/postgresql"
	travelSvc "golang-travel-api/internal/service"

	"github.com/gin-gonic/gin"
)

func initTravel(r *gin.RouterGroup) {
	repo := travelRepo.New()
	svc := travelSvc.New(repo)
	handler := travelHandler.New(svc)
	travel := r.Group("/travel")
	travel.GET("", handler.Get)
	travel.POST("", handler.Create)
	travel.GET(":travelID", handler.GetById)
	travel.PUT(":travelID", handler.Update)
	travel.DELETE(":travelID", handler.Delete)
}

func Init() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	initTravel(api)
	return r
}
