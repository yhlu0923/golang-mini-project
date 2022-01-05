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
	if n <= 5 {
		return "一場遊戲至少要有六顆石頭！"
	}
	nim_left[user_ip] = n
	k := Rnd(n / 2 - 2) + 2
	nim_lim[user_ip] = k
	return "開始新的 nim! 目前有 " + strconv.Itoa(n) + " 顆石頭.\n你每次可以拿 1 到 " + strconv.Itoa(k) + " 顆石頭\n你先手!"
}

func Play_nim(user_ip string, argv []string) string {
    if len(argv) > 1 && argv[1] == "--help" { // rule
        rule := ""
        rule += "使用 nim new <number> 來開始一個初始為 <number> 顆石頭的遊戲\n"
        rule += "使用 nim take <number> 取走 <number> 顆石頭\n"
        rule += "將石頭拿光者勝"
    }
	if len(argv) < 3 || (argv[1] != "new" && argv[1] != "take") {
		return "使用方法: nim new|take <number>"
	}
	n, err := strconv.Atoi(argv[2])
	if err != nil || n <= 0 {
		return "非法的數字!"
	}
	if argv[1] == "new" { // new game
		return New_game(user_ip, n)
	} else { // move
		number, flag := nim_left[user_ip]
		k, _ := nim_lim[user_ip]
		if flag {
			if number < n || k < n {
				return "非法的操作!"
			}
			re := "原本有 " + strconv.Itoa(number) + " 顆石頭.\n"
			number -= n
			re = re + "你拿了 " + strconv.Itoa(n) + " 顆石頭.\n"
			bn := Bot_move(number, k)
			if bn == -1 {
				re = re + "沒有石頭了，你贏了!!"
				delete(nim_left, user_ip)
				delete(nim_lim, user_ip)
			} else {
			    number -= bn
				re = re + "我拿了 " + strconv.Itoa(bn) + " 顆石頭.\n"
				if number == 0 {
				    re += "你輸了"
    				delete(nim_left, user_ip)
	    			delete(nim_lim, user_ip)
				} else {
        			re += "剩 " + strconv.Itoa(number) + " 顆石頭"
        			nim_left[user_ip] = number
        	    }
			}
			return re
		} else {
			return "目前沒有執行中的 nim, 試試使用 \"nim new <number>\"!"
		}
	}
	return "出現錯誤!"
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
