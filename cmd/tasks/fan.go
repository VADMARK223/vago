package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	workerCount = 3
	contextWait = 20 * time.Second
	workTimeout = 2 * time.Second
)

type Worker struct {
	id     int
	chanId int
	value  int
}

func RunFun() {
	ctx, cancel := context.WithTimeout(context.Background(), contextWait)
	defer func() {
		cancel()
	}()
	ch1 := createChannel(ctx, 1)
	ch2 := createChannel(ctx, 2)

	ch3 := make(chan int)
	go func() {
		ch3 <- 100
		close(ch3)
	}()

	outChannel := merge(ctx, ch1, ch2, ch3, nil)

	var sum int
	for v := range outChannel {
		fmt.Println("Получили из общего канала:", v)
		sum += v
	}

	fmt.Println("Выход из главного потока с результатом:", sum)
}

// Fan-out
func createChannel(ctx context.Context, chanId int) chan int {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go startWork(ctx, Worker{id: i, chanId: chanId, value: i}, &wg, ch)
	}

	go func() { // Параллельно ждем всех воркеров, чтобы потом закрыть горутину и остановить range
		wg.Wait()
		close(ch)
	}()

	return ch
}

// Fan-in
func merge(ctx context.Context, inChannels ...chan int) chan int {
	outChannels := make(chan int)
	wg := sync.WaitGroup{}

	for _, inChannel := range inChannels { // Важно идти по каналам, а не по индексам (_, inChannel)
		if inChannel == nil {
			continue
		}
		wg.Add(1)

		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				select {
				case <-ctx.Done():
					return
				case outChannels <- v:
				}
			}
		}(inChannel) // С версии 1.22 можно без прокидки.
	}

	go func() {
		wg.Wait()
		close(outChannels)
	}()

	return outChannels
}

func startWork(ctx context.Context, worker Worker, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	fmt.Println(" Начали работу:", worker.id, "Канал:", worker.chanId)

	select {
	case <-ctx.Done():
		fmt.Println("Отмена во время работы:", worker.id)
		return
	case <-time.After(workTimeout):
		fmt.Println("Отработал worker", worker.id)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Отмена перед отправкой", worker.id)
	case ch <- worker.value:
		fmt.Println(" Закончили работу worker", worker.id)
	}
}

// ==========================================
func runFan() {
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
