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

const (
	ContextUserID = "userID"
	ContextRole   = "role"
)

func CasbinEnforcer() *casbin.Enforcer {
	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
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

		if !strings.HasPrefix(authHeader, "") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "")

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

		userID, ok := claims["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id claim"})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role claim"})
			c.Abort()
			return
		}

		c.Set(ContextUserID, userID)
		c.Set(ContextRole, role)

		c.Next()
	}
}

// Casbin Middleware for RBAC
func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString(ContextRole)

		allowed, err := enforcer.Enforce(userRole, c.FullPath(), c.Request.Method)
		if err != nil {
			log.Println("Casbin enforcement error:", err)
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

	if !strings.HasPrefix(jwtToken, "") {
		return "", errors.New("invalid authorization header format")
	}

	tokenString := strings.TrimPrefix(jwtToken, "")

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

func GetRole(r *http.Request) (string, error) {
	jwtToken := r.Header.Get("Authorization")

	if jwtToken == "" || strings.Contains(jwtToken, "Basic") {
		return "unauthorized", nil
	}

	claims, err := token.ExtractClaim(jwtToken)
	if err != nil {
		log.Println("Error while extracting claims: ", err)
		return "unauthorized", err
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "unauthorized", errors.New("role claim not found")
	}
	return role, nil
}

// Error handlers
func InvalidToken(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "Invalid token",
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
