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
	value       uint8
	left, right *cell
}

var (
	program  = []string{}
	ip       = 0
	scanner  = bufio.NewScanner(os.Stdin)
	cellptr  = &cell{value: 0x00, left: nil, right: nil}
	loops    = []int{}
	tokenMap = map[string]actionFunc{
		"ôwo": left,
		"owô": right,
		"òwó": add,
		"ówò": sub,
		"OwO": print,
		"owo": read,
		"ÕwO": startLoop,
		"OwÕ": endLoop,
	}
)

func left() {
	advanceCellPtrLeft()
}

func right() {
	advanceCellPtrRight()
}

func add() {
	cellptr.value++
}

func sub() {
	cellptr.value--
}

func print() {
	for ; cellptr.value != 0 && cellptr.right != nil; cellptr = cellptr.right {
		if 0x00 <= cellptr.value && cellptr.value <= 0x7f {
			fmt.Print(string(cellptr.value))
		}
	}
	if cellptr.right == nil && cellptr.value != 0 {
		advanceCellPtrRight()
	}
}

func read() {
	fmt.Print("Wou want swome inpwut? ")
	scanned := scanner.Scan()
	if !scanned {
		throw("ówò The scwanner bwoke")
	}
	line := scanner.Text()
	for _, char := range line {
		if 0x00 <= char && char <= 0x7f {
			cellptr.value = uint8(char)
			advanceCellPtrRight()
		}
	}
}

func startLoop() {
	loops = append([]int{ip}, loops...)
}

func endLoop() {
	if len(loops) == 0 {
		throw("ónò Thwere is a `ÕwO` missing. (IP=%d)", ip)
	}
	if cellptr.value == 0x00 {
		loops = loops[1:]
	} else {
		ip = loops[0]
	}
}

func advanceCellPtrRight() {
	if cellptr.right == nil {
		cellptr.right = &cell{value: 0x00, left: cellptr, right: nil}
	}
	cellptr = cellptr.right
}

func advanceCellPtrLeft() {
	if cellptr.left == nil {
		cellptr.left = &cell{value: 0x00, left: nil, right: cellptr}
	}
	cellptr = cellptr.left
}

func throw(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	fmt.Println()
	os.Exit(1)
}

func execute(input string) {
	program = strings.Split(input, " ")
	for ; ip < len(program); ip++ {
		action, ok := tokenMap[program[ip]]
		if !ok {
			throw("uwu I dwon't know thwis twoken: %s", program[ip])
		}
		action()
	}
}

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println("𝓞𝔀𝓞 What pwogwam do you want to execwute?")
		scanned := scanner.Scan()
		if !scanned {
			throw("ówò The scwanner bwoke")
		}
		input := scanner.Text()
		execute(input)
		fmt.Println()
	case 2:
		input, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			throw("(´・ω・｀) There was a twinzy pwobwem weading your fwile: %s", err)
		}
		execute(string(input))
		fmt.Println()
	default:
		fmt.Println("Uswage: ./owo [file]")
	}
}
