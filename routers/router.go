package routers

import (
    "account-manager/controllers"
    "account-manager/middleware"
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

// SetupRouter sets up the Gin router with all routes
func SetupRouter() *gin.Engine {
    r := gin.Default()

    // User routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    // Apply middleware to protect the following routes
    authorized := r.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        // Transaction routes
        authorized.POST("/send", controllers.SendTransaction)
        authorized.POST("/withdraw", controllers.WithdrawTransaction)

        // Account routes
        authorized.POST("/accounts", controllers.CreateAccount)
        authorized.GET("/accounts", controllers.GetAccounts)
        authorized.GET("/accounts/:id/transactions", controllers.GetTransactionsByAccountID)
        authorized.GET("/transactions", controllers.GetTransactions)


        // Auto payment routes
        authorized.POST("/autopayment", controllers.CreateAutoPayment)
    }

    // Swagger documentation
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}