package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-production-order-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemComponent *[]dpfm_api_output_formatter.ItemComponent
	var itemComponentStockConfirmation *[]dpfm_api_output_formatter.ItemComponentStockConfirmation
	var itemComponentCosting *[]dpfm_api_output_formatter.ItemComponentCosting
	var itemOperations *[]dpfm_api_output_formatter.ItemOperations
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Item":
			item = c.itemCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemComponent":
			itemComponent = c.itemComponentCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemComponentStockConfirmation":
			itemComponentStockConfirmation = c.itemComponentStockConfirmationCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemComponentCosting":
			itemComponentCosting = c.itemComponentCostingCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemOperations":
			itemOperations = c.itemOperationsCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                         header,
		Item:                           item,
		ItemComponent:                  itemComponent,
		ItemComponentStockConfirmation: itemComponentStockConfirmation,
		ItemComponentCosting:           itemComponentCosting,
		ItemOperations:                 itemOperations,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemComponent *[]dpfm_api_output_formatter.ItemComponent
	var itemComponentStockConfirmation *[]dpfm_api_output_formatter.ItemComponentStockConfirmation
	var itemComponentCosting *[]dpfm_api_output_formatter.ItemComponentCosting
	var itemOperations *[]dpfm_api_output_formatter.ItemOperations
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "Item":
			item = c.itemUpdateSql(mtx, input, output, errs, log)
		case "ItemComponent":
			itemComponent = c.itemComponentUpdateSql(mtx, input, output, errs, log)
		case "ItemComponentStockConfirmation":
			itemComponentStockConfirmation = c.itemComponentStockConfirmationUpdateSql(mtx, input, output, errs, log)
		case "ItemComponentCosting":
			itemComponentCosting = c.itemComponentCostingUpdateSql(mtx, input, output, errs, log)
		case "ItemOperations":
			itemOperations = c.itemOperationsUpdateSql(mtx, input, output, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                         header,
		Item:                           item,
		ItemComponent:                  itemComponent,
		ItemComponentStockConfirmation: itemComponentStockConfirmation,
		ItemComponentCosting:           itemComponentCosting,
		ItemOperations:                 itemOperations,
	}

	return data
}

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID

	headerData := subfuncSDC.Message.Header
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "ProductionOrderHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemData := range *subfuncSDC.Message.Item {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "ProductionOrderItem", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponent {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemComponentData := range *subfuncSDC.Message.ItemComponent {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentData, "function": "ProductionOrderItemComponent", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Component Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentStockConfirmationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentStockConfirmation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemComponentStockConfirmationData := range *subfuncSDC.Message.ItemComponentStockConfirmation {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentStockConfirmationData, "function": "ProductionOrderItemComponentStockConfirmation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Component Stock Confirmation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentStockConfirmationCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentCostingCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentCosting {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemComponentCostingData := range *subfuncSDC.Message.ItemComponentCosting {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentCostingData, "function": "ProductionOrderItemComponentCosting", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Component Costing Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentCostingCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemOperationsCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperations {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemOperationData := range *subfuncSDC.Message.ItemOperations {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemOperationData, "function": "ProductionOrderItemOperations", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Operation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemOperationsCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) headerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	header := input.Header
	headerData := dpfm_api_processing_formatter.ConvertToHeaderUpdates(header)

	sessionID := input.RuntimeSessionID
	if headerIsUpdate(headerData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "ProductionOrderHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Header Data cannot update"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderUpdates(header)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	req := make([]dpfm_api_processing_formatter.ItemUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		itemData := *dpfm_api_processing_formatter.ConvertToItemUpdates(item)

		if itemIsUpdate(&itemData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "ProductionOrderItem", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Item Data cannot update"
				return nil
			}
		}
		req = append(req, itemData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponent {
	req := make([]dpfm_api_processing_formatter.ItemComponentUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemComponent := range item.ItemComponent {
			itemComponentData := *dpfm_api_processing_formatter.ConvertToItemComponentUpdates(itemComponent)

			if itemComponentIsUpdate(&itemComponentData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentData, "function": "ProductionOrderItemComponent", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Component Data cannot update"
					return nil
				}
			}
			req = append(req, itemComponentData)
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentStockConfirmationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentStockConfirmation {
	req := make([]dpfm_api_processing_formatter.ItemComponentStockConfirmationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemComponent := range item.ItemComponent {
			for _, itemComponentStockConfirmation := range itemComponent.ItemComponentStockConfirmation {
				itemComponentStockConfirmationData := *dpfm_api_processing_formatter.ConvertToItemComponentStockConfirmationUpdates(itemComponentStockConfirmation)

				if itemComponentStockConfirmationIsUpdate(&itemComponentStockConfirmationData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentStockConfirmationData, "function": "ProductionOrderItemComponentStockConfirmation", "runtime_session_id": sessionID})
					if err != nil {
						err = xerrors.Errorf("rmq error: %w", err)
						*errs = append(*errs, err)
						return nil
					}
					res.Success()
					if !checkResult(res) {
						output.SQLUpdateResult = getBoolPtr(false)
						output.SQLUpdateError = "Item Component Stock Confirmation Data cannot update"
						return nil
					}
				}
				req = append(req, itemComponentStockConfirmationData)
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentStockConfirmationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentCostingUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentCosting {
	req := make([]dpfm_api_processing_formatter.ItemComponentCostingUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemComponent := range item.ItemComponent {
			for _, itemComponentCosting := range itemComponent.ItemComponentCosting {
				itemComponentCostingData := *dpfm_api_processing_formatter.ConvertToItemComponentCostingUpdates(itemComponentCosting)

				if itemComponentCostingIsUpdate(&itemComponentCostingData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentCostingData, "function": "ProductionOrderItemComponentCosting", "runtime_session_id": sessionID})
					if err != nil {
						err = xerrors.Errorf("rmq error: %w", err)
						*errs = append(*errs, err)
						return nil
					}
					res.Success()
					if !checkResult(res) {
						output.SQLUpdateResult = getBoolPtr(false)
						output.SQLUpdateError = "Item Component Costing Data cannot update"
						return nil
					}
				}
				req = append(req, itemComponentCostingData)
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentCostingUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemOperationsUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperations {
	req := make([]dpfm_api_processing_formatter.ItemOperationsUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemOperation := range item.ItemOperations {
			itemOperationData := *dpfm_api_processing_formatter.ConvertToItemOperationsUpdates(itemOperation)

			if itemOperationIsUpdate(&itemOperationData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemOperationData, "function": "ProductionOrderItemOperations", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Operation Data cannot update"
					return nil
				}
			}
			req = append(req, itemOperationData)
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemOperationsUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func headerIsUpdate(header *dpfm_api_processing_formatter.HeaderUpdates) bool {
	productionOrder := header.ProductionOrder

	return !(productionOrder == 0)
}

func itemIsUpdate(item *dpfm_api_processing_formatter.ItemUpdates) bool {
	productionOrder := item.ProductionOrder
	productionOrderItem := item.ProductionOrderItem

	return !(productionOrder == 0 || productionOrderItem == 0)
}

func itemComponentIsUpdate(itemComponent *dpfm_api_processing_formatter.ItemComponentUpdates) bool {
	productionOrder := itemComponent.ProductionOrder
	productionOrderItem := itemComponent.ProductionOrderItem
	operations := itemComponent.Operations
	operationsItem := itemComponent.OperationsItem
	billOfMaterial := itemComponent.BillOfMaterial
	billOfMaterialItem := itemComponent.BillOfMaterialItem

	return !(productionOrder == 0 || productionOrderItem == 0 || operations == 0 || operationsItem == 0 || billOfMaterial == 0 || billOfMaterialItem == 0)
}

func itemComponentStockConfirmationIsUpdate(itemComponentStockConfirmation *dpfm_api_processing_formatter.ItemComponentStockConfirmationUpdates) bool {
	productionOrder := itemComponentStockConfirmation.ProductionOrder
	productionOrderItem := itemComponentStockConfirmation.ProductionOrderItem
	operations := itemComponentStockConfirmation.Operations
	operationsItem := itemComponentStockConfirmation.OperationsItem
	billOfMaterial := itemComponentStockConfirmation.BillOfMaterial
	billOfMaterialItem := itemComponentStockConfirmation.BillOfMaterialItem

	return !(productionOrder == 0 || productionOrderItem == 0 || operations == 0 || operationsItem == 0 || billOfMaterial == 0 || billOfMaterialItem == 0)
}

func itemComponentCostingIsUpdate(itemComponentCosting *dpfm_api_processing_formatter.ItemComponentCostingUpdates) bool {
	productionOrder := itemComponentCosting.ProductionOrder
	productionOrderItem := itemComponentCosting.ProductionOrderItem
	operations := itemComponentCosting.Operations
	operationsItem := itemComponentCosting.OperationsItem

	return !(productionOrder == 0 || productionOrderItem == 0 || operations == 0 || operationsItem == 0)
}

func itemOperationIsUpdate(itemOperation *dpfm_api_processing_formatter.ItemOperationsUpdates) bool {
	productionOrder := itemOperation.ProductionOrder
	productionOrderItem := itemOperation.ProductionOrderItem
	operations := itemOperation.Operations
	operationsItem := itemOperation.OperationsItem

	return !(productionOrder == 0 || productionOrderItem == 0 || operations == 0 || operationsItem == 0)
}
