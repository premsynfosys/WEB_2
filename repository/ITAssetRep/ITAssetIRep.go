package ITAssetRep

import (
	"context"
	"github.com/premsynfosys/AMS_WEB/models/CmnModel"
	"github.com/premsynfosys/AMS_WEB/models/ITAssetsmodel"
)

// ITAssetIntrfc explain method def...
type ITAssetIntrfc interface {
	ITAssetDelete(ctx context.Context, AssetID int) (error) 
	GetITAssetReqListByEmp(ctx context.Context, EmpID int) ([]*ITAssetsmodel.ITAssetReqList, error)
	ITAsset_Service_Request_Resolve(ctx context.Context, usr *ITAssetsmodel.ITAsset_service_request) error
	Get_ITAssetsHistory_ByAsset(ctx context.Context, AssetID int) ([]*ITAssetsmodel.ITAssetModel, error)
	ITAssetGroups_Create(ctx context.Context, usr *ITAssetsmodel.ITAssetGroup) error 
	GetEmployeeITAssetsByID(ctx context.Context, EmpID int,isCheckin int) ([]*ITAssetsmodel.ITAssetModel, error)
	CreateITAsset(ctx context.Context, usr *ITAssetsmodel.ITAssetModel) (int64, error)
	BulkCreateITAsset(ctx context.Context, usr []*ITAssetsmodel.ITAssetModel) error
	UpadteITAsset(ctx context.Context, mdl *ITAssetsmodel.ITAssetModel) error
	GetITAssets(ctx context.Context,LocID int) ([]*ITAssetsmodel.ITAssetModel, error)
	CreateITAssetsCheckout(ctx context.Context, usr *ITAssetsmodel.ITassetCheckout) error
	CreateITAssetsCheckIn(ctx context.Context, usr *ITAssetsmodel.ITassetCheckout) error
	GetITassetsFilesByID(ctx context.Context, id int) ([]*ITAssetsmodel.ITassetsFiles, error)
	CreateITAssetFiles(ctx context.Context, mdl *ITAssetsmodel.ITassetsFiles) error
	ITAssetsBulkEdit(ctx context.Context, mdl *ITAssetsmodel.ITAssetModel, ids *[]string) error
	GetCustomFields(ctx context.Context) (*ITAssetsmodel.ITAssetModel, error)
	ITAssetRetire(ctx context.Context, Rtr *CmnModel.Retire) error
	GetITAssetGroups(ctx context.Context) ([]*ITAssetsmodel.ITAssetGroup, error)
	CreateITAssetRequest(ctx context.Context, mdl []*ITAssetsmodel.ITAssetRequest) error
	GetITAssetReqList(ctx context.Context, ApprvrID int) ([]*ITAssetsmodel.ITAssetReqList, error)
	GetITAssetReqListByReqGroup(ctx context.Context, ReqGroupID int,ApproverID int) ([]*ITAssetsmodel.ITAssetReqList, error)
	ITAssetReqReject(ctx context.Context, mdl *ITAssetsmodel.ITAssetRequestApproval) (bool, error)
	ITAssetReqApprove(ctx context.Context, mdl []*ITAssetsmodel.ITAssetRequest) (bool, error)
	
	ITasset_services_Insert(ctx context.Context, usr *ITAssetsmodel.ITasset_services) error 
	ITasset_services_start_Update(ctx context.Context, itm *ITAssetsmodel.ITasset_services) error
	ITasset_services_Complete_Update(ctx context.Context, itm *ITAssetsmodel.ITasset_services) error
	ITasset_services_Extend_Update(ctx context.Context, itm *ITAssetsmodel.ITasset_services) error 
	GetITAssetservices_List(ctx context.Context,ITAssetID int) ([]*ITAssetsmodel.ITasset_services, error)
	ITAsset_Service_Request(ctx context.Context, usr *ITAssetsmodel.ITAsset_service_request) error
	GetITAsset_service_request_List(ctx context.Context, EmpID int) ([]*ITAssetsmodel.ITAsset_service_request, error)
	GetITAsset_Retired(ctx context.Context,LocID int,EmpID int) ([]*ITAssetsmodel.ITAssetModel, error)
	ITAssetReqForward(ctx context.Context, mdl *ITAssetsmodel.ITAssetRequestApproval) error
	ITAssetReq_ApprovalStatusList(ctx context.Context, ReqGroupID int) ([]*ITAssetsmodel.ITAssetRequestApproval, error)
}
