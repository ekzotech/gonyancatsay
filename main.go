package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

const nyanCat = `        \
         \
      +       o         +
          _-_-_-_-_-_-_,------,
          _-_-_-_-_-_-|   /\_/\
          _-_-_-_-_-_-~|__( ^ .^)  +
          _-_-_-_-_-_-_ ""  ""
    +          o         +        o`

var rainbowColors = []string{"31", "33", "32", "36", "34", "35"}

func main() {
	rainbow := flag.Bool("rainbow", false, "раскрасить вывод в цвета радуги")
	flag.Parse()

	message := strings.Join(flag.Args(), " ")
	if len(flag.Args()) == 0 {
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "nyancatsay:", err)
			os.Exit(1)
		}
		message = strings.TrimSuffix(string(input), "\n")
		message = strings.TrimSuffix(message, "\r")
	}

	fmt.Println(render(message, *rainbow))
}

func render(message string, rainbow bool) string {
	output := speechBubble(message) + "\n" + nyanCat
	if !rainbow {
		return output
	}

	lines := strings.Split(output, "\n")
	for i, line := range lines {
		lines[i] = "\x1b[" + rainbowColors[i%len(rainbowColors)] + "m" + line + "\x1b[0m"
	}
	return strings.Join(lines, "\n")
}

func speechBubble(message string) string {
	lines := strings.Split(message, "\n")
	width := 0
	for _, line := range lines {
		width = max(width, utf8.RuneCountInString(line))
	}

	result := []string{" " + strings.Repeat("_", width+2)}
	if len(lines) == 1 {
		result = append(result, "< "+lines[0]+" >")
	} else {
		for i, line := range lines {
			padding := strings.Repeat(" ", width-utf8.RuneCountInString(line))
			switch i {
			case 0:
				result = append(result, "/ "+line+padding+" \\")
			case len(lines) - 1:
				result = append(result, "\\ "+line+padding+" /")
			default:
				result = append(result, "| "+line+padding+" |")
			}
		}
	}
	result = append(result, " "+strings.Repeat("-", width+2))
	return strings.Join(result, "\n")
}
