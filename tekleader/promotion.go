package tekleader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var apiOffset int

type Promotion struct {
	Items []struct {
		Title    string `json:"title"`
		Login    string `json:"login"`
		Nom      string `json:"nom"`
		Prenom   string `json:"prenom"`
		Location string `json:"location"`
	} `json:"items"`
	Total int `json:"total"`
}

// AppendPromotion fetch students obstruct API offset
func AppendPromotion() *Promotion {
	// Init promotion structs
	aprom := &Promotion{}
	prom := aprom

	// Get apiOffset, abort if parameters are wrongs
	apiOffset = len(getPromotion(0).Items)
	if apiOffset < 1 {
		log.Fatal("Wrong parameters")
	}

	// Fetch and append data in new array
	for i := 0; ; i += apiOffset {
		prom = getPromotion(i)
		if len(prom.Items) == 0 {
			break
		}
		aprom.Items = append(aprom.Items, prom.Items...)
	}

	// Set total of students
	aprom.Total = prom.Total

	return aprom
}

// getPromotion fetch students using parameters given
func getPromotion(offset int) *Promotion {
	prom := Promotion{}
	req, err := http.NewRequest("GET",
		intraURL+
			AuthKey+
			"/user/filter/user"+
			jsonFormat+
			"&location=FR/"+Location+
			"&year="+strconv.Itoa(time.Now().Year())+
			"&course="+Course+"/classic"+
			"&active=true"+
			"&promo="+Promo+
			"&offset="+strconv.Itoa(offset), nil)
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

	json.Unmarshal(body, &prom)

	return &prom
}
