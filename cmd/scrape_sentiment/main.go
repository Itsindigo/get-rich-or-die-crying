package main

import (
	"fmt"

	"github.com/itsindigo/get-rich-or-die-crying/internal/app_config"
)

func main() {
	config := app_config.LoadConfig()

	fmt.Printf("Config: %v", config)
}
