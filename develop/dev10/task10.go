package main

/*
Утилита telnet

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
1.	Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
2.	Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
3.	При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var host string
var port int
var timeOutStr string

func copyTo(target io.Writer, source io.Reader) {
	buff := make([]byte, 1024)
	for {
		// чтение данных из источника
		n, err := source.Read(buff[:])
		if errors.Is(err, io.EOF) {
			fmt.Println("Connection is closed")
			os.Exit(0)
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		// запись данных в целевой writer
		_, err = target.Write(buff[:n])
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			return
		}
	}
}

func main() {
	// парсинг аргументов командной строки
	flag.StringVar(&host, "h", "127.0.0.1", "host for connection")
	flag.IntVar(&port, "p", -1, "port for connection")
	flag.StringVar(&timeOutStr, "t", "10s", "timeout connection")
	flag.Parse()

	if port < 0 || port > 65535 {
		log.Fatal("Port is not found")
	}

	timeOut, err := time.ParseDuration(timeOutStr)
	if err != nil {
		log.Fatal(err)
	}

	// создание подключения по tcp к хосту
	conn, err := net.DialTimeout("tcp", host+":"+fmt.Sprint(port), timeOut)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	// обработка нажатия Ctrl+D
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		os.Exit(0)
	}()

	// функция копирования ответа в стандартный поток вывода
	go copyTo(os.Stdout, conn)

	// функция копирования стандартного потока ввода для отправки на хост
	copyTo(conn, os.Stdin)
}
