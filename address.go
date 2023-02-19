// Package Address where you can find more information 
// about an address in the city. This package counts with
// Address time where you can write the ID or zipcode
// Notes about the delivery information and the date where 
//  the delivery took place or any explanation about what happened
// 
// METHODS 
// NewAddress will set a json file with more information
// about certain address, 
// 
// WriteNote will receive new information about an address
// collecting the time when it was written and Body for information.

package address

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Types Address & Notes

// Address comments and addresses info
type Address struct {
	ID           int     `json:"id"`
	StreetName   string  `json:"StreetName"`
	Neighborhood string  `json:"neighborhood"`
	AddressNotes []Notes `json:"notes"`
}

// Notes add notes to the address comments
type Notes struct {
	CreatedAt time.Time `json:"created_when"`
	Body      string    `json:"message"`
}


// METHODS

func (addr *Address) NewAddress() {
	content, err := json.Marshal(addr)
	if err != nil {
		fmt.Errorf("Something went wrong %s", err)
	}

	err = os.WriteFile("address.json", content, 0644)
}
