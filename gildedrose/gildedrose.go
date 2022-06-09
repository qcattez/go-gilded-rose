package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

type shopItem struct {
	Name            string
	SellIn, Quality int
}

func NewShopItem(item *Item) shopItem {
	return shopItem{Name: item.Name, SellIn: item.SellIn, Quality: item.Quality}
}

func (s shopItem) saveIn(item *Item) {
	item.Name = s.Name
	item.SellIn = s.SellIn
	item.Quality = s.Quality
}

type Common struct {
	item Item
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		if items[i].Name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
			items[i].SellIn = items[i].SellIn - 1
			if items[i].SellIn < 0 {
				items[i].Quality = 0
			}
			continue
		}

		if items[i].Name == "Aged Brie" {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
			}
			items[i].SellIn = items[i].SellIn - 1
			if items[i].SellIn < 0 {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
			continue
		}

		shopItem := NewShopItem(items[i])

		if shopItem.Quality > 0 {
			shopItem.Quality = shopItem.Quality - 1
		}

		shopItem.SellIn = shopItem.SellIn - 1

		if shopItem.SellIn < 0 {
			if shopItem.Quality > 0 {
				shopItem.Quality = shopItem.Quality - 1
			}
		}

		shopItem.saveIn(items[i])
	}

}
