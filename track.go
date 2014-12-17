package ukrpost

import (
	"net/url"
	"strings"
)

const trackEndpoint = `/barcodestatistic/barcodestatistic.asmx`

// CallTrackAPI invokes ukrpost track api method
func (s *Service) CallTrackAPI(meth string, args url.Values, out interface{}) error {
	return s.callAPI(trackEndpoint, meth, args, out)
}

type TrackInfo struct {
	Barcode         string `xml:"barcode"`
	Code            Int    `xml:"code,omitempty"`
	LastOfficeIndex string `xml:"lastofficeindex,omitempty"`
	LastOffice      string `xml:"lastoffice,omitempty"`
	EventDate       Date   `xml:"eventdate"`
	EventDesc       string `xml:"eventdescription"`
}

// Track returns delivery status of postal item with the given barcode
func (s *Service) Track(barcode string) (TrackInfo, error) {
	lang := s.Lang
	if lang == "ua" {
		lang = "uk"
	}
	vals := url.Values{
		"guid":    {s.Guid},
		"barcode": {barcode},
		"culture": {lang},
	}
	var status TrackInfo
	if err := s.CallTrackAPI(`GetBarcodeInfo`, vals, &status); err != nil {
		return TrackInfo{}, err
	}
	status.EventDesc = strings.Trim(status.EventDesc, " \n\r")
	status.LastOfficeIndex = strings.Trim(status.LastOfficeIndex, " ")
	return status, nil
}
