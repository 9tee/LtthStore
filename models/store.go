package models

type Store struct {
	CreateDate int64   `json:"create_date" xml:"create_date"`
	UpdateDate int64   `json:"update_date" xml:"update_date"`
	ID         int     `json:"id" xml:"id"`
	Name       string  `json:"name" xml:"name"`
	RateAvg    float32 `json:"rate_avg" xml:"rate_avg"`
	RateOne    float32 `json:"rate_one" xml:"rate_one"`
	RateTwo    float32 `json:"rate_two" xml:"rate_two"`
	RateThree  float32 `json:"rate_three" xml:"rate_three"`
	RateFour   float32 `json:"rate_four" xml:"rate_four"`
	RateFive   float32 `json:"rate_five" xml:"rate_five"`
}
