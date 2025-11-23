package dto

// OrdersMetadata provides summary statistics about a set of orders, including their total sum and count.
type OrdersMetadata struct {
	TotalSum int64 `json:"totalSum" example:"233700"`
	Count    int   `json:"count" example:"6"`
}
