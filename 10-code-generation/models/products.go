package models

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (products Products) Filter(predicate func(product Product) bool) Products {
	result := Products{}
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func (products Products) All(predicate func(product Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

func (products Products) Any(predicate func(product Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}
