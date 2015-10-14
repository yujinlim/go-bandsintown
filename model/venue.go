package model

type Venue struct {
  Name      string  `json:"name"`
  City      string  `json:"city"`
  Region    string  `json:"region"`
  Country   string  `json:"country"`
  Latitude  float32 `json:"latitude"`
  Longitude float32 `json:"longitude"`
}
