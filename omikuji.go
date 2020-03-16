// おみくじ
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	num    int    // 乱数格納用
	result string // おみくじ結果格納用
)

func main() {
	// 擬似乱数を発生させる
	rand.Seed(time.Now().UnixNano())

	// 変数numに0~10の値を代入する
	num = rand.Intn(10)

	switch {
	case num == 0:
		result = "超☆大吉"
	case num >= 1 && num <= 3:
		result = "大吉"
	case num >= 4 && num <= 6:
		result = "中吉"
	case num >= 7 && num <= 9:
		result = "小吉"
	default: // 上記のどれにも当てはまらないケース
		result = "大凶"
	}

	// 結果出力
	fmt.Println("あなたの今日の運勢は、", result, "です！")
}
