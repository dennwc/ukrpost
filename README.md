# ukrpost [![GoDoc](https://godoc.org/github.com/dennwc/ukrpost?status.png)](https://godoc.org/github.com/dennwc/ukrpost)

Go package for [UkrPost](http://ukrposhta.ua/) post tracking and index search API.

### Example usage

    // Create api client with default key
    post := ukrpost.New("")

    // Get status of postal item AA123456789BB
    status, err := post.Track("AA123456789BB")

    // Get info about last post office
    o, err := post.OfficeByIndex(status.LastOfficeIndex)
