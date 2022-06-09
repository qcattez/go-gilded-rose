package gildedrose

type backstagePass struct {
	item *Item
}

func newBackstagePass(item *Item) backstagePass {
	return backstagePass{item}
}

func (b backstagePass) updateQuality() {
	b.item.SellIn--

	if b.concertIsOver() {
		b.item.Quality = 0
		return
	}

	if b.concertIsHappeningInLessThan5days() {
		increaseQualityBy(b.item, 3)
		return
	}

	if b.concertIsHappeningInLessThan10days() {
		increaseQualityBy(b.item, 2)
		return
	}

	increaseQualityBy(b.item, 1)
}

func (b backstagePass) concertIsHappeningInLessThan10days() bool {
	return b.item.SellIn < 10
}

func (b backstagePass) concertIsHappeningInLessThan5days() bool {
	return b.item.SellIn < 5
}

func (b backstagePass) concertIsOver() bool {
	return b.item.SellIn < 0
}
