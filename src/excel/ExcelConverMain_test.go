package excel

import (
	"errors"
	"github.com/tealeg/xlsx/v3"
	"strconv"
	"strings"
	"testing"
)
import "fmt"

func TestExcel(t *testing.T) {
	wb, err := xlsx.OpenFile("C:\\Users\\71013\\Desktop\\2020-12-4全国新冠疫情数据统计表.xlsx")
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
	fmt.Println(typeDict)
	typeHandle(typeDict, sh)
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

func typeHandle(typeDict map[string]int, sh *xlsx.Sheet) {
	result := make(map[string]map[string]int)
	for k, v := range typeDict {
		result[k] = CumulativeCases(v, sh)
	}

	for k, v := range result {
		fmt.Println(k, v)
	}
}

var timeNum string = "44168"

func CumulativeCases(start int, sh *xlsx.Sheet) map[string]int {
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
	fmt.Println(cityWidth)
	fmt.Println(cities)
	cities = cities[2:]
	timeIndex := 0
	for i := start; i < sh.MaxRow; i++ {
		cell, _ := sh.Cell(i, 0)
		if cell.Value == timeNum {
			timeIndex = i
			break
		}
	}

	row, _ = sh.Row(timeIndex)
	result := make(map[string]int)
	for k, v := range cityIndex {
		temp := 0
		for j := 0; j <= cityWidth[k]; j++ {
			cell, _ := sh.Cell(timeIndex, v+j)
			i, _ := strconv.Atoi(cell.Value)
			temp = temp + i
		}
		result[k] = temp
	}
	for k, v := range result {
		fmt.Println(k + ":" + strconv.Itoa(v))
	}
	return result
}
