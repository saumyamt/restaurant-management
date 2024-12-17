package routes

import(
	controller "golang-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)
func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices",controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id"contcontroller.GetInvoice())
	incomingRoutes.POST("/invoices",concontroller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id",contrcontroller.UpdateInvoice())
	
}