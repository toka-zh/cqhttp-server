package pkg

import (
	"cqhttp-server/config"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func WallHavenCraw(target string) error {
	log.Println("pixiv craw running...")

	// 确认下载目录在不在
	_, err := os.Stat(config.WHPath)
	if err != nil {
		log.Println("download dir doesn't exist,recreate...")
		os.Mkdir(config.WHPath, 0777)
	}

	// 获取不能爬取的列表

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
	req.AddCookie(&http.Cookie{
		Name:     "_pk_id.1.01b8",
		Value:    "ed751ff6aadd26c5.1663683548.",
		HttpOnly: true,
	})
	req.AddCookie(&http.Cookie{
		Name:     "_pk_ses.1.01b8",
		Value:    "1",
		HttpOnly: true,
	})
	req.AddCookie(&http.Cookie{
		Name:     "remember_web_59ba36addc2b2f9401580f014c7f58ea4e30989d",
		Value:    "eyJpdiI6ImYwanVGZDFoTm93VjJ2VytVR1c3MWc9PSIsInZhbHVlIjoiNndabE9ZNnJxS0wyTmhkRG8xaVgwaHNYREY2Vkhya1RYZVROeWJ6VUVsRDlMXC8rSGZYTVlCU1E3clwvc1FVbkZSYUl5VThmTDJ2OUZlUWV2Y2g2RHplaHN0NnB4M2oyeXdIdFRjdVdUY1I1S1d0OWxcL1l1NjU3dlVqTjlTWmowTWo0QjNZQmxPckNObEVUeHR1S2pEVEFua212dEZVR2dLMEJYZTJrZm1Wb243cVwvazl3bFkrM1FsZXNXdEVjMzdOUiIsIm1hYyI6IjY1NjJkYWNmYWFiNzE1ZDVmNzg3ZjFiM2FiY2I3M2MzYTZlNzA2NWY3ZjQ4MTVkMzM1MGQ1OTNkYTFlOTFkYTMifQ==",
		HttpOnly: true,
	})
	req.AddCookie(&http.Cookie{
		Name:     "XSRF-TOKEN",
		Value:    "eyJpdiI6IlZ1M1ZRRnJcLzNabmlKd3U1OE01WkZRPT0iLCJ2YWx1ZSI6IlBVczl0VGxYU3NOMmlsTlorQ3VjRVc5eitvaGtCMVpoZjdyVFkwYjVNcHA0U3M2dXdmRXl0SU9tR0wxMDRTMGwiLCJtYWMiOiI3ODEzMzUyZDhjZDU1ZGE5NWY5YWY1Yzg4NTViYzAyOTU0OWM0ZGFjMDFhOGQ1N2E4NmNjMzc1YTMzOTFjYjFhIn0=",
		HttpOnly: true,
	})
	req.AddCookie(&http.Cookie{
		Name:     "wallhaven_session",
		Value:    "eyJpdiI6IkJnTnV2RVwvelN5NW5jOEpkV2tiUFB3PT0iLCJ2YWx1ZSI6IlRtb0gweEZrMUFhR0tOR25zckhydUlyZFB4WkFPbHhyYitiNDZRc3gyWWVaWmVoYlhtOVJhS3Z6Mnd6bE51UDQiLCJtYWMiOiI1OTYxNzg2N2UzMGQ3ZjY3YTBlMTIzZjc3MWYwOGI2YWQzOGE3ODk1ZTNlN2RlNDZhNmM5ZGViYzVlYzIwZGQ0In0=",
		HttpOnly: true,
	})
	do, err := client.Do(req)
	if err != nil {
		return err
	}
	all, _ := io.ReadAll(do.Body)
	body := string(all)

	// 转化为dom对象
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
	dom.Find(".preview").Each(
		func(i int, selection *goquery.Selection) {
			path, exists := selection.Attr("href")
			if !exists {
				return
			}
			//id, ok := last.Attr("data-id")
			//if !ok {
			//	return
			//}

			//存在性检查
			//filename := config.WHPath + path[strings.LastIndex(path, "/")+1:]
			//if pkg.FileExists(filename) {
			//	return
			//}

			parse, _ := url.Parse(path)
			split := strings.Split(parse.Path, "/")

			//存在性检查
			filename := config.WHPath + split[len(split)-1] + ".jpg"
			if FileExists(filename) {
				return
			}

			//path1 := "https://" + parse.Host + parse.Path[strings.Index(parse.Path, "/img-master"):]
			path1 := "https://w.wallhaven.cc/full/" + split[len(split)-1][:2] + "/wallhaven-" + split[len(split)-1] + ".jpg"
			req, _ := http.NewRequest("GET", path1, nil)
			req.AddCookie(&http.Cookie{
				Name:     "_pk_id.1.01b8",
				Value:    "ed751ff6aadd26c5.1663683548.",
				HttpOnly: true,
			})
			req.AddCookie(&http.Cookie{
				Name:     "_pk_ses.1.01b8",
				Value:    "1",
				HttpOnly: true,
			})
			req.AddCookie(&http.Cookie{
				Name:     "remember_web_59ba36addc2b2f9401580f014c7f58ea4e30989d",
				Value:    "eyJpdiI6ImYwanVGZDFoTm93VjJ2VytVR1c3MWc9PSIsInZhbHVlIjoiNndabE9ZNnJxS0wyTmhkRG8xaVgwaHNYREY2Vkhya1RYZVROeWJ6VUVsRDlMXC8rSGZYTVlCU1E3clwvc1FVbkZSYUl5VThmTDJ2OUZlUWV2Y2g2RHplaHN0NnB4M2oyeXdIdFRjdVdUY1I1S1d0OWxcL1l1NjU3dlVqTjlTWmowTWo0QjNZQmxPckNObEVUeHR1S2pEVEFua212dEZVR2dLMEJYZTJrZm1Wb243cVwvazl3bFkrM1FsZXNXdEVjMzdOUiIsIm1hYyI6IjY1NjJkYWNmYWFiNzE1ZDVmNzg3ZjFiM2FiY2I3M2MzYTZlNzA2NWY3ZjQ4MTVkMzM1MGQ1OTNkYTFlOTFkYTMifQ==",
				HttpOnly: true,
			})
			req.AddCookie(&http.Cookie{
				Name:     "XSRF-TOKEN",
				Value:    "eyJpdiI6IlZ1M1ZRRnJcLzNabmlKd3U1OE01WkZRPT0iLCJ2YWx1ZSI6IlBVczl0VGxYU3NOMmlsTlorQ3VjRVc5eitvaGtCMVpoZjdyVFkwYjVNcHA0U3M2dXdmRXl0SU9tR0wxMDRTMGwiLCJtYWMiOiI3ODEzMzUyZDhjZDU1ZGE5NWY5YWY1Yzg4NTViYzAyOTU0OWM0ZGFjMDFhOGQ1N2E4NmNjMzc1YTMzOTFjYjFhIn0=",
				HttpOnly: true,
			})
			req.AddCookie(&http.Cookie{
				Name:     "wallhaven_session",
				Value:    "eyJpdiI6IkJnTnV2RVwvelN5NW5jOEpkV2tiUFB3PT0iLCJ2YWx1ZSI6IlRtb0gweEZrMUFhR0tOR25zckhydUlyZFB4WkFPbHhyYitiNDZRc3gyWWVaWmVoYlhtOVJhS3Z6Mnd6bE51UDQiLCJtYWMiOiI1OTYxNzg2N2UzMGQ3ZjY3YTBlMTIzZjc3MWYwOGI2YWQzOGE3ODk1ZTNlN2RlNDZhNmM5ZGViYzVlYzIwZGQ0In0=",
				HttpOnly: true,
			})

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
	log.Println("finish")
	return nil
}
