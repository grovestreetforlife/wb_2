package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Visitor interface {
	VisitSushiBar(s *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
}

type Place interface {
	Accept(v Visitor)
}

type People struct {
}

func (v *People) VisitSushiBar(s *SushiBar) string {
	return s.BuySushi()
}

func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

type SushiBar struct {
}

func (s *SushiBar) Accept(v Visitor) {
	v.VisitSushiBar(s)
}

func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

type Pizzeria struct {
}

func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

/*
 Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс
этого объекта.
	Применимость:
 Когда нужно добавить новые операции к объектам разных классов, не изменяя их.
Когда нужно выполнить операцию над всеми элементами сложной структуры объектов.

	Плюсы и минусы:
 +Упрощает добавление операций, работающих со сложными структурами объектов.
 +Объединяет родственные операции в одном классе.
 +Посетитель может накапливать состояние при обходе структуры элементов.

 - Паттерн не оправдан, если иерархия элементов часто меняется.
- Может привести к нарушению инкапсуляции элементов.
*/
