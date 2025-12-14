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

func main() {
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
