package games

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Guess num

//生成规定范围内的整数
//设置起始数字范围，0开始,endNum截止
func CreateRandomNumber(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}

func GuessNumber() {

}

func GuessNumber_Continue(remain_message string, Flag_Game_GuessNum *bool, AnswerNum int) string {

	command, err := strconv.Atoi(string(remain_message)) //string to int,并作输入格式判断
	if err != nil {
		return "格式不對，請輸入\"猜數字 (數字)\""
	} else {

		if command == AnswerNum {
			*Flag_Game_GuessNum = false
			return fmt.Sprintf("恭喜你，答對了~, 答案就是%d", AnswerNum)
		} else if command < AnswerNum {
			return fmt.Sprintf("你輸入的數字(%d)小於生成的數字，别灰心!再来一次~", command)
		} else if command > AnswerNum {
			return fmt.Sprintf("你輸入的數字(%d)大於生成的數字，别灰心!再来一次~", command)
		}
	}
	return "Somethong went wrong"
}
