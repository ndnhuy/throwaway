package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
)

type ProvinceInfo struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	NameEn     string `json:"nameEn"`
	FullName   string `json:"fullName"`
	FullNameEn string `json:"fullNameEn"`
	CodeName   string `json:"codeName"`
}
type GetAllProvincesResponse struct {
	Data []*ProvinceInfo `json:"data"`
}

type DistrictInfo struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	NameEn       string `json:"nameEn"`
	FullName     string `json:"fullName"`
	FullNameEn   string `json:"fullNameEn"`
	CodeName     string `json:"codeName"`
	ProvinceCode string `json:"provinceCode"`
}
type GetDistrictsByProvinceCodeResponse struct {
	Data []*DistrictInfo `json:"data"`
}

func main() {
	convertCSVToJSON()
}

func crawlDists() {
	f, err := os.OpenFile("all.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	provinces := fetchProvinces()
	for _, p := range provinces {
		dists := fetchDistricts(p.Code)
		for _, d := range dists {
			time.Sleep(time.Duration(500) * time.Millisecond)
			fetchLocThenAppendToFile(f, p.Code, p.FullName, p.Name, d.FullName)
		}
	}
}

func convertCSVToJSON() {
	c, err := ioutil.ReadFile("all.txt")
	if err != nil {
		panic(err.Error())
	}
	f, err := os.OpenFile("all-json.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	lines := strings.Split(string(c), "\n")
	lines = lo.Filter(lines, func(l string, _ int) bool {
		return len(l) > 0
	})

	type Record struct {
		Code string `json:"code"`
		City string `json:"city"`
		Dist string `json:"district"`
		Lat  string `json:"lat"`
		Lng  string `json:"lng"`
	}
	validate := func(line string) error {
		vals := strings.Split(line, ",")
		if len(vals) < 5 {
			return errors.New("invalid line: " + line)
		}
		return nil
	}

	toRecord := func(line string, _ int) Record {
		vals := strings.Split(line, ",")
		code := vals[0]
		city := vals[1]
		dist := vals[2]
		lat := vals[3]
		lng := vals[4]

		r := Record{}
		r.Code = code
		r.City = city
		r.Dist = dist
		r.Lat = lat
		r.Lng = lng
		return r
	}
	for _, line := range lines {
		if err := validate(line); err != nil {
			panic(err.Error())
		}
	}

	records := lo.Map(lines, toRecord)
	cityGroups := lo.GroupBy(records, func(r Record) string {
		return r.City
	})

	distGroups := map[string]map[string]Record{}
	for city, records := range cityGroups {
		recordByDist := lo.Associate(records, func(r Record) (string, Record) {
			return r.Dist, r
		})
		distGroups[city] = recordByDist
	}

	j, _ := json.Marshal(distGroups)
	if _, err := f.WriteString(string(j) + "\n"); err != nil {
		panic(err.Error())
	}

	// for _, line := range lines {
	// 	vals := strings.Split(line, ",")
	// 	if len(vals) < 5 {
	// 		continue
	// 	}
	// 	code := vals[0]
	// 	city := vals[1]
	// 	dist := vals[2]
	// 	lat := vals[3]
	// 	lng := vals[4]

	// 	r := Record{}
	// 	r.Code = code
	// 	r.City = city
	// 	r.Dist = dist
	// 	r.Lat = lat
	// 	r.Lng = lng
	// 	j, _ := json.Marshal(r)
	// 	if _, err := f.WriteString(string(j) + "\n"); err != nil {
	// 		panic(err.Error())
	// 	}
	// }
}

func fetchLocThenAppendToFile(f *os.File, code string, cityFullName string, cityName string, dist string) {
	fmt.Println("search for location: [" + code + ", " + dist + ", " + cityFullName + ", " + "]")
	lat, lng, err := fetchLoc(dist + ", " + cityFullName)
	if err != nil {
		appendToFile(f, code, cityName, dist, "", err.Error())
		return
	}
	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lngStr := strconv.FormatFloat(lng, 'f', -1, 64)
	appendToFile(f, code, cityName, dist, latStr, lngStr)
}

func appendToFile(f *os.File, code string, city string, dist string, lat string, lng string) {
	s := fmt.Sprintf("%s,%s,%s,%s,%s\n", code, city, dist, lat, lng)
	if _, err := f.WriteString(s); err != nil {
		panic(err.Error())
	}
}

const AUTH = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJvbGFoLnZuIiwiZW1haWwiOiJtaW5oZGFuZy5vcmcxQG9sYWgudm4iLCJleHAiOjE3MDEwNzIwNjcsImlhdCI6MTY5MzI5NjA2NywiaWF1IjpmYWxzZSwiaXNzIjoiT0wiLCJqdGkiOiIiLCJuYmYiOjAsIm9pZCI6MTgsInNjb3BlcyI6InVzZXIiLCJzdWIiOiIxNjQzNDY4NzIyMTcxNTQ3NjQ4MDMwMSIsInVpZCI6IjE2NDM0Njg3MjIxNzE1NDc2NDgwMzAxIn0.g8LgiphbLrg1Y75ZAcaRcTOOsb_QyP-ynHvjaa6QOEo"

func fetchProvinces() []*ProvinceInfo {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8888/api/sms/provinces", nil)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Authorization", AUTH)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	resp := GetAllProvincesResponse{}
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		panic(err.Error())
	}
	// names := []string{}
	// for _, d := range resp.Data {
	// 	names = append(names, d.FullName)
	// }
	return resp.Data
}
func fetchDistricts(provinceCode string) []*DistrictInfo {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8888/api/sms/provinces/"+provinceCode+"/districts", nil)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Authorization", AUTH)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	resp := GetDistrictsByProvinceCodeResponse{}
	json.Unmarshal(resBody, &resp)
	// distNames := []string{}
	// for _, d := range resp.Data {
	// 	distNames = append(distNames, d.FullName)
	// }
	return resp.Data
}

func fetchLoc(query string) (float64, float64, error) {
	req, err := http.NewRequest(http.MethodGet, "https://maps.googleapis.com/maps/api/place/textsearch/json", nil)
	if err != nil {
		panic(err.Error())
	}
	q := req.URL.Query()
	q.Add("query", query)
	q.Add("key", "AIzaSyCv_E4ERMHktFrkZv6KzCcqvLxLC_vtrpw")
	req.URL.RawQuery = q.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	type Result struct {
		Geometry struct {
			Location struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	}
	type Resp struct {
		Results []Result `json:"results"`
	}

	obj := map[string]interface{}{}
	_ = json.Unmarshal(resBody, &obj)
	results := obj["results"].([]interface{})
	if len(results) == 0 {
		fmt.Println("search error: [" + query + "]")
		return 0, 0, errors.New(string(resBody))
	}
	rs := results[0].(map[string]interface{})
	geometry := rs["geometry"].(map[string]interface{})
	location := geometry["location"].(map[string]interface{})
	lat := location["lat"].(float64)
	lng := location["lng"].(float64)
	return lat, lng, nil
}
