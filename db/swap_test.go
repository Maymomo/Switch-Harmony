package db

import "testing"

func TestSwapTable(t *testing.T) {
	SwapTable(DetailTable, DetailTempTable)
	SwapTable(SummaryTempTable, SummaryTable)
}
