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
	var component *[]dpfm_api_output_formatter.Component
	var componentStockConfirmation *[]dpfm_api_output_formatter.ComponentStockConfirmation
	var componentCosting *[]dpfm_api_output_formatter.ComponentCosting
	var operation *[]dpfm_api_output_formatter.Operation
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Item":
			item = c.itemCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Component":
			component = c.componentCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ComponentStockConfirmation":
			componentStockConfirmation = c.componentStockConfirmationCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ComponentCosting ":
			componentCosting = c.componentCostingCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Operation":
			operation = c.operationCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                     header,
		Item:                       item,
		Component:                  component,
		ComponentStockConfirmation: componentStockConfirmation,
		ComponentCosting:           componentCosting,
		Operation:                  operation,
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
	var component *[]dpfm_api_output_formatter.Component
	var componentStockConfirmation *[]dpfm_api_output_formatter.ComponentStockConfirmation
	var componentCosting *[]dpfm_api_output_formatter.ComponentCosting
	var operation *[]dpfm_api_output_formatter.Operation
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "Item":
			item = c.itemUpdateSql(mtx, input, output, errs, log)
		case "Component":
			component = c.componentUpdateSql(mtx, input, output, errs, log)
		case "ComponentStockConfirmation":
			componentStockConfirmation = c.componentStockConfirmationUpdateSql(mtx, input, output, errs, log)
		case "ComponentCosting ":
			componentCosting = c.componentCostingUpdateSql(mtx, input, output, errs, log)
		case "Operation":
			operation = c.operationUpdateSql(mtx, input, output, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                     header,
		Item:                       item,
		Component:                  component,
		ComponentStockConfirmation: componentStockConfirmation,
		ComponentCosting:           componentCosting,
		Operation:                  operation,
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
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "ProductionOrderComponent", "runtime_session_id": sessionID})
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

func (c *DPFMAPICaller) componentCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Component {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, componentData := range *subfuncSDC.Message.Component {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": componentData, "function": "ProductionOrderComponent", "runtime_session_id": sessionID})
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

	data, err := dpfm_api_output_formatter.ConvertToComponentCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) componentStockConfirmationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ComponentStockConfirmation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, componentStockConfirmationData := range *subfuncSDC.Message.ComponentStockConfirmation {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": componentStockConfirmationData, "function": "ProductionOrderComponentStockConfirmation", "runtime_session_id": sessionID})
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

	data, err := dpfm_api_output_formatter.ConvertToComponentStockConfirmationCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) componentCostingCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ComponentCosting {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, componentCostingData := range *subfuncSDC.Message.ComponentCosting {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": componentCostingData, "function": "ProductionOrderComponentCosting ", "runtime_session_id": sessionID})
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

	data, err := dpfm_api_output_formatter.ConvertToComponentCostingCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) operationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Operation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, operationData := range *subfuncSDC.Message.Operation {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": operationData, "function": "ProductionOrderOperation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Operation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToOperationCreates(subfuncSDC)
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
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "OrdersHeader", "runtime_session_id": sessionID})
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
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "OrdersItem", "runtime_session_id": sessionID})
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

func (c *DPFMAPICaller) componentUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Component {
	req := make([]dpfm_api_processing_formatter.ComponentUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, component := range item.Component {
			componentData := *dpfm_api_processing_formatter.ConvertToComponentUpdates(component)

			if componentIsUpdate(&componentData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": componentData, "function": "ProductionOrderComponent", "runtime_session_id": sessionID})
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
			req = append(req, componentData)
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToComponentUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) componentStockConfirmationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ComponentStockConfirmation {
	req := make([]dpfm_api_processing_formatter.ComponentStockConfirmationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, component := range item.Component {
			for _, componentStockConfirmation := range component.ComponentStockConfirmation {
				componentStockConfirmationData := *dpfm_api_processing_formatter.ConvertToComponentStockConfirmationUpdates(componentStockConfirmation)

				if componentStockConfirmationIsUpdate(&componentStockConfirmationData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": componentStockConfirmationData, "function": "ProductionOrderComponentStockConfirmation", "runtime_session_id": sessionID})
					if err != nil {
						err = xerrors.Errorf("rmq error: %w", err)
						*errs = append(*errs, err)
						return nil
					}
					res.Success()
					if !checkResult(res) {
						output.SQLUpdateResult = getBoolPtr(false)
						output.SQLUpdateError = "Component Stock Confirmation Data cannot update"
						return nil
					}
				}
				req = append(req, componentStockConfirmationData)
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToComponentStockConfirmationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) componentCostingUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ComponentCosting {
	req := make([]dpfm_api_processing_formatter.ComponentCostingUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, component := range item.Component {
			for _, componentCosting := range component.ComponentCosting {
				componentCostingData := *dpfm_api_processing_formatter.ConvertToComponentCostingUpdates(componentCosting)

				if componentCostingIsUpdate(&componentCostingData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": componentCostingData, "function": "ProductionOrderComponentCosting", "runtime_session_id": sessionID})
					if err != nil {
						err = xerrors.Errorf("rmq error: %w", err)
						*errs = append(*errs, err)
						return nil
					}
					res.Success()
					if !checkResult(res) {
						output.SQLUpdateResult = getBoolPtr(false)
						output.SQLUpdateError = "Component Costing Data cannot update"
						return nil
					}
				}
				req = append(req, componentCostingData)
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToComponentCostingUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) operationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Operation {
	req := make([]dpfm_api_processing_formatter.OperationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, component := range item.Component {
			for _, operation := range component.Operation {
				operationData := *dpfm_api_processing_formatter.ConvertToOperationUpdates(operation)

				if operationIsUpdate(&operationData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": operationData, "function": "ProductionOrderOperation", "runtime_session_id": sessionID})
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
				req = append(req, operationData)
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToOperationUpdates(&req)
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

func componentIsUpdate(component *dpfm_api_processing_formatter.ComponentUpdates) bool {
	productionOrder := component.ProductionOrder
	productionOrderItem := component.ProductionOrderItem
	productionOrderSequence := component.ProductionOrderSequence
	productionOrderOperation := component.ProductionOrderOperation
	orderInternalBillOfOperations := component.OrderInternalBillOfOperations

	return !(productionOrder == 0 || productionOrderItem == 0 || productionOrderSequence == "" || productionOrderOperation == "" || orderInternalBillOfOperations == "")
}

func componentStockConfirmationIsUpdate(componentStockConfirmation *dpfm_api_processing_formatter.ComponentStockConfirmationUpdates) bool {
	productionOrder := componentStockConfirmation.ProductionOrder
	productionOrderItem := componentStockConfirmation.ProductionOrderItem
	productionOrderSequence := componentStockConfirmation.ProductionOrderSequence
	oroductionOrderOperation := componentStockConfirmation.ProductionOrderOperation
	orderInternalBillOfOperations := componentStockConfirmation.OrderInternalBillOfOperations
	componentProduct := componentStockConfirmation.ComponentProduct

	return !(productionOrder == 0 || productionOrderItem == 0 || productionOrderSequence == "" || oroductionOrderOperation == "" || orderInternalBillOfOperations == "" || componentProduct == "")
}

func componentCostingIsUpdate(componentCosting *dpfm_api_processing_formatter.ComponentCostingUpdates) bool {
	productionOrder := componentCosting.ProductionOrder
	productionOrderItem := componentCosting.ProductionOrderItem
	productionOrderSequence := componentCosting.ProductionOrderSequence
	oroductionOrderOperation := componentCosting.ProductionOrderOperation
	orderInternalBillOfOperations := componentCosting.OrderInternalBillOfOperations
	componentProduct := componentCosting.ComponentProduct

	return !(productionOrder == 0 || productionOrderItem == 0 || productionOrderSequence == "" || oroductionOrderOperation == "" || orderInternalBillOfOperations == "" || componentProduct == "")
}

func operationIsUpdate(operation *dpfm_api_processing_formatter.OperationUpdates) bool {
	productionOrder := operation.ProductionOrder
	productionOrderItem := operation.ProductionOrderItem
	productionOrderSequence := operation.ProductionOrderSequence
	oroductionOrderOperation := operation.ProductionOrderOperation
	orderInternalBillOfOperations := operation.OrderInternalBillOfOperations
	orderIntBillOfOperationsItem := operation.OrderIntBillOfOperationsItem

	return !(productionOrder == 0 || productionOrderItem == 0 || productionOrderSequence == "" || oroductionOrderOperation == "" || orderInternalBillOfOperations == "" || orderIntBillOfOperationsItem == 0)
}
