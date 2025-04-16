package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employees struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EmpID        string             `bson:"emp_id" json:"emp_id"`
	FirstName     string             `bson:"first_name" json:"first_name"`
	LastName      string             `bson:"last_name" json:"last_name"`
	Email         string             `bson:"email" json:"email"`
	MobileNO      string             `bson:"mobile_no" json:"mobile_no"`
	Gender        string             `bson:"gender" json:"gender"`
	Age           int                `bson:"age" json:"age"`
	EmployeeType  string             `bson:"employee_type" json:"employee_type"`
	Department    string             `bson:"department" json:"department"`
}
