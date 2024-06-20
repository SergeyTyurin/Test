package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Anagram(strInput []string) map[string][]string {
	mapTemp := make(map[string][]string, len(strInput))
	mapOutput := make(map[string][]string)
	for _, elem := range strInput {
		// сортировка символов в каждом слове, отсортированная последовательность - ключ анаграммы
		newElem := strings.Split(strings.ToLower(elem), "")
		sort.Strings(newElem)
		newStr := strings.Join(newElem, "")
		mapTemp[newStr] = append(mapTemp[newStr], strings.ToLower(elem))
	}
	for _, val := range mapTemp {
		mapElem := make(map[string]struct{})
		sliceElem := make([]string, 0, len(val))

		// добавление слов в map позволяет убрать дублирование слов
		for _, el := range val {
			mapElem[el] = struct{}{}
		}

		// добавление анаграм, без дублей в результирующзий слайс
		for keyMap := range mapElem {
			sliceElem = append(sliceElem, keyMap)
		}

		// сортировка слайса. Первый элемент слайса - ключ результирующей map анаграм
		if len(sliceElem) > 1 {
			sort.Strings(sliceElem)
			mapOutput[val[0]] = sliceElem
		}

	}
	return mapOutput
}

func main() {
	mapRes := Anagram([]string{"мааал", "пятак", "НоГа", "слиток", "пятка", "тяпка", "листок", "ламаа", "столик", "рука", "нога", "столик"})
	fmt.Println(mapRes)
}
