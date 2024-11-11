package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pabloespinosa12/tiny-route/internal/database"
)

func GetDatabaseFromContext(c *gin.Context) *database.DatabaseService {
	val, _ := c.Get("database")
	if val == nil {
		panic("Database service not found in context")
	}

	db, ok := val.(*database.DatabaseService)
	if !ok {
		panic("Database service not found in context")
	}

	return db
}

func DatabaseProvider(db *database.DatabaseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
		c.Next()
	}
}
