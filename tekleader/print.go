package tekleader

import (
	"fmt"
	str "strings"
	"time"

	"github.com/fatih/color"
)

func PrintLeader(students *SortStudents) {
	total := len(*students)
	format := ""
	logins := splitLogins(Find)

	// Iterate on all sorted students
	for i, student := range *students {

		// If Find flag is set, print only the student you want to find
		for x, login := range logins {
			if login == "" || (login != "" && student.Login == login) {

				if login != "" {
					logins = append(logins[:x], logins[x+1:]...)
				}
				// If Race is set, print only the login and don't send new request on /user
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
	}

	// Students not exist
	if len(logins[0]) != 0 {
		for _, login := range logins {
			color.Red("[WARN] %s not found.", login)
		}
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
				break
			}
		}
	}
}

func splitLogins(logins string) []string {
	return str.Split(logins, ",")
}
