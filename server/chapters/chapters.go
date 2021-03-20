package chapters

import (
	"github.com/fchazal/noteworthy/server/notes"
	"github.com/fchazal/noteworthy/server/resources"
	"github.com/fchazal/noteworthy/server/utils"
	uuid "github.com/satori/go.uuid"
)

// CHAPTERS ////////////////////////////////////////////////////////////////////

type Chapter struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Parent string `json:"parent_id"`
	notes.NoteContainer
	resources.ResourceContainer
}

func New(name string, parent_id string) *Chapter {
	return &Chapter{
		uuid.NewV4().String(),
		name,
		utils.ToPath(name),
		"",
		notes.NewContainer(),
		resources.NewContainer(),
	}
}

// ALL CHAPTERS ////////////////////////////////////////////////////////////////

var Chapters *map[string]*Chapter
var Save *func()

func addChapter(c *Chapter) {
	(*Chapters)[c.Id] = c
	(*Save)()
}

func removeChapter(c *Chapter) {
	delete(*Chapters, c.Id)
	(*Save)()
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
	removeChapter(item)
	c.Chapters = utils.RemoveItem(c.Chapters, item.Id)
}
