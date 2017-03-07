package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spiegel-im-spiegel/pf"
)

func main() {
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 2 {
		fmt.Fprintln(os.Stderr, "年月を指定してください")
		return
	}
	args := make([]int, 2)
	for i := 0; i < 2; i++ {
		num, err := strconv.Atoi(argsStr[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		args[i] = num
	}
	d, err := pf.GetPremiumFriday(args[0], time.Month(args[1]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(d)
}
