package shop

type AgedBrie struct {
	item *Item
}

func NewAgedBrie(item *Item) AgedBrie {
	return AgedBrie{item}
}

func (agedBrie AgedBrie) UpdateQuality() {
	agedBrie.item.SellIn--
	agedBrie.item.increaseQualityBy(1)
}

func (agedBrie AgedBrie) ToItem() *Item {
	return agedBrie.item
}
