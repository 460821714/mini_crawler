// @Time : 2018/5/23 20:07
// @Author : minigeek
package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"bufio"

	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// limit fetch speed.
var rateLimiter = time.Tick(10 * time.Millisecond)

// Fetch returns contents that get response from specially url.
func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("url is %s ,wrong status code is %d", url, resp.StatusCode)
	}
	newReader := bufio.NewReader(resp.Body)
	return ioutil.ReadAll(transform.NewReader(newReader, determineEncoding(newReader).NewDecoder()))
}

// determineEncoding returns encodind from r with automatic det.
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		//log.Printf("fetcher err:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
