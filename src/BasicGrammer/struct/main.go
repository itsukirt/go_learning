package main

import "fmt"

// 自定义类型：通过type关键字
type NewInt int

// 类型别名：只在代码中存在，编译完成时不存在MyInt类型。
type MyInt = int

// 结构体
type person struct {
	name, city string
	age        int8
}

func main() {
	// 自定义类型与类型别名
	var a NewInt
	var b MyInt
	fmt.Printf("type of a: %T\n", a) // type of a: main.NewInt
	fmt.Printf("type of b: %T\n", b) // type of b: int

	// 结构体实例化
	var p1 person
	p1.name = "HAKUNO"
	p1.city = "Xi'an"
	p1.age = 24

	fmt.Printf("p1 = %v\n", p1)  // p1 = {HAKUNO Xi'an 24}
	fmt.Printf("p1 = %#v\n", p1) // p1 = main.person{name:"HAKUNO", city:"Xi'an", age:24}

	// 匿名结构体：临时数据结构等场景使用。
	var user struct {
		Name string
		Age  int
	}
	user.Name = "Teemo"
	user.Age = 7
	fmt.Printf("%#v\n", user) // struct { Name string; Age int }{Name:"Teemo", Age:7}

	// 指针类型结构体
	var p2 = new(person)
	p2.name = "YaSuo"
	p2.city = "艾欧尼亚"
	p2.age = 35
	fmt.Printf("%T\n", p2)       // *main.person
	fmt.Printf("p2 = %#v\n", p2) // p2 = &main.person{name:"YaSuo", city:"艾欧尼亚", age:35}

	// 取结构体地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)     // *main.person
	fmt.Printf("p3=%#v\n", p3) // p3=&main.person{name:"", city:"", age:0}
	p3.name = "ZOE"
	p3.age = 19
	p3.city = "NULL"
	fmt.Printf("p3=%#v\n", p3) // p3=&main.person{name:"ZOE", city:"NULL", age:19}
}
