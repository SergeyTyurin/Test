package main

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {

	fmt.Println("Shell is running\nFor exit enter \\quit")

	// чтение из стандартного потока ввода
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		if strings.EqualFold(txt, "\\quit") {
			break
		}
		// разбиение пайпа команд
		cmdArr := strings.Split(txt, "|")
		//обработка каждой команды
		for _, el := range cmdArr {
			txtCmd := strings.Split(el, " ")

			if len(txtCmd) == 0 {
				continue
			}

			switch txtCmd[0] {
			case "cd":
				// переход в директорию
				var dirname string
				if len(txtCmd) < 2 {
					var err error
					dirname, err = os.UserHomeDir()
					if err != nil {
						log.Fatal(err)
					}

				} else {
					dirname = txtCmd[1]
				}
				os.Chdir(dirname)
				newDir, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Current Working Direcotry: %s\n", newDir)
			case "pwd":
				// вывод на экран текущей директории
				newDir, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Current Working Direcotry: %s\n", newDir)
			case "echo":
				// вывод на экран строки, следующей за echo
				fmt.Println(strings.Join(txtCmd[1:], " "))
			case "kill":
				// поиск и удаление процесса по id используя модуль "github.com/shirou/gopsutil/v3/process"
				processes, err := process.Processes()
				if err != nil {
					log.Fatal(err)
				}
				pid, err := strconv.ParseInt(txtCmd[1], 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				var procFound bool
				for _, p := range processes {
					if int32(pid) == p.Pid {
						p.Kill()
						procFound = true
						break
					}
				}
				if !procFound {
					fmt.Println("process not found")
				}
			case "ps":
				// отображение списка процессов, используя модуль "github.com/shirou/gopsutil/v3/process"
				processes, err := process.Processes()
				if err != nil {
					log.Fatal(err)
				}
				w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
				for _, p := range processes {
					username, err := p.Username()
					if err != nil {
						log.Fatal(err)
					}
					createTime, err := p.CreateTime()
					if err != nil {
						log.Fatal(err)
					}

					cmd, err := p.Cmdline()
					if err != nil {
						log.Fatal(err)
					}
					fmt.Fprintf(w, "%s\t%d\t%s\t%s\t\n", username, p.Pid, time.Unix(0, createTime*int64(time.Millisecond)), cmd[:min(len(cmd), 10)])
				}
				w.Flush()
			default:
				fmt.Println("Command not found")
			}
		}
	}
}
