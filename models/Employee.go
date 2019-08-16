package models

import (
	"context"
	"fmt"
)

// Employee ...
type Employee struct {
	ID       int
	Name     string
	Location string
}

// GetAllEmployees ...
func (db *DB) GetAllEmployees() ([]*Employee, error) {
	// call db and get
	ctx := context.Background()
	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	emps := make([]*Employee, 0)

	// Iterate through the result set.
	for rows.Next() {
		emp := new(Employee)
		// Get values from row.
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Location)
		if err != nil {
			return nil, err
		}
		emps = append(emps, emp)

	}
	return emps, nil
}

//GetAllEmployeesWithSomeBusinessLogic ...
func GetAllEmployeesWithSomeBusinessLogic(env *Env) ([]*Employee, error) {
	emps, err := env.DB.GetAllEmployees()
	if err != nil {
		//log.Println(err)
		return nil, err
	}
	// adding dummy business logic to prove a point
	if len(emps) >= 3 {
		return emps[:1], nil
	}
	return emps, nil

}
