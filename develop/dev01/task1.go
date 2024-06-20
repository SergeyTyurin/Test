package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

var ntpServer = "0.beevik-ntp.pool.ntp.org"

func Time() error {
	// получение ответа от ntp сервера
	response, err := ntp.Query(ntpServer)
	if err != nil {
		// запись в стандартный поток ошибок
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	// вывод в стандартный вывод текущего времени и уточненного времени
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Add(response.ClockOffset))
	return nil
}
