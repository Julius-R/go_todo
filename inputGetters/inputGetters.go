package inputGetters

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetString() (string, error) {
	scan := bufio.NewReader(os.Stdin)

	value, err := scan.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(value), nil
}

func GetInt() (int, error) {
	strVal, err := GetString()
	if err != nil {
		return 0, err
	}
	intVal, intErr := strconv.Atoi(strVal)
	if intErr != nil {
		return 0, err
	}
	return intVal, nil

}
