package main

import (
	"context"
	"sync"
)

func main() {

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
