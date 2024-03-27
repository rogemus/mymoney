package models

import "time"

type Budget struct {
	ID          int
	Uuid        string
	Created     time.Time
	Description string
	Title       string
}

var budgets = []Budget{
	{
		ID:          1,
		Uuid:        "3d955632-f398-4b80-b009-23707944af1e",
		Created:     time.Now(),
		Description: "ARCHITAX",
		Title:       "Elit dolor cillum elit aute do aliquip esse. Nostrud id eu ut labore eiusmod non. Labore est deserunt duis exercitation anim do id duis consequat incididunt. Labore quis quis laboris mollit ad exercitation proident cupidatat dolore labore. Enim anim deserunt culpa nostrud occaecat cillum labore elit.\r\n",
	},
	{
		ID:          2,
		Uuid:        "44e45856-db0f-428f-aeb4-d2291109c2de",
		Created:     time.Now(),
		Description: "FUTURITY",
		Title:       "In et consequat sit tempor in sit laboris. Qui amet eiusmod minim labore. Ex non do exercitation nisi sunt ipsum. Dolor sint sunt quis officia ea consectetur nulla laboris cillum tempor et magna magna.\r\n",
	},
	{
		ID:          3,
		Uuid:        "abe060c9-faa3-4923-9474-12063fbc90f7",
		Created:     time.Now(),
		Description: "CYTREX",
		Title:       "Cupidatat magna ad non aliquip aliquip in ut adipisicing. Sit sunt occaecat culpa esse occaecat anim cillum duis amet ea veniam. Commodo ad occaecat excepteur duis id amet culpa est quis. Anim excepteur est cillum pariatur incididunt consequat sint veniam sint qui esse occaecat ut amet. Exercitation qui incididunt proident non minim ut adipisicing deserunt ullamco sunt minim enim nisi incididunt. Ad irure duis non exercitation non elit veniam veniam dolor nostrud sint. Eu veniam esse ullamco qui ut aute do et quis.\r\n",
	},
}

func (db *Database) GetBudget(id int) Budget {
	return budgets[1]
}

func (db *Database) GetBudgets() []Budget {
	return budgets
}
