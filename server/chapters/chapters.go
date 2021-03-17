package chapters

import (
	"github.com/fchazal/noteworthy/server/notes"
	"github.com/fchazal/noteworthy/server/resources"
	uuid "github.com/satori/go.uuid"
)

// CHAPTERS ////////////////////////////////////////////////////////////////////

type Chapter struct {
	Id   string
	Name string
	notes.NoteContainer
	resources.ResourceContainer
}

func New(name string) *Chapter {
	return &Chapter{
		uuid.NewV4().String(),
		name,
		notes.NewContainer(),
		resources.NewContainer(),
	}
}

// ALL CHAPTERS ////////////////////////////////////////////////////////////////

var Chapters *map[string]*Chapter

func addChapter(c *Chapter) {
	(*Chapters)[c.Id] = c
}

func removeChapter(c *Chapter) {
	delete(*Chapters, c.Id)
}

// CHAPTER CONTAINER ///////////////////////////////////////////////////////////

type ChapterContainer struct {
	Chapters []string `json:"chapters"`
}

func NewContainer() ChapterContainer {
	return ChapterContainer{
		[]string{},
	}
}
func (c *ChapterContainer) AddChapter(item *Chapter) {
	addChapter(item)
	c.Chapters = append(c.Chapters, item.Id)
}

func (c *ChapterContainer) RemoveChapter(item *Chapter) {
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

	removeChapter(item)
	c.Chapters = findAndDelete(c.Chapters, item.Id)
}
