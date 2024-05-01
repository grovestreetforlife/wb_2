Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Вывод: от 1 до 8 в случайном порядке, потом всё время 0, с новой строки.
asChan создает канал и запускает горутину, которая отправляет в канал числа из переданного списка с случайной задержкой до 1 секунды, а затем закрывает канал.

merge создает новый канал и запускает горутину, которая бесконечно читает из каналов a и b и отправляет прочитанные значения в новый канал. Она использует оператор select, который выбирает первый доступный канал для чтения, поэтому значения из a и b могут поступать в любом порядке, в зависимости от того, как быстро они становятся доступными. Но чтение с закрытого канала приводит к получению нулевых для типа передаваемого в канал (для чисел это 0).

Чтобы избежать, можно использовать ok, которая говорит о том, что прочитана переданная переменная или в канале пусто:
case v, ok := <-a:
	if ok {
		c <- v
	}

```