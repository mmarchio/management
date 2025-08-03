package models

type PromptTemplate struct {
	Model
	ID 			string 	`form:"id" json:"id"`
	Name		string  `form:"name" json:"name"`
	Template 	string 	`form:"template" json:"template"`
	Vars 		string 	`form:"vars" json:"vars"`
}