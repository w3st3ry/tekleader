package tekleader

import (
	"sort"
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

type SortStudent struct {
	Login string
	Gpa   string
}

type SortStudents []SortStudent

func SortPromotion() *SortStudents {
	// Get promotion only and set the total
	oldProm := AppendPromotion()
	total := oldProm.Total

	newProm := SortStudents{}

	input := make(chan string, total)
	result := make(chan SortStudent, total)
	quitChan := make(chan bool)

	// Setting up status bar and fetch students logins
	bar := pb.StartNew(total).Prefix("* Fetching student informations |")
	processed := 0
	for _, student := range oldProm.Items {
		input <- student.Login
	}

	// Init loop to stop workers
	go func(done *int, total int) {
		for *done != total {
			time.Sleep(10 * time.Millisecond)
		}
		quitChan <- true
	}(&processed, total)

	// Fetching gpa from login, warn to your socket
	for i := 0; i < total; i++ {
		go func() {
			for {
				login := <-input
				std := GetStudent(login)
				if std.Gpa[0].Gpa == "n/a" {
					std.Gpa[0].Gpa = "0.00"
				}
				result <- SortStudent{Login: std.Login, Gpa: std.Gpa[0].Gpa}
			}
		}()
		go func(done *int) {
			for {
				res := <-result
				newProm = append(newProm, res)
				bar.Increment()
				*done += 1
			}
		}(&processed)
	}

	// Little hack to stop workers and continue
	<-quitChan

	bar.FinishPrint("\n * Next step * \n")

	// Sort students given
	sort.Sort(newProm)

	return &newProm
}

// Collection of methods that satisfies sort.Interface
func (gpa SortStudents) Len() int {
	return len(gpa)
}

func (gpa SortStudents) Less(i, j int) bool {
	return gpa[i].Gpa < gpa[j].Gpa
}

func (gpa SortStudents) Swap(i, j int) {
	gpa[i], gpa[j] = gpa[j], gpa[i]
}
