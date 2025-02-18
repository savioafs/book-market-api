package utils

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math/rand"

	"strings"
	"time"
)

func NewID() (string, error) {
	id, err := uuid.NewV6()
	if err != nil {
		return "", errors.New("err to create new ID")
	}

	return id.String(), nil
}

func NewCode(prefix string) (string, error) {
	namePrefix := strings.ToUpper(prefix[:3])
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomDigits := rand.Intn(1000)

	formattedDigits := fmt.Sprintf("%03d", randomDigits)

	code := fmt.Sprintf("%s%s", namePrefix, formattedDigits)

	return code, nil
}
