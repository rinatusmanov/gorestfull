package maker

import "github.com/rinatusmanov/gorestfull/types"

// generatedOption простой генератор опций
type generatedOption struct {
	TypeData  types.TInitOptionType
	ValueData interface{}
}

func (g *generatedOption) Type() types.TInitOptionType {
	return g.TypeData
}

func (g *generatedOption) Value() interface{} {
	return g.ValueData
}


// TOptionTypeUri генерирует Option для смены стандартного енднпоинта на новый указаный в параметре
func TOptionTypeUri(uri string) types.Option {
	return &generatedOption{
		TypeData:  uriSet,
		ValueData: uri,
	}
}
