package main

import (
	"github.com/durpintm/user-management/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

}
