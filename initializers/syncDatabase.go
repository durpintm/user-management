package initializers

import (
	"github.com/durpintm/user-management/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.InvitationCode{})
}
