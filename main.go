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

	draw_picture "local-packages"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {
	// initialize our databases

	// initialize a line bot
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

////////////////////////////////////////////////////////////////
//////////////// Function of parsing picture ///////////////////
////////////////////////////////////////////////////////////////

// func get_html(url string) string {
// 	fmt.Println("Fetch Url", url)
// 	client := &http.Client{}
// 	request, _ := http.NewRequest("GET", url, nil)
// 	request.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
// 	response, err := client.Do(request)
// 	if err != nil {
// 		fmt.Println("Http get err:", err)
// 		return ""
// 	}
// 	if response.StatusCode != 200 {
// 		fmt.Println("Http status:", response.StatusCode)
// 		return ""
// 	}
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("ReadAll get err:", err)
// 		return ""
// 	}
// 	response.Body.Close()
// 	return string(body)
// }

// func parse(body string) string {
// 	body = strings.Replace(body, "\n", "", -1)
// 	img_reg := regexp.MustCompile(`<img class=(.*?)>`)
// 	src_reg := regexp.MustCompile(`src="(.*?)"`)
// 	img_url := src_reg.FindAllStringSubmatch(img_reg.FindAllStringSubmatch(body, -1)[1][1], -1)[0][1]
// 	return img_url
// }

////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////

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

				function_type := message.Text[0:3]
				switch function_type {
				case "抽":
					search := message.Text[3:]
					html_body := draw_picture.Get_html("https://www.google.com/search?q=" + search + "&tbm=isch")
					img_url := draw_picture.Parse(html_body)

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(img_url, img_url)).Do(); err != nil {
						log.Print(err)
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
