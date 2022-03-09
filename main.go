package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

type Req struct {
	groupNum int // 几组题目
	itemNum int  // 每组多少题
	rowNum int32  // 几列
	isFilter int32  // 是否去重
	numeral Numeral
}

type Numeral struct {
	Min int32
	Max int32
}

func main() {
	GetData(Req{
		groupNum: 10,
		itemNum: 10,
		rowNum: 2,
		isFilter:1, // 目前只支持去重
		numeral: Numeral{
			Min: 1,
			Max: 10,
		},
	})
}

func GetData(req Req)  {
	for i := 0; i < req.groupNum; i++ {
		createGroup(req)
	}
}

func createGroup(req Req)  {
	itemMap := map[string]string{}

	// 生成一组题目
	for i := 0; i < req.itemNum; i++ {
		// 生成一题
		item := createItem(req)
		itemKey := MD5(item)
		if req.isFilter == 1 {
			if _, ok := itemMap[itemKey];ok {
				//fmt.Printf("error item:%v\n",item)
				i--
				continue
			}
			itemMap[itemKey] = item
		}
	}

	// 排版
	// 打印标题
	printTitle(req.numeral.Max)
	i := 0
	//fmt.Printf("total:%v\n",len(itemMap))
	for _, v := range itemMap {
		if i != 0 && i%int(req.rowNum) == 0 {
			fmt.Println()
		}
		fmt.Print(v)
		i ++
	}
	println(1)
}

func createItem(req Req) string {
	// 运算符
	symbol := getSymbol()
	// 随机数字A
	left := RandInt(req.numeral.Min, req.numeral.Max)
	// 随机运算符
	symbolIndex := RandInt(0, int32(len(symbol)))
	// 随机数字B
	right := RandInt(req.numeral.Min, req.numeral.Max)
	// 当运算符为减法，且A小于B时，调整位置
	if symbol[symbolIndex] == "-" && left < right {
		tmpRight := left
		left = right
		right = tmpRight
	}

	return fmt.Sprintf("%v %v %v = \t", left, symbol[symbolIndex], right)
}

func printTitle(maxNum int32) {
	fmt.Println(fmt.Sprintf("（%v 以内加减法）",maxNum))
	fmt.Println(fmt.Sprintf("打卡时间：%v\t  用时：%v\t正确率：%v","_____","_____","_____"))
}

func getSymbol() map[int32]string {
	symbol := make(map[int32]string, 0)
	symbol[0] = "+"
	symbol[1] = "-"
	symbol[2] = "+"
	symbol[3] = "-"
	symbol[4] = "-"

	return symbol
}

func println(num int)  {
	for i := 0; i < num; i++ {
		fmt.Println()
	}
}
// RandInt ...
func RandInt(min, max int32) int32 {
	/*if min == 0 && max == 0 || min >= max {
		fmt.Errorf("err")
		return max
	}*/
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(max-min) + min
}

func MD5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}