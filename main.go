package main

import "mnc-authentication/database"

func main() {
	// Initialize Database
	database.Connect("root@tcp(localhost:3306)/mnc_db?parseTime=true")
	database.Migrate()
}