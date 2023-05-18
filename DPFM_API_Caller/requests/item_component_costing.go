package requests

type ItemComponentCosting struct {
	ProductionOrder     int      `json:"ProductionOrder"`
	ProductionOrderItem int      `json:"ProductionOrderItem"`
	Operations          int      `json:"Operations"`
	OperationsItem      int      `json:"OperationsItem"`
	BillOfMaterial      int      `json:"BillOfMaterial"`
	BillOfMaterialItem  int      `json:"BillOfMaterialItem"`
	ComponentProduct    string   `json:"ComponentProduct"`
	Currency            *string  `json:"Currency"`
	CostingAmount       *float32 `json:"CostingAmount"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}
