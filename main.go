package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"./rtk_cal"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	for {
		fmt.Println("┌---------------------------------")
		fmt.Println("│  输入 1 开始斐波那契数列计算")
		fmt.Println("│  输入 q 退出")
		fmt.Println("└---------------------------------")

		inputReader = bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err != nil {
			fmt.Println("Error occurd:", err)
		}

		switch input {
		case "1":
			fmt.Println(" - 启动进行斐波那契数列计算...")
			fib()
		case "q", "Q":
			os.Exit(0)
		default:
			fmt.Printf("您输入的参数 \"%v\" 不可用，请重新输入。\n", input)
		}
	}
}

func fib() {
	fmt.Println(" - 请输入需要计算第几项？")
	inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		fmt.Println("Error occurd:", err)
	}

	timeout := make(chan bool, 1)
	ch := make(chan uint64, 1)
	if num, err := strconv.Atoi(input); err != nil {
		fmt.Printf("【出错】你输入的好像不是一个合法数字。\n")
		fib()
		return
	} else {
		go func() {
			time.Sleep(1e9 * 3)
			timeout <- true
		}()
		go rtk_cal.Getfibonacci(num, ch)
		select {
		case res := <-ch:
			fmt.Printf("斐波那契数列的第 %v 项的值是：%v\n", num, res)
		case <-timeout:
			fmt.Println("【请注意】！！！！啊哦！计算超时了！！！！")
		}

	}

}
