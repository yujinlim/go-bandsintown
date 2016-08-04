package model

import "time"

type Event struct {
	Name            string   `json:"name"`
	Title           string   `json:"title"`
	Datetime        DateTime `json:"datetime"`
	TicketUrl       string   `json:"ticket_url"`
	TicketType      string   `json:"ticket_type"`
	TicketStatus    string   `json:"ticket_status"`
	FacebookRSVPUrl string   `json:"facebook_rsvp_url"`
	Description     string   `json:"description"`
	Artists         []Artist `json:"artists"`
	Venue           Venue    `json:"venue"`
}

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	current := string(data[1 : len(data)-1])
	t1, err := time.Parse("2006-01-02T15:04:05", current)

	if err != nil {
		return err
	}

	*t = DateTime{t1}

	return nil
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len("2006-01-02T15:04:05")+2)
	b = append(b, '"')
	b = t.AppendFormat(b, "2006-01-02T15:04:05")
	b = append(b, '"')
	return b, nil
}
