package printer

import (
	"fmt"

	"github.com/sjlleo/netflix-verify/verify"
)

const (
	AUTHOR        = "@sjlleo"
	VERSION       = "v3.0"
	RED_PREFIX    = "\033[1;31m"
	GREEN_PREFIX  = "\033[1;32m"
	YELLOW_PREFIX = "\033[1;33m"
	BLUE_PREFIX   = "\033[1;34m"
	PURPLE_PREFIX = "\033[1;35m"
	CYAN_PREFIX   = "\033[1;36m"
	RESET_PREFIX  = "\033[0m"
)

func Print(fr verify.FinalResult) {
	printVersion()
	printResult("4", fr.Res[1])
	fmt.Println()
	printResult("6", fr.Res[2])
}

func printVersion() {
	fmt.Println("**NetFlix 解锁检测小工具 " + VERSION + " By " + CYAN_PREFIX + AUTHOR + RESET_PREFIX + "**")
}

func printResult(ipVersion string, vResponse verify.VerifyResponse) {
	fmt.Printf("[IPv%s]\n", ipVersion)
	switch code := vResponse.StatusCode; {
	case code < -1:
		fmt.Println("网络可能没有正常配置IPv" + ipVersion + "，或者没有IPv" + ipVersion + "网络接入")
		fmt.Println(vResponse.CountryName)
	case code == -1:
		fmt.Println("Netflix在出口IP所在的国家不提供服务")
		fmt.Println(vResponse.CountryName)
	case code == 0:
		fmt.Println("Netflix在出口IP所在的国家提供服务，但是您的IP疑似代理，无法正常使用服务")
		fmt.Println(vResponse.CountryName)
	case code == 1:
		fmt.Println("出口IP可以使用Netflix，但仅可看Netflix自制剧")
		fmt.Println(vResponse.CountryName)
	case code == 2:
		fmt.Println("出口IP完整解锁Netflix，支持非自制剧的观看")
		fmt.Println(vResponse.CountryName)
	case code == 3:
		fmt.Println("出口IP无法观看此电影")
		fmt.Println(vResponse.CountryName)
	case code == 4:
		fmt.Println("出口IP可以观看此电影")
		fmt.Println(vResponse.CountryName)
	default:
		fmt.Println("解锁检测失败，请稍后重试")
		fmt.Println(vResponse.CountryName)
	}
}
