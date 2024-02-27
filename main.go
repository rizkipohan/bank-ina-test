package main

import (
	"bank-ina-test/delivery"
	"bank-ina-test/domain"
	"bank-ina-test/repository"
	"bank-ina-test/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DbUser      = "root"
	DbPassword  = ""
	DbName      = "db_bank_ina"
	DbHost      = "localhost"
	DbPort      = "3306"
	DbQueryAttr = "parseTime=true"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", DbUser, DbPassword, DbHost, DbPort, DbName, DbQueryAttr)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&domain.User{},
		&domain.Task{},
	)

	if err != nil {
		panic("failed to migrate database")
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// Initialize use cases
	userUsecase := usecase.UserUsecase{UserRepository: *userRepo}
	taskUsecase := usecase.TaskUsecase{TaskRepository: *taskRepo}

	// Initialize handlers
	userHandler := delivery.UserHandler{Usecase: userUsecase}
	taskHandler := delivery.TaskHandler{Usecase: taskUsecase}

	// User endpoints
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.POST("/", userHandler.RegisterUser)
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Task endpoints
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.Use(Auth())
		taskRoutes.POST("/", taskHandler.CreateTask)
		taskRoutes.GET("/", taskHandler.GetAllTasks)
		taskRoutes.GET("/:id", taskHandler.GetTaskByID)
		taskRoutes.PUT("/:id", taskHandler.UpdateTask)
		taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
	}

	r.Run(":8080")
}
