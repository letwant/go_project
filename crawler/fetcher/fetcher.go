package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"zhenai-crawler/crawler/common/reporter"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

const UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36"

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", UserAgent)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		reporter.ReportError("HttpClient请求出错", err)
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
