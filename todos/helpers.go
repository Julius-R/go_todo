package todos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getString() string {
	scan := bufio.NewReader(os.Stdin)

	value, err := scan.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return strings.TrimSpace(value)
}

func getInt() int {
	val, err := strconv.Atoi(getString())

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return val

}
