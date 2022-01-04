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

func CreateRandomNumber(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}

/* GrabThirty */
func GrabThirty(user_ip string, argv []string) string {

	if game_info, ok := InfoMap[user_ip]; ok {

		replyNumLen := len(argv) - 1
		if !(replyNumLen >= 1 && replyNumLen <= 3) {
			return "Wrong length of reply number, it should be between 1 and 3"
		}

		var reply_nums []int
		for i := 1; i < len(argv); i++ {
			tmp_num, err := strconv.Atoi(string(argv[i])) //string to int,并作输入格式判断
			if err != nil {
				return "格式不對，請輸入\"guessnumber (數字)...(數字)\""
			}
			reply_nums = append(reply_nums, tmp_num)
		}

		for i := 0; i < replyNumLen; i++ {
			if reply_nums[i] != game_info.LastRecordNumber+1 {
				return "You must reply continuous numbers"
			}
		}
		game_info.LastRecordNumber = reply_nums[replyNumLen-1]

		msg := Bot_move(game_info, replyNumLen)
		return msg
	}

	var game_info GameInfo
	game_info.TargetNumber = 30
	game_info.LastRecordNumber = 0
	game_info.user_ip = user_ip
	InfoMap[user_ip] = game_info

	return "Game start, you go first"
}

func Bot_move(game_info GameInfo, n int) string {

	for i := 0; i < len(goodnumber); i++ {
		if goodnumber[i] == game_info.LastRecordNumber {
			game_info.LastRecordNumber += 1
			if game_info.LastRecordNumber == game_info.TargetNumber {
				delete(InfoMap, game_info.user_ip)
				return fmt.Sprintf("You win the game!, I grabbed %d, I losed", game_info.TargetNumber)
			}
			return string(game_info.LastRecordNumber)
		}
	}

	tmp_str := ""
	for i := 0; i < 4-n; i++ {
		game_info.LastRecordNumber += 1
		tmp_str += string(game_info.LastRecordNumber) + " "
	}
	return tmp_str
}
