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
	"strings"

	draw_picture "local-packages/draw-picture"
	games "local-packages/games"
	grabthirty "local-packages/grabthirty"
	nim "local-packages/nim"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func get_ip(r *http.Request) string {
	info := fmt.Sprint(*r)
	idx := strings.Index(info, "X-Forwarded-For:") + 17
	var ip string
	for ; info[idx] != ']'; idx++ {
		ip = ip + string(info[idx])
	}
	return ip
}

func parse_command(command string) []string {
	command = command + " "
	command = strings.ToLower(command)
	command = strings.Replace(command, "\t", " ", -1)

	var arg []string
	last := -1
	for {
		idx := strings.Index(command[last+1:], " ")
		if idx == -1 {
			break
		}
		idx += last + 1
		if idx-last > 1 {
			arg = append(arg, command[last+1:idx])
		}
		last = idx
	}
	return arg
}

func main() {
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

				argv := parse_command(message.Text)
				client_ip := get_ip(r)
				switch argv[0] {
				case "抽":
					search := argv[1]
					html_body := draw_picture.Get_html("https://www.google.com/search?q=" + search + "&tbm=isch")
					img_url := draw_picture.Parse(html_body)

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(img_url, img_url)).Do(); err != nil {
						log.Print(err)
					}

				case "猜數字", "guessnumber", "GuessNumber":
					msg := games.GuessNumber(client_ip, argv)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				case "nim", "Nim":
					msg := nim.Play_nim(client_ip, argv)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				case "搶三十", "搶30", "grabthirty":
					msg := grabthirty.GrabThirty(client_ip, argv)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				default:
					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply_string)).Do(); err != nil {
					// 	log.Print(err)
					// }
				}

			}
		}
	}
}
