package resources

import (
	"github.com/fchazal/noteworthy/server/utils"
	uuid "github.com/satori/go.uuid"
)

// RESOURCES ///////////////////////////////////////////////////////////////////

type Resource struct {
	Id   string
	Name string
	Path string
}

func New(name string) *Resource {
	return &Resource{
		uuid.NewV4().String(),
		name,
		"",
	}
}

// ALL RESOURCES ///////////////////////////////////////////////////////////////

var Resources *map[string]*Resource

func addResource(n *Resource) {
	(*Resources)[n.Id] = n
}

func removeResource(n *Resource) {
	delete(*Resources, n.Id)
}

// RESOURCE CONTAINER //////////////////////////////////////////////////////////

type ResourceContainer struct {
	Resources []string `json:"resource"`
}

func NewContainer() ResourceContainer {
	return ResourceContainer{
		[]string{},
	}
}
func (c *ResourceContainer) AddChapter(item *Resource) {
	addResource(item)
	c.Resources = append(c.Resources, item.Id)
}

func (c *ResourceContainer) RemoveChapter(item *Resource) {
	removeResource(item)
	c.Resources = utils.RemoveItem(c.Resources, item.Id)
}
