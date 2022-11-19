package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	a, err := unpackString("a4bc2d5e")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(a)
	os.Exit(0)
}

func unpackString(s string) (string, error) {

	var newStr strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); {
		var numStr strings.Builder

		if !unicode.IsDigit(runes[i]) {
			// ! это буква !
			j := i + 1
			for ; j < len(runes); j++ {

				if unicode.IsDigit(runes[j]) {
					// ! это число !
					numStr.WriteRune(runes[j])
				} else {
					break
				}

			}

			count, err := strconv.Atoi(numStr.String())

			if err != nil {
				newStr.WriteRune(runes[i])
			} else {
				for u := 0; u < count; u++ {
					newStr.WriteRune(runes[i])

				}
			}
			i = j

		} else {
			return newStr.String(), errors.New("Incorrect str")
			// ! ошибка не корректная строка
		}

	}

	return newStr.String(), nil
}
