package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"regexp"
)

func main() {


	printCityList(all)
}

func printCityList(conttents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(conttents, -1)
	for _, m := range matches{
		fmt.Printf("City: %s, URL:%s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}


