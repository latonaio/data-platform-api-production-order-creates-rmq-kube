package requests

type ItemOperationCosting struct {
    ProductionOrder	    int	        `json:"ProductionOrder"`
    ProductionOrderItem	int	        `json:"ProductionOrderItem"`
    Operations    	    int	        `json:"Operations"`
    OperationsItem    	int	        `json:"OperationsItem"`
    OperationsID    	int	        `json:"OperationsID"`
    Currency	        string	    `json:"Currency"`
    CostingAmount	    float32	    `json:"CostingAmount"`
    CreationDate	    string	    `json:"CreationDate"`
    CreationTime	    string	    `json:"CreationTime"`
    LastChangeDate	    string	    `json:"LastChangeDate"`
    LastChangeTime	    string	    `json:"LastChangeTime"`
    IsReleased	        *bool	    `json:"IsReleased"`
    IsLocked	        *bool	    `json:"IsLocked"`
    IsCancelled	        *bool   	`json:"IsCancelled"`
    IsMarkedForDeletion	*bool	    `json:"IsMarkedForDeletion"`
}
