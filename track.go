package ukrpost

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"strings"
)

const trackEndpoint = `/barcodestatistic/barcodestatistic.asmx`

// CallTrackAPI invokes ukrpost track api method
func (s *Service) CallTrackAPI(meth string, args url.Values, out interface{}) error {
	resp, err := http.Get(baseUrl + trackEndpoint + "/" + meth + "?" + args.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return xml.NewDecoder(resp.Body).Decode(out)
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
	vals := url.Values{
		"guid":    {s.Guid},
		"barcode": {barcode},
		"culture": {s.Lang},
	}
	var status TrackInfo
	if err := s.CallTrackAPI(`GetBarcodeInfo`, vals, &status); err != nil {
		return TrackInfo{}, err
	}
	status.EventDesc = strings.Trim(status.EventDesc, " \n\r")
	status.LastOfficeIndex = strings.Trim(status.LastOfficeIndex, " ")
	return status, nil
}
