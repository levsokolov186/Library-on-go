package UI

import (
	"Library/UI/Authentication"
	"Library/UI/ShelvesUI"
)

func StartMainMenu() {
	Authentication.Entrance()
	ShelvesUI.PrintShelves()
}
