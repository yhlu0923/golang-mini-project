package nim

import (
    "strings"
	"strconv"
	"math/rand"
	"time"
)

var left map[string]int
var lim map[string]int

func new_game(user_ip string, n int) string {
    left.delete(user_ip)
    lim.delete(lim)
    if n <= 4 {
        return "Invalid number!\nThe new game number must greater than 4!"
    }
    left[user_ip] = n
    lim[user_ip] = rnd(n / 2 - 2) + 2
    return "New nim game, new number is: " + strconv.Itoa(n) + ".\nYou take first!\n"
}

func play_nim(user_ip string, argv []string) string {
    if len(argv) < 3 || (argv[1] != "new" || argv[1] != "take") {
        return "Usage: nim new|take <number>"
    }
    n, err := strconv.Atoi(argv[1])
    if err != nil || n <= 0 {
        return "Invalid number!"
    }
    if argv[1] == "new" { // new game
        return new_game(n)
    } else {
        number, flag := left[user_ip]
        k, _ := lim[user_ip]
        if flag {
            if number < n || number < k {
                return "invalid move!"
            }
            re := "The number was " + strconv.Itoa(number) + ".\n"
            number -= n;
            re = re + "You take" + strconv.Itoa(n) + ".\n"
            bn = bot_move(number, k)
            if bn == -1 {
                re := "I can't move, You WIN!!"
            } else {
                re = "I take " + strconv.Itoa(bn) + "."
            }
            return re;
        } else {
            return "No nim game exist, try nim new <number>!"
        }
    }
    return "Some error occur!"
}

func bot_move(n int, k int) {
    if n % k != 0 {
        return n % k
    }
    return rnd(k) + 1
}

func rnd(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}
