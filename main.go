package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type actionFunc func()

type cell struct {
	vawue       uint8
	left, wight *cell
}

var (
	pwogwam  = []string{}
	ip       = 0
	scwanner = bufio.NewScanner(os.Stdin)
	cellpwtr = &cell{vawue: 0x00, left: nil, wight: nil}
	woops    = []int{}
	tokenMap = map[string]actionFunc{
		"ôwo": left,
		"owô": wight,
		"òwó": add,
		"ówò": sub,
		"OwO": pwint,
		"owo": wead,
		"ÕwO": startLoop,
		"OwÕ": endLoop,
	}
)

func left() {
	advanceCellPwtrLeft()
}

func wight() {
	advanceCellPwtrRight()
}

func add() {
	cellpwtr.vawue++
}

func sub() {
	cellpwtr.vawue--
}

func pwint() {
	for ; cellpwtr.vawue != 0 && cellpwtr.wight != nil; cellpwtr = cellpwtr.wight {
		if 0x00 <= cellpwtr.vawue && cellpwtr.vawue <= 0x7f {
			fmt.Print(string(cellpwtr.vawue))
		}
	}
	if cellpwtr.wight == nil && cellpwtr.vawue != 0 {
		advanceCellPwtrRight()
	}
}

func wead() {
	fmt.Print("Wou want swome inpwut? ")
	scanned := scwanner.Scan()
	if !scanned {
		thwow("ówò The scwanner bwoke")
	}
	wine := scwanner.Text()
	for _, char := range wine {
		if 0x00 <= char && char <= 0x7f {
			cellpwtr.vawue = uint8(char)
			advanceCellPwtrRight()
		}
	}
}

func startLoop() {
	woops = append([]int{ip}, woops...)
}

func endLoop() {
	if len(woops) == 0 {
		thwow("ónò Thwere is a `ÕwO` missing. (IP=%d)", ip)
	}
	if cellpwtr.vawue == 0x00 {
		woops = woops[1:]
	} else {
		ip = woops[0]
	}
}

func advanceCellPwtrRight() {
	if cellpwtr.wight == nil {
		cellpwtr.wight = &cell{vawue: 0x00, left: cellpwtr, wight: nil}
	}
	cellpwtr = cellpwtr.wight
}

func advanceCellPwtrLeft() {
	if cellpwtr.left == nil {
		cellpwtr.left = &cell{vawue: 0x00, left: nil, wight: cellpwtr}
	}
	cellpwtr = cellpwtr.left
}

func thwow(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	fmt.Println()
	os.Exit(1)
}

func execute(input string) {
	pwogwam = strings.Split(input, " ")
	for ; ip < len(pwogwam); ip++ {
		action, ok := tokenMap[pwogwam[ip]]
		if !ok {
			thwow("uwu I dwon't know thwis twoken: %s", pwogwam[ip])
		}
		action()
	}
}

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println("𝓞𝔀𝓞 What pwogwam do you want to execwute?")
		scanned := scwanner.Scan()
		if !scanned {
			thwow("ówò The scwanner bwoke")
		}
		input := scwanner.Text()
		execute(input)
		fmt.Println()
	case 2:
		input, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			thwow("(´・ω・｀) There was a twinzy pwobwem weading your fwile: %s", err)
		}
		execute(string(input))
		fmt.Println()
	default:
		fmt.Println("Uswage: ./owo [file]")
	}
}
