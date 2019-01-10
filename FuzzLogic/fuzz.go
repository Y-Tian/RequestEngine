package FuzzLogic

import (
	"RequestEngine/RequestLogic"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func MakeFuzz(filename string) {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intReplacer(scanner.Text())
	}

	Check(scanner.Err())
}

func intFuzz(s string) {
	_ = RequestLogic.MakeRequest(s)
}

func intReplacer(s string) {
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	minMaxInts := []int{MaxInt, MinInt}

	const target = "[int]"
	for i := 0; i < strings.Count(s, target); i++ {
		for j := 0; j < len(minMaxInts); j++ {
			replacedString := strings.Replace(s, target, string(minMaxInts[j]), i+1)
			fmt.Printf(replacedString)
			intFuzz(replacedString)
		}
	}
}

func stringFuzz() {
	return
}
