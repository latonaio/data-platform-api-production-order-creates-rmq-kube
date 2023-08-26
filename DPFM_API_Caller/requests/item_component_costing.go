package requests

type ItemComponentCosting struct {
	ProductionOrder     int      `json:"ProductionOrder"`
	ProductionOrderItem int      `json:"ProductionOrderItem"`
	BillOfMaterial      int      `json:"BillOfMaterial"`
	BillOfMaterialItem  int      `json:"BillOfMaterialItem"`
	Currency            string 	 `json:"Currency"`
	CostingAmount       float32	 `json:"CostingAmount"`
	CreationDate        string   `json:"CreationDate"`
	CreationTime        string   `json:"CreationTime"`
	LastChangeDate      string   `json:"LastChangeDate"`
	LastChangeTime      string   `json:"LastChangeTime"`
	IsReleased          *bool    `json:"IsReleased"`
	IsLocked            *bool    `json:"IsLocked"`
	IsCancelled         *bool    `json:"IsCancelled"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}
