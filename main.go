// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	draw_picture "local-packages/draw-picture"
	games "local-packages/games"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

var (
	flag_Game_GuessNum bool
	endNum             int
	answerNum          int
)

func init_all() {
	flag_Game_GuessNum = false
	endNum = 200
	answerNum = 0
}

func main() {
	init_all()
	port := os.Getenv("PORT")
	// // initialize our databases
	// games.InitializeGames()

	// initialize a line bot
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				quota, err := bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				reply_string := message.ID + ":" + message.Text + " OK! remain message:" + strconv.FormatInt(quota.Value, 10)

				find_space := strings.Index(message.Text, " ")
				function_type := message.Text[0:find_space]
				remain_message := message.Text[find_space+1:]

				switch function_type {
				case "抽":
					search := remain_message
					html_body := draw_picture.Get_html("https://www.google.com/search?q=" + search + "&tbm=isch")
					img_url := draw_picture.Parse(html_body)

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(img_url, img_url)).Do(); err != nil {
						log.Print(err)
					}

				case "猜數字":
					var tmp_str string
					if !flag_Game_GuessNum { // new game
						answerNum = games.CreateRandomNumber(endNum)
						flag_Game_GuessNum = true
						tmp_str = fmt.Sprintf("請輸入數字，範圍為: 0-%d", endNum)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tmp_str)).Do(); err != nil {
							log.Print(err)
						}
					} else { // continue game
						command, err := strconv.Atoi(string(remain_message)) //string to int,并作输入格式判断
						if err != nil {
							tmp_str = fmt.Sprintf("格式不對，請輸入數字")
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tmp_str)).Do(); err != nil {
								log.Print(err)
							}
						} else {

							tmp_str = fmt.Sprintf("你你輸入的數字:", command)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tmp_str)).Do(); err != nil {
								log.Print(err)
							}

							if command == answerNum {
								flag_Game_GuessNum = false
								tmp_str = fmt.Sprintf("恭喜你，答對了~")
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tmp_str)).Do(); err != nil {
									log.Print(err)
								}
							} else if command < answerNum {
								tmp_str = fmt.Sprintf("你輸入的數字小於生成的數字，别灰心!再来一次~")
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tmp_str)).Do(); err != nil {
									log.Print(err)
								}
							} else if command > answerNum {
								tmp_str = fmt.Sprintf("你輸入的數字大於生成的數字，别灰心!再来一次~")
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(tmp_str)).Do(); err != nil {
									log.Print(err)
								}
							}
						}
					}

				default:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply_string)).Do(); err != nil {
						log.Print(err)
					}
				}

			}
		}
	}
}
