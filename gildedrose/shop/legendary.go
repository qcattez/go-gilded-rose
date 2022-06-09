package shop

type Legendary struct {
	item *Item
}

func NewLegendary(item *Item) Legendary {
	return Legendary{item}
}

func (legendary Legendary) UpdateQuality() {}

func (legendary Legendary) ToItem() *Item {
	return legendary.item
}
