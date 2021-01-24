package types

// IMaker Основной интерфейс библиотеки
type IMaker interface {
	// Crud создает create(MethodPost), read(MethodGet), update(MethodPut), delete(MethodDelete) методы для работы с сущностью
	Crud(model interface{}, options ...Option) IMaker
	// Error возвращает ошибку выполнения предыдущей операции Crud
	Error() (err error)
}

// TInitOptionType Тип создан для индификации опций
type TInitOptionType string

// Option опции для IMaker.Crud метода
type Option interface {
	Type() TInitOptionType
	Value() interface{}
}
