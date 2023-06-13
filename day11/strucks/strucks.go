package main

import (
	"fmt"
)

type product struct {
	name, catagory string
	Price          float64
}

type OrderProduct struct {
	product    product //embedded fields
	totalorder int     //regular fields
}

func initProduct() {
	//literal value

	prod := product{
		name:     "Laptop",
		catagory: "Electronic",
		Price:    250,
	}

	prod.name = "Dell Laptop"
	prod.Price = 355.98
	fmt.Println(prod)

	//default zero value
	var tv product
	tv.name = "sharp"
	meja := product{}
	fmt.Println(tv, meja)
}

func initProductPointer() {
	// built-in new function,result pointer struct
	var phone = new(product)
	phone.name = "Samsung"
	phone.catagory = "HP"
	phone.Price = 235.98

	fmt.Println(phone)

	//using &
	var oppo = &product{}
	oppo.name = "oppo"
	oppo.catagory = "HP"
	oppo.Price = 185

	fmt.Println(oppo)
}

func initOrderProduct() {
	myOrder := OrderProduct{
		product:    product{"Iphone", "HP", 355},
		totalorder: 1,
	}
	fmt.Println(myOrder)
}

func compareStruck() {

	oppo := product{
		name: "Oppo", catagory: "HP", Price: 185,
	}

	Iphone := product{
		name: "Oppo", catagory: "HP", Price: 185,
	}

	fmt.Println("result : ", oppo == Iphone)

}

// anonymous struct
func printStruct(val struct {
	name, catagory string
	Price          float64
}) {

	fmt.Println("name: ", val.name, "catagory:", val.catagory, "price:", val.Price)
}

func initSliceProduct() []product {

	oppo := product{
		name: "Oppo", catagory: "HP", Price: 185,
	}

	Iphone := product{
		name: "Oppo", catagory: "HP", Price: 185,
	}

	fmt.Println("result : ", oppo == Iphone)

	products := []product{oppo, Iphone}
	products = append(products,
		product{name: "samsung", catagory: "HP", Price: 256})

	return products
}

func FindProductByName(products []product, search string) product {
	for _, v := range products {
		if v.name == search {

			return v
		}
	}
	return product{}

}
func main() {
	//initProduct()
	//initProductPointer()
	//initOrderProduct()
	//compareStruck()
	//fmt.Println(FindProductByName(initSliceProduct(), "samsung"))

}
