package inits

import (
	"UserManagementAPI/utils"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func SetPolicies(db *gorm.DB) *casbin.Enforcer {
	utils.WriteLog("AppLog.txt")

	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Println("Init failed: Casbin adapter.")
		log.Println(err.Error())
	} else {
		log.Println("Init sucesssful: Casbin adapter.")
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		log.Println("Init failed: Casbin enforcer.")
		log.Println(err.Error())
	} else {
		log.Println("Init sucesssful: Casbin enforcer.")
	}

	//add policy
	if hasPolicy := enforcer.HasPolicy("admin", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("admin", "report", "read")
	}
	if hasPolicy := enforcer.HasPolicy("admin", "report", "write"); !hasPolicy {
		enforcer.AddPolicy("admin", "report", "write")
	}
	if hasPolicy := enforcer.HasPolicy("public", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("public", "report", "read")
	}
	return enforcer
}
