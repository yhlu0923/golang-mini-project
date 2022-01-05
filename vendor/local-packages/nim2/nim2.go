package nim2

import (
	"math/rand"
	"strconv"
	"time"
	"sort"
)

var nim_left map[string]int
var nim_lim map[string]map[int]bool

func Init() {
    nim_left = make(map[string]int)
    nim_lim = make(map[string]map[int]bool)
}

func New_game(user_ip string, n int) string {
	delete(nim_left, user_ip)
	delete(nim_lim, user_ip)
	if n <= 5 {
		return "一場遊戲至少要有六顆石頭！"
	}
	if n > 1000 {
	    return "一場遊戲至多能有 1000 顆石頭！"
	}
	nim_left[user_ip] = n
	lim := n / 2
	if lim > 10 {
	    lim = 10
	}
	arr := make(map[int]bool)
	amount := Rnd(lim - 2) + 2
	for i := 0; i < amount; i++ {
	    now := Rnd(lim) + 1
	    _, has := arr[now]
	    if has {
	        i--
	    } else {
	        arr[now] = true
        }
	}
	nim_lim[user_ip] = arr
	return_msg := "開始新的 nim! 目前有 " + strconv.Itoa(n) + " 顆石頭.\n你每次可以拿"
	var tmp []int
	for i, _ := range arr {
	    tmp = append(tmp, i)
	}
	sort.Slice(tmp, func(x, y int) bool {return tmp[x] < tmp[y]})
	for _, i := range tmp {
	    return_msg += " " + strconv.Itoa(i)
    }
	return_msg += " 顆石頭\n你先手!"
	return return_msg
}

func Play_nim(user_ip string, argv []string) string {
    if len(argv) > 1 && argv[1] == "--help" { // rule
        rule := ""
        rule += "1. 使用 \"nim new <number>\" 來開始一個初始為 <number> 顆石頭的遊戲\n"
        rule += "2. 使用 \"nim take <number>\" 取走 <number> 顆石頭\n"
        rule += "3. 無法拿石頭者敗"
        return rule
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
    		_, can_take := k[n]
			if number < n || can_take == false {
				return "非法的操作!"
			}
			re := "原本有 " + strconv.Itoa(number) + " 顆石頭.\n"
			number -= n
			re = re + "你拿了 " + strconv.Itoa(n) + " 顆石頭.\n"
			bn := Bot_move(number, k)
			if bn == -1 {
				re = re + "沒有辦法再拿石頭了，你贏了!!"
				delete(nim_left, user_ip)
				delete(nim_lim, user_ip)
			} else {
			    number -= bn
				re = re + "我拿了 " + strconv.Itoa(bn) + " 顆石頭.\n"
				can_take_precheck := false
				for i, _ := range k {
				    if i <= number {
                        can_take_precheck = true
				        break
				    }
				}
				if can_take_precheck == false {
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

func Bot_move(n int, k map[int]bool) int {
	var dp []int
	re := -1
	dp = append(dp, 0)
	for i := 1; i <= n; i++ {
	    var SG [11]bool
	    for j, _ := range k {
	        if j <= i {
	            SG[dp[i - j]] = true
                if i == n && dp[i - j] == 0 && re < j{
                    re = j
                }
            }
	        // fmt.Println(i, j)
	    }
	    for j := 0; j <= 10; j++ {
	        if SG[j] == false {
	            dp = append(dp, j)
	            break
	        }
	    }
	}
	if re != -1 {
	    return re
    }
	
	// rnd move
	rnd_sz := 0
	rnd_mx := 0
	for j, _ := range k {
	    if j <= n {
	        rnd_sz++
	        if j <= rnd_mx {
	            rnd_mx = j
	        }
	    }
	}
	rnd_idx := Rnd(rnd_sz)
	for j, _ := range k {
	    if j <= n {
	        rnd_idx--
	        if rnd_idx == 0 {
	            return j
	        }
	    }
	}
	return -1
}

func Rnd(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}
