package middlewares

import (
	"UserManagementAPI/utils"
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// Authorize determines if current user has been authorized to take an action on an object.
func VerifyAccess(obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	utils.WriteLog("AppLog.txt")
	return func(context *gin.Context) {
		// Get current user/subject
		sub, existed := context.Get("user")
		// uid := user.(*model.User).UID

		if !existed {
			log.Println(sub)
			context.AbortWithStatusJSON(401, gin.H{"message": "User hasn't logged in yet."})
			return
		}

		// Load policy from Database
		err := enforcer.LoadPolicy()
		if err != nil {
			log.Println("Casbin failed: Load policy.")
			log.Println(err.Error())
			context.AbortWithStatusJSON(500, gin.H{"message": "Failed to load policy from DB"})
			return
		}

		// Casbin enforces policy
		ok, err := enforcer.Enforce(fmt.Sprint(sub), obj, act)

		if err != nil {
			log.Println("Casbin failed: Enforce policy.")
			log.Println(err.Error())
			context.AbortWithStatusJSON(500, gin.H{"message": "Error occurred when authorizing user."})
			return
		}

		if !ok {
			log.Println("Access denied: Unauthorized attempt.")
			context.AbortWithStatusJSON(403, gin.H{"message": "You are not authorized"})
			return
		}
		context.Next()
	}
}
