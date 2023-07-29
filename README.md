# Wrapper logger

## Описание
Данный модуль предназначается для логирования.  
Модуль является оберткой над [uber-go/zap](https://github.com/uber-go/zap).
- trimmer - функционал для удаления полей, которые задаются при использование данного функционала, которые не должны отображаться в логах.  
- shortener - функционал для логирования определенных полей, которые задаются при использование данной функции.
## Пример использования
Wrapper:
- Использование через конструктор.
``` go
package main

import (
	"log"

	"go.uber.org/zap"
	"github.com/requiemofthesouls/logger"
)

func main() {
	var (
		l logger.Wrapper
		err     error
	)
	if l, err = logger.New(logger.Config{
		Address:     "localhost:65000",
		Level:       "info",
		Encoding:    "console",
		Caller:      true,
		Stacktrace:  "error",
		Development: false,
	}, []logger.Field{
		logger.String("service", "example"),
	}); err != nil {
		log.Fatal(err)
	}
	
	l.Info("Starting service", logger.String("address", "localhost:80"))
}
```
- Использование через definitions.
``` go
package main

import (
	"log"

	"github.com/requiemofthesouls/container"
	"go.uber.org/zap"
	"github.com/requiemofthesouls/logger"
)

func main() {
	var l logger.Wrapper
	if err := container.Container.Fill(loggerCont.DIWrapper, &l); err != nil {
		log.Fatal(err)
	}

	l.Info("Starting service", logger.String("address", "localhost:80"))
}
```
Trimmer:
- Example.
``` go
package main

import (
	"log"

	"github.com/requiemofthesouls/logger/trimmer"
)

func main() {
	var s = trimmer.New(map[string][]string{
		"example_handler_name": {
			"example_field_name_1",
			"example_field_name_2",
		},
	})
	log.Println(string(s.Trim("example_handler_name", []byte("{\"example_field_name_1\":1, \"example_field_name_3\":2}"))))
}
```
- Output.
``` json
{
   "example_field_name_1":"TRIMMED_CONTENT",
   "example_field_name_3":2
}
```
Shortener:
- Example.
``` go
package main

import (
	"log"

	"github.com/requiemofthesouls/logger/shortener"
)

func main() {
	var s = shortener.New(map[string][]string{
		"example_handler_name": {
			"example_field_name_1",
			"example_field_name_2",
		},
	})
	log.Println(string(s.Shorten("example_handler_name", []byte("{\"example_field_name_1\":1, \"example_field_name_3\":2}"))))
}
```
- Output.
``` json
{
   "example_field_name_1":1
}
```
## Пример конфигурации
``` yaml
logger:
  level: info             # Уровень логера.
  encoding: json          # Формат логов.
  caller: true            # Аннотирования каждого сообщения именем файла.
  stacktrace: error       # Трассировка стека для заданного уровня.
  address: localhost:5110 # Адресс для сохранения логов в logstash по udp. Указывать не обязательно. 
  development: true       # Перевод логера в режим разработки. Дополнительно регистрирует panics.
```
## Зависимости от модулей
- [config](https://github.com/requiemofthesouls/config/-/blob/main/README.md)  
- [container](https://github.com/requiemofthesouls/container/-/blob/main/README.md)
