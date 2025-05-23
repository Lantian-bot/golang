package object

import "fmt"

//2.题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
//为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者。

// Person 结构体包含人员基本信息
type Person struct {
	Name string
	Age  int
}

// Employee 结构体通过组合方式继承Person，并添加员工编号
type Employee struct {
	Person     // 匿名嵌入Person结构体
	EmployeeID string
}

// PrintInfo 方法输出员工完整信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工信息：\n姓名: %s\n年龄: %d\n工号: %s\n", e.Name, e.Age, e.EmployeeID)
}
