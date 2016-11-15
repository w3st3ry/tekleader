package tekleader

import (
	"fmt"

	_ "github.com/fatih/color"
)

func PrintLeader(students *SortStudents) {
	total := len(*students)
	for i, student := range *students {
		std := GetStudent(student.Login)
		fmt.Printf("[%d/%d] - %s \t| Gpa : %s\n",
			total-i,
			total,
			std.Title,
			std.Gpa[0].Gpa)
	}
}
