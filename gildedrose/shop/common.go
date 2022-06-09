package shop

type Common struct {
	item *Item
}

func NewCommon(item *Item) *Common {
	return &Common{item}
}

func (common *Common) UpdateQuality() {
	common.item.SellIn--
	common.item.decreaseQualityBy(1)
}

func (common Common) ToItem() *Item {
	return common.item
}
