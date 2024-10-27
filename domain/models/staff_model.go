package models

type Staff struct {
	StaffID 		string `json:"staff_id"`
	Name		 	string `json:"staff_name"`
	Email 			string `json:"staff_email"`
	Password 		string `json:"staff_password"`
}