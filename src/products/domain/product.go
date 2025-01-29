package domain

type Product struct {
	//minuscula significa que es privado, mayuscula significa que es publico
	Id int32 
	Name string
	Price float32
}

//crear una estrategia manera en la cual podamos crear objetos para manipular de una forma más organizada

//newProduct es un método que nos permite crear un nuevo producto, funciona como un CONSTRUCTOR
func NewProduct(Name string, Price float32) *Product {
	//creamos un nuevo objeto de tipo Product, *Product es el valor de retorno
	return &Product{Id: 1, Name: Name, Price: Price}
}

func (p *Product) GetID() int32 {
	return p.Id
}

func (p *Product) SetID(Id int32) {
	p.Id = Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(Name string) {
	p.Name = Name
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) SetPrice(Price float32) {
	p.Price = Price
}

