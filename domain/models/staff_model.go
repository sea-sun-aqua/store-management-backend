package models

type Staff struct {
	StaffID 		string `json:"staff_id" db:"staff_id"`
	Name		 	string `json:"staff_name" db:"staff_name"`
	Email 			string `json:"staff_email" db:"staff_email"`
	Password 		string `json:"staff_password" db:"staff_password"`
}