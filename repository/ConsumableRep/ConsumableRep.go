package ConsumableRep

import (
	//ConsumableModel "github.com/premsynfosys/AMS_WEB/models/ConsumableModel"

	"context"
	"encoding/json"
	"fmt"

	CmnModel "github.com/premsynfosys/AMS_WEB/models/CmnModel"
	"github.com/premsynfosys/AMS_WEB/models/ConsumableModel"
	"github.com/premsynfosys/AMS_WEB/utils"
)

//APIRepo ..
type APIRepo struct {
	APIConn string
}

// NewAPIRepo retunrs implement of post repository interface
func NewAPIRepo(api string) ConsumableIntrfc {
	return &APIRepo{
		APIConn: api,
	}
}

//CreateConsumable ..
func (m *APIRepo) CreateConsumable(ctx context.Context, mdl *ConsumableModel.Consumables) error {
	url := fmt.Sprintf(m.APIConn + "/Consumables/CreateConsumable")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//UpdateConsumable ..
func (m *APIRepo) UpdateConsumable(ctx context.Context, mdl *ConsumableModel.Consumables) error {
	url := fmt.Sprintf(m.APIConn + "/Consumables/UpdateConsumable")
	j, err := json.Marshal(mdl)
	_, err = utils.PostRequest(url, j)
	return err
}

//ConsumableOprtnsAddStock ..
func (m *APIRepo) ConsumableOprtnsAddStock(ctx context.Context, mdl *ConsumableModel.ConsumableOprtns) error {
	url := fmt.Sprintf(m.APIConn + "/Consumables/PostConsumableOprtnsAddStock")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//PostConsumableOprtnsRemovestock ..
func (m *APIRepo) PostConsumableOprtnsRemovestock(ctx context.Context, mdl *ConsumableModel.ConsumableOprtns) error {
	url := fmt.Sprintf(m.APIConn + "/Consumables/PostConsumableOprtnsRemovestock")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//GetConsumableGroups ..
func (m *APIRepo) GetConsumableGroups(ctx context.Context) ([]*ConsumableModel.ConsumableGroup, error) {
	url := fmt.Sprintf(m.APIConn + "/Consumables/GetConsumableGroups")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ConsumableModel.ConsumableGroup, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetConsumableList ..
func (m *APIRepo) GetConsumableList(ctx context.Context, LocID int) ([]*ConsumableModel.Consumables, error) {
	url := fmt.Sprintf(m.APIConn+"/Consumables/GetConsumableList/%d", LocID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ConsumableModel.Consumables, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetConsumableMasterList ..
func (m *APIRepo) GetConsumableMasterList(ctx context.Context) ([]*ConsumableModel.Consumablemaster, error) {
	url := fmt.Sprintf(m.APIConn + "/Consumables/GetConsumableMasterList")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ConsumableModel.Consumablemaster, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetConsumableListByID ..
func (m *APIRepo) GetConsumableListByID(ctx context.Context, IDConsumable int) (*ConsumableModel.Consumables, error) {
	url := fmt.Sprintf(m.APIConn+"/Consumables/GetConsumableListByID/%d", IDConsumable)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := ConsumableModel.Consumables{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

//GetConsumableOprtnListByID ..
func (m *APIRepo) GetConsumableOprtnListByID(ctx context.Context, IDConsumable int) ([]*ConsumableModel.Consumables, error) {
	url := fmt.Sprintf(m.APIConn+"/Consumables/GetConsumableOprtnListByID/%d", IDConsumable)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ConsumableModel.Consumables, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetConsumableOprtnList ..
func (m *APIRepo) GetConsumableOprtnList(ctx context.Context) ([]*ConsumableModel.Consumables, error) {
	url := fmt.Sprintf(m.APIConn + "/Consumables/GetConsumableOprtnList")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ConsumableModel.Consumables, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type consumableBulkEdit struct {
	Consumable ConsumableModel.Consumables
	Ids        []string
}

//ConsumableBulkEdit ..
func (m *APIRepo) ConsumableBulkEdit(ctx context.Context, usr *ConsumableModel.Consumables, idss *[]string) error {
	url := fmt.Sprintf(m.APIConn + "/Consumables/ConsumablesBulkEdit")

	mdl := consumableBulkEdit{
		Consumable: *usr,
		Ids:        *idss,
	}
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//CnsmblRetire ..
func (m *APIRepo) CnsmblBulkRetire(ctx context.Context, mdl []*CmnModel.InWardOutWardAsset) error {
	url := fmt.Sprintf(m.APIConn + "/Consumables/CnsmblBulkRetire")
	j, err := json.Marshal(mdl)
	_, err = utils.PostRequest(url, j)
	return err
}

func (m *APIRepo) CheckDuplicateAssetEntry(ctx context.Context, MasterID int, LocID int) (*int, error) {
	url := fmt.Sprintf(m.APIConn+"/Consumables/CheckDuplicateAssetEntry/%d/%d", MasterID, LocID)

	bytes, err := utils.PutRequest(url, nil)
	if err != nil {
		return nil, err
	}
	var id int
	err = json.Unmarshal(bytes, &id)
	if err != nil {
		return nil, err
	} else {
		return &id, nil
	}

}

func (m *APIRepo) ConsumablemasterInsert(ctx context.Context, mdl *ConsumableModel.Consumablemaster) error {
	url := fmt.Sprintf(m.APIConn + "/ConsumablemasterInsert")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}
func (m *APIRepo) Check_Unique_Consumable(ctx context.Context, ConsumableName string) (*string, error) {
	url := fmt.Sprintf(m.APIConn + "/Check_Unique_Consumable/" + ConsumableName)

	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	var name string
	err = json.Unmarshal(bytes, &name)
	if err != nil {
		return nil, err
	} else {
		return &name, nil
	}

}


func (m *APIRepo) GetVendorsByConsumable(ctx context.Context,ConsumableID int) ([]*CmnModel.VendorsAssetDetails, error) {
	url := fmt.Sprintf(m.APIConn + "/GetVendorsByConsumable/%d",ConsumableID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.VendorsAssetDetails, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) ConsumableBulkDelete(ctx context.Context, idss []string) error {
	url := fmt.Sprintf(m.APIConn + "/ConsumableBulkDelete")
	j, err := json.Marshal(idss)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

func (m *APIRepo) ConsumableDelete(ctx context.Context, AssetID int) (error) {
	url := fmt.Sprintf(m.APIConn+"/ConsumableDelete/%d", AssetID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return  err
	}
	err = json.Unmarshal(bytes, nil)
	return  nil
}

func (m *APIRepo) GetConsumableMastersByVendors(ctx context.Context,VendorID int) ([]*CmnModel.VendorsAssetDetails, error) {
	url := fmt.Sprintf(m.APIConn + "/GetConsumableMastersByVendors/%d",VendorID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.VendorsAssetDetails, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


//BulkCreateConsumables ..
func (m *APIRepo) BulkCreateConsumables(ctx context.Context, mdl []*ConsumableModel.Consumables) error {
	url := fmt.Sprintf(m.APIConn + "/Consumables/BulkCreate")
	j, err := json.Marshal(mdl)
	_, err = utils.PostRequest(url, j)
	return err

}

