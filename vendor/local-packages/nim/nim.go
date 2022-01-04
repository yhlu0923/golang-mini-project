package nim

import (
	"math/rand"
	"strconv"
	"time"
)

var nim_left map[string]int
var nim_lim map[string]int

func Init() {
    nim_left = make(map[string]int)
    nim_lim = make(map[string]int)
}

func New_game(user_ip string, n int) string {
	delete(nim_left, user_ip)
	delete(nim_lim, user_ip)
	if n <= 4 {
		return "Invalid number!\nThe new game number must greater than 4!"
	}
	nim_left[user_ip] = n
	k := Rnd(n / 2 - 2) + 2
	nim_lim[user_ip] = k
	return "New nim game! New number is: " + strconv.Itoa(n) + ".\nYour can take number from 1 to " + strconv.Itoa(k) + "\nYou take first!"
}

func Play_nim(user_ip string, argv []string) string {
	if len(argv) < 3 || (argv[1] != "new" && argv[1] != "take") {
		return "Usage: nim new|take <number>"
	}
	n, err := strconv.Atoi(argv[2])
	if err != nil || n <= 0 {
		return "Invalid number!"
	}
	if argv[1] == "new" { // new game
		return New_game(user_ip, n)
	} else { // move
		number, flag := nim_left[user_ip]
		k, _ := nim_lim[user_ip]
		if flag {
			if number < n || k < n {
				return "Invalid move!"
			}
			re := "The number was " + strconv.Itoa(number) + ".\n"
			number -= n
			re = re + "You take " + strconv.Itoa(n) + ".\n"
			bn := Bot_move(number, k)
			if bn == -1 {
				re = re + "I can't move, You WIN!!"
				delete(nim_left, user_ip)
				delete(nim_lim, user_ip)
			} else {
			    number -= bn
				re = re + "I take " + strconv.Itoa(bn) + ".\n"
				if number == 0 {
				    re += "You lose"
    				delete(nim_left, user_ip)
	    			delete(nim_lim, user_ip)
				} else {
        			re += strconv.Itoa(number) + " left"
        			nim_left[user_ip] = number
        	    }
			}
			return re
		} else {
			return "No nim game exist, try \"nim new <number>\"!"
		}
	}
	return "Some error occur!"
}

func Bot_move(n int, k int) int {
	var re int
	if n == 0 {
		re = -1
	} else if n % (k + 1) != 0 {
		re = n % (k + 1)
	} else {
		re = Rnd(k) + 1
	}
	return re
}

func Rnd(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}
