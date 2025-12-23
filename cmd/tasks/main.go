package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const d = 300 * time.Millisecond

func main() {
	fmt.Printf("➡️ \033[93m%s\033[0m\n", "Собираю твой гавнокод. Подожди...")
	time.Sleep(d)
	fmt.Printf("➡️ \033[93m%s\033[0m\n", "Я видел говнокод, но это - событие")
	time.Sleep(d)
	fmt.Printf("➡️ \033[93m%s\033[0m\n", "Ты это писал трезвым?")
	time.Sleep(d)
	fmt.Printf("➡️ \033[93m%s\033[0m\n", "Такой хуйни я давно не компилил")
	time.Sleep(d)
	fmt.Printf("➡️ \033[93m%s\033[0m\n", "Даже если это заработает - тебе нельзя этим гордиться")
	time.Sleep(d)
	panic("Опять ошибка: автор кода - долбоёб!!!")
}

func fanOut(ctx context.Context, in <-chan int, count int) <-chan int {
	out := make(chan int) // Выходной канал результатов
	var wg sync.WaitGroup
	wg.Add(count) // Ожидаем завершение всех workers

	for i := 0; i < count; i++ { // В цикле создаем по одной горутине на worker
		go func() {
			defer wg.Done() // Уменьшаем счетчик по завершению горутины

			for {
				select {
				case <-ctx.Done(): // Контекст дал сигнал на остановку
					return // Выходим из цикла
				case val, ok := <-in: // Блокирующе ждем данных из входного канала
					if !ok { // Канал закрыт
						return // Выходим из цикла
					}

					select {
					case <-ctx.Done():
					case out <- val * val: // В выходной канал складываем вычисления
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()  // Ожидаем завершения всех workers
		close(out) // Закрываем канал с результатами
	}()

	return out
}
