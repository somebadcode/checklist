package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/somebadcode/checklist/jsonchecklist"
)

type Checklist interface {
	Check(item string)
	IsChecked(item string) bool
	Uncheck(item string)
}

var (
	ErrWrongNumberOfArguments = errors.New("wrong number of arguments")
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments")
		return
	}

	checklist, err := jsonchecklist.New(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer checklist.Flush()

	command := strings.ToLower(os.Args[2])

	switch command {
	case "check":
		err = check(checklist, os.Args[3:])
		if err != nil {
			fmt.Println(err)
		}
	case "ischecked":
		if err = isChecked(checklist, os.Args[3:]); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("unknown command", command)
	}
}

func check(checklist Checklist, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly 1 argument: %w", ErrWrongNumberOfArguments)
	}

	checklist.Check(args[0])

	return nil
}

func isChecked(checklist Checklist, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly 1 argument: %w", ErrWrongNumberOfArguments)
	}

	fmt.Printf("%s: %v\n", args[0], checklist.IsChecked(args[0]))

	return nil
}
