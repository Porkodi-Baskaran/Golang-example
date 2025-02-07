package DB

import (
	"database/sql"
	"fmt"
)

type StudentDetails struct {
	ID      int32
	Name    string
	Class   string
	Address string
}

func RetrieveDatas(db *sql.DB, name string) ([]StudentDetails, error) {

	var Studdetails []StudentDetails

	rows, err := db.Query("Select ID,Name from student where name=?", name)

	if err != nil {
		return nil, fmt.Errorf("STUDENTDETAILS %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var stud StudentDetails
		// err := rows.Scan(&stud.ID, &stud.Name, &stud.Class, &stud.Address)
		err := rows.Scan(&stud.ID, &stud.Name)
		if err != nil {
			return nil, fmt.Errorf("STUDENTDETAILS %q: %v", name, err)
		}
		Studdetails = append(Studdetails, stud)

	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("STUDENTDETAILS %q: %v", name, err)
	}
	return Studdetails, nil
}
