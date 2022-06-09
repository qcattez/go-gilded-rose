package gildedrose

type legendary struct {
	item *Item
}

func newLegendary(item *Item) legendary {
	return legendary{item}
}

func (l legendary) updateQuality() {}
