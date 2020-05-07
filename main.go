package main

import (
	_ "github.com/Los-Crackitos/Excelante/docs"
	"github.com/Los-Crackitos/Excelante/router"
)

// @title Excelante API
// @version 0.1.0
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /api/v1
func main() {
	router.CreateRouter()
}
