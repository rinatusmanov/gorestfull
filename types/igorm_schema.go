package types

// IGormSchema необходим для дальнейшего развития библиотеки
type IGormSchema interface {
	// Table название таблицы связанной с сущностью
	Table() string
}
