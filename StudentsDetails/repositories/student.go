package repositories

import (
	"database/sql"
	"errors"
	config "example/StudentsDetails/Config"
	"example/StudentsDetails/models"
)

func GetStudentDetails() ([]models.StudentDetails, error) {
	rows, err := config.DB.Query("select * from student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []models.StudentDetails{}
	for rows.Next() {
		var student models.StudentDetails
		if err := rows.Scan(&student.ID, &student.Name, &student.Class, &student.Address); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func GetStudentDetailsbyID(id int) (models.StudentDetails, error) {
	var student models.StudentDetails
	row := config.DB.QueryRow("select * from student where ID=?", id)
	err := row.Scan(&student.ID, &student.Name, &student.Class, &student.Address)

	if err != nil {
		if err == sql.ErrNoRows {
			return student, errors.New("student not found")
		}
		return student, err
	}
	return student, nil
}

func CreateStudent(student models.StudentDetails) error {
	_, err := config.DB.Exec("INSERT INTO student (Name,  Class, Address) VALUES (?, ?, ?)", student.Name, student.Class, student.Address)
	return err
}

func UpdateStudent(student models.StudentDetails) error {
	_, err := config.DB.Exec("UPDATE student SET Name = ?, Class = ?, Address= ? WHERE id = ?", student.Name, student.Class, student.Address, student.ID)
	return err
}

func DeleteStudent(id int) error {
	_, err := config.DB.Exec("DELETE FROM student WHERE id = ?", id)
	return err
}

func RegisterUser(user models.LoginUser) error {

	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := config.DB.Exec(query, user.Username, user.Password)
	return err

}

// func LoginUser(user models.LoginUser, loginvals models.LoginUser) error {
// 	query := "SELECT id, username, password FROM users WHERE username = ?"
// 	row := config.DB.QueryRow(query, loginvals.Username)
// 	error := row.Scan(&user.UID, &user.Username, &user.Password)
// 	return error

// }
