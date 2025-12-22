package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Контекст для сигнала остановки
	defer cancel()                                                           // В конце обязательно отменяем контекст

	in := make(chan int) // Входной канал, который надо распараллелить

	go func() { // Имитируем медленную запись во входной канал
		for i := 2; i <= 5; i++ {
			time.Sleep(1 * time.Second)
			in <- i
		}

		close(in) // Когда все записали, закрываем его
	}()

	out := fanOut(ctx, in, workerCount)
	for result := range out { // Блокирующе читаем результаты
		fmt.Println("Result:", result)
	}
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
