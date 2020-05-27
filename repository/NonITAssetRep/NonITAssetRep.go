package NonITAssetRep

import (
	"context"
	"encoding/json"
	"fmt"

	//"github.com/premsynfosys/AMS_WEB/models/CmnModel"
	"github.com/premsynfosys/AMS_WEB/models/NonITAssets_mdl"
	"github.com/premsynfosys/AMS_WEB/utils"
	//ITAssetRep "github.com/premsynfosys/AMS_WEB/repository/ITAssetRep"
)

// NewAPIRepo retunrs implement of post repository interface
func NewAPIRepo(api string) NonITAssetIntrfc {
	return &APIRepo{
		APIConn: api,
	}
}

//APIRepo ..
type APIRepo struct {
	APIConn string
}

//GetNonITAssetMasterList ..
func (m *APIRepo) GetNonITAssetMasterList(ctx context.Context) ([]*NonITAssets_mdl.NonITAssets_Master, error) {
	url := fmt.Sprintf(m.APIConn + "/GetNonITAssetMasterLists")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*NonITAssets_mdl.NonITAssets_Master, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetNonITAssetLists ..
func (m *APIRepo) GetNonITAssetLists(ctx context.Context,LocID int) ([]*NonITAssets_mdl.NonITAssets, error) {
	url := fmt.Sprintf(m.APIConn + "/GetNonITAssetLists/%d",LocID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*NonITAssets_mdl.NonITAssets, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//CreateNonITAsset ..
func (m *APIRepo) CreateNonITAsset(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets) error {
	url := fmt.Sprintf(m.APIConn + "/CreateNonITAsset")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//GetNonITAssetListTS_BY_AssetID ..
func (m *APIRepo) GetNonITAssetList_BY_AssetID(ctx context.Context, AssetID int) (*NonITAssets_mdl.NonITAssets, error) {

	url := fmt.Sprintf(m.APIConn+"/GetNonITAssetList_BY_AssetID/%d", AssetID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := NonITAssets_mdl.NonITAssets{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

//PostNonITAssets_oprtns_AddStock ..
func (m *APIRepo) PostNonITAssets_oprtns_AddStock(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_Oprtns) error {
	url := fmt.Sprintf(m.APIConn + "/PostNonITAssets_oprtns_AddStock")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//PostNonITAssets_oprtns_Removestock ..
func (m *APIRepo) PostNonITAssets_oprtns_Removestock(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_Oprtns) error {
	url := fmt.Sprintf(m.APIConn + "/PostNonITAssets_oprtns_Removestock")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//PostNonITAssetEdit ..
func (m *APIRepo) PostNonITAssetEdit(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets) error {
	url := fmt.Sprintf(m.APIConn + "/PostNonITAssetEdit")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//PostNonITAssets_CheckOut ..
func (m *APIRepo) PostNonITAssets_CheckOut(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_checkout_checkin) error {
	url := fmt.Sprintf(m.APIConn + "/PostNonITAssets_CheckOut")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

func (m *APIRepo) Check_Unique_NonITAsset(ctx context.Context, NonITAssetName string) (*string, error) {
	url := fmt.Sprintf(m.APIConn + "/Check_Unique_NonITAsset/" + NonITAssetName)

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

func (m *APIRepo) CheckDuplicateNonITAssetEntry(ctx context.Context, MasterID int, LocID int) (*int, error) {
	url := fmt.Sprintf(m.APIConn+"/CheckDuplicateNonITAssetEntry/%d/%d", MasterID, LocID)

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

func (m *APIRepo) NonITAssetemasterInsert(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_Master) error {
	url := fmt.Sprintf(m.APIConn + "/NonITAssetemasterInsert")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

func (m *APIRepo) GetNonITAssetGroups(ctx context.Context) ([]*NonITAssets_mdl.NonITAssets_Groups, error) {

	url := fmt.Sprintf(m.APIConn + "/GetNonITAssetGroups")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*NonITAssets_mdl.NonITAssets_Groups{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) GetNonITAssetOprtnListByID(ctx context.Context, IDNonITAsset int) ([]*NonITAssets_mdl.NonITAssets, error) {
	url := fmt.Sprintf(m.APIConn+"/GetNonITAssetOprtnListByID/%d", IDNonITAsset)

	bytes, err := utils.PutRequest(url, nil)
	if err != nil {
		return nil, err
	}
	data := []*NonITAssets_mdl.NonITAssets{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}

}

func (m *APIRepo) CreateNonITAssetRequest(ctx context.Context, mdl []*NonITAssets_mdl.NonITAssetRequest) error {
	url := fmt.Sprintf(m.APIConn + "/CreateNonITAssetRequest")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

func (m *APIRepo) GetNonITAssetReqList(ctx context.Context, Apprvrid int) ([]*NonITAssets_mdl.NonITAssetReqList, error) {
	url := fmt.Sprintf(m.APIConn+"/GetNonITAssetReqList/%d", Apprvrid)
	bytes, err := utils.PutRequest(url, nil)
	if err != nil {
		return nil, err
	}
	data := []*NonITAssets_mdl.NonITAssetReqList{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}

}
func (m *APIRepo) NonITAssetReq_ApprovalStatusList(ctx context.Context, ReqGroupID int) ([]*NonITAssets_mdl.NonITAssetRequestApproval, error) {
	url := fmt.Sprintf(m.APIConn+"/NonITAssetReq_ApprovalStatusList/%d", ReqGroupID)
	bytes, err := utils.PutRequest(url, nil)
	if err != nil {
		return nil, err
	}
	data := []*NonITAssets_mdl.NonITAssetRequestApproval{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}

}





func (m *APIRepo) NonITAssetReqReject(ctx context.Context, mdl *NonITAssets_mdl.NonITAssetRequestApproval) (bool, error) {
	var res bool
	res = false
	url := fmt.Sprintf(m.APIConn + "/NonITAssetReqReject")
	j, err := json.Marshal(mdl)
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(bytes, &res)
	return res, err
}


//NonITAssetReqForward ..
func (m *APIRepo) NonITAssetReqForward(ctx context.Context, mdl *NonITAssets_mdl.NonITAssetRequestApproval) error {
	url := fmt.Sprintf(m.APIConn + "/NonITAssetReqForward")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)

	return err
}

func (m *APIRepo) NonITAssetReqListByReqGroup(ctx context.Context, ReqGroupID int, ApproverID int) ([]*NonITAssets_mdl.NonITAssetReqList, error) {
	url := fmt.Sprintf(m.APIConn+"/NonITAssetReqListByReqGroup/%d/%d", ReqGroupID, ApproverID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*NonITAssets_mdl.NonITAssetReqList, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func (m *APIRepo) NonITAssetReqApprove(ctx context.Context, mdl []*NonITAssets_mdl.NonITAssetRequest) error {
	url := fmt.Sprintf(m.APIConn + "/NonITAssetReqApprove")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)

	return err
}


func (m *APIRepo) GetNonITAssetCheckinDetails(ctx context.Context, LocID int) ([]*NonITAssets_mdl.NonITAssets_checkout_checkin, error) {
	url := fmt.Sprintf(m.APIConn+"/GetNonITAssetCheckinDetails/%d", LocID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*NonITAssets_mdl.NonITAssets_checkout_checkin, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) NonITAssetCheckin(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_checkin) error {
	url := fmt.Sprintf(m.APIConn + "/NonITAssetCheckin")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)

	return err
}

func (m *APIRepo) Getnonitassets_checkinByID(ctx context.Context, CheckinID int) ([]*NonITAssets_mdl.NonITAssets_checkin, error) {
	url := fmt.Sprintf(m.APIConn+"/Getnonitassets_checkinByID/%d", CheckinID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*NonITAssets_mdl.NonITAssets_checkin, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) GetNonITAssetCheckinDetailsByAsset(ctx context.Context, IDNonITAsset int) ([]*NonITAssets_mdl.NonITAssets_checkout_checkin, error) {
	url := fmt.Sprintf(m.APIConn+"/GetNonITAssetCheckinDetailsByAsset/%d", IDNonITAsset)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*NonITAssets_mdl.NonITAssets_checkout_checkin, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) GetNonITAssetCheckinDetailsByEmp(ctx context.Context, EmpID int) ([]*NonITAssets_mdl.NonITAssets_checkout_checkin, error) {
	url := fmt.Sprintf(m.APIConn+"/GetNonITAssetCheckinDetailsByEmp/%d", EmpID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*NonITAssets_mdl.NonITAssets_checkout_checkin, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func (m *APIRepo) GetNonITAssetReqListByEmp(ctx context.Context, EmpID int) ([]*NonITAssets_mdl.NonITAssetReqList, error) {
	url := fmt.Sprintf(m.APIConn+"/GetNonITAssetReqListByEmp/%d", EmpID)
	bytes, err := utils.PutRequest(url, nil)
	if err != nil {
		return nil, err
	}
	data := []*NonITAssets_mdl.NonITAssetReqList{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}

}