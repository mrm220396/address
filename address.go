package address

type Address struct {
	ID int `json:id`
	StreetName string `json:streetname`
	Neighborhood string `json:neighborhood`
	AddressNotes []Notes `json:notes`
}

type Notes struct {
	createdAt time.Time `json:created_when`
	Body string 
}

