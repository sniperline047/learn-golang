//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello       = "hello"
	CmdBye         = "bye"
	CmdMandalorian = "mandalorian"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		lineCount    int
		commandCount int
	)

	for scanner.Scan() {
		scannedLine := scanner.Text()
		fmt.Println("You entered:", scanner.Text())
		if strings.ToLower(scannedLine) == "q" {
			break
		} else {
			scannedLine = strings.TrimSpace(scannedLine)

			switch scannedLine {
			case CmdHello:
				commandCount += 1
				fmt.Println("Hello human!")
			case CmdBye:
				commandCount += 1
				fmt.Println("Chao!")
			case CmdMandalorian:
				commandCount += 1
				fmt.Println("This is the way!")
			}

			if scannedLine != "" {
				lineCount += 1
			}
		}
	}

	fmt.Printf("You entered %v commands!\n", commandCount)
	fmt.Printf("You entered %v lines!\n", lineCount)
}
