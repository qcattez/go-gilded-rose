package gildedrose

type common struct {
	item *Item
}

func newCommon(item *Item) common {
	return common{item}
}

func (c common) updateQuality() {
	c.item.SellIn--
	decreaseQualityBy(c.item, 1)
}
