package main

import (
	"log"
	"resume/cv"
	"resume/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(sqlite.Open("resume.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	resumeRepository := cv.NewRepository(db)

	resumeService := cv.NewService(resumeRepository)

	resumeHandler := handler.NewResumeHandler(resumeService)

	router := gin.Default()
	router.Use(CORSMiddleware())

	api := router.Group("api/v1")

	api.POST("/create_resume", resumeHandler.CreateResume)
	api.GET("/list_resume", resumeHandler.ListResume)
	api.GET("/get_by_id/:id", resumeHandler.GetResume)
	api.PUT("/update_resume/:id", resumeHandler.UpdateResume)
	api.DELETE("/delete_resume/:id", resumeHandler.DeleteResume)

	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
