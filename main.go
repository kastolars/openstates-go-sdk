package main

import (
	"fmt"
	"openstates-go-sdk/jurisdictions"
)

func main() {
	apiKey := "9d4f7630-8a15-4334-a677-b9700cb50081"
	p := jurisdictions.NewProvider(apiKey)

	jl, err := p.ListJurisdictions(jurisdictions.State, true, true, 1, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", jl)

	j, err := p.GetJurisdictionDetails("ocd-jurisdiction/country:us/state:tx/place:abilene/government", false, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", j)
}
