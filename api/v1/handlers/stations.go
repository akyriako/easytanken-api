package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ApiKey         string
	stationsUrl    = "https://creativecommons.tankerkoenig.de/json/list.php?lat=%s&lng=%s&rad=%s&sort=%s&type=%s&apikey=%s"
	stationByIdUrl = "https://creativecommons.tankerkoenig.de/json/detail.php?id=%s&apikey=%s"
)

func GetStationsInProximity(w http.ResponseWriter, r *http.Request) {

	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("lng")
	radius := r.URL.Query().Get("rad")
	sortBy := r.URL.Query().Get("sort")
	fuelType := r.URL.Query().Get("type")

	if len(lat) == 0 || len(long) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(radius) == 0 {
		radius = "4"
	}

	if len(sortBy) == 0 {
		sortBy = "price"
	}

	if len(fuelType) == 0 {
		sortBy = "diesel"
	}

	requestUrl := fmt.Sprintf(stationsUrl, lat, long, radius, sortBy, fuelType, ApiKey)

	responseBody, httpStatusCode, err := executeHttpGetRequest(requestUrl)
	if err != nil {
		w.WriteHeader(httpStatusCode)
		fmt.Fprint(w, err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	fmt.Fprint(w, responseBody)
}

func GetStationById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/stations/"):]

	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestUrl := fmt.Sprintf(stationByIdUrl, id, ApiKey)

	responseBody, httpStatusCode, err := executeHttpGetRequest(requestUrl)
	if err != nil {
		w.WriteHeader(httpStatusCode)
		fmt.Fprint(w, err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	fmt.Fprint(w, responseBody)
}

func executeHttpGetRequest(requestUrl string) (string, int, error) {
	response, err := http.Get(requestUrl)
	if err != nil {
		return "", response.StatusCode, err
	}

	defer response.Body.Close()

	responseBody, err := getResponseBodyAsString(response)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return responseBody, response.StatusCode, nil
}

func getResponseBodyAsString(response *http.Response) (string, error) {
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return replaceLicense(string(responseBody)), nil
}
