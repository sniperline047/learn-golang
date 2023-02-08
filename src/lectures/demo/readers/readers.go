package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		values, err := r.ReadString(' ')
		n := strings.TrimSpace(values)
		if n == "" {
			continue
		}
		num, convErr := strconv.Atoi(n)

		if convErr != nil {
			fmt.Println("conversion error", err)
		} else {
			sum += num
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Input error ocurred:", err)
		}
	}

	fmt.Println("Sum is:", sum)
}
