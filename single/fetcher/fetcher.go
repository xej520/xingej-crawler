package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetch(url string) ([]byte, error)  {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	//获取 网页的源码格式，GBK，utf-8
	e := determineEncoding(resp.Body)

	//将urf-8代码，转换成 GBK格式
	//因为网页是由GBK格式的，而GO语言默认是UTF-8格式，因此需要转换
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	//all, err := ioutil.ReadAll(resp.Body)
	return ioutil.ReadAll(utf8Reader)

}

//自动检测，从网页里获取的源码的 格式 是GBK,还是UTF 之类的
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		//如果读取失败的话，还是可以继续读取的，
		//给定默认的值，UTF8
		log.Printf("Fetcher error : %v", err)
		return unicode.UTF8 //
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}