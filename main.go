package main

import (
	"fmt"

	"github.com/ericivander/element-arena/entity"
)

const openQuote = "```"

func main() {
	for i, elementsItem1 := range entity.AllElementsItems {
		for j, elementsItem2 := range entity.AllElementsItems {
			if i >= j || elementsItem1.Conflict(elementsItem2) {
				continue
			}

			for k, elementsItem3 := range entity.AllElementsItems {
				if j >= k || elementsItem1.Conflict(elementsItem3) || elementsItem2.Conflict(elementsItem3) {
					continue
				}

				var elementsItemsCombination = []*entity.ElementsItem{elementsItem1, elementsItem2, elementsItem3}

				fmt.Println(openQuote)
				for _, elementsItem := range elementsItemsCombination {
					for _, upgradedElementsItem := range entity.AllUpgradedElementsItems {
						if upgradedElementsItem.ContainsElementsItem(elementsItem) {
							upgradedElementsItem.Print()
						}
					}
				}
				fmt.Println(openQuote)

				fmt.Println()
			}
		}
	}
}
