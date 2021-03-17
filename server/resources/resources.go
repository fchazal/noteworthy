package resources

import uuid "github.com/satori/go.uuid"

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

	removeResource(item)
	c.Resources = findAndDelete(c.Resources, item.Id)
}
