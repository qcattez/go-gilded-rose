package shop

type Conjured struct {
	item *Item
}

func NewConjured(item *Item) *Conjured {
	return &Conjured{item}
}

func (conjured *Conjured) UpdateQuality() {
	conjured.item.SellIn--
	conjured.item.decreaseQualityBy(2)
}

func (conjured Conjured) ToItem() *Item {
	return conjured.item
}
