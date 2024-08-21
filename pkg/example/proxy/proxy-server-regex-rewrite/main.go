package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	target := "http://localhost:8080" // 대상 서버 URL (Tumblebug 서버)
	url, err := url.Parse(target)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// 정규 표현식 경로 재작성 설정
	regexRewrite := map[*regexp.Regexp]string{
		regexp.MustCompile(`^/beetle/ns$`):            "/tumblebug/ns",
		regexp.MustCompile(`^/beetle/ns/([^/]+)$`):    "/tumblebug/ns/$1",
		regexp.MustCompile(`^/beetle/ns/([^/]+)/mci`): "",
	}

	e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: url,
			},
		}),
		RegexRewrite: regexRewrite,
		ModifyResponse: func(res *http.Response) error {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}

			log.Printf("Response from target: %s", string(body))

			res.Body = io.NopCloser(bytes.NewReader(body))
			return nil
		},
	}))

	// 무시할 경로에 대한 핸들러 추가
	e.Any("/beetle/ns/:nsId/mci/*", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "Not Found")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
