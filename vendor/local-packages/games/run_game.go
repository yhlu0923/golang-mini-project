package games

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/* Guess num */

var InfoMap map[string]GameInfo

type GameInfo struct {
	AnswerNum int
}


func Init() {
	InfoMap = make(map[string]GameInfo)
}

func CreateRandomNumber(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}

func GuessNumber(user_ip string, argv []string) string {
    if len(argv) == 1 {
        return fmt.Sprintf("使用 \"gn --help\" 獲得更多資訊")
    }
	if argv[1] == "--help" {
	    var help_msg string
	    help_msg += "1. 使用 gn, guessnumber, 猜數字 當作前綴\n"
	    help_msg += "2. 使用 \"gn new <number>\" 來開始新的遊戲\n"
	    help_msg += "3. 使用 \"gn <number>\" 來猜數字"
		return help_msg
	}
	
    var EndNum int = 10

	if argv[1] == "new" {
		if len(argv) >= 3 {
			num, err := strconv.Atoi(string(argv[2])) //string to int,并作输入格式判断
			if err != nil {
				return "new 後面的要接數字"
			}
			EndNum = int(num)
		}
		delete(InfoMap, user_ip)
		var game_info GameInfo
		game_info.AnswerNum = CreateRandomNumber(EndNum) + 1
		InfoMap[user_ip] = game_info
		return fmt.Sprintf("請輸入數字，範圍為: 1-%d", EndNum)
	}

	if game_info, ok := InfoMap[user_ip]; ok {

		if len(argv) == 1 {
			return "格式不對，請輸入\"猜數字 (數字)\""
		}
		command, err := strconv.Atoi(string(argv[1])) //string to int,并作输入格式判断
		if err != nil {
			return "格式不對，請輸入\"猜數字 (數字)\""
		} else {

			if command == game_info.AnswerNum {
				delete(InfoMap, user_ip)
				return fmt.Sprintf("恭喜你，答對了~, 答案就是 %d", game_info.AnswerNum)
			} else if command < game_info.AnswerNum {
				return fmt.Sprintf("你輸入的數字(%d)小於生成的數字，别灰心!再来一次~", command)
			} else if command > game_info.AnswerNum {
				return fmt.Sprintf("你輸入的數字(%d)大於生成的數字，别灰心!再来一次~", command)
			}
		}
		return "Somethong went wrong"
	}

	var game_info GameInfo
	game_info.AnswerNum = CreateRandomNumber(EndNum)
	InfoMap[user_ip] = game_info
	return fmt.Sprintf("請輸入數字，範圍為: 1-%d", EndNum)
}
