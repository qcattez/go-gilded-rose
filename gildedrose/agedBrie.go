package gildedrose

type agedBrie struct {
	item *Item
}

func newAgedBrie(item *Item) agedBrie {
	return agedBrie{item}
}

func (a agedBrie) updateQuality() {
	a.item.SellIn--
	increaseQualityBy(a.item, 1)
}
