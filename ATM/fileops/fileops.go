package fileops

import(
	"fmt"
	"os"
	"strconv"
	"errors"
)

func WriteBalanceToFile(fileName string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func ReadBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		// fmt.Println("Error reading balance file:", err)
		return 1000.0, errors.New("failed to read balance file")
	}
	balanceTect := string(data)
	balance, err := strconv.ParseFloat(balanceTect, 64)
	if err != nil {
		return 1000.0, errors.New("failed to find balance in file")
	}
	return balance, nil
}