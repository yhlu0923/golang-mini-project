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
	grabthirty "local-packages/grabthirty"
	nim "local-packages/nim"
	nim2 "local-packages/nim2"

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
	command = strings.Replace(command, "\t", " ", -1)

	var argv []string
	last := -1
	for {
		idx := strings.Index(command[last+1:], " ")
		if idx == -1 {
			break
		}
		idx += last + 1
		if idx-last > 1 {
			argv = append(argv, command[last+1:idx])
		}
		last = idx
	}
	if len(argv) != 0 {
		argv[0] = strings.ToLower(argv[0])
	}
	return argv
}

func initialization() {
	nim.Init()
	nim2.Init()
	games.Init()
	grabthirty.Init()
}

func main() {
	port := os.Getenv("PORT")
	// // initialize our databases
	initialization()
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
				case "help":
				    var help_msg string
				    help_msg += "????????????????????????:\n"
				    help_msg += "?????????:\n"
				    help_msg += "1. guessnumber\n"
				    help_msg += "2. nim\n"
				    help_msg += "3. nim2\n"
				    help_msg += "4. grabthirty\n"
				    help_msg += "?????????:\n"
				    help_msg += "1. ???\n"
				    help_msg += "2. random\n"
				    help_msg += "??????????????????????????????????????? \"--help\" ???????????????????????????????????????"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(help_msg)).Do(); err != nil {
						log.Print(err)
					}
				case "???":
				    if len(argv) == 1 {
				        msg := "?????? \"??? <message>\" ??????????????????"
					    if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						    log.Print(err)
					    }
				    } else {
				        var search string
				        for idx, i := range argv {
                            if idx > 1 {
                                search += "+"
                            }
				            if idx != 0 {
				                search += i
                            }
				        }
					    html_body := draw_picture.Get_html("https://tw.images.search.yahoo.com/search/images?p=" + search)
                        img_url := draw_picture.Parse(html_body)
                        if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(img_url, img_url)).Do(); err != nil {
						    log.Print(err)
					    }
                    }
				case "?????????", "guessnumber", "GuessNumber", "gn":
					msg := games.GuessNumber(client_ip, argv)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				case "nim":
					msg := nim.Play_nim(client_ip, argv)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				case "nim2":
					msg := nim2.Play_nim(client_ip, argv)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				case "?????????", "???30", "grabthirty", "gt":
					msg := grabthirty.GrabThirty(client_ip, argv)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				case "random":
					if argv[1] == "--help" {
						msg := "1. Use random as the first argument\n2. Followed with a number which is the end of range you desire.\n3. Get the random number!"
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
							log.Print(err)
						}
						return
					}

					msg := ""
					endNum, err := strconv.Atoi(string(argv[1])) //string to int,????????????????????????
					if err != nil {
						msg = "????????????????????????\"Random (??????)\""
					} else if endNum <= 0 {
					    msg = "??????????????????"
					}else {
						msg = fmt.Sprintf("Random number from 1-%d is \"%d\"", endNum, games.CreateRandomNumber(endNum) + 1)
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				default:
					reply_string := "????????????????????? \"help\" ??????????????????"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply_string)).Do(); err != nil {
						log.Print(err)
					}
				}

			}
		}
	}
}
