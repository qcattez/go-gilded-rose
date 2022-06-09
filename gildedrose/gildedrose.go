package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const MinimumQuality = 0
const MaximumQuality = 50

func decreaseQualityBy(item *Item, quantity int) {
	item.Quality -= quantity

	if item.SellIn < 0 {
		item.Quality -= quantity
	}

	if item.Quality < MinimumQuality {
		item.Quality = MinimumQuality
	}
}

func increaseQualityBy(item *Item, quantity int) {
	item.Quality += quantity

	if item.SellIn < 0 {
		item.Quality += quantity
	}

	if item.Quality > MaximumQuality {
		item.Quality = MaximumQuality
	}
}

type shopItem interface {
	updateQuality()
}

func newShopItem(item *Item) shopItem {
	if item.Name == "Sulfuras, Hand of Ragnaros" {
		return newLegendary(item)
	}
	if item.Name == "Aged Brie" {
		return newAgedBrie(item)
	}
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		return newBackstagePass(item)
	}
	return newCommon(item)
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		shopItem := newShopItem(items[i])
		shopItem.updateQuality()
	}
}
