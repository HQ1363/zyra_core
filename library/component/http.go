package component

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**
* http uniform deal with
* https://www.bookstack.cn/read/go42/content-42_36_http.md
**/
func InvokeHttpRequest(method, reqAddr string, params map[string]string, headers map[string]string, data map[string]string, cookies []*http.Cookie) string {
	client := &http.Client{Timeout: 30 *time.Second}
	u, _ := url.Parse(reqAddr)
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	var payload = make(url.Values, 0)
	if data != nil {
		for key, value := range data {
			payload.Set(key, value)
		}
	}
	req, _ := http.NewRequest(strings.ToUpper(method), u.String(), strings.NewReader(payload.Encode()))
	if cookies != nil && len(cookies) > 0 {
		for i := 0; i < len(cookies); i++ {
			req.AddCookie(cookies[i])
		}
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return ""
	} else {
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("http parse data error: ", err.Error())
		}
		fmt.Printf("url: %s, status code: %d, result: %s\n", resp.Request.URL, resp.StatusCode, result)
		return string(result)
	}
}
