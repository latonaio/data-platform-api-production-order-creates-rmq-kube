package requests

type ComponentCosting struct {
	ProductionOrder               int      `json:"ProductionOrder"`
	ProductionOrderItem           int      `json:"ProductionOrderItem"`
	ProductionOrderSequence       string   `json:"ProductionOrderSequence"`
	ProductionOrderOperation      string   `json:"ProductionOrderOperation"`
	OrderInternalBillOfOperations string   `json:"OrderInternalBillOfOperations"`
	ComponentProduct              string   `json:"ComponentProduct"`
	Currency                      *string  `json:"Currency"`
	CostingAmount                 *float32 `json:"CostingAmount"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}
