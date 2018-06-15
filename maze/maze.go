package main

import (
	"os"
	"fmt"
)

func readMaze(filename string) [][]int {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col, rowend int

	// fmt.Fscanf函数
	// 遇到%d就尽量解析一个数字，
	// 遇到空格就结束而自动跳过，
	// 然后再解析下一个参数，
	// 遇到换行就自动停止

	// windows下换行符是\r\n，会导致每个行末读进一个0。
	// 可在读完一行后，配合fmt.Fscanln使用
	// 可在每行后面加一个Scanln，或者使用bufio里的readstring
	// 此处用fmt.Fscanf换行符赋给其他变量的方法

	// 先读取文件第一行的数据，即用指针取地址来获取行数值（row）和列数值（col）
	fmt.Fscanf(file, "%d %d", &row, &col)
	// 再将第一行末尾的换行符（windows系统特有）剔除
	_, err = fmt.Fscanf(file, "%d", &rowend)

	// 声明一个以[]int为元素的slice，数量值为row
	maze := make([][]int, row)

	for i := range maze {
		// 声明每个slice（即每一行数据），数量值为col
		maze[i] = make([]int, col)
		for j := range maze[i] {
			//获取每行每列的数据
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
		// 剔除每行末尾的换行符
		_, err = fmt.Fscanf(file, "%d", &rowend)
	}
	return maze
}

// 点结构
type point struct {
	i, j int
}

// 探索方向
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

// 定义看到下一节点的坐标的方法
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

//
func (p point) at(grid [][]int) (int, bool) {

	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int,	start, end point) [][]int {
	// 声明一个以[]int为元素的slice，行号值为len(maze)
	// 同readMaze()函数中的切片maze类似
	// 代表从起始节点point{0, 0}走了多少格才到达该节点
	steps := make([][]int, len(maze))
	//
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// 把起点就加入队列
	Q := []point{start}

	// 探索路径
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {    // 发现终点就退出
			break
		}
		//
		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != start

			val, ok := next.at(maze)
			if !ok || val == 1 {     // 撞墙
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {       // 已走过
				continue
			}

			if next == start {
				continue
			}


			curSteps, _ := cur.at(steps)
			// 探索
			steps[next.i][next.j] =
				curSteps + 1
			// 更新队列
			Q = append(Q, next)
		}
	}
	return steps
}


func main() {
	maze := readMaze("basic/maze/maze.in")

	// 打印maze图
	fmt.Println("It's original:")
	//按行
	for _, row := range maze {
		//按列
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}

	// 走迷宫
	// 起始节点point{0, 0}
	// 终止节点point{len(maze) -1, len(maze[0]) - 1})
	steps := walk(maze, point{0, 0},
		point{len(maze) -1, len(maze[0]) - 1})

	// 打印路径
	fmt.Println("It's the road:")
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}
