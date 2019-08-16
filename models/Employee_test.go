package models

import (
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDB struct{}

func (mdb *mockDB) GetAllEmployees() ([]*Employee, error) {
	emps := make([]*Employee, 0)
	emps = append(emps, &Employee{ID: 1, Location: "l1", Name: "n1"})
	emps = append(emps, &Employee{ID: 2, Location: "l2", Name: "n2"})
	return emps, nil
}

func TestGetAllEmployeesWithSomeBusinessLogic(t *testing.T) {
	env := Env{DB: &mockDB{}}
	emps, err := GetAllEmployeesWithSomeBusinessLogic(&env)
	assert.Nil(t, err)
	assert.NotNil(t, emps)
	expected := []*Employee{&Employee{ID: 1, Location: "l1", Name: "n1"}, &Employee{ID: 2, Location: "l2", Name: "n2"}}
	assert.Equal(t, expected, emps)

}

type mock3DB struct{}

func (mdb *mock3DB) GetAllEmployees() ([]*Employee, error) {
	emps := make([]*Employee, 0)
	emps = append(emps, &Employee{ID: 1, Location: "l1", Name: "n1"})
	emps = append(emps, &Employee{ID: 2, Location: "l2", Name: "n2"})
	emps = append(emps, &Employee{ID: 3, Location: "l3", Name: "n3"})
	return emps, nil
}

func TestGetAllEmployeesWithSomeBusinessLogicMock3(t *testing.T) {
	env := Env{DB: &mock3DB{}}
	emps, err := GetAllEmployeesWithSomeBusinessLogic(&env)
	assert.Nil(t, err)
	assert.NotNil(t, emps)
	expected := []*Employee{&Employee{ID: 1, Location: "l1", Name: "n1"}}
	assert.Equal(t, expected, emps)

}

type mockDBError struct{}

func (mdb *mockDBError) GetAllEmployees() ([]*Employee, error) {
	return nil, errors.New("Error occured connecting")
}

func TestGetAllEmployeesWithSomeBusinessLogicError(t *testing.T) {
	env := Env{DB: &mockDBError{}}
	emps, err := GetAllEmployeesWithSomeBusinessLogic(&env)
	assert.Nil(t, emps)
	assert.NotNil(t, err)

}
