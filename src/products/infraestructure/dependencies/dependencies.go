package dependencies

import (
	"fmt"
	"mundo/src/core"
	"mundo/src/products/application"
	"mundo/src/products/infraestructure"
	"mundo/src/products/infraestructure/http/controllers"
)

var (mySQL infraestructure.MySQL)

func Init() {
	db, err := core.InitMySQL()
	if err != nil {
		fmt.Println("Error de servidor")
		return
	}

	// defer db.Close()
	mySQL = *infraestructure.NewMySQL(db)
}

//CASOS DE USO

func CreateProductController() *controllers.CreateProductController {
	ucCreateProduct := application.NewCreateProduct(&mySQL)

	return controllers.NewCreateProductController(ucCreateProduct)
}

func GetAllProductController() *controllers.GetAllProductController {
	ucGetAllProducts := application.NewGetAllProduct(&mySQL)

	return controllers.NewGetAllProductController(ucGetAllProducts)
}

func UpdateProductController() *controllers.UpdateProductByIDController {
	ucUpdateProduct := application.NewUpdateProductByID(&mySQL)

	return controllers.NewUpdateProductByIDController(ucUpdateProduct)
}

func DeleteProductController() *controllers.DeleteProductByIDController {
	ucDeleteProduct := application.NewDeleteProductByID(&mySQL)

	return controllers.NewDeleteProductByIDController(ucDeleteProduct)
}
