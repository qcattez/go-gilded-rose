package shop

type BackstagePass struct {
	item *Item
}

func NewBackstagePass(item *Item) BackstagePass {
	return BackstagePass{item}
}

func (backstagePass BackstagePass) UpdateQuality() {
	backstagePass.item.SellIn--

	if backstagePass.isConcertOver() {
		backstagePass.item.Quality = 0
		return
	}

	if backstagePass.isConcertHappeningInLessThan5days() {
		backstagePass.item.increaseQualityBy(3)
		return
	}

	if backstagePass.IsConcertHappeningInLessThan10days() {
		backstagePass.item.increaseQualityBy(2)
		return
	}

	backstagePass.item.increaseQualityBy(1)
}

func (backstagePass BackstagePass) IsConcertHappeningInLessThan10days() bool {
	return backstagePass.item.SellIn < 10
}

func (backstagePass BackstagePass) isConcertHappeningInLessThan5days() bool {
	return backstagePass.item.SellIn < 5
}

func (backstagePass BackstagePass) isConcertOver() bool {
	return backstagePass.item.SellIn < 0
}

func (backstagePass BackstagePass) ToItem() *Item {
	return backstagePass.item
}
