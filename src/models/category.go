package models

type Category struct {
	ID   int
	Name string
}

type CategoryMap map[int]Category

type CategoryList []Category
