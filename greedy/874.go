package main

import (
// "fmt"
)

// 874. 模拟行走机器人
// 机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：
//     -2：向左转 90 度
//     -1：向右转 90 度
//     1 <= x <= 9：向前移动 x 个单位长度
// 在网格上有一些格子被视为障碍物。
// 第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])
// 机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。
// 返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。
// 提示：
//     0 <= commands.length <= 10000
//     0 <= obstacles.length <= 10000
//     -30000 <= obstacle[i][0] <= 30000
//     -30000 <= obstacle[i][1] <= 30000
//     答案保证小于 2 ^ 31
// https://leetcode-cn.com/problems/walking-robot-simulation/
func main() {
	// fmt.Println(robotSim([]int{4, -1, 3}, nil)) //25
	// fmt.Println(robotSim([]int{4, -1, 4, -2, 4},
	// [][]int{
	// 	{2, 4},
	// })) //65
}

// func robotSim(commands []int, obstacles [][]int) (area int) {
// 	n := len(commands)
// 	if n == 0 {
// 		return 0
// 	}
// 	xy := [2]int{0, 0}
// 	obstaclesX := make(map[int][]int) // 以x轴为key统计障碍物
// 	obstaclesY := make(map[int][]int) // 以y轴为key统计障碍物
// 	for _, v := range obstacles {
// 		obstaclesX[v[0]] = append(obstaclesX[v[0]], v[1])
// 		obstaclesY[v[1]] = append(obstaclesY[v[1]], v[0])
// 	}

// 	towards := 1 // 1北 2东 3南 4西
// 	for i := 0; i < n; i++ {
// 		if commands[i] < 0 {
// 			towards = turn(towards, commands[i])
// 		} else {
// 			switch towards {
// 			case 1:
// 				d := getDist(obstaclesX[xy[0]], xy[1], commands[i])
// 				xy[1] += d
// 			case 2:
// 				d := getDist(obstaclesY[xy[1]], xy[0], commands[i])
// 				xy[0] += d
// 			case 3:
// 				xy[1] -= d
// 			case 4:
// 				xy[0] -= d
// 			}
// 		}
// 		// 更新面积
// 		area = getMax(area, xy[0]*xy[0]+xy[1]*xy[1])
// 	}
// 	return area
// }
// func getDist(obstacles []int, start, distance int) (dist int) {
// 	if start > end {
// 		start, end = end, start
// 	}
// 	for i := 1; i <= dist; i++ {
// 		for _, o := range obstacles {
// 			if o == start+i {
// 				return i - start - 1
// 			}
// 		}
// 	}
// 	return distance
// }

// // 1北 2东 3南 4西
// func turn(towards, diff int) int {
// 	if diff == -2 {
// 		towards--
// 		if towards == 0 {
// 			towards = 4
// 		}
// 	} else if diff == -1 {
// 		towards++
// 		if towards == 5 {
// 			towards = 1
// 		}
// 	}
// 	return towards
// }

// func getMax(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
