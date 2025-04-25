package main

import (
	"basic/task3/db"
	"fmt"
)

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

CREATE TABLE employees (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    department VARCHAR(100) NOT NULL,
    salary DECIMAL(10, 2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/

type Employee struct {
	ID         uint    `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func getEmployee() {
	// 1. 连接数据库
	conn := db.ConnectMySQL()
	defer conn.Close()
	// 2. 部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
	query1 := "select * from employees where department = '技术部'"
	var employees []Employee
	err := conn.Select(&employees, query1)
	if err != nil {
		fmt.Println(err)
	}
	for _, employees := range employees {
		fmt.Println(employees)
	}

	// 3. 工资最高的员工信息，并将结果映射到一个 Employee 结构体中
	emp := new(Employee)
	query2 := "select * from employees order by salary desc limit 1"
	err = conn.Get(emp, query2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(emp)
}

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

// Book 结构体严格对应数据库表结构
type Book struct {
	ID     int     `db:"id"`     // 明确对应数据库列名
	Title  string  `db:"title"`  // 使用小写标签
	Author string  `db:"author"` // 与数据库列名完全一致
	Price  float64 `db:"price"`  // 使用float64对应DECIMAL类型
}

func getBooks() {
	// 1. 连接数据库
	conn := db.ConnectMySQL()
	defer conn.Close()
	// 2. 查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全
	var books []Book
	err := conn.Select(&books, "SELECT id, title, author, price FROM books WHERE price >50")
	if err != nil {
		fmt.Println(err)
	}
	for _, employees := range books {
		fmt.Println(employees)
	}

}

func main() {
	getBooks()
}
