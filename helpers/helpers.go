package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetString() string {
	scan := bufio.NewReader(os.Stdin)

	value, err := scan.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return strings.TrimSpace(value)
}

func GetInt() int {
	val, err := strconv.Atoi(GetString())

	// Todo: Gracefully handle this error
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return val

}
