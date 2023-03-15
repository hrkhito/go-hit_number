package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/* 乱数を作成するときの最初の設定値のことをシード値という
	コンピュータで乱数を生成する場合、ある数値をシード値として用い、
	その値を基にして特定の演算により乱数を生成する。
	このため同じシード値を使用した場合、発生する乱数(結果)はまったく同じものになる。
	*/

	// 乱数を生成するために使うシード値というものがデフォルトでは固定されているので毎回変えるための処理
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(10) + 1
	count := 0

	for {
		var guess int

		fmt.Print("Your guess? ")
		// 入力された値をguess変数に代入する
		fmt.Scanf("%v", &guess)
		count++

		if answer == guess {
			fmt.Printf("Bingo! It took %v guesses!\n", count)
			break
		} else if answer > guess {
			fmt.Println("The answer is higher!")
		} else {
			fmt.Println("The answer is lower!")
		}
	}
}