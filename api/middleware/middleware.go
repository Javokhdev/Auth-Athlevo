package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	_ "auth-athlevo/api/docs"
	"auth-athlevo/api/token"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// func Middleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		url := (ctx.Request.URL.Path)

// 		if strings.Contains(url, "swagger") || (url == "/auth/login") || (url == "/auth/register") {
// 			ctx.Next()
// 			return
// 		}
// 		ctx.Next()
// 	}
// }

// func JWTMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
// 			c.Abort()
// 			return
// 		}

// 		// if !strings.HasPrefix(authHeader, "Bearer ") {
// 		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
// 		// 	c.Abort()
// 		// 	return
// 		// }

// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// 		valid, err := t.ValidateToken(tokenString)
// 		if err != nil || !valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := t.ExtractClaim(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims", "details": err.Error()})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("claims", claims)
// 		c.Next()
// 	}
// }

// func GetUserId(r *http.Request) (string, error) {
// 	jwtToken := r.Header.Get("Authorization")

// 	if jwtToken == "" || strings.Contains(jwtToken, "Basic") {
// 		return "unauthorized", nil
// 	}

// 	// if !strings.HasPrefix(jwtToken, "Bearer ") {
// 	// 	return "unauthorized", errors.New("invalid authorization header format")
// 	// }

// 	// tokenString := strings.TrimPrefix(jwtToken, "Bearer ")

// 	claims, err := t.ExtractClaim(jwtToken)
// 	if err != nil {
// 		log.Println("Error while extracting claims: ", err)
// 		return "unauthorized", err
// 	}

// 	userID, ok := claims["user_id"].(string)
// 	if !ok {
// 		return "unauthorized", errors.New("user_id claim not found")
// 	}
// 	return userID, nil
// }

// func InvalidToken(c *gin.Context) {
// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 		"error": "Invalid token !!!",
// 	})
// }

// func RequirePermission(c *gin.Context) {
// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 		"error": "Permission denied",
// 	})
// }

// func RequireRefresh(c *gin.Context) {
// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 		"error": "Access token expired",
// 	})
// }

// Initialize Casbin Enforcer
func CasbinEnforcer() *casbin.Enforcer {
	e, err := casbin.NewEnforcer("path/to/casbin_model.conf", "path/to/casbin_policy.csv")
	if err != nil {
		log.Fatalf("Failed to initialize Casbin enforcer: %v", err)
	}
	return e
}

// JWT Middleware for token validation
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		valid, err := token.ValidateToken(tokenString)
		if err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		claims, err := token.ExtractClaim(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims", "details": err.Error()})
			c.Abort()
			return
		}

		// Extract and set userID and role in context
		userID, _ := claims["user_id"].(string)
		role, _ := claims["role"].(string)

		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}

// Casbin Middleware for RBAC
func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role") // Retrieve role from context

		allowed, err := enforcer.Enforce(userRole, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enforcing access control"})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Utility to get user ID from request
func GetUserId(r *http.Request) (string, error) {
	jwtToken := r.Header.Get("Authorization")

	if jwtToken == "" || strings.Contains(jwtToken, "Basic") {
		return "", errors.New("unauthorized")
	}

	if !strings.HasPrefix(jwtToken, "Bearer ") {
		return "", errors.New("invalid authorization header format")
	}

	tokenString := strings.TrimPrefix(jwtToken, "Bearer ")

	claims, err := token.ExtractClaim(tokenString)
	if err != nil {
		log.Println("Error while extracting claims:", err)
		return "", err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id claim not found")
	}
	return userID, nil
}

// Error handlers
func InvalidToken(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "Invalid token !!!",
	})
}

func RequirePermission(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "Permission denied",
	})
}

func RequireRefresh(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Access token expired",
	})
}
