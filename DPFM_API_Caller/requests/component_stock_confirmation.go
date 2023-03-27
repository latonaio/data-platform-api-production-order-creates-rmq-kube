package requests

type ComponentStockConfirmation struct {
	ProductionOrder                 int      `json:"ProductionOrder"`
	ProductionOrderItem             int      `json:"ProductionOrderItem"`
	ProductionOrderSequence         string   `json:"ProductionOrderSequence"`
	ProductionOrderOperation        string   `json:"ProductionOrderOperation"`
	OrderInternalBillOfOperations   string   `json:"OrderInternalBillOfOperations"`
	ComponentProduct                string   `json:"ComponentProduct"`
	ComponentProductRequirementDate *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime *string  `json:"ComponentProductRequirementTime"`
	InventoryStockType              *string  `json:"InventoryStockType"`
	InventorySpecialStockType       *string  `json:"InventorySpecialStockType"`
	AvailableProductStock           *float32 `json:"AvailableProductStock"`
	IsMarkedForDeletion             *bool    `json:"IsMarkedForDeletion"`
}
