package aggregate

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Category) CreateCategory(id int, name string) *Category {
	return &Category{
		Id:   id,
		Name: name,
	}
}

func (c *Category) UpdateCategory(name string) *Category {
	c.Name = name

	return c
}
