package merchantservices

type Service struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

// Sample data.
var SampleData = []Service{
	{
		ID:          1,
		Name:        "Car Oil Change",
		Description: "Complete engine oil change and filter replacement",
		Price:       59.99,
		Category:    "Car Repair",
	},
	{
		ID:          2,
		Name:        "Bike Tune-up",
		Description: "Full bicycle tune-up including gear and brake adjustment",
		Price:       45.00,
		Category:    "Bike Repair",
	},
	{
		ID:          3,
		Name:        "Car Brake Repair",
		Description: "Brake pad replacement and system check",
		Price:       129.99,
		Category:    "Car Repair",
	},
}
