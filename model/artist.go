package model

type Artist struct {
  Name                string `json:"name"`
  ImageUrl            string `json:"image_url"`
  ThumbUrl            string `json:"thumb_url"`
  FacebookTourUrl     string `json:"facebook_tour_dates_url"`
  UpcomingEventsCount uint   `json:"upcoming_events_count"`
  TrackerCount        uint   `json:"tracker_count"`
  MBID                string `json:"mbid"`
  URL                 string `json:"url"`
  Website             string `json:"website"`
}
