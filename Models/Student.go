package Models

import (
	"fmt"
	"gin-framework/Config"
	_ "github.com/go-sql-driver/mysql"
)
func GetAllStudents(st *[]Student) (err error) {
	if err = Config.DB.Find(st).Error; err != nil {
		return err
	}
	return nil
}
func CreateStudent(st *Student) (err error) {
	if err = Config.DB.Create(st).Error; err != nil {
		return err
	}
	return nil
}
func GetStudentByID(st *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(st).Error; err != nil {
		return err
	}
	return nil
}
func UpdateStudent(st *Student, id string) (err error) {
	fmt.Println(st)
	Config.DB.Save(st)
	return nil
}
func DeleteStudent(st *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(st)
	return nil
}
