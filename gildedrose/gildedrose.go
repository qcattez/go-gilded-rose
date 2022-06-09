package gildedrose

import (
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/goblin"
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/shop"
)

func GetShopItemFromGoblinItem(item *goblin.Item) *shop.Item {
	return &shop.Item{
		Name:    item.Name,
		SellIn:  item.SellIn,
		Quality: item.Quality,
	}
}

func GetGoblinItemFromShopItem(item *shop.Item) *goblin.Item {
	return &goblin.Item{
		Name:    item.Name,
		SellIn:  item.SellIn,
		Quality: item.Quality,
	}
}

func UpdateQuality(items []*goblin.Item) {
	for i := 0; i < len(items); i++ {
		shopItem := GetShopItemFromGoblinItem(items[i])

		updatableItem := shop.NewUpdatableItem(shopItem)
		updatableItem.UpdateQuality()

		*items[i] = *GetGoblinItemFromShopItem(updatableItem.ToItem())
	}
}
