package api

import (
	"github.com/gin-gonic/gin"
	"go-todo/api/controller"
	"go-todo/api/middleware"
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
		auth.POST("/password/change", middleware.AuthMiddleWare(), controller.ChangePassword)
		auth.POST("/password/forgot", controller.ForgotPassword)
		auth.POST("/password/reset", controller.ResetPassword)

		// User routes
		users := v1.Group("/users")
		users.GET("/:id", middleware.AuthMiddleWare(), controller.GetUser)
		users.PUT("/:id", middleware.AuthMiddleWare(), controller.UpdateUser)

		// To-dos routes
		todo := v1.Group("/todo")
		todo.POST("", middleware.AuthMiddleWare(), controller.CreateTodo)
		todo.GET("/:id", middleware.AuthMiddleWare(), controller.GetTodo)
		todo.PUT("/:id", middleware.AuthMiddleWare(), controller.UpdateTodo)
		todo.DELETE("/:id", middleware.AuthMiddleWare(), controller.DeleteTodo)

		// User to-do
		v1.GET("/user_todo/:id", middleware.AuthMiddleWare(), controller.GetUserTodo)
	}
}

func handleIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
