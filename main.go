// bands This is a BandsInTown golang api package that supports getting artist and artist's events
package bands

import (
	"fmt"
	"log"
	"time"

	"github.com/bradfitz/latlong"
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

// Client client struct that stores api key and others properties
type Client struct {
	API_KEY string
}

// parseEvents parse event to grab timezone
func parseEvents(events []model.Event) []model.Event {
	var parsedEvents []model.Event

	for _, event := range events {
		tz := latlong.LookupZoneName(float64(event.Venue.Latitude), float64(event.Venue.Longitude))
		loc, err := time.LoadLocation(tz)
		if err != nil {
			continue
		}
		log.Println(event.Datetime.Time)
		event.Datetime.Time = time.Date(event.Datetime.Time.Year(), event.Datetime.Time.Month(), event.Datetime.Time.Day(), event.Datetime.Time.Hour(), event.Datetime.Time.Minute(), event.Datetime.Time.Second(), event.Datetime.Time.Nanosecond(), loc)
		log.Println(event.Datetime.Time)
		parsedEvents = append(parsedEvents, event)
	}

	return parsedEvents
}

// New create new bandsintown api client
func New(key string) *Client {
	m := Client{key}
	return &m
}

// GetArtist get artist information based on artist name
func (c *Client) GetArtist(name string) (model.Artist, error) {
	var artist model.Artist
	url := fmt.Sprintf("%s/%s?app_id=%s&api_version=%s&format=json", URL, name, c.API_KEY, VERSION)

	if err := get(url, &artist); err != nil {
		return artist, err
	}

	trace("artist %s", artist)

	return artist, nil
}

// GetArtistEvents get artist's events by name
func (c Client) GetArtistEvents(name string) ([]model.Event, error) {
	var events []model.Event
	url := fmt.Sprintf("%s/%s/events?app_id=%s&api_version=%s&format=json", URL, name, c.API_KEY, VERSION)

	if err := get(url, &events); err != nil {
		trace("error %s", err)
		return events, err
	}

	trace("events %d", len(events))
	events = parseEvents(events)

	return events, nil
}
