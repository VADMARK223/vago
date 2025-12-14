package main

import (
	"context"
	"fmt"
	"time"
)

const (
	workTime = 2 * time.Second
	waitTime = 1 * time.Second
)

func getText() string {
	time.Sleep(workTime)
	return "VADMARK"
}

func Run() {
	ch := make(chan string, 1)
	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()

	go func(readCh chan<- string) {
		fmt.Println("Стартуем горутину.")
		result := getText()
		select {
		case readCh <- result:
			fmt.Println("Получен результат в горутине:", result)
		case <-ctx.Done(): // Защита от блокировки на отправку, если главный поток уже вышел по cxt.Done()
			fmt.Println("Контекст выполнился, завершаем горутину.")
		}
	}(ch)

	select {
	case msg := <-ch:
		fmt.Println("Результат получен в главном потоке:", msg)
	case <-ctx.Done():
		fmt.Println("Контекст завершился: ", ctx.Err())
	}

	fmt.Println("Выходим из главного потока.")
}
