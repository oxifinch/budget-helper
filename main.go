package main

import (
	"budget-helper/database"
)

const (
	debug = true
	port  = ":8080"
)

func main() {
	db := database.NewDatabase()
	if debug {
		db.Seed()
	}

	// TESTING: Display output as JSON and save to file.
	// br := budgets.NewBudgetRepo(db)
	// b, err := br.Get(1)
	// if err != nil {
	// 	log.Fatalf("main: %v\n", err)
	// }

	// json, err := json.MarshalIndent(b, "", "    ")
	// if err != nil {
	// 	log.Fatalf("Get: %v\n", err)
	// }
	// fmt.Printf("\nBudget struct: \n")
	// fmt.Printf("%s\n", string(json))
	// err = os.WriteFile("./data-budget.json", json, 0666)
	// if err != nil {
	// 	log.Fatalf("handleDashboard: %v\n", err)
	// }

	server := NewServer(port, db)
	server.Run()
}
