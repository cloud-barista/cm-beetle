package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	target := "http://localhost:8080" // target server url
	url, err := url.Parse(target)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: url,
			},
		}),
		Skipper: func(c echo.Context) bool {
			// Skip url patterns that start with /beetle/ns/mcis
			path := c.Request().URL.Path
			if strings.HasPrefix(path, "/beetle/ns/") {
				parts := strings.Split(path, "/")
				if len(parts) > 3 && parts[3] == "mcis" {
					return true
				}
			}
			return false
		},
		Rewrite: map[string]string{
			"/beetle/ns":   "/tumblebug/ns",
			"/beetle/ns/*": "/tumblebug/ns/$1",
		},
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

	e.Logger.Fatal(e.Start(":1323"))
}
