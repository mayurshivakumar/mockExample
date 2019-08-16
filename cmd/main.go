package main

import (
	"log"
	"mockExample/models"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	db, err := models.NewDB("sqlserver")
	if err != nil {
		log.Panic(err)
	}
	env := models.Env{DB: db}
	emps, err := models.GetAllEmployeesWithSomeBusinessLogic(&env)
	if err != nil {
		log.Panic(err)
	}

	for _, emp := range emps {
		log.Println(emp)
	}

}
