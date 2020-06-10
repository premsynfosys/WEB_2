package ITAssetRep

import (
	"context"
	"encoding/json"
	"fmt"

	CmnModel "github.com/premsynfosys/AMS_WEB/models/CmnModel"
	ITAssetsmodel "github.com/premsynfosys/AMS_WEB/models/ITAssetsmodel"
	utils "github.com/premsynfosys/AMS_WEB/utils"
	//ITAssetRep "github.com/premsynfosys/AMS_WEB/repository/ITAssetRep"
)

// NewAPIRepo retunrs implement of post repository interface
func NewAPIRepo(api string) ITAssetIntrfc {
	return &APIRepo{
		APIConn: api,
	}
}

//APIRepo ..
type APIRepo struct {
	APIConn string
}

//GetEmployeeITAssetsByID ..
func (m *APIRepo) GetEmployeeITAssetsByID(ctx context.Context, EmpID int, isCheckin int) ([]*ITAssetsmodel.ITAssetModel, error) {
	url := fmt.Sprintf(m.APIConn+"/GetEmployeeITAssetsByID?EmpID=%d&isCheckin=%d", EmpID, isCheckin)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetModel, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//CreateITAsset ..
func (m *APIRepo) CreateITAsset(ctx context.Context, mdl *ITAssetsmodel.ITAssetModel) (int64, error) {
	url := fmt.Sprintf(m.APIConn + "/itassets/createitassets")
	j, err := json.Marshal(mdl)
	if err != nil {
		return 0, err
	}
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return 0, err
	}
	var data map[string]int64
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return 0, err
	} else {
		return data["id"], nil
	}
}

//CreateITAssetRequest ..
func (m *APIRepo) CreateITAssetRequest(ctx context.Context, mdl []*ITAssetsmodel.ITAssetRequest) (err error) {
	url := fmt.Sprintf(m.APIConn + "/itassets/CreateRequest")
	j, err := json.Marshal(mdl)
	if err != nil {
		return
	}
	_, err = utils.PostRequest(url, j)
	return
}

//ITAssetRetire ..
func (m *APIRepo) ITAssetRetire(ctx context.Context, mdl *CmnModel.Retire) error {
	url := fmt.Sprintf(m.APIConn + "/itassets/ITAssetRetire")
	j, err := json.Marshal(mdl)
	_, err = utils.PostRequest(url, j)
	return err
}

//BulkCreateITAsset ..
func (m *APIRepo) BulkCreateITAsset(ctx context.Context, mdl []*ITAssetsmodel.ITAssetModel) error {
	url := fmt.Sprintf(m.APIConn + "/itassets/Bulkcreateitassets")
	j, err := json.Marshal(mdl)
	_, err = utils.PostRequest(url, j)
	return err

}

//UpadteITAsset ..
func (m *APIRepo) UpadteITAsset(ctx context.Context, mdl *ITAssetsmodel.ITAssetModel) error {
	url := fmt.Sprintf(m.APIConn + "/itassets/Updateitassets")
	j, err := json.Marshal(mdl)
	_, err = utils.PostRequest(url, j)
	return err
}

//GetITAssets ..
func (m *APIRepo) GetITAssets(ctx context.Context, LocID int) ([]*ITAssetsmodel.ITAssetModel, error) {
	url := fmt.Sprintf(m.APIConn+"/itassets?LocID=%d", LocID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetModel, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetCustomFields ..
func (m *APIRepo) GetCustomFields(ctx context.Context, id int, Mod string) (*ITAssetsmodel.ITAssetModel, error) {
	url := fmt.Sprintf(m.APIConn+"/GetCustomFields/%d/"+Mod, id)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := ITAssetsmodel.ITAssetModel{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

//GetITAssetGroups ..
func (m *APIRepo) GetITAssetGroups(ctx context.Context) ([]*ITAssetsmodel.ITAssetGroup, error) {
	url := fmt.Sprintf(m.APIConn + "/itassets/Groups")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetGroup, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetITassetsFilesByID ..
func (m *APIRepo) GetITassetsFilesByID(ctx context.Context, id int) ([]*ITAssetsmodel.ITassetsFiles, error) {
	url := fmt.Sprintf(m.APIConn+"/GetITassetsFilesByID/%d", id)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITassetsFiles, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetITAssetReqList ..
func (m *APIRepo) GetITAssetReqList(ctx context.Context, ApprvrID int) ([]*ITAssetsmodel.ITAssetReqList, error) {
	url := fmt.Sprintf(m.APIConn+"/itassets/GetITAssetReqList/%d", ApprvrID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetReqList, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetITAssetReqListByReqGroup ..
func (m *APIRepo) GetITAssetReqListByReqGroup(ctx context.Context, ReqGroupID int, ApproverID int) ([]*ITAssetsmodel.ITAssetReqList, error) {
	url := fmt.Sprintf(m.APIConn+"/itassets/GetITAssetReqListByReqGroup/%d/%d", ReqGroupID, ApproverID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetReqList, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//ITAssetReqReject ..
func (m *APIRepo) ITAssetReqReject(ctx context.Context, mdl *ITAssetsmodel.ITAssetRequestApproval) (bool, error) {
	var res bool
	res = false
	url := fmt.Sprintf(m.APIConn + "/itassets/ITAssetReqReject")
	j, err := json.Marshal(mdl)
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(bytes, &res)
	return res, err
}

//ITAssetReqApprove ..
func (m *APIRepo) ITAssetReqApprove(ctx context.Context, mdl []*ITAssetsmodel.ITAssetRequest) (bool, error) {
	var res bool
	res = false
	url := fmt.Sprintf(m.APIConn + "/itassets/ITAssetReqApprove")
	j, err := json.Marshal(mdl)
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(bytes, &res)
	return res, err
}

//CreateITAssetFiles ..
func (m *APIRepo) CreateITAssetFiles(ctx context.Context, mdl *ITAssetsmodel.ITassetsFiles) error {
	url := fmt.Sprintf(m.APIConn + "/CreateITAssetFiles")
	j, err := json.Marshal(mdl)
	_, err = utils.PostRequest(url, j)
	return err

}

//CreateITAssetsCheckout ..
func (m *APIRepo) CreateITAssetsCheckout(ctx context.Context, usr *ITAssetsmodel.ITassetCheckout) error {
	url := fmt.Sprintf(m.APIConn + "/CreateITAssetCheckout")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//CreateITAssetsCheckIn ..
func (m *APIRepo) CreateITAssetsCheckIn(ctx context.Context, usr *ITAssetsmodel.ITassetCheckout) error {
	url := fmt.Sprintf(m.APIConn + "/CreateITAssetsCheckIn")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//CreateITAssetsCheckIn ..
func (m *APIRepo) ITAssetGroups_Create(ctx context.Context, usr *ITAssetsmodel.ITAssetGroup) error {
	url := fmt.Sprintf(m.APIConn + "/ITAssetGroups_Create")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

type iTAssetsBulkEdit struct {
	ITAssetModel ITAssetsmodel.ITAssetModel
	Ids          []string
}

//ITAssetsBulkEdit ..
func (m *APIRepo) ITAssetsBulkEdit(ctx context.Context, usr *ITAssetsmodel.ITAssetModel, idss *[]string) error {
	url := fmt.Sprintf(m.APIConn + "/itassets/ITAssetsBulkEdit")

	mdl := iTAssetsBulkEdit{
		ITAssetModel: *usr,
		Ids:          *idss,
	}
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//ITasset_services_Insert" ..
func (m *APIRepo) ITasset_services_Insert(ctx context.Context, usr *ITAssetsmodel.ITasset_services) error {
	url := fmt.Sprintf(m.APIConn + "/ITasset_services_Insert")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//ITasset_services_start_Update" ..
func (m *APIRepo) ITasset_services_start_Update(ctx context.Context, usr *ITAssetsmodel.ITasset_services) error {
	url := fmt.Sprintf(m.APIConn + "/ITasset_services_start_Update")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//ITasset_services_Extend_Update" ..
func (m *APIRepo) ITasset_services_Extend_Update(ctx context.Context, usr *ITAssetsmodel.ITasset_services) error {
	url := fmt.Sprintf(m.APIConn + "/ITasset_services_Extend_Update")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//ITasset_services_Complete_Update" ..
func (m *APIRepo) ITasset_services_Complete_Update(ctx context.Context, usr *ITAssetsmodel.ITasset_services) error {
	url := fmt.Sprintf(m.APIConn + "/ITasset_services_Complete_Update")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//GetITAssetservices_List ..
func (m *APIRepo) GetITAssetservices_List(ctx context.Context, ITAssetID int) ([]*ITAssetsmodel.ITasset_services, error) {
	url := fmt.Sprintf(m.APIConn+"/GetITAssetservices_List/%d", ITAssetID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITasset_services, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//ITAsset_Service_Request" ..
func (m *APIRepo) ITAsset_Service_Request(ctx context.Context, usr *ITAssetsmodel.ITAsset_service_request) error {
	url := fmt.Sprintf(m.APIConn + "/ITAsset_Service_Request")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

func (m *APIRepo) ITAsset_Service_Request_Resolve(ctx context.Context, usr *ITAssetsmodel.ITAsset_service_request) error {
	url := fmt.Sprintf(m.APIConn + "/ITAsset_Service_Request_Resolve")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//GetITAsset_service_request_List ..
func (m *APIRepo) GetITAsset_service_request_List(ctx context.Context, EmpID int) ([]*ITAssetsmodel.ITAsset_service_request, error) {
	url := fmt.Sprintf(m.APIConn+"/GetITAsset_service_request_List/%d", EmpID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAsset_service_request, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetITAssets ..
func (m *APIRepo) GetITAsset_Retired(ctx context.Context, LocID int, EmpID int) ([]*ITAssetsmodel.ITAssetModel, error) {
	url := fmt.Sprintf(m.APIConn+"/GetITAsset_Retired?LocID=%d&EmpID=%d", LocID, EmpID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetModel, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) Get_ITAssetsHistory_ByAsset(ctx context.Context, AssetID int) ([]*ITAssetsmodel.ITAssetModel, error) {
	url := fmt.Sprintf(m.APIConn+"/Get_ITAssetsHistory_ByAsset/%d", AssetID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetModel, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) ITAssetReq_ApprovalStatusList(ctx context.Context, ReqGroupID int) ([]*ITAssetsmodel.ITAssetRequestApproval, error) {
	url := fmt.Sprintf(m.APIConn+"/ITAssetReq_ApprovalStatusList/%d", ReqGroupID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetRequestApproval, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) ITAssetReqForward(ctx context.Context, mdl *ITAssetsmodel.ITAssetRequestApproval) error {
	url := fmt.Sprintf(m.APIConn + "/ITAssetReqForward")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)

	return err
}

func (m *APIRepo) GetITAssetReqListByEmp(ctx context.Context, EmpID int) ([]*ITAssetsmodel.ITAssetReqList, error) {
	url := fmt.Sprintf(m.APIConn+"/GetITAssetReqListByEmp/%d", EmpID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetReqList, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) ITAssetDelete(ctx context.Context, AssetID int) error {
	url := fmt.Sprintf(m.APIConn+"/ITAssetDelete/%d", AssetID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, nil)
	return nil
}

func (m *APIRepo) GetITAssetToCheckoutToITasset(ctx context.Context, LocID int, AssetID int) ([]*ITAssetsmodel.ITAssetModel, error) {
	url := fmt.Sprintf(m.APIConn+"/GetITAssetToCheckoutToITasset/%d/%d", LocID, AssetID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*ITAssetsmodel.ITAssetModel, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
