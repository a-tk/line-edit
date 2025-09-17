package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/a-tk/go-datastructures/gap_buffer"
	"golang.org/x/term"
	"io"
	"os"
	"strconv"
)

func run(gb *gap_buffer.GapBuffer) {

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {
			panic(err)
		}
	}(int(os.Stdin.Fd()), oldState)

	termIn := term.NewTerminal(os.Stdin, "")

	display(gb.String(), gb.Cursor())

	buf := make([]byte, 1)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}
		b := buf[0]

		switch b {
		case 3:
			display("^C\n", 0)
			return

		case 4: //
			display("^D\n", 0)
			return

		case 8, 127: // delete/backspace
			gb.DeleteRune()

		case 10: // newline
			continue

		case ':':
			// switch this up, check next char. If no numbers, also accept :: as a single colon
			cmd, _ := termIn.ReadLine()
			// at this point it should be a number, or : or any other commands to implement
			if len(cmd) == 0 {
				display(fmt.Sprintf("Invalid metacommand: %q\n", cmd), 0)
			} else if cmd[0] == ':' {
				if !gb.InsertRune(rune(':')) {
					display("Buffer full\n", 0)
				}
			} else if cmd[0] == 'q' {
				return
			} else if cmd[0] == 'b' {
				gb.MoveCursor(gb.Begin())
			} else if cmd[0] == 'e' {
				gb.MoveCursor(gb.End())
			} else {
				n, err := strconv.Atoi(cmd)
				if err == nil {
					if !gb.MoveCursor(n) {
						display(fmt.Sprintf("Invalid cursor move to %d\n", n), 0)
					}
				} else {
					display(fmt.Sprintf("Invalid metacommand: %q\n", cmd), 0)
				}
			}
		default:
			if b >= 32 && b <= 126 {
				if !gb.InsertRune(rune(b)) {
					display("Buffer full\n", 0)
				}
			}
		}

		display(gb.String(), gb.Cursor())
	}
}

func display(s string, cursor int) {
	fmt.Print("\r\033[2K")
	fmt.Print(s)
	fmt.Printf("\r\033[%dC", cursor)
}

func main() {

	debug := flag.Bool("d", false, "debug mode: view the gap")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, err := reader.ReadString('\n')

	for err != io.EOF {
		text = text[:len(text)-1]

		gb := gap_buffer.New(text, *debug)

		run(gb)

		display(gb.String(), gb.Cursor())
		fmt.Println()

		fmt.Print("> ")
		text, err = reader.ReadString('\n')

	}
}
