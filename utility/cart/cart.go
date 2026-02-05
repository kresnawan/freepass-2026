package cart

import (
	"canteen/internal/models"
)

func MergeCartItem(items []models.CartItem) []models.CartItem {
	mergedMap := make(map[int]int)

	for _, item := range items {
		mergedMap[item.MenuId] += item.Quantity
	}

	var finalItems []models.CartItem
	for id, totalQty := range mergedMap {
		finalItems = append(finalItems, models.CartItem{
			MenuId:   id,
			Quantity: totalQty,
		})
	}

	return finalItems
}
