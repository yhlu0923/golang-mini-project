package grabthirty

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var InfoMap map[string]GameInfo

type GameInfo struct {
	user_ip          string
	TargetNumber     int
	LastRecordNumber int
}

var EndNum int = 10
var goodnumber []int = []int{1, 5, 9, 13, 17, 21, 25, 29}

func Init() {
	InfoMap = make(map[string]GameInfo)
}

func CreateRandomNumber(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}

/* GrabThirty */
func GrabThirty(user_ip string, argv []string) string {

	if argv[1] == "--help" {
		return fmt.Sprintf("1. 使用gt, grabthirty, 搶30 當作前綴\n2. 從一開始往上喊，一次最多喊三個數字，最後搶到30的人就輸囉～\n3.剩下的根據指示操作就行嚕～")
	}

	if argv[1] == "new" {
		delete(InfoMap, user_ip)
		var game_info GameInfo
		game_info.TargetNumber = 30
		game_info.LastRecordNumber = 0
		game_info.user_ip = user_ip
		InfoMap[user_ip] = game_info

		return "New game start, you go first"
	}

	if game_info, ok := InfoMap[user_ip]; ok {

		replyNumLen := len(argv) - 1
		if !(replyNumLen >= 1 && replyNumLen <= 3) {
			return "Wrong length of reply number, it should be between 1 and 3"
		}

		var reply_nums []int
		for i := 0; i < replyNumLen; i++ {
			tmp_num, err := strconv.Atoi(string(argv[i+1])) //string to int,并作输入格式判断
			if err != nil {
				return "格式不對，請輸入\"guessnumber (數字)...(數字)\""
			}
			reply_nums = append(reply_nums, tmp_num)
		}

		tmp_num := game_info.LastRecordNumber
		for i := 0; i < replyNumLen; i++ {
			if reply_nums[i] != tmp_num+1 {
				return fmt.Sprintf("You must reply continuous numbers, start from %d", game_info.LastRecordNumber)
			} else {
				tmp_num = tmp_num + 1
			}
		}

		game_info.LastRecordNumber = reply_nums[replyNumLen-1]
		InfoMap[user_ip] = game_info

		if reply_nums[replyNumLen-1] == game_info.TargetNumber {
			return fmt.Sprintf("I win the game!, I grabbed %d, you losed", game_info.TargetNumber)
		}

		msg := Bot_move(user_ip, reply_nums)
		return msg
	}

	var game_info GameInfo
	game_info.TargetNumber = 30
	game_info.LastRecordNumber = 0
	game_info.user_ip = user_ip
	InfoMap[user_ip] = game_info

	return "Game start, you go first"
}

func Bot_move(user_ip string, reply_nums []int) string {

	n := len(reply_nums)
	game_info := InfoMap[user_ip]
	flag_num := 0

	for i := 0; i < len(goodnumber); i++ {
		if goodnumber[i] == game_info.LastRecordNumber {
			// game_info.LastRecordNumber = reply_nums[len(reply_nums)-1] + 1
			// InfoMap[user_ip] = game_info
			flag_num = reply_nums[len(reply_nums)-1] + 1

			if flag_num == game_info.TargetNumber {
				delete(InfoMap, game_info.user_ip)
				return fmt.Sprintf("You win the game!, I grabbed %d, I losed", game_info.TargetNumber)
			}
			// return fmt.Sprintf("%d", game_info.LastRecordNumber)
		}
	}

	tmp_str := "You choose:"
	for i := 0; i < n; i++ {
		tmp_str += fmt.Sprintf(" %d", reply_nums[i])
	}

	tmp_str += "\nI choose:"
	if flag_num != 0 {
		game_info.LastRecordNumber = flag_num
		tmp_str += fmt.Sprintf(" %d", flag_num)
	} else {
		tmp_num := 0
		for i := 0; i < len(goodnumber); i++ {
			if goodnumber[i]-game_info.LastRecordNumber > 0 {
				tmp_num = goodnumber[i] - game_info.LastRecordNumber
				break
			}
		}

		for i := 0; i < tmp_num; i++ {
			game_info.LastRecordNumber += 1
			tmp_str += fmt.Sprintf(" %d", game_info.LastRecordNumber)
		}
	}

	InfoMap[game_info.user_ip] = game_info
	return tmp_str
}
