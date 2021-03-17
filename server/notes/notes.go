package notes

import (
	"time"

	"github.com/fchazal/noteworthy/server/resources"
	uuid "github.com/satori/go.uuid"
)

// NOTE ////////////////////////////////////////////////////////////////////////

type Note struct {
	Id      string
	Name    string
	Content string
	Created time.Time
	Updated time.Time
	resources.ResourceContainer
}

func New(name string) *Note {
	return &Note{
		uuid.NewV4().String(),
		name,
		"",
		time.Time{},
		time.Time{},
		resources.NewContainer(),
	}
}

// ALL NOTES ///////////////////////////////////////////////////////////////////

var Notes *map[string]*Note

func addNote(n *Note) {
	(*Notes)[n.Id] = n
}

func removeNote(n *Note) {
	delete(*Notes, n.Id)
}

// NOTE CONTAINER //////////////////////////////////////////////////////////////

type NoteContainer struct {
	Notes []string `json:"notes"`
}

func NewContainer() NoteContainer {
	return NoteContainer{
		[]string{},
	}
}
func (c *NoteContainer) AddChapter(item *Note) {
	addNote(item)
	c.Notes = append(c.Notes, item.Id)
}

func (c *NoteContainer) RemoveChapter(item *Note) {
	findAndDelete := func(s []string, e string) []string {
		x := 0
		for _, i := range s {
			if i != e {
				s[x] = i
				x++
			}
		}
		return s[:x]
	}

	removeNote(item)
	c.Notes = findAndDelete(c.Notes, item.Id)
}
