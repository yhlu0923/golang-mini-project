// Function of parsing picture

package draw_picture

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func Get_html(url string) string {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if response.StatusCode != 200 {
		fmt.Println("Http status:", response.StatusCode)
		return ""
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ReadAll get err:", err)
		return ""
	}
	response.Body.Close()
	return string(body)
}

func Parse(body string) string {
	body = strings.Replace(body, "\n", "", -1)
	img_reg := regexp.MustCompile(`<img class=(.*?)>`)
	src_reg := regexp.MustCompile(`src="(.*?)"`)
	img_url := src_reg.FindAllStringSubmatch(img_reg.FindAllStringSubmatch(body, -1)[Rnd(5) + 1][1], -1)[0][1]
	return img_url
}

func Rnd(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}
