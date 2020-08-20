package ddd

type LayerSpec struct {
	name      string
	comment   string
	api       []StructOrInterfaceOrFunc
	factories []*FuncOrTypeSpecs
}

// UseCases can only ever import DomainCore API
func UseCases(api []StructOrInterface, factories []FuncOrStruct) *LayerSpec {
	return NewLayer("UseCases", "...all the use cases of the domain.", api, factories)
}

// DomainCore has never any dependencies to any other layer.
func DomainCore(api []StructOrInterface, factories []FuncOrStruct) *LayerSpec {
	return NewLayer("DomainCore", "...all the core domain API of the domain.", api, factories)
}



func NewLayer(name, comment string, api []StructOrInterfaceOrFunc, factories ...*FuncOrTypeSpecs) *LayerSpec {
	return &LayerSpec{
		name:      name,
		comment:   comment,
		api:       api,
		factories: factories,
	}
}
