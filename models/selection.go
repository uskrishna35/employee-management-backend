package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Selection struct {
    ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    OrganizationID   string             `bson:"org_id" json:"org_id"`
    FacilityID       string             `bson:"facility_id" json:"facilityId"`
    CampID           string             `bson:"camp_id" json:"campId"`
    EmployeeID       string             `bson:"employee_id" json:"employeeId"`
    DepartmentID     string             `bson:"department_id" json:"departmentId"`
    OrganizationName string             `bson:"organization_name" json:"organizationName`
    FacilityName     string             `bson:"facility_name" json:"facilityName"`
    CampName         string             `bson:"camp_name" json:"campName"`
    EmployeeName     string             `bson:"employee_name" json:"employeeName"`
    DepartmentName   string             `bson:"department_name" json:"departmentName"`
    CreatedAt        primitive.DateTime `bson:"created_at" json:"createdAt"`
}




