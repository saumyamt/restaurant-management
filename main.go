package main

import ()
	"os"
	"github.com/gin-gonic/gin"
	"golang-restaurant-management/database"
	"golang-restaurant-management/routes"
	"golang-restaurant-management/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.foodCollection = database.OpenCollection(database.Client,"food")

func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port = "8000"
	}
	router := gin.New()
	router.User(gin.Logger())
	routes.UserRoutes(router)
	routes.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":"+port)

}