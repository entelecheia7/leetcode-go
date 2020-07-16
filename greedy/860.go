package main

import (
	"fmt"
)

// 860. 柠檬水找零
// 在柠檬水摊上，每一杯柠檬水的售价为 5 美元。
// 顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。
// 每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。
// 注意，一开始你手头没有任何零钱。
// 如果你能给每位顾客正确找零，返回 true ，否则返回 false 。
// 提示：
//     0 <= bills.length <= 10000
//     bills[i] 不是 5 就是 10 或是 20
// https://leetcode-cn.com/problems/lemonade-change/
func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20})) //true
}

// 找零时，总是试图先支出最大的面值
func lemonadeChange(bills []int) bool {
	n := len(bills)
	if n == 0 {
		return true
	} else if bills[0] != 5 {
		return false
	}
	count5, count10 := 1, 0
	for i := 1; i < n; i++ {
		if bills[i] == 5 { // 不找零
			count5++
		} else if bills[i] == 10 { // 找5块
			if count5 == 0 {
				return false
			}
			count10++
			count5--
		} else { // 找15
			if count10 > 0 && count5 > 0 {
				count10--
				count5--
			} else if count5 >= 3 {
				count5 -= 3
			} else {
				return false
			}
		}
	}
	return true
}
