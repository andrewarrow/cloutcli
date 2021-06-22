package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFromIn() string {
	fmt.Print("Enter text . on new line to end\n\n")

	buff := []string{}
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "." {
			break
		}
		buff = append(buff, strings.Replace(text, "\"", "\\\"", -1))
	}

	return strings.Join(buff, "\\n")
}
