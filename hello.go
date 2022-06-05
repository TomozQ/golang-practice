package main

import ("fmt")

func main() {
	// 変数の宣言と代入を同時に行っている。　型の指定は不要
	a,b,c := 100, 200, 300
	fmt.Print("total:")
	fmt.Println(a+b+c)
}