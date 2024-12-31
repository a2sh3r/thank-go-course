package main

// не меняйте импорты, они нужны для проверки
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

var errInsufficientFunds = errors.New("insufficient funds")

var errInvalidAmount = errors.New("invalid transaction amount")

var errInvalidOverdraft = errors.New("expect overdraft >= 0")

var errOverdraftExceeded = errors.New("balance cannot exceed overdraft")

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

func main() {
	var acc account
	var trans []int
	var err error
	if acc, trans, err = parseInput(); err == nil {
		fmt.Printf("-> %v %v", acc, trans)
	} else {
		fmt.Printf("-> %v", err)
	}
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
	accSrc, transSrc := readInput()
	acc, err := parseAccount(accSrc)
	if err != nil {
		return account{}, []int{}, err
	}
	trans, err := parseTransactions(transSrc)
	if err != nil {
		return account{}, []int{}, err
	}
	return acc, trans, nil
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0])
	if err != nil {
		return account{}, err
	}
	overdraft, err := strconv.Atoi(parts[1])
	if err != nil {
		return account{}, err
	}
	if overdraft < 0 {
		return account{}, errInvalidOverdraft
	}
	if balance < -overdraft {
		return account{}, errOverdraftExceeded
	}
	return account{balance, overdraft}, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		if t, err := strconv.Atoi(s); err == nil {
			trans[idx] = t
		} else {
			return []int{}, err
		}
	}
	return trans, nil
}
