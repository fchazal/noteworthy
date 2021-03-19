package notes

import (
	"time"

	"github.com/fchazal/noteworthy/server/resources"
	"github.com/fchazal/noteworthy/server/utils"
	uuid "github.com/satori/go.uuid"
)

// NOTE ////////////////////////////////////////////////////////////////////////

type Note struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Path    string `json:"-"`
	Parent  string `json:"parent_id"`
	Created time.Time
	Updated time.Time
	resources.ResourceContainer
}

func New(name string) *Note {
	return &Note{
		uuid.NewV4().String(),
		name,
		"",
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
	removeNote(item)
	c.Notes = utils.RemoveItem(c.Notes, item.Id)
}
