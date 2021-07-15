package model

type Employee struct {
	ID        string `json:"id" gorm:"<-;column:id"`
	Firstname string `json:"firstname" gorm:"<-;column:firstname"`
	Lastname  string `json:"lastname" gorm:"<-;column:lastname"`
	Salary    int    `json:"salary" gorm:"<-;column:salary"`
}
