package jurisdictions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type JurisdictionClassification string

const (
	State        JurisdictionClassification = "state"
	Municipality JurisdictionClassification = "municipality"
	Country      JurisdictionClassification = "country"
)

type Post struct {
	Label              string `json:"label"`
	Role               string `json:"role"`
	DivisionId         string `json:"division_id"`
	MaximumMemberships int    `json:"maximum_memberships"`
}

type Chamber struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Classification string `json:"classification"`
	Districts      []Post `json:"districts"`
}

type LegislativeSession struct {
	Identifier     string `json:"identifier"`
	Name           string `json:"name"`
	Classification string `json:"classification"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
}

type Jurisdiction struct {
	Id                  string                     `json:"id"`
	Name                string                     `json:"name"`
	Classification      JurisdictionClassification `json:"classification"`
	DivisionId          string                     `json:"division_id"`
	Url                 string                     `json:"url"`
	Organizations       []Chamber                  `json:"organizations"`
	LegislativeSessions []LegislativeSession       `json:"legislative_sessions"`
}

type PaginationMeta struct {
	PerPage    int `json:"per_page"`
	Page       int `json:"page"`
	MaxPage    int `json:"max_page"`
	TotalItems int `json:"total_items"`
}

type JurisdictionList struct {
	Results    []Jurisdiction `json:"results"`
	Pagination PaginationMeta `json:"pagination"`
}

const baseUrl = "https://v3.openstates.org"

const jurisdictionsEndpoint = "/jurisdictions"

type JurisdictionProvider struct {
	apiKey  string
	baseUrl string
}

func NewProvider(apiKey string) JurisdictionProvider {
	return JurisdictionProvider{apiKey, baseUrl}
}

func (p *JurisdictionProvider) ListJurisdictions(classification JurisdictionClassification, includeOrganizations bool, includeLegislativeSessions bool, page int, perPage int) (JurisdictionList, error) {
	jurisdictionList := JurisdictionList{}
	url := p.baseUrl + jurisdictionsEndpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return jurisdictionList, err
	}

	q := req.URL.Query()
	q.Add("apikey", p.apiKey)
	q.Add("classification", string(classification))
	if includeOrganizations {
		q.Add("include", "organizations")
	}
	if includeLegislativeSessions {
		q.Add("include", "legislative_sessions")
	}
	q.Add("page", strconv.Itoa(page))
	q.Add("per_page", strconv.Itoa(perPage))
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return jurisdictionList, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return jurisdictionList, err
	}

	err = json.Unmarshal(body, &jurisdictionList)
	if err != nil {
		return jurisdictionList, err
	}

	return jurisdictionList, nil
}

func (p *JurisdictionProvider) GetJurisdictionDetails(jurisdictionId string, includeOrganizations bool, includeLegislativeSessions bool) (Jurisdiction, error) {
	jurisdiction := Jurisdiction{}
	url := p.baseUrl + jurisdictionsEndpoint + "/" + jurisdictionId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return jurisdiction, err
	}
	q := req.URL.Query()
	q.Add("apikey", p.apiKey)
	if includeOrganizations {
		q.Add("include", "organizations")
	}
	if includeLegislativeSessions {
		q.Add("include", "legislative_sessions")
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return jurisdiction, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return jurisdiction, err
	}

	err = json.Unmarshal(body, &jurisdiction)
	if err != nil {
		return jurisdiction, err
	}

	return jurisdiction, nil
}
