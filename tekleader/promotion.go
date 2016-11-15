package tekleader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const apiOffset = 48

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

func AppendPromotion() *Promotion {
	aprom := &Promotion{}
	prom := aprom
	for i := 0; ; i += apiOffset {
		prom = getPromotion(i)
		if len(prom.Items) == 0 {
			break
		}
		aprom.Items = append(aprom.Items, prom.Items...)
	}
	aprom.Total = prom.Total
	return aprom
}

func getPromotion(offset int) *Promotion {
	prom := Promotion{}
	req, err := http.NewRequest("GET",
		intraURL+
			AuthKey+
			"/user/filter/user"+
			jsonFormat+
			"&location=FR/LYN"+
			"&year=2016"+
			"&course=bachelor/classic"+
			"&active=true"+
			"&promo=tek3"+
			"&offset="+
			strconv.Itoa(offset), nil)
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
