package nim

import (
	"strconv"
	"math/rand"
	"time"
)

var left map[string]int
var lim map[string]int

func new_game(user_ip string, n int) string {
    delete(left, user_ip)
    delete(lim, user_ip)
    if n <= 4 {
        return "Invalid number!\nThe new game number must greater than 4!"
    }
    left[user_ip] = n
    k := rnd(n / 2 - 2) + 2
    lim[user_ip] = k
    return "New nim game, new number is: " + strconv.Itoa(n) + ".\nYour can take number from 1 to " + strconv.Itoa(k) + "\nYou take first!\n"
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
        return new_game(user_ip, n)
    } else { // move
        number, flag := left[user_ip]
        k, _ := lim[user_ip]
        if flag {
            if number < n || number < k {
                return "Invalid move!"
            }
            re := "The number was " + strconv.Itoa(number) + ".\n"
            number -= n;
            re = re + "You take" + strconv.Itoa(n) + ".\n"
            bn := bot_move(number, k)
            if bn == -1 {
                re = re + "I can't move, You WIN!!"
            } else {
                re = re + "I take " + strconv.Itoa(bn) + "."
            }
            return re;
        } else {
            return "No nim game exist, try nim new <number>!"
        }
    }
    return "Some error occur!"
}

func bot_move(n int, k int) int {
    var re int
    if n == 0 {
        re = -1;
    } else if n % k != 0 {
        re = n % k
    } else {
        re = rnd(k) + 1
    }
    return re
}

func rnd(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}
