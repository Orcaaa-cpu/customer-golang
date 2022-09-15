package main

import (
	"github.com/Orcaaa-cpu/customer-golang/controllers"
	"github.com/Orcaaa-cpu/customer-golang/db"
	"github.com/labstack/echo"
)

func main() {

	db.CreateCon()
	e := echo.New()

	e.GET("/customer", controllers.GetCustomer)
	e.POST("/customer", controllers.PostCustomer)
	e.PUT("/customer", controllers.UpdateCustomer)
	e.DELETE("/customer", controllers.DeleteCustomer)

	e.Logger.Fatal(e.Start(":2000"))
}
