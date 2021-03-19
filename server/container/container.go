package container

type Container struct {
	items []*string
}

type ChapterContainer struct {
	Items []string `json:"chapters"`
}
