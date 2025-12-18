package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			select {
			case ch1 <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		for i := 4; i < 6; i++ {
			select {
			case ch2 <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	out := fanIn(ctx, ch1, ch2)

	sum := 0
	for v := range out {
		sum += v
	}
	fmt.Println("result:", sum)
}
func fanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {
		// Проходимся по всем входным каналам
		go func(c <-chan int) {
			// Запускаем отдельную горутину для КАЖДОГО входного канала.
			// ВАЖНО: передаём ch как аргумент (c), чтобы избежать ловушки замыкания.
			defer wg.Done()
			// Гарантируем, что при любом выходе из горутины
			// счётчик WaitGroup будет уменьшен ровно один раз.
			for {
				// Бесконечный цикл — выходим только по return:
				// 1) входной канал закрыт
				// 2) контекст отменён
				select {
				case v, ok := <-c:
					// Пытаемся получить значение из входного канала.
					// Горутина блокируется здесь, пока:
					// - не придёт значение
					// - или канал не закроется
					// - или не сработает ctx.Done() (см. ниже)
					if !ok {
						// Если канал закрыт — корректно завершаем горутину.
						// defer wg.Done() отработает автоматически.
						return
					}
					// Значение успешно получено из входного канала.
					// Теперь нужно отправить его в общий выходной канал.
					select {
					case out <- v:
					// Успешно отправили значение в выходной канал.
					// Если out небуферизированный — здесь возможна блокировка,
					// пока consumer (main) не выполнит чтение.
					case <-ctx.Done():
						// Контекст отменён во время попытки отправки в out.
						// Чтобы не зависнуть на send — выходим из горутины.
						return
					}
				case <-ctx.Done():
					// Контекст отменён во время ожидания данных из входного канала.
					// Немедленно завершаем горутину, предотвращая утечку.
					return
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
