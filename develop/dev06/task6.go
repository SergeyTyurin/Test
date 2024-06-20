package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	Separator bool
	Column    []string
}

type Tab []Row

var fields []int
var delimiter string
var separated bool

func cut(tab Tab) Tab {
	var tabNew Tab
	for _, elTab := range tab {
		if !separated && !elTab.Separator {
			// если не указан флаг "только строки с разделителем" и не найден сепаратор, то добавляем первую колонку(вся строка)
			row := Row{}
			row.Column = append(row.Column, elTab.Column[0])
			tabNew = append(tabNew, row)
		} else if elTab.Separator {
			row := Row{}
			// если указан сепаратор, то добавляем только указанные колонки
			for _, el := range fields {

				if el > len(elTab.Column) {
					break
				}
				row.Column = append(row.Column, elTab.Column[el-1])
			}
			tabNew = append(tabNew, row)
		}
	}
	return tabNew
}

func parseFields(fieldsStr string) {
	fieldsTempStr := strings.Split(fieldsStr, "-")
	var prevEl int
	// добавление в список интервала
	for i, el := range fieldsTempStr {
		arr := strings.Split(el, ",")
		if i > 0 {
			nextEl, err := strconv.Atoi(arr[0])
			if err != nil {
				panic(err)
			}
			// добавление в список значений между левой и правой границей интервала
			for i := prevEl + 1; i < nextEl; i++ {
				fields = append(fields, i)
			}
		}

		// добавление остальных значений, перечисленных через ,
		for _, sliceEl := range arr {
			curEl, err := strconv.Atoi(sliceEl)
			if err != nil {
				panic(err)
			}
			fields = append(fields, curEl)
		}
		prevEl = fields[len(fields)-1]
	}
}

func main() {
	var tab Tab
	var fieldsStr string

	flag.StringVar(&fieldsStr, "f", "1-10,11,12,13-14,16", "'fields' - выбрать поля (колонки)")
	flag.StringVar(&delimiter, "d", "\t", "'delimiter' - использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "'separated' - только строки с разделителем")
	flag.Parse()

	// парсинг строки fields. с учетом перечислений и интервалов
	parseFields(fieldsStr)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := sc.Text()
		row := Row{}
		// задание флага наличия сепаратора
		if strings.Contains(text, delimiter) {
			row.Separator = true
		}
		row.Column = append(row.Column, strings.Split(sc.Text(), delimiter)...)
		tab = append(tab, row)
	}

	// вызов функции обрезки строки
	tabNew := cut(tab)

	// вывод на экран результата выполнения функции cut
	for i, el := range tabNew {
		if i > 0 {
			fmt.Println("")
		}
		if len(el.Column) == 0 {
			fmt.Print("")
			continue
		}
		for l, elRow := range el.Column {
			if l > 0 {
				fmt.Print(" ")
			}
			fmt.Print(elRow)
		}
	}
	fmt.Println("")
}
