package excel

import (
	"errors"
	"github.com/tealeg/xlsx/v3"
	"strconv"
	"strings"
	"testing"
	"time"
)
import "fmt"

func TestExcel(t *testing.T) {
	//44168   2020/12/3
	var timeNum string = "44168"
	excelHandle(timeNum)
}

func excelHandle(timeNum string) {
	result := make(map[string]map[string]string, 16)
	wb, err := xlsx.OpenFile("D:\\Study\\Go\\Golang-Book-Rank\\src\\excel\\2020-12-4全国新冠疫情数据统计表.xlsx")
	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	fmt.Println("Sheets in this file:")
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name)
	}
	sh, ok := wb.Sheet["全国累计、新增、疑似、治愈"]
	if !ok {
		panic(errors.New("Sheet not found"))
	}
	typeDict := make(map[string]int)
	for j := 0; j < sh.MaxRow; j++ {

		cell, _ := sh.Cell(j, 0)
		//fmt.Println(cell.Value)
		if len(cell.Value) < 20 && strings.Contains(cell.Value, "病例") || strings.Contains(cell.Value, "治愈") {
			typeDict[cell.Value] = j
		}
	}
	typeHandle(typeDict, sh, result, timeNum)

	deadSh, ok := wb.Sheet["死亡人数统计"]
	if !ok {
		panic(errors.New("Sheet not found"))
	}
	deadHeadIndex := 0
	for i := 0; i < deadSh.MaxRow; i++ {
		cell, _ := deadSh.Cell(i, 0)
		if strings.Contains(cell.Value, "累计死亡人数") {
			deadHeadIndex = i
			break
		}
	}
	deadHeadRow, _ := deadSh.Row(deadHeadIndex)
	var deadList []string
	deadMap := make(map[string]int)
	for i := 1; i < deadSh.MaxCol; i++ {
		cell := deadHeadRow.GetCell(i)
		deadList = append(deadList, cell.Value)
		deadMap[cell.Value] = i
	}

	var deadCase = func(start int, sh *xlsx.Sheet) map[string]string {
		result := make(map[string]string)
		for _, v := range deadList {
			cell, _ := deadSh.Cell(start, deadMap[v])
			result[v] = cell.Value
		}
		return result
	}
	deadName := "累计死亡人数"
	for i := 0; i < deadSh.MaxRow; i++ {
		cell, _ := deadSh.Cell(i, 0)
		if cell.Value == timeNum {
			result[deadName] = deadCase(i, deadSh)
			deadName = "新增死亡人数"
		}
	}

	maybeSh, ok := wb.Sheet["累计接触者"]
	var maybeList []string
	maybeMap := make(map[string]int)
	maybeIndex := 0
	for i := 0; i < maybeSh.MaxRow; i++ {
		cell, _ := maybeSh.Cell(i, 0)
		if strings.Contains(cell.Value, "各省累计接触者") {
			maybeIndex = i
			break
		}
	}

	for i := 1; i < maybeSh.MaxCol; i++ {
		cell, _ := maybeSh.Cell(maybeIndex, i)
		if cell.Value != "" && !strings.Contains(cell.Value, "贵州境外输入") && !strings.Contains(cell.Value, "广东") && !strings.Contains(cell.Value, "广西境外输入") && !strings.Contains(cell.Value, "吉林境外输入") {
			maybeList = append(maybeList, cell.Value)
		}
	}
	temp := 0
	for i := 1; i < maybeSh.MaxCol; i++ {
		cell, _ := maybeSh.Cell(maybeIndex+1, i)
		if cell.Value == "累计" || cell.Value == "新增接受" {
			maybeMap[maybeList[temp]] = i
			temp++
		}

	}
	var maybeCase = func(start int, sh *xlsx.Sheet) map[string]string {
		result := make(map[string]string)
		for _, v := range maybeList {
			cell, _ := maybeSh.Cell(start, maybeMap[v])
			result[v] = cell.Value
		}
		return result
	}
	for i := 0; i < deadSh.MaxRow; i++ {
		cell, _ := maybeSh.Cell(i, 0)
		if cell.Value == timeNum {
			result["累计接触"] = maybeCase(i, deadSh)
		}
	}

	for k, v := range result {
		fmt.Println(k, v)
	}

	writeSheet(result, timeNum)
}

var cityList = []string{"湖北",
	"广东",
	"北京",
	"上海",
	"浙江",
	"天津",
	"台湾",
	"河南",
	"重庆",
	"四川",
	"山东",
	"云南",
	"湖南",
	"澳门",
	"江西",
	"辽宁",
	"海南",
	"安徽",
	"福建",
	"贵州",
	"山西",
	"宁夏",
	"广西",
	"河北",
	"黑龙江",
	"香港",
	"江苏",
	"吉林",
	"内蒙古",
	"陕西",
	"新疆",
	"甘肃",
	"青海",
	"西藏",
	"全国"}
var headerList = strings.Split("主键id,省,市,县/区,统计时间,新增病例,累积病例,在治重症病例,危重症病例,治愈病例,新增死亡病例,累积死亡病例,备注,国家,疑似病例,新增疑似病例,新增治愈人数,累计接触人数,新增接触人数", ",")
var headT = strings.Split("id,province,city,area,statistical_time,new_num,total_num,treatment_num,critical_num,cure_num,new_dead_num,dead_num,remark,country,suspect_num,new_suspect_num,new_cure_num,contact_num,new_contact_num", ",")
var writeMap = map[string]int{"新增病例": 5, "累计病例": 6, "累计治愈人数": 9, "新增死亡人数": 10, "累计死亡人数": 11, "现有疑似病例": 14, "新增疑似病例": 15, "新增治愈人数": 16, "累计接触": 17}

func writeSheet(result map[string]map[string]string, timeNum string) {
	wb := xlsx.NewFile()
	sh, _ := wb.AddSheet(timeFormat)
	for i := 0; i < 40; i++ {
		row := sh.AddRow()
		for _, _ = range headerList {
			row.AddCell()
		}
	}

	for i, v := range headerList {
		cell, _ := sh.Cell(0, i)
		cell.SetString(v)
	}
	for i, v := range headT {
		cell, _ := sh.Cell(1, i)
		cell.SetString(v)
	}

	for i, v := range cityList {
		cell, _ := sh.Cell(2+i, 1)
		cell.SetString(v)
		cell, _ = sh.Cell(2+i, 2)
		cell.SetString(v)
		cell, _ = sh.Cell(2+i, 4)
		cell.SetString(timeFormat2)
		cell, _ = sh.Cell(2+i, 13)
		cell.SetString("中国")
	}

	for i, v := range cityList {
		for k, value := range writeMap {
			temp := v
			if temp == "全国" && k == "新增病例" {
				temp = "全国网报"
			}
			if temp == "全国" && (k == "累计病例" || k == "新增治愈人数" || k == "累计治愈人数") {
				temp = "网报全国量"
			}
			if temp == "全国" && (k == "新增死亡人数" || k == "累计死亡人数") {
				temp = "网报全国"
			}
			if temp == "青海" {
				temp = "青海省"
			}
			cell, _ := sh.Cell(2+i, value)
			cell.SetString(result[k][temp])
		}
	}

	//sh = wb.Sheet[timeNum]
	//sh.ForEachRow(func(r *xlsx.Row) error {
	//	r.ForEachCell(func(c *xlsx.Cell) error {
	//		fmt.Print(c.Value)
	//		return nil
	//	})
	//	return nil
	//})

	err := wb.Save("C:\\Users\\mayantao\\Desktop\\test1.xlsx")
	if err != nil {
		fmt.Println("err:", err)
	}
}

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Cell value:", value)
	}
	return err
}

func rowVisitor(r *xlsx.Row) error {
	return r.ForEachCell(cellVisitor)
}

func typeHandle(typeDict map[string]int, sh *xlsx.Sheet, result map[string]map[string]string, timeNum string) {
	for k, v := range typeDict {
		caseResult := CumulativeCases(v, sh, timeNum)
		result[k] = caseResult
	}
}

var timeFormat string
var timeTime time.Time
var timeFormat2 string

func CumulativeCases(start int, sh *xlsx.Sheet, timeNum string) map[string]string {
	cityWidth := make(map[string]int)
	cityIndex := make(map[string]int)
	var cities []string = []string{}
	row, _ := sh.Row(start)
	index := 0
	var cityHandle = func(c *xlsx.Cell) error {
		cityWidth[c.Value] = c.HMerge
		cities = append(cities, c.Value)
		cityIndex[c.Value] = index
		index++
		return nil
	}
	row.ForEachCell(cityHandle)
	cities = cities[2:]
	timeIndex := 0
	for i := start; i < sh.MaxRow; i++ {
		cell, _ := sh.Cell(i, 0)
		if cell.Value == timeNum {
			timeIndex = i
			timeTime, _ = cell.GetTime(false)

			timeFormat = timeTime.Format("01月02日")
			timeFormat2 = timeTime.Format("2006/01/02")
			break
		}
	}

	row, _ = sh.Row(timeIndex)
	result := make(map[string]string)
	for k, v := range cityIndex {
		temp := 0
		for j := 0; j <= cityWidth[k]; j++ {
			cell, _ := sh.Cell(timeIndex, v+j)
			i, _ := strconv.Atoi(cell.Value)
			temp = temp + i
		}
		result[k] = strconv.Itoa(temp)
	}
	return result
}
