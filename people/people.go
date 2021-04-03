package people

import (
	"net/http"
	"openstates-go-sdk/pagination"
	"strconv"
)

type CurrentRole struct {
	Title             string      `json:"title"`
	OrgClassification string      `json:"org_classification"`
	District          interface{} `json:"district"`
	DivisionId        string      `json:"division_id"`
}

type CompactJurisdiction struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Classification string `json:"classification"`
}

type AltIdentifier struct {
	Identifier string `json:"identifier"`
	Scheme     string `json:"scheme"`
}

type AltName struct {
	Name string `json:"name"`
	Note string `json:"note"`
}

type Link struct {
	Url  string `json:"url"`
	Note string `json:"note"`
}

type Office struct {
	Name    string `json:"name"`
	Fax     string `json:"fax"`
	Voice   string `json:"voice"`
	Address string `json:"address"`
}

type Person struct {
	Id               string              `json:"id"`
	Name             string              `json:"name"`
	Party            string              `json:"party"`
	CurrentRole      CurrentRole         `json:"current_role"`
	Jurisdiction     CompactJurisdiction `json:"jurisdiction"`
	GivenName        string              `json:"given_name"`
	FamilyName       string              `json:"family_name"`
	Image            string              `json:"image"`
	Email            string              `json:"email"`
	Gender           string              `json:"gender"`
	BirthDate        string              `json:"birth_date"`
	DeathDate        string              `json:"death_date"`
	Extras           map[string]string   `json:"extras"`
	CreatedAt        string              `json:"created_at"`
	UpdatedAt        string              `json:"updated_at"`
	OpenstatesUrl    string              `json:"openstates_url"`
	OtherIdentifiers []AltIdentifier     `json:"other_identifiers"`
	OtherNames       []AltName           `json:"other_name"`
	Links            []Link              `json:"links"`
	Sources          []Link              `json:"sources"`
	Offices          []Office            `json:"offices"`
}

type PeopleList struct {
	Results    []Person                  `json:"results"`
	Pagination pagination.PaginationMeta `json:"pagination"`
}

const baseUrl = "https://v3.openstates.org"

const peopleEndpoint = "/people"
const peopleGeoEndpoint = "/people.geo"

type PeopleProvider struct {
	apiKey  string
	baseUrl string
}

func NewProvider(apiKey string) PeopleProvider {
	return PeopleProvider{apiKey, baseUrl}
}

type FilterType int

const (
	Jurisdiction FilterType = iota
	Name         FilterType = iota
	District     FilterType = iota
)

func (f FilterType) String() string {
	return [...]string{"jurisdiction", "name", "district"}[f]
}

type OrgClassification int

const (
	Legislature OrgClassification = iota
	Executive   OrgClassification = iota
	Lower       OrgClassification = iota
	Upper       OrgClassification = iota
	Government  OrgClassification = iota
)

func (c OrgClassification) String() string {
	return [...]string{"legislature", "executive", "lower", "upper", "government "}[c]
}

type Inclusion int

const (
	OtherNames       Inclusion = iota
	OtherIdentifiers Inclusion = iota
	Links            Inclusion = iota
)

func (i Inclusion) String() string {
	return [...]string{"other_names", "other_identifiers", "links"}[i]
}

func (p *PeopleProvider) ListPeople(filters map[FilterType]string, ids []string, orgClassification OrgClassification, inclusions []Inclusion, page uint, perPage uint) (PeopleList, error) {
	peopleList := PeopleList{}
	url := p.baseUrl + peopleEndpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return peopleList, err
	}

	q := req.URL.Query()
	q.Add("apikey", p.apiKey)
	for filterType, filterValue := range filters {
		q.Add(filterType.String(), filterValue)
	}
	for _, id := range ids {
		q.Add("id", id)
	}
	q.Add("org_classification", orgClassification.String())

	inclusionMap := make(map[Inclusion]struct{}, 3)
	for _, inc := range inclusions {
		if _, ok := inclusionMap[inc]; ok {
			continue
		}
		inclusionMap[inc] = struct{}{}
		q.Add("include", inc.String())
	}

	q.Add("page", strconv.Itoa(int(page)))
	q.Add("per_page", strconv.Itoa(int(perPage)))

	return peopleList, nil
}

func (p *PeopleProvider) ListPeopleGeo(lat int, lng int, inclusions []Inclusion) (PeopleList, error) {
	peopleList := PeopleList{}
	url := p.baseUrl + peopleGeoEndpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return peopleList, err
	}

	q := req.URL.Query()
	q.Add("apikey", p.apiKey)
	q.Add("lat", strconv.Itoa(lat))
	q.Add("lng", strconv.Itoa(lng))
	inclusionMap := make(map[Inclusion]struct{}, 3)
	for _, inc := range inclusions {
		if _, ok := inclusionMap[inc]; ok {
			continue
		}
		inclusionMap[inc] = struct{}{}
		q.Add("include", inc.String())
	}

	return peopleList, nil
}
