package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-production-order-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *subfuncSDC.Message.Item {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToComponentCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Component, error) {
	components := make([]Component, 0)

	for _, data := range *subfuncSDC.Message.Component {
		component, err := TypeConverter[*Component](data)
		if err != nil {
			return nil, err
		}

		components = append(components, *component)
	}

	return &components, nil
}

func ConvertToComponentStockConfirmationCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ComponentStockConfirmation, error) {
	componentStockConfirmations := make([]ComponentStockConfirmation, 0)

	for _, data := range *subfuncSDC.Message.ComponentStockConfirmation {
		componentStockConfirmation, err := TypeConverter[*ComponentStockConfirmation](data)
		if err != nil {
			return nil, err
		}

		componentStockConfirmations = append(componentStockConfirmations, *componentStockConfirmation)
	}

	return &componentStockConfirmations, nil
}

func ConvertToComponentCostingCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ComponentCosting, error) {
	componentCostings := make([]ComponentCosting, 0)

	for _, data := range *subfuncSDC.Message.ComponentCosting {
		componentCosting, err := TypeConverter[*ComponentCosting](data)
		if err != nil {
			return nil, err
		}

		componentCostings = append(componentCostings, *componentCosting)
	}

	return &componentCostings, nil
}

func ConvertToOperationCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Operation, error) {
	operationes := make([]Operation, 0)

	for _, data := range *subfuncSDC.Message.Operation {
		operation, err := TypeConverter[*Operation](data)
		if err != nil {
			return nil, err
		}

		operationes = append(operationes, *operation)
	}

	return &operationes, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *itemUpdates {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToComponentUpdates(componentUpdates *[]dpfm_api_processing_formatter.ComponentUpdates) (*[]Component, error) {
	components := make([]Component, 0)

	for _, data := range *componentUpdates {
		component, err := TypeConverter[*Component](data)
		if err != nil {
			return nil, err
		}

		components = append(components, *component)
	}

	return &components, nil
}

func ConvertToComponentStockConfirmationUpdates(componentStockConfirmationUpdates *[]dpfm_api_processing_formatter.ComponentStockConfirmationUpdates) (*[]ComponentStockConfirmation, error) {
	componentStockConfirmations := make([]ComponentStockConfirmation, 0)

	for _, data := range *componentStockConfirmationUpdates {
		componentStockConfirmation, err := TypeConverter[*ComponentStockConfirmation](data)
		if err != nil {
			return nil, err
		}

		componentStockConfirmations = append(componentStockConfirmations, *componentStockConfirmation)
	}

	return &componentStockConfirmations, nil
}

func ConvertToComponentCostingUpdates(componentCostingUpdates *[]dpfm_api_processing_formatter.ComponentCostingUpdates) (*[]ComponentCosting, error) {
	componentCostings := make([]ComponentCosting, 0)

	for _, data := range *componentCostingUpdates {
		componentCosting, err := TypeConverter[*ComponentCosting](data)
		if err != nil {
			return nil, err
		}

		componentCostings = append(componentCostings, *componentCosting)
	}

	return &componentCostings, nil
}

func ConvertToOperationUpdates(operationUpdates *[]dpfm_api_processing_formatter.OperationUpdates) (*[]Operation, error) {
	operationes := make([]Operation, 0)

	for _, data := range *operationUpdates {
		operation, err := TypeConverter[*Operation](data)
		if err != nil {
			return nil, err
		}

		operationes = append(operationes, *operation)
	}

	return &operationes, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
