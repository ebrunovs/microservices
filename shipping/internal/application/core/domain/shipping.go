package domain

type ShippingItem struct {
	ProductCode 	string
	Quantity 		int32
}

type Shipping struct {
	ID 				int64
	OrderID 		int64
	Items			[]ShippingItem
	DeliveryDays	int32
}

func NewShipping(orderID int64, items []ShippingItem) Shipping {
	totalQuantity := int32(0)
	for _, item := range items{
		totalQuantity += item.Quantity
	}

	deliveryDays := int32(1) + (totalQuantity / 5)

	return Shipping{
		OrderID:		orderID,
		Items:			items,
		DeliveryDays:	deliveryDays,
	}
}