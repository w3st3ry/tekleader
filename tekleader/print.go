package tekleader

import (
	"fmt"

	_ "github.com/fatih/color"
)

func PrintLeader(students *SortStudents) {
	total := len(*students)
	var format string
	for i, student := range *students {
		if Race {
			format = fmt.Sprintf("%s \t\t| GPA: %s\n", student.Login, student.Gpa)
		} else {
			std := GetStudent(student.Login)
			format = fmt.Sprintf("%s \t\t| %s | %s credits | GPA: %s\n", std.Title,
				std.Location,
				std.Credits,
				std.Gpa[0].Gpa)
		}
		fmt.Printf("[%d/%d] - %s", total-i, total, format)
	}
}
