package rtk_cal

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var inputReader *bufio.Reader
var input string
var err error
var rou float64
var length float64
var weight float64

var AreaerI Areaer

const (
	CSrou  = 7850  //碳钢管的密度是7850kg/m3
	CSrou2 = 15700 //测试用管，碳钢管密度的两倍

)

type Areaer interface {
	Area() float64 //求面积，返回单位平方米
}

type RoundPipe struct { //定义圆管结构体
	diameter  float64 // 管道外径，单位mm
	thickness float64 // 管道壁厚，单位mm
}

type RectPipe struct {
	long      float64 // 管道长边长度，单位：mm
	short     float64 // 管道短边长度，单位mm
	thickness float64 // 管道壁厚，单位mm
}

func (rp *RoundPipe) Area() float64 { //求面积，返回单位平方米
	return (math.Pi / 4) * ((rp.diameter/1000)*(rp.diameter/1000) - ((rp.diameter-rp.thickness*2)/1000)*((rp.diameter-rp.thickness*2)/1000))
}

func (rp *RectPipe) Area() float64 {
	return (rp.long/1000)*(rp.short/1000) - ((rp.long-rp.thickness*2)/1000)*((rp.short-rp.thickness*2)/1000)
}

func GetPipeWeight() {
LABEL1:
	fmt.Println("┌---------【请选择管道材质】---------")
	fmt.Println("│  输入 1 ：碳钢管")
	fmt.Println("│  输入 2 ：碳钢管密度两倍的管子（测试用）")
	fmt.Println("│  输入 3 ：自定义管材密度")
	fmt.Println("│  输入 q ：退回到主界面")
	fmt.Println("└---------------------------------")

	inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err != nil {
		fmt.Println("Error occurd:", err)
	}

	switch input {
	case "1":
		rou = CSrou
	case "2":
		rou = CSrou2
	case "3":
		setrou()
	case "q", "Q":
		return
	default:
		fmt.Printf("【出错】您输入的参数 \"%v\" 不可用，请重新输入。\n", input)
		goto LABEL1
	}

LABEL2:
	fmt.Println("┌----------请选择管道材质----------")
	fmt.Println("│  输入 1 ：圆管")
	fmt.Println("│  输入 2 ：方管")
	fmt.Println("│  输入 q ：退回到主界面")
	fmt.Println("└---------------------------------")

	inputReader = bufio.NewReader(os.Stdin)
	input, err = inputReader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err != nil {
		fmt.Println("Error occurd:", err)
	}

	switch input {
	case "1":
		RoundPipeWeightCal()
	case "2":
		RectPipeWeightCal()
	case "q", "Q":
		return
	default:
		fmt.Printf("【出错】您输入的参数 \"%v\" 不可用，请重新输入。\n", input)
		goto LABEL2
	}
	fmt.Printf("【结果】管道重量是：%.3f kg\n", weight)
}

func RoundPipeWeightCal() {
	rp1 := new(RoundPipe)
	AreaerI = Areaer(rp1)
	for {
		fmt.Println(" - 请依次输入：管道外径(mm) 壁厚(mm) 管长(m)")
		fmt.Scanln(&rp1.diameter, &rp1.thickness, &length)
		if rp1.thickness*2 > rp1.diameter {
			fmt.Println("【出错】管径或壁厚输入错误，请重新输入w")
		} else {
			weight = AreaerI.Area() * rou * length
			return
		}
	}

}

func RectPipeWeightCal() {
	rp2 := new(RectPipe)
	AreaerI = Areaer(rp2)
	for {
		fmt.Println(" - 请依次输入：长边(mm) 短边(mm) 壁厚(mm) 管长(m)")
		fmt.Scanln(&rp2.long, &rp2.short, &rp2.thickness, &length)
		if rp2.thickness*2 > rp2.long || rp2.thickness*2 > rp2.short {
			fmt.Println("【出错】边长或壁厚输入错误，请重新输入w")
		} else {
			weight = AreaerI.Area() * rou * length
			return
		}
	}
}

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
