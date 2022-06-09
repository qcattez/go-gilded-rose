package shop

import "strings"

const MinimumQuality = 0
const MaximumQuality = 50
const LegendaryQuality = 80

type Item struct {
	Name            string
	SellIn, Quality int
}

func (item Item) isConjured() bool {
	return strings.Split(item.Name, " ")[0] == "Conjured"
}

func (item Item) isLegendary() bool {
	return item.Quality == LegendaryQuality && !item.isConjured()
}

type UpdatableItem interface {
	UpdateQuality()
	ToItem() *Item
}

func NewUpdatableItem(item *Item) UpdatableItem {
	if item.isLegendary() {
		return NewLegendary(item)
	}
	if item.Name == "Aged Brie" {
		return NewAgedBrie(item)
	}
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		return NewBackstagePass(item)
	}
	if item.isConjured() {
		return NewConjured(item)
	}
	return NewCommon(item)
}

func (item *Item) decreaseQualityBy(quantity int) {
	item.Quality -= quantity

	if item.SellIn < 0 {
		item.Quality -= quantity
	}

	if item.Quality < MinimumQuality {
		item.Quality = MinimumQuality
	}
}

func (item *Item) increaseQualityBy(quantity int) {
	item.Quality += quantity

	if item.SellIn < 0 {
		item.Quality += quantity
	}

	if item.Quality > MaximumQuality {
		item.Quality = MaximumQuality
	}
}
