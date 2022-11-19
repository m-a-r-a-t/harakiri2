package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
!=== Базовая задача ===
!
!Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
!Использовать библиотеку https://github.com/beevik/ntp.
!Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
!
!Программа должна быть оформлена с использованием как go module.
!Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
!Программа должна проходить проверки go vet и golint.
*/

func main() {

	time, err := getCurrentTime()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	fmt.Println(time)
	os.Exit(0)
}

func getCurrentTime() (time.Time, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	return time, err
}
