package requests

type ItemOperations struct {
	ProductionOrder                      int      `json:"ProductionOrder"`
	ProductionOrderItem                  int      `json:"ProductionOrderItem"`
	Operations                           int      `json:"Operations"`
	OperationsItem                       int      `json:"OperationsItem"`
	Sequence                             int      `json:"Sequence"`
	OperationsText                       *string  `json:"OperationsText"`
	SequenceText                         *string  `json:"SequenceText"`
	OperationIsReleased                  *bool    `json:"OperationIsReleased"`
	OperationIsPartiallyConfirmed        *bool    `json:"OperationIsPartiallyConfirmed"`
	OperationIsConfirmed                 *bool    `json:"OperationIsConfirmed"`
	OperationIsClosed                    *bool    `json:"OperationIsClosed"`
	OperationIsMarkedForDeletion         *bool    `json:"OperationIsMarkedForDeletion"`
	ProductionPlant                      *string  `json:"ProductionPlant"`
	WorkCenter                           *int     `json:"WorkCenter"`
	OperationErlstSchedldExecStrtDte     *string  `json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme     *string  `json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDate     *string  `json:"OperationErlstSchedldExecEndDate"`
	OperationErlstSchedldExecEndTme      *string  `json:"OperationErlstSchedldExecEndTme"`
	OperationActualExecutionStartDate    *string  `json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime    *string  `json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate      *string  `json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime      *string  `json:"OperationActualExecutionEndTime"`
	ErlstSchedldExecDurnInWorkdays       *string  `json:"ErlstSchedldExecDurnInWorkdays"`
	OperationActualExecutionDays         *string  `json:"OperationActualExecutionDays"`
	OperationUnit                        *string  `json:"OperationUnit"`
	OperationPlannedTotalQuantity        *float32 `json:"OperationPlannedTotalQuantity"`
	OperationTotalConfirmedYieldQuantity *float32 `json:"OperationTotalConfirmedYieldQuantity"`
}
