package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Row []string

var clmnSort int
var isReverse bool
var isNotDouble bool
var isNumeric bool

type Tab struct {
	rows []Row
}

func (t Tab) Len() int {
	return len(t.rows)
}

func (t Tab) Swap(i int, j int) {
	t.rows[i], t.rows[j] = t.rows[j], t.rows[i]
}

func (t Tab) Less(i int, j int) bool {
	// сортировка по всем колонкам строки минимальной длины, если номер колонки больше мин. длины строки
	if len(t.rows[i]) <= clmnSort || len(t.rows[j]) <= clmnSort {
		for l := 0; l < min(len(t.rows[i]), len(t.rows[j])); l++ {
			if t.rows[i][l] != t.rows[j][l] {
				return check(t.rows[i][l], t.rows[j][l])
			}
		}
		return len(t.rows[i]) < len(t.rows[j])
	} else {
		// иначе: сортировка по колонке
		if t.rows[i][clmnSort] != t.rows[j][clmnSort] {
			return check(t.rows[i][clmnSort], t.rows[j][clmnSort])
		} else {
			// если элементы в колонке равны, то сравниваем строки от начала, кроме этой колонки
			for l := 0; l < min(len(t.rows[i]), len(t.rows[j])); l++ {
				if l != clmnSort {
					if t.rows[i][l] != t.rows[j][l] {
						return check(t.rows[i][l], t.rows[j][l])
					}

				}
			}
		}
		return len(t.rows[i]) < len(t.rows[j])
	}
}

func check(el1 string, el2 string) bool {
	// лексиграфическое сравнение для строк, если не задан флаг сортировки по числам
	if !isNumeric {
		return el1 < el2
	} else {
		// иначе парсим и сравниваем два числа.
		// по аналогии с утилитой ос. Если первый элемент не число - true, если второй элемент не число - false
		// если оба элемента не числа или оба элемента числа, то результат = первый элемент< второй элемент
		x, errX := strconv.ParseInt(el1, 10, 64)
		y, errY := strconv.ParseInt(el2, 10, 64)

		if errX != nil && errY != nil {
			return el1 < el2
		} else if errX != nil {
			return true
		} else if errY != nil {
			return false
		}

		return x < y
	}
}

func NewTab(f io.Reader) Tab {
	tab := Tab{}                   // структура, реализующая метод sort.Interface, содержащая строки
	m := make(map[string]struct{}) // структура для исключения дублирования строк
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		_, found := m[sc.Text()]
		if !(isNotDouble && found) {
			tab.rows = append(tab.rows, strings.Split(sc.Text(), " "))
			m[sc.Text()] = struct{}{}
		}
	}
	return tab
}

func SortTab(tab Tab) {
	if isReverse {
		sort.Sort(sort.Reverse(tab))
	} else {
		sort.Sort(tab)
	}
}

func main() {
	var filename string
	// определение и парсинг аргументов командной строки
	flag.IntVar(&clmnSort, "k", 0, "number of column for sort")
	flag.BoolVar(&isReverse, "r", false, "is reverse sort")
	flag.BoolVar(&isNotDouble, "u", false, "sort with unique rows")
	flag.BoolVar(&isNumeric, "n", false, "compare according to string numerical value")
	flag.StringVar(&filename, "f", "test.txt", "filename for read")
	flag.Parse()

	// открытие файла с текстом для сортировки
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tab := NewTab(f) // структура, реализующая метод sort.Interface, содержащая строки
	SortTab(tab)     // сортировка в обратной или прямой последовательности, в зависимости от флага

	// создание файла с суффиксом _out для записи результата
	filenameOutArr := strings.Split(filename, ".")
	filenameOut := strings.Join([]string{filenameOutArr[0], "_out.", filenameOutArr[1]}, "")

	f, err = os.Create(filenameOut)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// запись результата работы программы
	for i, line := range tab.rows {
		if i > 0 {
			_, err := f.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err := f.WriteString(strings.Join(line[:], " "))
		if err != nil {
			log.Fatal(err)
		}
	}
}
