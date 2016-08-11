package model

import "time"

type Event struct {
	ID              int64    `json:"id"`
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
		t1, err = time.Parse(time.RFC3339, current)
		if err != nil {
			return err
		}
	}

	*t = DateTime{t1}

	return nil
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, time.RFC3339)
	b = append(b, '"')
	return b, nil
}
