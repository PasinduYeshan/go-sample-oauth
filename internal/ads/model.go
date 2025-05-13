package ads

type Ad struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Sample data.
var SampleData = []Ad{
	{ID: 1, Title: "Ad 1", Description: "Description for Ad 1", Price: 100.00},
	{ID: 2, Title: "Ad 2", Description: "Description for Ad 2", Price: 200.00},
	{ID: 3, Title: "Ad 3", Description: "Description for Ad 3", Price: 300.00},
}
