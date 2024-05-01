package main

import (
	"bytes"
	"fmt"
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

func unpack(str string) (string, error) {
	var buf bytes.Buffer
	stringRune := []rune(str)

	for i := 0; i < len(stringRune); i++ {

		if stringRune[i] == '\\' {
			i++
			if i >= len(stringRune) {
				break
			}
			buf.WriteRune(stringRune[i])
			continue
		}

		if unicode.IsDigit(stringRune[i]) {
			if i == 0 {
				return "", fmt.Errorf("некорректная строка")
			}

			count := int(stringRune[i] - '0')

			for k := i; k < len(stringRune)-1 && unicode.IsDigit(stringRune[k+1]); k++ {
				count = count*10 + int(stringRune[i]-'0')
			}

			for j := 1; j < count; j++ {
				buf.WriteRune(stringRune[i-1])
			}

		} else {
			buf.WriteRune(stringRune[i])
		}
	}
	return buf.String(), nil
}

func main() {
	res, err := unpack("a5b2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
