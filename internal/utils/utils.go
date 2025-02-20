package utils

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"

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

func NewCode(prefix string, sizeCode int) (string, error) {

	namePrefix := strings.ToUpper(prefix[:sizeCode/2])
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomDigits := rand.Intn(1000)

	sizeCodeStr := strconv.Itoa(sizeCode / 2)

	formattedDigits := fmt.Sprintf("%0"+sizeCodeStr+"d", randomDigits)

	code := fmt.Sprintf("%s%s", namePrefix, formattedDigits)

	return code, nil
}

func Retry(attempts int, sleep time.Duration, f func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		for i > 0 {
			fmt.Printf("retrying after error: %s", err)
			time.Sleep(sleep)
		}
	}

	err = f()
	if err != nil {
		return nil
	}

	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
