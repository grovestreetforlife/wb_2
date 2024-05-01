package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Car struct {
	color      string
	engineType string
	hasSunroof bool
}

type CarBuilder interface {
	SetColor(color string) CarBuilder
	SetEngineType(engineType string) CarBuilder
	SetSunroof(hasSunroof bool) CarBuilder
	Build() *Car
}

func NewCarBuilder() CarBuilder {
	return &carBuilder{
		car: &Car{},
	}
}

type carBuilder struct {
	car *Car
}

func (cb *carBuilder) SetColor(color string) CarBuilder {
	cb.car.color = color
	return cb
}

func (cb *carBuilder) SetEngineType(engineType string) CarBuilder {
	cb.car.engineType = engineType
	return cb
}

func (cb *carBuilder) SetSunroof(hasSunroof bool) CarBuilder {
	cb.car.hasSunroof = hasSunroof
	return cb
}

func (cb *carBuilder) Build() *Car {
	return cb.car
}

type Director struct {
	builder CarBuilder
}

func (d *Director) ConstructCar(color, engineType string, hasSunroof bool) *Car {
	d.builder.SetColor(color).SetEngineType(engineType).SetSunroof(hasSunroof)
	return d.builder.Build()
}

/*
 Паттерн Builder определяет процесс поэтапного построения сложного продукта.
После того как будет построена последняя его часть, продукт можно использовать.

	Применимость:
 Когда мы хотим избавиться от «телескопического конструктора».
 Когда нам нужно собирать сложные составные объекты.

	Плюсы и минусы:
 +Позволяет варьировать внутреннее представление продукта.
 +Инкапсулирует код для построения и представления.
 +Обеспечивает контроль над этапами процесса построения.

 -Для каждого типа продукта должен быть создан отдельный ConcreteBuilder.
 -Классы Builder должны быть мутабельными.
 -Может затруднять/усложнять инъекцию зависимостей.

*/
