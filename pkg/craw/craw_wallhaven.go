package craw

import (
	"cqhttp-server/pkg"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func WallHavenCraw(target, filePath string) error {

	// 确认下载目录在不在
	_, err := os.Stat(filePath)
	if err != nil {
		os.Mkdir(filePath, 0777)
	}

	client := &http.Client{}

	// 获取链接内容
	req, _ := http.NewRequest(http.MethodGet, target, nil)
	AddWHCookie(req)
	// 获取内容
	do, err := client.Do(req)
	// 如果请求错误使用代理重新获取
	if err != nil {
		log.Println("connect failed,switch to proxy")
		proxy, _ := url.Parse("http://localhost:7890")
		transport := &http.Transport{Proxy: http.ProxyURL(proxy)}
		client = &http.Client{Transport: transport}

		do, err = client.Do(req)
		if err != nil {
			return errors.New("failed connect")
		}
	}

	metaList, _ := io.ReadAll(do.Body)
	previewList := string(metaList)

	// 转化为dom对象
	previewSelector, _ := goquery.NewDocumentFromReader(strings.NewReader(previewList))
	previewSelector.Find(".preview").Each(
		func(_ int, selection *goquery.Selection) {
			preview, exists := selection.Attr("href")
			if !exists {
				return
			}

			detailUrl, _ := url.Parse(preview)
			req.URL = detailUrl

			detailRes, _ := client.Do(req)
			metaDetail, _ := io.ReadAll(detailRes.Body)
			detailData := string(metaDetail)
			detailSelector, _ := goquery.NewDocumentFromReader(strings.NewReader(detailData))
			detailSelector.Find("#wallpaper").Each(
				func(_ int, selection *goquery.Selection) {
					fileURL, exists := selection.Attr("src")
					if !exists {
						return
					}

					req, _ := http.NewRequest(http.MethodGet, fileURL, nil)
					AddWHCookie(req)

					parse, _ := url.Parse(fileURL)
					split := strings.Split(parse.Path, "/")
					filename := split[len(split)-1]
					if pkg.FileExists(filename) {
						return
					}

					data, err := client.Do(req)
					if err != nil {
						log.Println(req.URL, " failed", err.Error())
						return
					}

					readAll, err := io.ReadAll(data.Body)
					if err != nil {
						log.Println(req.URL, " failed", err.Error())
						return
					}

					err = os.WriteFile(filePath+filename, readAll, 0666)
					if err != nil {
						log.Println(req.URL, " failed", err.Error())
						return
					}
					log.Println(req.URL, " success")
				})
		},
	)
	log.Println("finish")
	return nil
}

func AddWHCookie(req *http.Request) {
	req.Header.Set("Cookie", "_pk_id.1.01b8=ed751ff6aadd26c5.1663683548.; remember_web_59ba36addc2b2f9401580f014c7f58ea4e30989d=eyJpdiI6ImYwanVGZDFoTm93VjJ2VytVR1c3MWc9PSIsInZhbHVlIjoiNndabE9ZNnJxS0wyTmhkRG8xaVgwaHNYREY2Vkhya1RYZVROeWJ6VUVsRDlMXC8rSGZYTVlCU1E3clwvc1FVbkZSYUl5VThmTDJ2OUZlUWV2Y2g2RHplaHN0NnB4M2oyeXdIdFRjdVdUY1I1S1d0OWxcL1l1NjU3dlVqTjlTWmowTWo0QjNZQmxPckNObEVUeHR1S2pEVEFua212dEZVR2dLMEJYZTJrZm1Wb243cVwvazl3bFkrM1FsZXNXdEVjMzdOUiIsIm1hYyI6IjY1NjJkYWNmYWFiNzE1ZDVmNzg3ZjFiM2FiY2I3M2MzYTZlNzA2NWY3ZjQ4MTVkMzM1MGQ1OTNkYTFlOTFkYTMifQ==; _pk_ses.1.01b8=1; XSRF-TOKEN=eyJpdiI6IitJcEJKK3lrZkpkTkNpTHFaTEtWUHc9PSIsInZhbHVlIjoibFwvckh5RDViTlhFSzRKTEZDNzJMODlMdWV3MTZrbGhLb05sMGNsWmNkZkpHb3pZcEpEd1VKQnVMTXNoWWVQb3IiLCJtYWMiOiIwNGMxNWQ0OTJjNWZhNzMzODNhNjIwZjc5YjI3OGJiMmE5ZjUzZTRiZmJjMDUzYjAyZmMzZDkzY2RiODVjZTkwIn0=; wallhaven_session=eyJpdiI6IlhTUzBMcHVPTlNOdTVJdnJGREoyQ3c9PSIsInZhbHVlIjoiOFZEUTNKUnhBRWhOb2NjbEVcL28xNDdsTjBNR3Q3cUZ5ZU9QZnVmY1pHUnU4RTNBeWRpMEwyZno0RlRcL3Budlg5IiwibWFjIjoiMTgxNDFmOWM4NDgxMzU5OGNmMWQ5OWMyNjc0MzkzYzdiNTc1NDNlN2MzOTJmODY1NjkwZjQ1Yzc4NmY3ZmNlYSJ9")
}
