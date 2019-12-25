package api

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/controller"
	"net/http"
)

func (s *Server) configureRoutes() {
	v1 := s.Router.Group("/ap1/v1")
	{
		// Handle index
		v1.GET("", handleIndex())

		// Auth routes
		auth := v1.Group("/auth")
		auth.POST("/register", controller.RegisterAccount)
		auth.POST("/login", controller.Login)
		auth.GET("/logout", controller.Logout)
		auth.POST("/password/change", controller.ChangePassword)
		auth.POST("/password/forgot", controller.ForgotPassword)
		auth.POST("/password/reset", controller.ResetPassword)

		// User routes
		users := v1.Group("/users")
		//users.POST("") // create user
		users.GET("/:id", controller.GetUser)
		users.PUT("/:id", controller.UpdateUser)

		// To-dos routes
		todo := v1.Group("/todo")
		todo.POST("", controller.CreateTodo)
		todo.GET("/:id", controller.GetTodo)
		todo.PUT("/:id", controller.UpdateTodo)
		todo.DELETE("/:id", controller.DeleteTodo)

		v1.GET("/user_todo/:id", controller.GetUserTodo)
	}
}

func handleIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
