package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/AdolphKevin/pachong/transport"

	"github.com/gocolly/colly/extensions"

	"github.com/gocolly/colly"
)

func main() {
	transportUrl := "https://www.tdict.com/tdictajax/tdictajax_new.asp?action=query&stype=dmyd&sclass=0&q="

	f, err := os.OpenFile("./transport.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	doneProvinceMap()
	//doneCityMap()
	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// After making a request get "url" from
	// the context of the request
	c.OnResponse(func(r *colly.Response) {
		var st []*City
		err := json.Unmarshal(r.Body, &st)
		if err != nil {
			fmt.Printf("转数据结构失败:%v\n", err)
			return
		}
		// 获取请求的英语
		q := r.Request.URL.Query()
		queryEnglish := make([]string, 1)
		queryEnglish = q["q"]

		if len(st[0].Name) == 0 {
			// 尝试解析，是否被封杀了IP
			var errMsg []*ErrorMsg
			err := json.Unmarshal(r.Body, &errMsg)
			if err != nil {
				fmt.Printf("转数据结构失败:%v\n", err)
				return
			}
			if errMsg[0].ErrMsg == "Sorry,you are denied! Problem Id:951" {
				fmt.Println("IP已被封杀")
				//os.Exit(0)
				return
			}
			// 查不到内容
			w := bufio.NewWriter(f)
			lineStr := fmt.Sprintf("ProvinceMap[\"%s\"] = \"%s\"", queryEnglish[0], "QueryFalse")
			fmt.Fprintln(w, lineStr)
			w.Flush()
			return
		}
		for _, value := range st {
			chineseName := removeUnnecessaryString(value.Name[0].ChineseName)
			// 直接文本记录
			w := bufio.NewWriter(f)
			lineStr := fmt.Sprintf("ProvinceMap[\"%s\"] = \"%s\"", queryEnglish[0], chineseName)
			fmt.Fprintln(w, lineStr)
			w.Flush()

			fmt.Printf("value:%v\n", value)
		}
	})

	for _, english := range transport.RemainProvince {
		// 判断是否已经记录
		if _, ok := provinceMap[english]; ok {
			continue
		}
		// 将请求参数中的空格用%20代替
		english = strings.Replace(english, " ", "%20", -1)
		// 拼接URL
		url := fmt.Sprintf("%s%s", transportUrl, english)
		fmt.Printf("%s,", url)
		// 随机用户代理
		extensions.RandomUserAgent(c)
		// 发起请求
		err = c.Visit(url)
		if err != nil {
			fmt.Println("c.Visit err:", err)
			break
		}
		// 创建定时器
		time.Sleep(time.Second * 4)
	}

	// 结束后来个结尾语
	w := bufio.NewWriter(f)
	lineStr := fmt.Sprintf("%s", "已经结束")
	fmt.Fprintln(w, lineStr)
	w.Flush()
}

func removeUnnecessaryString(name string) string {
	// 移除省、市、区以及括号内的字符串
	name = strings.Replace(name, "省", "", -1)
	name = strings.Replace(name, "市", "", -1)
	name = strings.Replace(name, "区", "", -1)
	name = strings.Replace(name, "县", "", -1)
	var newName strings.Builder
	for _, char := range name {
		if char == '(' {
			break
		}
		newName.WriteString(fmt.Sprintf("%c", char))
	}
	return newName.String()
}

type City struct {
	Name []CityName `json:"juku"`
}

type CityName struct {
	Explain     string `json:"cls1"`
	ChineseName string `json:"cntxt"`
	EnglishName string `json:"entxt"`
}

type ErrorMsg struct {
	ErrMsg string `errmsg`
}

// 获取一个随机时间
func getRandomInt() (result int64) {
	result = 30
	for i := 0; i < 1000; i++ {
		s := rand.NewSource(time.Now().UnixNano())
		result = s.Int63() % 256
		if result < 30 || result > 50 {
			continue
		} else {
			break
		}
	}
	return result
}
