package pattern

/*
	Реализовать паттерн «фабричный метод».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
type action string

const (
	A action = "A"
	B action = "B"
)

type Creator interface {
	CreateProduct(action action) Product
}

type Product interface {
	Use() string
}

type ConcreteCreator struct{}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) CreateProduct(action action) Product {
	var product Product

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	default:
		print("Unknown Action")
	}

	return product
}

type ConcreteProductA struct {
	action string
}

func (p *ConcreteProductA) Use() string {
	return p.action
}

type ConcreteProductB struct {
	action string
}

func (p *ConcreteProductB) Use() string {
	return p.action
}

/*
 Паттерн Factory Method полезен, когда система должна оставаться легко расширяемой
путем добавления объектов новых типов. По этому, если перед разработчиком стоят
не четкие требования для продукта или не ясен способ организации взаимодействия между
продуктами - он идеально подойдёт.

 Применяется для создания объектов с определенным интерфейсом, реализации которого
предоставляются потомками. Другими словами, есть базовый абстрактный класс фабрики,
который говорит, что каждая его наследующая фабрика должна реализовать такой-то метод
для создания своих продуктов.

	Применимость:
 -Классу заранее неизвестно, объекты каких подклассов ему нужно создавать.
 -Класс спроектирован так, чтобы объекты, которые он создаёт, специфицировались подклассами.
 -Класс делегирует свои обязанности одному из нескольких вспомогательных подклассов,
и планируется локализовать знание о том, какой класс принимает эти обязанности на себя.

	Плюсы и минусы:
 +Упрощает добавление новых продуктов.
 +Упрощает добавление новых продуктов.
 +Уменьшает зависимости между клиентом и фабрикой.

 -Увеличение количества кода.
 -Необходимо создавать наследника Creator для каждого нового типа продукта (ConcreteProduct).
*/
