package main // 一个目录下需要在main 包中才能独立运行,下面需要有个 main 函数

import "fmt" // 这是包名, 下面函数中可以调用包内的方法

func main() { // main 函数,每个go程序都需要这个入口函数
	fmt.Println("hello world")
}
