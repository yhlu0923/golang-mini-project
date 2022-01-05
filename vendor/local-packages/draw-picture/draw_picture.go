package draw_picture

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
	"math/rand"
)

func Get_html (url string) string {
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
	response.Body.Close()
	return string(body)
}

func Parse(body string) string {
	body = strings.Replace(body, "\n", "", -1)
	img_reg := regexp.MustCompile(`<a class="img"    href='(.*?)'aria`)
	select_idx := img_reg.FindAllStringSubmatch(body, -1)[Rnd(10)][1]
	
	body2 := get_html("https://tw.images.search.yahoo.com/" + select_idx)
	body2 = strings.Replace(body2, "\n", "", -1)
	
	src_reg := regexp.MustCompile(`<img src=('|")(.*?)&.*?('|")(.*?)>`)
	img_url := src_reg.FindAllStringSubmatch(body2, -1)[0][2]
	
	return img_url
}

func Rnd(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}
