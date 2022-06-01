package router

import (
	"budget-helper/database"
	"testing"
)

func TestBudgetCategoriesAllocated(t *testing.T) {
	bc1Payments := []database.Payment{
		{Amount: 200.00},
		{Amount: 100.00},
		{Amount: 300.00},
	}

	bc1 := database.BudgetCategory{
		Allocated: 500.00,
		Payments:  bc1Payments,
	}

	bc2Payments := []database.Payment{
		{Amount: 50.00},
	}

	bc2 := database.BudgetCategory{
		Allocated: 200.00,
		Payments:  bc2Payments,
	}

	budget := database.Budget{
		Allocated: 1000.00,
		BudgetCategories: []database.BudgetCategory{
			bc1,
			bc2,
		},
	}

	tables := struct {
		bcAllocated     float64
		bcSpent         float64
		bufAllocated    float64
		bufSpent        float64
		spentPercent    int
		bc1SpentPercent int
		bc2SpentPercent int
	}{
		bcAllocated:     700.00,
		bcSpent:         650.00,
		bufAllocated:    300.00,
		bufSpent:        100.00,
		spentPercent:    92,
		bc1SpentPercent: 120,
		bc2SpentPercent: 25,
	}

	calcBCAllocated := BudgetCategoriesAllocated(&budget)
	if calcBCAllocated != tables.bcAllocated {
		t.Errorf("calcBCAllocated was incorrect, got: %v, want: %v.\n", calcBCAllocated, tables.bcAllocated)
	}

	calcBCSpent := BudgetCategoriesSpent(&budget)
	if calcBCSpent != tables.bcSpent {
		t.Errorf("calcBCSpent was incorrect, got: %v, want: %v.\n", calcBCSpent, tables.bcSpent)
	}

	calcBufAllocated := BudgetBufferAllocated(&budget)
	if calcBufAllocated != tables.bufAllocated {
		t.Errorf("calcBufAllocated was incorrect, got: %v, want: %v.\n", calcBufAllocated, tables.bufAllocated)
	}

	calcBufSpent := BudgetBufferSpent(&budget)
	if calcBufSpent != tables.bufSpent {
		t.Errorf("calcBufSpent was incorrect, got: %v, want: %v.\n", calcBufSpent, tables.bufSpent)
	}

	calcSpentPercent := BudgetPercentageSpent(&budget)
	if calcSpentPercent != tables.spentPercent {
		t.Errorf("calcSpentPercent was incorrect, got: %v, want: %v.\n", calcSpentPercent, tables.spentPercent)
	}

	calcBC1SpentPercent := BudgetCategoryPercentageSpent(&bc1)
	if calcBC1SpentPercent != tables.bc1SpentPercent {
		t.Errorf("calcBC1SpentPercent was incorrect, got: %v, want: %v.\n", calcBC1SpentPercent, tables.bc1SpentPercent)
	}

	calcBC2SpentPercent := BudgetCategoryPercentageSpent(&bc2)
	if calcBC2SpentPercent != tables.bc2SpentPercent {
		t.Errorf("calcBC2SpentPercent was incorrect, got: %v, want: %v.\n", calcBC2SpentPercent, tables.bc2SpentPercent)
	}

	// TODO: Test BudgetTotalPercentageSpent
}
