package main

import (
	"fmt"
	"time"
	"github.com/inancgumus/screen"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {
	screen.MoveTopLeft()
	now := time.Now()
	hour, min, sec := now.Hour(), now.Minute(), now.Second()

	fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

	// [8][5]string
	clock := [...]placeholder{
		// extract the digits: 17 becomes, 1 and 7 respectively
		digits[hour/10], digits[hour%10],
		colon,
		digits[min/10], digits[min%10],
		colon,
		digits[sec/10], digits[sec%10],
	}

	for line := range clock[0] {
		for digit := range clock {
			next := clock[digit][line]
			fmt.Print(next, "  ")
		}
		fmt.Println()
	}
	time.Sleep(time.Second)
  }
}