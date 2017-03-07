package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/spiegel-im-spiegel/pf"
)

func main() {
	flag.Parse()
	argsStr := flag.Args()
	ln := len(argsStr)
	if ln < 2 {
		fmt.Fprintln(os.Stderr, "年月（日）を指定してください")
		return
	}
	args := make([]int, ln)
	for i := 0; i < ln; i++ {
		num, err := strconv.Atoi(argsStr[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		args[i] = num
	}
	if ln == 2 {
		dt := pf.NewYearMonth(args[0], args[1])
		d, err := dt.GetPremiumFriday()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(d)
		return
	}
	dt := pf.NewDate(args[0], args[1], args[2])
	fmt.Println(dt.IsPremiumFriday())
}
