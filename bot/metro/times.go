package metro

type Time struct {
	Id       string    `json:"id"`
	Pier     string    `json:"pier"`
	Hour     string    `json:"hour"`
	Arrivals []Arrival `json:"arrivals"`
	Dest     string    `json:"dest"`
	Exit     string    `json:"exit"`
	UT       string    `json:"ut"`
}

type Arrival struct {
	Train string `json:"train"`
	Time  string `json:"time"`
}
