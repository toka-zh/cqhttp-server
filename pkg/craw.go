package pkg

import (
	"cqhttp-server/internal/pkg"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func Craw(target string) error {
	//getRobots(target)
	log.Println("pixiv craw starting...")

	// 获取排行版的所有图片
	resp, err := http.Get(target)
	if err != nil {
		return err
	}
	all, _ := io.ReadAll(resp.Body)
	body := string(all)

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
			filename := "./download/" + path[strings.LastIndex(path, "/")+1:]
			if pkg.FileExists(filename) {
				return
			}

			parse, err := url.Parse(path)
			path1 := "https://" + parse.Host + parse.Path[strings.Index(parse.Path, "/img-master"):]
			req, _ := http.NewRequest("GET", path1, nil)
			req.Header.Set("Referer", "https://www.pixiv.net/artworks/"+id)

			data, err := (&http.Client{}).Do(req)
			if err != nil {
				return
			}

			readAll, err := io.ReadAll(data.Body)
			if err != nil {
				return
			}

			err = os.WriteFile(filename, readAll, 0666)
			if err != nil {
				return
			}
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
