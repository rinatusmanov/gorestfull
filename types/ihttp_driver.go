package types

// IHttpDriver добавляет новый роутинг
type IHttpDriver interface {
	// SetHandler добавляет новый роут, pattern - роут,
	// fun - функция необходимая для запуска по данному роуту
	// функция возвращает значения, 	err - ошибка которая возникает при добавления роута
	SetHandler(pattern string, fun THandler) (err error)
}

type THandler func(rw IReadWriter)
