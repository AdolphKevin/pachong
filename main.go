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
	doneCityMap()
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
		for _, value := range st {
			for _, v := range value.Name {
				// 判断字符串是否包含地名
				if mapVal, ok := transport.CityMap[v.EnglishName]; ok {
					// 去重
					if mapVal == "" {
						transport.CityMap[v.EnglishName] = v.ChineseName
						w := bufio.NewWriter(f)
						lineStr := fmt.Sprintf("cityMap[\"%s\"] = \"%s\"", v.EnglishName, v.ChineseName)
						fmt.Fprintln(w, lineStr)
						w.Flush()
					}
				}
			}
			fmt.Printf("value:%v\n", value)
		}
	})

	for english, _ := range transport.CityMap {
		// 将请求参数中的空格用%20代替
		english = strings.Replace(english, " ", "%20", -1)
		// 拼接URL
		url := fmt.Sprintf("%s%s", transportUrl, english)
		fmt.Printf("%s,", url)
		// 随机用户代理
		extensions.RandomUserAgent(c)
		fmt.Println()
		// 发起请求
		err = c.Visit(url)
		if err != nil {
			fmt.Println("c.Visit err:", err)
			break
		}
		// 创建定时器
		time.Sleep(time.Second * 4)
	}

	//	结束语
	w := bufio.NewWriter(f)
	lineStr := fmt.Sprintf("%s", "已经结束")
	fmt.Fprintln(w, lineStr)
	w.Flush()
}

type City struct {
	Name []CityName `json:"juku"`
}

type CityName struct {
	Explain     string `json:"cls1"`
	ChineseName string `json:"cntxt"`
	EnglishName string `json:"entxt"`
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
