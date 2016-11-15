package tekleader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Students []Student

type Student struct {
	Login    string `json:"login"`
	Title    string `json:"title"`
	Promo    int    `json:"promo"`
	Location string `json:"location"`
	Credits  int    `json:"credits"`
	Gpa      []struct {
		Gpa   string `json:"gpa"`
		Cycle string `json:"cycle"`
	} `json:"gpa"`
}

func GetStudent(login string) *Student {
	student := Student{}

	req, err := http.NewRequest("GET",
		intraURL+
			AuthKey+
			"/user/"+
			login+
			"/"+jsonFormat, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := getResRequest(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &student)

	return &student
}

func GetStudents(logins []string) *Students {
	students := Students{}

	for _, login := range logins {
		res := GetStudent(login)
		students = append(students, *res)
	}

	return &students
}
