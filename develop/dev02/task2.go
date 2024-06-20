package task2

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

import (
	"fmt"
	"unicode"
)

func addElem(repeatCnt int, elemToAdd rune, res []rune) []rune {
	if repeatCnt >= 0 {
		for l := 0; l < repeatCnt; l++ {
			res = append(res, elemToAdd)
		}
	} else {
		res = append(res, elemToAdd)
	}
	return res
}

func Unpack(str string) (string, error) {
	res := []rune{}
	repeatCnt := -1   // число повторений
	isEscape := false // флаг экранированного элемента
	escape := rune('\u005c')
	var elemToAdd rune // элемент, который нужно добавить

	if len(str) == 0 {
		return "", nil
	}

	for i, elem := range str {
		// проверка, что первый элемент не цифра
		if unicode.IsDigit(elem) && i == 0 {
			return "", fmt.Errorf("invalid string: %s", str)
		}

		// если текущий элемент экранирован, то добавляем в результат
		if isEscape {
			if elemToAdd != 0 {
				res = addElem(repeatCnt, elemToAdd, res)
			}
			elemToAdd = elem
			isEscape = false
			continue
		}

		// если элемент = "\", то следующий элемент будет экранирован
		if elem == escape {
			isEscape = true
			continue
		}

		// в остальных случаях, если элемент - цифра, то составляем число
		if unicode.IsDigit(elem) {
			if repeatCnt == -1 {
				repeatCnt = int(elem - '0')
			} else {
				repeatCnt = repeatCnt*10 + int(elem-'0')
			}
			continue
		}

		// если есть элемент для записи, то добавляем с учетом числа повторений
		if elemToAdd != 0 {
			res = addElem(repeatCnt, elemToAdd, res)
		}
		repeatCnt = -1
		elemToAdd = elem

	}

	// если последний элемент = "\", то это ошибка
	if isEscape {
		return "", fmt.Errorf("invalid string: %s", str)
	}

	// добавления элемента с последней итерации
	res = addElem(repeatCnt, elemToAdd, res)

	return string(res), nil
}
