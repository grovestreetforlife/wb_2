package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».

Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Facade_pattern
*/
type CPU struct{}

func (c *CPU) Signal() {
	fmt.Println("Ping..")
}

type HardDrive struct{}

func (h *HardDrive) Read() {
	fmt.Println("Reading..")
}

type Memory struct{}

func (m Memory) Load() {
	fmt.Println("Loading..")
}

type Computer struct {
	cpu       *CPU
	hardDrive *HardDrive
	memory    *Memory
}

func (c *Computer) Start() {
	c.cpu.Signal()
	c.hardDrive.Read()
	c.memory.Load()
}

/*
	Применимость:
 Даёт возможность: предоставить простой или урезанный интерфейс к сложной подсистеме и
может выполнять дополнительную функциональность до/после пересылки запроса;
уменьшить количество зависимостей между клиентом и сложной системой.

	Плюсы и минусы:
 +Изолирует клиентов от сложной подсистемы.
 +Дополнительная функциональность до/после пересылки запроса.

 -Интерфейс может сильно вырасти и привязаться ко всем структурам программы.
*/
