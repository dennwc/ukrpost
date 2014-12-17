package ukrpost

import (
	"fmt"
	"net/url"
	"strings"
)

const indexEndpoint = `/ServIndexnp/servindexnp.asmx`

// CallIndexAPI invokes ukrpost index api method
func (s *Service) CallIndexAPI(meth string, args url.Values, out interface{}) error {
	return s.callAPI(indexEndpoint, meth, args, out)
}

// Office represents a post office
type Office struct {
	Index      string `xml:"index"`
	Address    string `xml:"address"`
	Phone      string `xml:"phone"`
	City       string `xml:"city_name"`
	Region     string `xml:"region_name"`
	Province   string `xml:"oblast_name"`
	Filial     string `xml:"postfilial_name"`
	FilialFull string `xml:"postfilial_fullname"`
	Schedule   string `xml:"schedule"`
	Number     string `xml:"number"`
}

// GetUrl returns website url for the post office
func (o Office) GetUrl() string {
	return fmt.Sprintf(baseUrl+"/postindex_new/details.aspx?postfilial=%s", o.Number)
}

// OfficeByIndex searches for post office by post index
func (s *Service) OfficeByIndex(index string) (Office, error) {
	index = strings.TrimSpace(index)
	if index == "" {
		return Office{}, fmt.Errorf("no post index provided")
	}

	vals := url.Values{
		"codeclient": {s.Guid},
		"index":      {index},
	}

	var diff struct {
		Obj Office `xml:"diffgram>NewDataSet>tbVPZ"`
	}
	if err := s.CallIndexAPI(`GetVPZByIndex`, vals, &diff); err != nil {
		return Office{}, err
	}
	if diff.Obj.Index == "" {
		diff.Obj.Index = index
	}
	return diff.Obj, nil
}

type City struct {
	Id       string `xml:"id"`
	Name     string `xml:"city"`
	Region   string `xml:"region"`
	District string `xml:"district"`
}

// CityByIndex returns city by post index
func (s *Service) CityByIndex(index string) (City, error) {
	vals := url.Values{
		"codeclient": {s.Guid},
		"index":      {index},
	}

	var diff struct {
		Obj City `xml:"diffgram>NewDataSet>tbCity"`
	}
	if err := s.CallIndexAPI(`GetCityByIndex`, vals, &diff); err != nil {
		return City{}, err
	}
	return diff.Obj, nil
}

type Region struct {
	Name string `xml:"name"`
}

// Regions lists all regions
func (s *Service) Regions() ([]Region, error) {
	vals := url.Values{
		"codeclient": {s.Guid},
	}

	var diff struct {
		Objs []Region `xml:"diffgram>NewDataSet>tbRegions"`
	}
	if err := s.CallIndexAPI(`GetRegionList`, vals, &diff); err != nil {
		return nil, err
	}
	return diff.Objs, nil
}

type District struct {
	Name string `xml:"region_name"`
}

// DistrictsByRegion lists districts for the region
func (s *Service) DistrictsByRegion(region string, prefix string) ([]District, error) {
	vals := url.Values{
		"codeclient": {s.Guid},
		"region":     {region},
		"prefix":     {prefix},
	}

	var diff struct {
		Objs []District `xml:"diffgram>NewDataSet>tbDistricts"`
	}
	if err := s.CallIndexAPI(`GetDistrictListByRegion`, vals, &diff); err != nil {
		return nil, err
	}
	return diff.Objs, nil
}
