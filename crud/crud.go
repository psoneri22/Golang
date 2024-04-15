package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/gorilla/mux"
// )

// // Employee struct represents the structure of an employee
// type Employee struct {
// 	ID     int
// 	Name   string
// 	Age    int
// 	Salary float64
// }

// var db *sql.DB

// func main() {
// 	var err error
// 	// Open database connection
// 	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/golang")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	// Create table
// 	createTable(db)

// 	//create emplioyee
// 	staticEmployee := Employee{Name: "John", Age: 30, Salary: 50000}

// 	// Initialize router
// 	r := mux.NewRouter()

// 	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
// 		// Serve the HTML file
// 		http.ServeFile(w, r, "index.html")
// 	})
// 	// Define routes
// 	r.HandleFunc("/employees", getEmployees).Methods("GET")
// 	r.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
// 	r.HandleFunc("/employees", createEmployee(db, staticEmployee)).Methods("POST")
// 	r.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
// 	r.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")

// 	// Start server
// 	fmt.Println("Server is listening on port 8080...")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// func createTable(db *sql.DB) {
// 	query := `
//         CREATE TABLE IF NOT EXISTS employees (
//             id INT AUTO_INCREMENT PRIMARY KEY,
//             name VARCHAR(50) NOT NULL,
//             age INT,
//             salary DOUBLE
//         )`
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println("Table created successfully")
// }

// // getEmployees retrieves all employees
// func getEmployees(w http.ResponseWriter, r *http.Request) {
// 	// Implement code to fetch all employees from the database
// 	// and return JSON response
// }

// // getEmployee retrieves a single employee by ID
// func getEmployee(w http.ResponseWriter, r *http.Request) {
// 	// Implement code to fetch an employee by ID from the database
// 	// and return JSON response
// }

// // createEmployee creates a new employee
// func createEmployee(db *sql.DB, emp Employee) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Insert employee into database
// 		query := "INSERT INTO employees (name, age, salary) VALUES (?, ?, ?)"
// 		_, err := db.Exec(query, emp.Name, emp.Age, emp.Salary)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		// Write success response
// 		w.WriteHeader(http.StatusCreated)
// 		fmt.Fprintln(w, "Employee inserted successfully")
// 	}
// }

// // updateEmployee updates an existing employee
// func updateEmployee(w http.ResponseWriter, r *http.Request) {
// 	// Implement code to parse JSON request body and update an employee
// 	// in the database
// }

// // deleteEmployee deletes an employee by ID
// func deleteEmployee(w http.ResponseWriter, r *http.Request) {
// 	// Implement code to delete an employee by ID from the database
// }
