package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	url := "https://m.qidian.com/majax/rank/yuepiaolist?_csrfToken=yOYgIBQMyWxfSQIFmFcanGrSC19FscXSY9qzQXKe&gender=male&pageNum=1&catId=12&yearmonth=202009"
	resp, err := http.Get(url)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	//fmt.Printf("%s", content)
	temp, err := zhToUnicode(content)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	//fmt.Printf("%s", temp)
	var data = make(map[string]interface{})
	if err := json.Unmarshal(temp, &data); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}

	fmt.Printf("code : %v , msg : %s \n", data["code"], data["msg"])

	mainData := data["data"].(map[string]interface{})

	records := mainData["records"].([]interface{})

	for _, v := range records {
		book := v.(map[string]interface{})
		fmt.Printf(" name : %s , rankNum : %v \n", book["bName"], book["rankNum"])
	}

}

/*
将json中的unicode转为汉字
*/
func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
