package ipinfo

import (
	"encoding/json"
	"net/http"
)

type Info struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
}

func Query(q string) (Info, error) {
	var info Info
	req, err := http.NewRequest("GET", "https://ipinfo.io/"+q, nil)
	if err != nil {
		return info, err
	}
	req.Header.Set("User-Agent", "curl")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return info, err
	}
	err = json.NewDecoder(res.Body).Decode(&info)
	res.Body.Close()
	if err != nil {
		return info, err
	}
	return info, nil
}
