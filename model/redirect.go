package model

import "time"

type Redirect struct {
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	ShortURL  string    `json:"short_url"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *Redirect) GetDomain() {
	r.Domain = "http://127.0.0.1:9998/"
}
