package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Handler interface {
	SendRequest(message int) string
}

type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type ConcreteHandlerC struct {
	next Handler
}

func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

/*
 Паттерн Chain Of Responsibility позволяет избежать привязки объекта-отправителя запроса к
объекту-получателю запроса, при этом давая шанс обработать этот запрос нескольким объектам.
Получатели связываются в цепочку, и запрос передается по цепочке, пока не будет обработан каким-то объектом.
По сути это цепочка обработчиков, которые по очереди получают запрос, а затем решают, обрабатывать его или нет.
Если запрос не обработан, то он передается дальше по цепочке. Если же он обработан, то паттерн сам решает
передавать его дальше или нет. Если запрос не обработан ни одним обработчиком, то он просто теряется.
	Применяемость:
 Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно, какие
конкретно запросы будут приходить и какие обработчики для них понадобятся.
 Когда набор объектов, способных обработать запрос, должен задаваться динамически.
 Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	Плюсы и минусы:
 +Уменьшает связность и увеличивает гибкость обработки запросов.

 -Запрос может остаться никем не обработанным.
*/
