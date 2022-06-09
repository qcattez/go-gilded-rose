package gildedrose_test

import (
	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateQuality", func() {
	It("decreases the quality", func() {
		// Given
		item := gildedrose.Item{Name: "Bread", SellIn: 3, Quality: 5}
		expectedItem := gildedrose.Item{Name: "Bread", SellIn: 2, Quality: 4}

		// When
		gildedrose.UpdateQuality([]*gildedrose.Item{&item})

		// Then
		Expect(item).To(Equal(expectedItem))
	})

	It("decreases twice as fast the quality when the sellIn is negative", func() {
		// Given
		item := gildedrose.Item{Name: "Bread", SellIn: 0, Quality: 5}
		expectedItem := gildedrose.Item{Name: "Bread", SellIn: -1, Quality: 3}

		// When
		gildedrose.UpdateQuality([]*gildedrose.Item{&item})

		// Then
		Expect(item).To(Equal(expectedItem))
	})

	It("never decreases the quality under 0", func() {
		// Given
		item := gildedrose.Item{Name: "Bread", SellIn: 3, Quality: 0}
		expectedItem := gildedrose.Item{Name: "Bread", SellIn: 2, Quality: 0}

		// When
		gildedrose.UpdateQuality([]*gildedrose.Item{&item})

		// Then
		Expect(item).To(Equal(expectedItem))
	})

	It("Age Brie items increase quality as the sellIn decreases", func() {
		// Given
		item := gildedrose.Item{Name: "Aged Brie", SellIn: 3, Quality: 2}
		expectedItem := gildedrose.Item{Name: "Aged Brie", SellIn: 2, Quality: 3}

		// When
		gildedrose.UpdateQuality([]*gildedrose.Item{&item})

		// Then
		Expect(item).To(Equal(expectedItem))
	})

	It("never increases the quality above 50", func() {
		// Given
		item := gildedrose.Item{Name: "Aged Brie", SellIn: 3, Quality: 50}
		expectedItem := gildedrose.Item{Name: "Aged Brie", SellIn: 2, Quality: 50}

		// When
		gildedrose.UpdateQuality([]*gildedrose.Item{&item})

		// Then
		Expect(item).To(Equal(expectedItem))
	})

	It("update the quality twice as fast when sellIn is negative", func() {
		// Given
		agedBrie := gildedrose.Item{Name: "Aged Brie", SellIn: 0, Quality: 30}
		bread := gildedrose.Item{Name: "Bread", SellIn: 0, Quality: 30}
		expectedAgedBrie := gildedrose.Item{Name: "Aged Brie", SellIn: -1, Quality: 32}
		expectedBread := gildedrose.Item{Name: "Bread", SellIn: -1, Quality: 28}

		// When
		gildedrose.UpdateQuality([]*gildedrose.Item{&agedBrie, &bread})

		// Then
		Expect(agedBrie).To(Equal(expectedAgedBrie))
		Expect(bread).To(Equal(expectedBread))
	})

	It("never updates sellIn and quality of Sulfuras", func() {
		// Given
		item := gildedrose.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80}
		expectedItem := item

		// When
		gildedrose.UpdateQuality([]*gildedrose.Item{&item})

		// Then
		Expect(item).To(Equal(expectedItem))
	})

	Context("Backstage passes to a TAFKAL80ETC concert", func() {
		It("increases the quality by 1 of the backstage pass when sellIn is above 10", func() {
			// Given
			item := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 11, Quality: 0}
			expectedItem := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 1}

			// When
			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			// Then
			Expect(item).To(Equal(expectedItem))
		})

		It("increases the quality by 2 of the backstage pass when sellIn is 10", func() {
			// Given
			item := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 0}
			expectedItem := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 9, Quality: 2}

			// When
			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			// Then
			Expect(item).To(Equal(expectedItem))
		})

		It("increases the quality by 2 of the backstage pass when sellIn is above 5 and less than 10", func() {
			// Given
			item := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 6, Quality: 0}
			expectedItem := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 2}

			// When
			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			// Then
			Expect(item).To(Equal(expectedItem))
		})

		It("increases the quality by 3 of the backstage pass when sellIn is 5", func() {
			// Given
			item := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 0}
			expectedItem := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 4, Quality: 3}

			// When
			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			// Then
			Expect(item).To(Equal(expectedItem))
		})

		It("increases the quality by 3 of the backstage pass when sellIn is above 0 and less than 5", func() {
			// Given
			item := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 1, Quality: 0}
			expectedItem := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 3}

			// When
			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			// Then
			Expect(item).To(Equal(expectedItem))
		})

		It("update the quality to 0 of the backstage pass when sellIn is 0", func() {
			// Given
			item := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 0, Quality: 10}
			expectedItem := gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -1, Quality: 0}

			// When
			gildedrose.UpdateQuality([]*gildedrose.Item{&item})

			// Then
			Expect(item).To(Equal(expectedItem))
		})
	})
})
