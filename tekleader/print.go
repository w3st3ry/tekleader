package tekleader

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func PrintLeader(students *SortStudents) {
	total := len(*students)
	var format string
	for i, student := range *students {
		if Race {
			format = fmt.Sprintf("%s \t\t| GPA: %s\n", student.Login, student.Gpa)
		} else {
			std := GetStudent(student.Login)
			format = fmt.Sprintf("%s \t\t| %s | %d credits | GPA: %s\n", std.Title,
				std.Location,
				std.Credits,
				std.Gpa[0].Gpa)
		}
		fmt.Printf("[%d/%d] - %s", total-i, total, format)
	}
}

func PrintStatus(persistent bool) {
	for {
		err := extendedTimeoutRequest()
		t := time.Now()
		clock := fmt.Sprintf("[%s] - ", t.Format(time.Stamp))

		if err != nil {
			color.Red(clock + "Intranet is down... ./fixbocal.exe :noel:\n")
		} else {
			color.Green(clock + "Intranet is alive :hap:")
			time.Sleep(time.Second * time.Duration(Timeout))
			if !persistent {
				fmt.Println("\n")
				break
			}
		}
	}
}
