// para las repositories se deben usar interfaces
package domain



type IProduct interface {
	//En las interfaces se pregunta uno ¿Que necesitan hacer los actores sobre mis casos de uso?
	//En las interfaces van los metodos de HUs
	Save(product Product) error
	GetAll() ([]Product, error)
	DeleteByID(id int) error
	UpdateByID(id int, product Product) error
}
