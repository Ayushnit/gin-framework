package Models

type Student struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"fName"`
	LastName  string `json:"lName"`
	DOB       string `json:"dob"`
	Address   string `json:"address"`
	Subject   string `json:"subject"`
	Marks     uint   `json:"marks"`
}
func (b *Student) TableName() string {
	return "student"
}
