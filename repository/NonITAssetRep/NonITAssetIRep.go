package NonITAssetRep

import (
	"context"
	//"github.com/premsynfosys/AMS_WEB/models/CmnModel"
	"github.com/premsynfosys/AMS_WEB/models/NonITAssets_mdl"
)

// NonITAssetIntrfc explain method def...
type NonITAssetIntrfc interface {
	GetNonITAssetMasterList(ctx context.Context) ([]*NonITAssets_mdl.NonITAssets_Master, error)
	CreateNonITAsset(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets) error
	GetNonITAssetLists(ctx context.Context,LocID int) ([]*NonITAssets_mdl.NonITAssets, error)
	GetNonITAssetList_BY_AssetID(ctx context.Context,AssetID int) (*NonITAssets_mdl.NonITAssets, error)
	PostNonITAssetEdit(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets) error 
	PostNonITAssets_oprtns_AddStock(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_Oprtns) error
	PostNonITAssets_oprtns_Removestock(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_Oprtns) error
	PostNonITAssets_CheckOut(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_checkout_checkin) error
	Check_Unique_NonITAsset(ctx context.Context, NonITAssetName string) (*string, error)
	CheckDuplicateNonITAssetEntry(ctx context.Context, MasterID int, LocID int) (*int, error)
	NonITAssetemasterInsert(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_Master) error 
	GetNonITAssetGroups(ctx context.Context) ([]*NonITAssets_mdl.NonITAssets_Groups, error)
	GetNonITAssetOprtnListByID(ctx context.Context, IDNonITAsset int) ([]*NonITAssets_mdl.NonITAssets, error)
	CreateNonITAssetRequest(ctx context.Context, mdl []*NonITAssets_mdl.NonITAssetRequest) error
	GetNonITAssetReqList(ctx context.Context, Apprvrid int) ([]*NonITAssets_mdl.NonITAssetReqList, error) 
	NonITAssetReq_ApprovalStatusList(ctx context.Context, ReqGroupID int) ([]*NonITAssets_mdl.NonITAssetRequestApproval, error)
	NonITAssetReqReject(ctx context.Context, mdl *NonITAssets_mdl.NonITAssetRequestApproval) (bool, error)
	NonITAssetReqForward(ctx context.Context, mdl *NonITAssets_mdl.NonITAssetRequestApproval) error
	NonITAssetReqApprove(ctx context.Context, mdl []*NonITAssets_mdl.NonITAssetRequest) error
	NonITAssetReqListByReqGroup(ctx context.Context, ReqGroupID int,ApproverID int) ([]*NonITAssets_mdl.NonITAssetReqList, error)
	GetNonITAssetCheckinDetails(ctx context.Context, LocID int) ([]*NonITAssets_mdl.NonITAssets_checkout_checkin, error)
	NonITAssetCheckin(ctx context.Context, mdl *NonITAssets_mdl.NonITAssets_checkin) error
	Getnonitassets_checkinByID(ctx context.Context, CheckinID int) ([]*NonITAssets_mdl.NonITAssets_checkin, error)
	GetNonITAssetCheckinDetailsByEmp(ctx context.Context, EmpID int) ([]*NonITAssets_mdl.NonITAssets_checkout_checkin, error)
	GetNonITAssetCheckinDetailsByAsset(ctx context.Context, IDNonITAsset int) ([]*NonITAssets_mdl.NonITAssets_checkout_checkin, error) 
}
