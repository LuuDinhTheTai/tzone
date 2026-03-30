package middleware

import (
	"net/http"

	"github.com/LuuDinhTheTai/tzone/internal/service"
	"github.com/LuuDinhTheTai/tzone/util/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RBACAuth(permissionService *service.PermissionService) gin.HandlerFunc {
	return func(c *gin.Context) {

		userIDStr, exists := c.Get("user_id")
		if !exists {
			response.Error(c, http.StatusUnauthorized, "unauthorized: missing user identity", nil)
			c.Abort()
			return
		}

		userID, err := uuid.Parse(userIDStr.(string))
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "invalid user identity", nil)
			c.Abort()
			return
		}

		// Information needed to check permission
		action := c.Request.Method // e.g., "GET", "POST", "PUT", "DELETE"
		resource := c.FullPath()   // e.g., "/admin/users", matches the exact Gin route pattern

		hasPerm, err := permissionService.HasPermission(userID, action, resource)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "failed to verify permissions", nil)
			c.Abort()
			return
		}

		if !hasPerm {
			response.Error(c, http.StatusForbidden, "forbidden: you do not have permission to access this resource", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
