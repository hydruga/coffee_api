package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var ProductList = []*Product{
	{
		ID:          1,
		Name:        "Drip Coffee",
		Description: "Flavor of the day, by our finest artisinal rosters.",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Beautiful rich crema from freshly roasted beans filled with c02",
	},
}

// Database logic here.
