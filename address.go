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

// Variables and constants

var counter int

// Types Address & Notes

// Address comments and addresses info
type Address struct {
	ID           int     `json:"id"`
	StreetName   string  `json:"StreetName"`
	Housenumber  int     `json:"number"`
	Neighborhood string  `json:"neighborhood"`
	AddressNotes []Notes `json:"notes"`
}

// ViewHistory shows all notes history
func (addr *Address) ViewHistory() {
	for _, i := range addr.AddressNotes {
		fmt.Printf("(%s) : %v\n", i.CreatedAt.Format("2006-01-02 15:04"), i.Body)
	}
}

// NewAddress receives a pointer from Address type
// the whole information must be written already in
// the variable which exports Address or it'll return an error
func (addr *Address) NewAddress() error {
	counter++
	addr.ID = counter
	content, err := json.Marshal(addr)
	if err != nil {
		fmt.Errorf("Something went wrong %s", err)
	}

	addrJsonFileName := fmt.Sprintf("%v%v_address.json", addr.Housenumber, addr.StreetName)

	err = os.WriteFile(addrJsonFileName, content, 0644)
	return err
}

// Notes add notes to the address comments
type Notes struct {
	Order     int       `json:"note_id"`
	CreatedAt time.Time `json:"created_when"`
	Body      string    `json:"message"`
}

// WriteNote receives a pointer of Notes, defines its time and add a message as Boby
func (nt *Notes) WriteNote(message string) *Notes {
	nt.CreatedAt = time.Now()
	nt.Body = message
	return nt
}

// DeleteNote will delete a given note
func DeleteNote(index int, addr Address) Address {
	notes := addr.AddressNotes
	copy(notes[index:], notes[index+1:])
	notes[len(notes)-1] = Notes{}
	notes = notes[:len(notes)-1]

	addr.NewAddress()
	addr.AddressNotes = notes
	return addr
}
