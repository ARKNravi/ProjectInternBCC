package main

import (
	"ProjectBuahIn/buah"
	"ProjectBuahIn/controllers"
	"ProjectBuahIn/handler"
	"ProjectBuahIn/initializer"
	"ProjectBuahIn/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDb()
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/projectbcc3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Error")
	}

	buahRepository := buah.NewRepository(db)
	buahService := buah.NewService(buahRepository)
	buahHandler := handler.NewBuahHandler(buahService)
	orderHandler := handler.NewOrderHandler()

	r := gin.Default()
	v1 := r.Group("/v1")
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	v1.GET("/buahs", buahHandler.GetBuahs)
	v1.GET("/buahs/:id", buahHandler.GetBuah)
	v1.GET("/buahnama/:nama", buahHandler.GetNamaBuah)
	v1.POST("/buahs", buahHandler.CreateBuah)
	v1.PUT("/buahs/:id", buahHandler.UpdateBuah)
	v1.GET("/user/:id", buahHandler.GetUser)
	v1.DELETE("/buahs/:id", buahHandler.DeleteBuah)
	v1.POST("/product/:nama/quantity/:quantity", orderHandler.OrderProduct)

	//v2 := r.Group("/v2")

	r.Run(":8090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
