package models

type Trademark struct {
	ID   int
	Name string
}

type TrademarkMap map[int]Trademark

type TrademarkList []Trademark
