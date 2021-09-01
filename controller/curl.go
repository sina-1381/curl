package controller

import (
	"bytes"
	"curl/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dextools(c *gin.Context) {
	var input validation.Curl
	validation.CheckValidate(&input , c)
	url := "https://www.dextools.io/back/user/login"
	var jsonStr = []byte(`{"id":"`+input.Id+`","password":"`+input.Password+`"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("referer", "https://www.dextools.io/app/bsc")
	req.Header.Set("authority", "www.dextools.io")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("authorization", "Bearer null")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4621.4 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("origin", "https://www.dextools.io/")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("accept-language", "en-US,en;q=0.9,fa-IR;q=0.8,fa;q=0.7")
	req.Header.Set("cookie", "_ga=GA1.2.230045494.1630129507; _gid=GA1.2.230103365.1630306947")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	c.JSON(http.StatusOK , gin.H{
		"x-auth" : resp.Header.Get("X-Auth"),
	})
}
