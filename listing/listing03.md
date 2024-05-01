Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Первый вывод nil
Второй вывод false

Структура пустого интерфейса в *рефлексии:

type emptyInterface struct {
	typ  *abi.Type
	word unsafe.Pointer
}

Структура интерфейса в рантайме:

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

Инициализация с присвоением на 12 строке присвоило !динамический! тип в структуре emptyInterface полю typ - os.PathError.
Таким образом word в emptyInterface == nil, но тип интерфейса != nil.
Поэтому при выводе мы видим nil.
```
