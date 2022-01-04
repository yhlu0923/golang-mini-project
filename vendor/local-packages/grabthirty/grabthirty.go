package grabthirty

import (
	"strconv"
)

var left map[string]int
var lim map[string]int

func New_game(user_ip string, n int) string {
	return "New nim game, new number is: " + strconv.Itoa(n) + ".\nYour can take number from 1 to " + strconv.Itoa(k) + "\nYou take first!\n"
}

/* GrabThirty */
func GrabThirty(user_ip string, argv []string) string {

}
