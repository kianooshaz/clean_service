package app

import "github.com/kianooshaz/clean_service/core/contract/interfaces"

func userRouting(controllers interfaces.IUserController) {
	e.POST("/users", controllers.Create)
	e.GET("/users/:id", controllers.Get)
	e.PUT("/users", controllers.Update)
	e.PATCH("/users", controllers.Update)
	e.DELETE("/users/:id", controllers.Delete)
	e.GET("/users", controllers.FindAll)
}
