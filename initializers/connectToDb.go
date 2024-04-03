package initializers

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	fmt.Println("Connect to DB")

}
