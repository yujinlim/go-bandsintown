// This is a BandsInTown golang api package that supports getting artist and artist's events
package bands

import (
	"fmt"
	"github.com/tj/go-debug"
	"github.com/yujinlim/go-bandsintown/model"
)

var trace = debug.Debug("bands:log")

const (
	API_ROUTE   = "http://api.bandsintown.com/"
	ARTIST_PATH = "artists"
	EVENTS_PATH = "events"
	VERSION     = "2.0"
	URL         = API_ROUTE + ARTIST_PATH
)

type ArtistApi interface {
	GetArtist() model.Artist
	GetArtistEvents() []model.Event
}

type Client struct {
	API_KEY string
}

// create new bandsintown api client
func New(key string) *Client {
	m := Client{key}
	return &m
}

// get artist information based on artist name
func (c *Client) GetArtist(name string) (model.Artist, error) {
	var artist model.Artist
	url := fmt.Sprintf("%s/%s?app_id=%s&api_version=%s&format=json", URL, name, c.API_KEY, VERSION)
	if err := get(url, &artist); err != nil {
		return artist, err
	}

	trace("artist %s", artist)

	return artist, nil
}

// get artist's events by name
func (c Client) GetArtistEvents(name string) ([]model.Event, error) {
	var events []model.Event
	url := fmt.Sprintf("%s/%s/events?app_id=%s&api_version=%s&format=json&date=all", URL, name, c.API_KEY, VERSION)

	if err := get(url, &events); err != nil {
		trace("error %s", err)
		return events, err
	}

	trace("events %d", len(events))

	return events, nil
}
