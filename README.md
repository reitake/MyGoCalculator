# MyGoCalculator

- 2019.03.25：
  * 项目建立
  * 写了斐波那契计算功能：
    * 建立一个缓存map用来存每次要求计算的斐波那契数列的结果（但不是计算过程中的每一项），比如这次计算了fib(99)，下次在要求求f(99)时，直接出存在map里的答案。
    * 用 `time` 和 chan 设定求值5秒的 timeout。
    * 待优化：让map存下计算数列时每一项的值，比如第一次计算了f(99)，那下次求99项之前的项都能立刻调出上次计算的结果（虽然现在感觉已经比较快了）。另外，求99项之后的可以从99项开始算？
  * 写了个算管重的功能：
    * 用接口写了求圆管、方管的面积
    * 待优化：在给用户自定义密度时，如何更好地判断用户输入的值是数字且大于0？
    * 目前实现方法如下：（不能用fmt.Scan，不然用户输入非期望的内容时，可能出错）
```go
func setrou() {
	for {
		fmt.Println(" - 请输入自定义的管材密度(kg/m3)：")
		input := bufio.NewReader(os.Stdin)
		_, err := fmt.Fscan(input, &rou)
		if err != nil {
			fmt.Println("【出错】只能输入数字哦~！")
		} else {
			if rou <= 0 {
				fmt.Println("【出错】密度要大于0哦w")
			} else {
				fmt.Printf(" - 已设定管道密度 = %v\n", rou)
				return
			}
		}
	}
}
```
