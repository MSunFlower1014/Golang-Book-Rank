package excel

import (
	"bytes"
	"encoding/json"
	"fmt"
	dbUtil "github.com/MSunFlower1014/Golang-Book-Rank/src/db"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	dbUtil.InitDB()
	defer dbUtil.DB.Close()
	for true {
		yearMont := time.Now().Format("200601")
		yearMonthDay := time.Now().Format("20060102")
		hasSave := dbUtil.Count(yearMonthDay)
		if hasSave {
			fmt.Printf("has saved, yearMonth : %s , next day execute\n", yearMont)
			time.Sleep(time.Duration(24) * time.Hour)
			continue
		}
		for i := 1; i <= 5; i++ {
			saveBookRank(strconv.Itoa(i), yearMont)
			time.Sleep(time.Duration(2) * time.Second)
		}

		fmt.Printf("day : %s , execute end,save books end ,next day execute\n", yearMonthDay)
		time.Sleep(time.Duration(24) * time.Hour)
	}
}

func saveBookRank(pageNum, month string) {
	var buffer bytes.Buffer
	buffer.WriteString("https://m.qidian.com/majax/rank/yuepiaolist?_csrfToken=yOYgIBQMyWxfSQIFmFcanGrSC19FscXSY9qzQXKe&gender=male&pageNum=")
	buffer.WriteString(pageNum)
	buffer.WriteString("&catId=-1&yearmonth=")
	buffer.WriteString(month)
	url := buffer.String()
	buffer.Reset()
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
	now := time.Now()
	yearMonthDay := time.Now().Format("20060102")

	for _, v := range records {
		book := v.(map[string]interface{})
		bookStruct := dbUtil.Book{book["bid"].(string), book["bName"].(string),
			book["bAuth"].(string), book["desc"].(string), book["cat"].(string),
			int(book["catId"].(float64)), book["bName"].(string), book["rankCnt"].(string),
			int(book["rankNum"].(float64)), month, yearMonthDay, now}
		if len(bookStruct.Desc) > 1000 {
			bookStruct.Desc = bookStruct.Desc[0:1000]
		}
		dbUtil.InsertBook(&bookStruct)
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
