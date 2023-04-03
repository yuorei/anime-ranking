package application

import (
	"fmt"
	"strconv"
	"strings"
)

func JudgeID(kind string, id string) (int, error) {
	splitID := strings.Split(id, ":")
	if kind != splitID[0] {
		return 0, fmt.Errorf("IDの種類が異なります")
	}

	i, err := strconv.Atoi(splitID[1])
	if err != nil {
		return 0, err
	}

	return i, nil
}
