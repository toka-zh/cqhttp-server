package pkg

import (
	"cqhttp-server/config"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func PixivCraw(target string) error {
	log.Println("pixiv craw running...")

	// 确认下载目录在不在
	_, err := os.Stat(config.SavePath)
	if err != nil {
		log.Println("download dir doesn't exist,recreate...")
		os.Mkdir(config.SavePath, 0777)
	}

	// 获取不能爬取的列表
	// getRobots(target)

	client := &http.Client{}

	// 如果请求错误使用代理
	testReq, _ := http.NewRequest(http.MethodGet, target, nil)
	if _, err := client.Do(testReq); err != nil {
		log.Println("connect failed,try to use proxy...")
		proxy, _ := url.Parse("http://localhost:7890")
		transport := &http.Transport{Proxy: http.ProxyURL(proxy)}
		client = &http.Client{Transport: transport}
	}

	// 获取链接内容
	req, _ := http.NewRequest(http.MethodGet, target, nil)
	do, err := client.Do(req)
	if err != nil {
		return err
	}
	all, _ := io.ReadAll(do.Body)
	body := string(all)

	// 转化为dom对象
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
	dom.Find("._layout-thumbnail").Each(
		func(i int, selection *goquery.Selection) {
			last := selection.Children().Last()
			path, exists := last.Attr("data-src")
			if !exists {
				return
			}
			id, ok := last.Attr("data-id")
			if !ok {
				return
			}

			//存在性检查
			filename := config.SavePath + path[strings.LastIndex(path, "/")+1:]
			if FileExists(filename) {
				return
			}

			parse, err := url.Parse(path)
			path1 := "https://" + parse.Host + parse.Path[strings.Index(parse.Path, "/img-master"):]
			req, _ := http.NewRequest("GET", path1, nil)
			req.Header.Set("Referer", "https://www.pixiv.net/artworks/"+id)

			data, err := client.Do(req)
			if err != nil {
				log.Println(req.URL, " failed")
				return
			}

			readAll, err := io.ReadAll(data.Body)
			if err != nil {
				log.Println(req.URL, " failed")
				return
			}

			err = os.WriteFile(filename, readAll, 0666)
			if err != nil {
				log.Println(req.URL, " failed")
				return
			}
			log.Println(req.URL, " success")
		},
	)
	return nil
}

func getRobots(target string) {
	u, err := url.Parse(target)
	if err != nil {
		return
	}
	// todo http链接处理
	robot := u.Scheme + "://" + u.Host + "/robots.txt"

	resp, err := http.Get(robot)
	if err != nil {
		return
	}
	all, _ := io.ReadAll(resp.Body)
	s := string(all)

	split := strings.Split(s, "\n")

	// 新建匹配器
	compile, _ := regexp.Compile("^Disallow")
	for _, v := range split {
		if !compile.MatchString(v) {
			continue
		}
	}
}
