package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/premsynfosys/AMS_WEB/handler/cmnhandler"
	"github.com/premsynfosys/AMS_WEB/handler/cnsmblhndlr"
	"github.com/premsynfosys/AMS_WEB/handler/itassethndlr"
	"github.com/premsynfosys/AMS_WEB/handler/nonitassetshndlr"
	"github.com/premsynfosys/AMS_WEB/utils"
)

//ConsumablesRoutings ..s
func ConsumablesRoutings(r *mux.Router, hndl *cnsmblhndlr.IConsumables) {
	r.HandleFunc("/GetConsumableMastersByVendors/{VendorID}", utils.AuthRequired(hndl.GetConsumableMastersByVendors))
	r.HandleFunc("/ConsumableDelete/{AssetID}", utils.AuthRequired(hndl.ConsumableDelete))
	r.HandleFunc("/ConsumableBulkDelete", utils.AuthRequired(hndl.ConsumableBulkDelete))
	r.HandleFunc("/Consumables/Consumables", utils.AuthRequired(hndl.Consumable))
	r.HandleFunc("/GetVendorsByConsumable/{ConsumableID}", utils.AuthRequired(hndl.GetVendorsByConsumable))
	r.HandleFunc("/Consumables/ConsumableEdit/{IDConsumables}", utils.AuthRequired(hndl.ConsumableEdit))
	r.HandleFunc("/Consumables/PostCreateConsumable", utils.AuthRequired(hndl.PostCreateConsumable))
	r.HandleFunc("/Consumables/PostConsumableOprtnsAddStock", utils.AuthRequired(hndl.PostConsumableOprtnsAddStock))
	r.HandleFunc("/Consumables/PostConsumableOprtnsRemovestock", utils.AuthRequired(hndl.PostConsumableOprtnsRemovestock))
	r.HandleFunc("/Consumables/GetConsumableList/{LocID}", utils.AuthRequired(hndl.GetConsumableList))
	r.HandleFunc("/Consumables/GetConsumableMasterList", utils.AuthRequired(hndl.GetConsumableMasterList))
	r.HandleFunc("/Consumables/ConsumableList", utils.AuthRequired(hndl.ConsumableList))
	r.HandleFunc("/Consumables/ConsumableView/{IDConsumables}", utils.AuthRequired(hndl.ConsumableView))
	r.HandleFunc("/Consumables/GetConsumableListByID/{IDConsumables}", utils.AuthRequired(hndl.GetConsumableListByID))
	r.HandleFunc("/Consumables/GetConsumableOprtnListByID/{IDConsumables}", utils.AuthRequired(hndl.GetConsumableOprtnListByID))
	r.HandleFunc("/Consumables/ConsumableOprtnListByID/{IDConsumables}", utils.AuthRequired(hndl.ConsumableOprtnListByID))
	r.HandleFunc("/Consumables/ConsumableBulkEdit", utils.AuthRequired(hndl.ConsumableBulkEdit))

	r.HandleFunc("/Consumables/GetConsumableOprtnList", utils.AuthRequired(hndl.GetConsumableOprtnList))
	r.HandleFunc("/Consumables/ConsumableHistory", utils.AuthRequired(hndl.ConsumableHistory))
	r.HandleFunc("/Consumables/CnsmblRetire", utils.AuthRequired(hndl.CnsmblRetire))
	r.HandleFunc("/Consumables/CheckDuplicateAssetEntry/{MasterID:[0-9]+}/{LocID:[0-9]+}", utils.AuthRequired(hndl.CheckDuplicateAssetEntry))
	r.HandleFunc("/AddConsumable_Partial", utils.AuthRequired(hndl.ConsumableAdd_Partial))
	r.HandleFunc("/GetConsumableGroups", utils.AuthRequired(hndl.GetConsumableGroups))
	r.HandleFunc("/ConsumablemasterInsert", utils.AuthRequired(hndl.ConsumablemasterInsert))
	r.HandleFunc("/Check_Unique_Consumable/{ConsumableName}", utils.AuthRequired(hndl.Check_Unique_Consumable))
}

//ITAssetsRoutings ..
func ITAssetsRoutings(r *mux.Router, hndl *itassethndlr.IITAsset) {
	r.HandleFunc("/GetITAssetToCheckoutToITasset/{LocID}/{AssetID}", utils.AuthRequired(hndl.GetITAssetToCheckoutToITasset))
	r.HandleFunc("/ITAssetDelete/{AssetID}", utils.AuthRequired(hndl.ITAssetDelete))
	r.HandleFunc("/MyITAssetRequestList", utils.AuthRequired(hndl.MyITAssetRequestList))
	r.HandleFunc("/ITAssetReqForward", utils.AuthRequired(hndl.ITAssetReqForward))
	r.HandleFunc("/GetITAssetReqListByEmp/{EmpID}",  utils.AuthRequired(hndl.GetITAssetReqListByEmp))
	r.HandleFunc("/ITAsset_Service_Request_Resolve", utils.AuthRequired(hndl.ITAsset_Service_Request_Resolve))
	r.HandleFunc("/ITAssets/view/{id:[0-9]+}", utils.AuthRequired(hndl.GetITAssetsByID))
	r.HandleFunc("/ITAssets/create", utils.AuthRequired(hndl.CreateITAsset))
	r.HandleFunc("/GetITAssetGroups", utils.AuthRequired(hndl.GetITAssetGroups))
	r.HandleFunc("/GetITAsset_History_ByAssetID/{AssetID}", utils.AuthRequired(hndl.GetITAsset_History_ByAssetID))
	r.HandleFunc("/ITAssetReq_ApprovalStatusList/{ReqGroupID}", utils.AuthRequired(hndl.ITAssetReq_ApprovalStatusList))
	r.HandleFunc("/ITAssetGroups_Create", utils.AuthRequired(hndl.ITAssetGroups_Create))
	r.HandleFunc("/ITAssetHistory/{AssetID}", utils.AuthRequired(hndl.ITAssetHistory))
	r.HandleFunc("/ITAssetRetiredList", utils.AuthRequired(hndl.ITAssetRetiredList))
	r.HandleFunc("/ITAssetMaintenanceList", utils.AuthRequired(hndl.ITAssetMaintenanceList))
	r.HandleFunc("/GetITAsset_Retired", utils.AuthRequired(hndl.GetITAsset_Retired))
	r.HandleFunc("/ITAssets", utils.AuthRequired(hndl.ITAssets))
	r.HandleFunc("/ITAssets/edit/{id:[0-9]+}", utils.AuthRequired(hndl.GetITAssetsEditByID))
	r.HandleFunc("/ITAssets/BulkEdit", utils.AuthRequired(hndl.ITAssetsBulkEdit))
	r.HandleFunc("/itassets/ITAssetRetire", utils.AuthRequired(hndl.ITAssetRetire))
	r.HandleFunc("/GetITAssets", utils.AuthRequired(hndl.GetITAssets))
	r.HandleFunc("/ITAssets/{id:[0-9]+}/checkout", utils.AuthRequired(hndl.ITAssetsCheckout))
	r.HandleFunc("/ITAssets/ITAssetView_Files/{id}", utils.AuthRequired(hndl.ITAssetViewFiles))
	r.HandleFunc("/ITAssets/GetITAssetView_Files/{id}", utils.AuthRequired(hndl.GetITassetsFilesByID))
	r.HandleFunc("/ITAssets/UploadFile", utils.AuthRequired(hndl.UploadFiles))
	r.HandleFunc("/ITAssets/ReadExcel", utils.AuthRequired(hndl.ITAssetReadExcel))
	r.HandleFunc("/ITAssets/Request", utils.AuthRequired(hndl.ITAssetRequest))
	r.HandleFunc("/ITAssets/PostITAssetRequest", utils.AuthRequired(hndl.PostITAssetRequest))
	r.HandleFunc("/ITAssets/RequestList", utils.AuthRequired(hndl.ITAssetRequestList))
	r.HandleFunc("/itassets/GetITAssetReqList/{Apprvrid:[0-9]+}", utils.AuthRequired(hndl.GetITAssetReqListByApprover))
	r.HandleFunc("/itassets/ITAssetReqReject", utils.AuthRequired(hndl.ITAssetReqReject))
	r.HandleFunc("/itassets/ITAssetReqApprove", utils.AuthRequired(hndl.ITAssetReqApprove))
	r.HandleFunc("/itassets/GetITAssetReqListByReqGroup/{ReqGroupID:[0-9]+}/{ApproverID:[0-9]+}", utils.AuthRequired(hndl.GetITAssetReqListByReqGroup))
	r.HandleFunc("/CreateITAssetsCheckIn", utils.AuthRequired(hndl.CreateITAssetsCheckIn))
	r.HandleFunc("/Customfields/{id}/{Mod}", utils.AuthRequired(hndl.Customfields))
	r.HandleFunc("/GetCustomFields/{id}/{Mod}", utils.AuthRequired(hndl.GetCustomFields))
	r.HandleFunc("/GetEmployeeITAssetsByID", utils.AuthRequired(hndl.GetEmployeeITAssetsByID))
	r.HandleFunc("/PrintQr", utils.AuthRequired(hndl.PrintQr))
	r.HandleFunc("/QrGenerator", utils.AuthRequired(hndl.QrGenerator))
	r.HandleFunc("/employee/{id}/ITassets", utils.AuthRequired(hndl.EmployeeITAssetsByID))
	r.HandleFunc("/ITAsset_Service/{ITAssetID}", utils.AuthRequired(hndl.ITAsset_Service))
	r.HandleFunc("/ITasset_services_Insert", utils.AuthRequired(hndl.ITasset_services_Insert))
	r.HandleFunc("/ITasset_services_Extend_Update", utils.AuthRequired(hndl.ITasset_services_Extend_Update))
	r.HandleFunc("/ITasset_services_Complete_Update", utils.AuthRequired(hndl.ITasset_services_Complete_Update))
	r.HandleFunc("/ITasset_services_start_Update", utils.AuthRequired(hndl.ITasset_services_start_Update))
	r.HandleFunc("/ITAssetView_Maintenance/{ITAssetID}", utils.AuthRequired(hndl.ITAssetView_Maintenance))
	r.HandleFunc("/GetITAssetservices_List/{ITAssetID:[0-9]+}", utils.AuthRequired(hndl.GetITAssetservices_List))

	r.HandleFunc("/GetITAssetservices_List_ByLoc/{LocID:[0-9]+}", utils.AuthRequired(hndl.GetITAssetservices_List_ByLoc))
	r.HandleFunc("/ITAsset_Service_Request/{ITAssetID}", utils.AuthRequired(hndl.ITAsset_Service_Request))

	r.HandleFunc("/ITAsset_Service_Request_List", utils.AuthRequired(hndl.ITAsset_Service_Request_List))
	r.HandleFunc("/GetITAsset_service_request_List/{EmpID:[0-9]+}", utils.AuthRequired(hndl.GetITAsset_service_request_List))
}

//CommonRoutings ..
func CommonRoutings(r *mux.Router, hndl *cmnhandler.ICommonrep) {
	// hub := utils.NewHub()
	// go hub.Run()
	// r.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	// 	utils.ServeWs(hub, w, r)
	// })
	r.HandleFunc("/Requisition", utils.AuthRequired(hndl.Requisition))
	r.HandleFunc("/RequisitionReceivedStock", utils.AuthRequired(hndl.RequisitionReceivedStock))
	r.HandleFunc("/POStatusChange/{IDPO}/{Status}", utils.AuthRequired(hndl.POStatusChange))
	r.HandleFunc("/RequisitionStatusChange/{ID}/{Status}", utils.AuthRequired(hndl.RequisitionStatusChange))
	r.HandleFunc("/PurchaseOrders_RequestsUpdate", utils.AuthRequired(hndl.PurchaseOrders_RequestsUpdate))
	r.HandleFunc("/Requisition_RequestsUpdate", utils.AuthRequired(hndl.Requisition_RequestsUpdate))
	r.HandleFunc("/POApprovalDetails", utils.AuthRequired(hndl.POApprovalDetails))
	r.HandleFunc("/DeleteVendors", utils.AuthRequired(hndl.DeleteVendors))
	r.HandleFunc("/RequisitionApprovalDetails", utils.AuthRequired(hndl.RequisitionApprovalDetails))
	r.HandleFunc("/POReqForward", utils.AuthRequired(hndl.POReqForward))
	r.HandleFunc("/POReqApproved", utils.AuthRequired(hndl.POReqApproved))
	r.HandleFunc("/RequisitionReqApproved", utils.AuthRequired(hndl.RequisitionReqApproved))
	r.HandleFunc("/RequisitionReqForward", utils.AuthRequired(hndl.RequisitionReqForward))
	r.HandleFunc("/PO_Edit/{IDPO}", utils.AuthRequired(hndl.PO_Edit))
	r.HandleFunc("/Requisition_Edit/{ID}", utils.AuthRequired(hndl.Requisition_Edit))
	r.HandleFunc("/POReqRejected", utils.AuthRequired(hndl.POReqRejected))
	r.HandleFunc("/RequisitionReqRejected", utils.AuthRequired(hndl.RequisitionReqRejected))
	r.HandleFunc("/GetPODetailsByApprover/{ApprvrID}", utils.AuthRequired(hndl.GetPODetailsByApprover))
	r.HandleFunc("/GetRequisitionDetailsByApprover/{ApprvrID}", utils.AuthRequired(hndl.GetRequisitionDetailsByApprover))
	r.HandleFunc("/POAssetDetailsByIDPO/{IDPO}", utils.AuthRequired(hndl.POAssetDetailsByIDPO))
	r.HandleFunc("/RequisitionAssetDetailsByID/{ID}", utils.AuthRequired(hndl.RequisitionAssetDetailsByID))
	r.HandleFunc("/GetRequisitionHistoryByReqID/{ReqID}", utils.AuthRequired(hndl.GetRequisitionHistoryByReqID))
	r.HandleFunc("/RequisitionHistoryDetails_Partial/{ReqID}", utils.AuthRequired(hndl.RequisitionHistoryDetails_Partial))
	r.HandleFunc("/PO_ApprovalStatusList/{IDPO}", utils.AuthRequired(hndl.PO_ApprovalStatusList))
	r.HandleFunc("/Requisition_ApprovalStatusList/{ID}", utils.AuthRequired(hndl.Requisition_ApprovalStatusList))
	r.HandleFunc("/PurchaseOrderView/{IDPO}", utils.AuthRequired(hndl.PurchaseOrderView))
	r.HandleFunc("/RequisitionView/{ID}", utils.AuthRequired(hndl.RequisitionView))
	r.HandleFunc("/GetPODetailsByReqstrID/{ReqstrID}", utils.AuthRequired(hndl.GetPODetailsByReqstrID))
	r.HandleFunc("/GetRequisitionDetailsByReqstrID/{ReqstrID}", utils.AuthRequired(hndl.GetRequisitionDetailsByReqstrID))
	r.HandleFunc("/GetPurchaseOrderUniqueID", utils.AuthRequired(hndl.GetPurchaseOrderUniqueID))
	r.HandleFunc("/PurchaseOrder", utils.AuthRequired(hndl.PurchaseOrder))
	r.HandleFunc("/PurchaseOrders_RequestsInsert", utils.AuthRequired(hndl.PurchaseOrders_RequestsInsert))
	r.HandleFunc("/Requisition_RequestsInsert", utils.AuthRequired(hndl.Requisition_RequestsInsert))
	r.HandleFunc("/PurchaseOrderDetails", utils.AuthRequired(hndl.PurchaseOrderDetails))
	r.HandleFunc("/RequisitionDetails", utils.AuthRequired(hndl.RequisitionDetails))
	r.HandleFunc("/GetMultiLevelApproval_Map", utils.AuthRequired(hndl.GetMultiLevelApproval_Map))
	r.HandleFunc("/InwardOutwardReqForward", utils.AuthRequired(hndl.InwardOutwardReqForward))
	r.HandleFunc("/IWOW_ApprovalStatusList/{IDinwardoutward}", utils.AuthRequired(hndl.IWOW_ApprovalStatusList))
	r.HandleFunc("/MultiLevelApproval_config", utils.AuthRequired(hndl.MultiLevelApproval_config))
	r.HandleFunc("/emp", utils.AuthRequired(hndl.Emp))
	r.HandleFunc("/MultiApproval", utils.AuthRequired(hndl.MultiApproval))
	r.HandleFunc("/MyDetails", utils.AuthRequired(hndl.MyDetails))
	r.HandleFunc("/User_ActivityLog/{EmpID}", utils.AuthRequired(hndl.User_ActivityLog))
	r.HandleFunc("/Activivty_Log_List/{EmpID}", utils.AuthRequired(hndl.Activivty_Log_List))
	r.HandleFunc("/Authorization_Create", utils.AuthRequired(hndl.Authorization_Create))
	r.HandleFunc("/GetAuthorizationList_ByRole/{RoleID}", utils.AuthRequired(hndl.GetAuthorizationList_ByRole))
	r.HandleFunc("/Authorization", utils.AuthRequired(hndl.Authorization))
	r.HandleFunc("/ResendActivationLink/{EmpID}", utils.AuthRequired(hndl.Resend_Activation_Link))
	r.HandleFunc("/", hndl.Login)
	r.HandleFunc("/logout", hndl.LogOut)
	r.HandleFunc("/GetNotifications", utils.AuthRequired(hndl.GetNotifications))
	r.HandleFunc("/dashboard", utils.AuthRequired(hndl.Dashboard))
	r.HandleFunc("/MyDashBoard", utils.AuthRequired(hndl.MyDashBoard))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.PathPrefix("/AppFiles/").Handler(http.StripPrefix("/AppFiles/", http.FileServer(http.Dir("./AppFiles"))))
	r.HandleFunc("/Location", utils.AuthRequired(hndl.Location))
	r.HandleFunc("/status/{typeName}", utils.AuthRequired(hndl.GetStatus))
	r.HandleFunc("/Vendors", utils.AuthRequired(hndl.Vendors))
	r.HandleFunc("/GetCountries", utils.AuthRequired(hndl.GetCountries)).Methods("GET")
	r.HandleFunc("/GetStatesByCountry/{id:[0-9]+}", utils.AuthRequired(hndl.GetStatesByCountry))
	r.HandleFunc("/GetLocations", utils.AuthRequired(hndl.GetLocations))
	r.HandleFunc("/GetVendors", utils.AuthRequired(hndl.GetVendors))
	r.HandleFunc("/GetRoles", utils.AuthRequired(hndl.GetRoles))
	r.HandleFunc("/GetEducations", utils.AuthRequired(hndl.GetEducations))
	r.HandleFunc("/GetDesignations", utils.AuthRequired(hndl.GetDesignations))
	r.HandleFunc("/GetInWardDetailsByEmp", utils.AuthRequired(hndl.GetInWardDetailsByEmp))
	//r.HandleFunc("/GenerateUniqueID/{modulename}", utils.AuthRequired(hndl.GenerateUniqueID))
	r.HandleFunc("/employees", utils.AuthRequired(hndl.GetEmployees))
	r.HandleFunc("/UsersList", utils.AuthRequired(hndl.UsersList))
	r.HandleFunc("/UserCreate", hndl.UserCreate)
	r.HandleFunc("/ResetPassword", hndl.ResetPassword)
	r.HandleFunc("/Send_ResetPasswordLink/{empid}", hndl.Send_ResetPasswordLink)
	r.HandleFunc("/User_Inactive/{UserID}", utils.AuthRequired(hndl.User_Inactive))
	r.HandleFunc("/User_Active/{UserID}", utils.AuthRequired(hndl.User_Active))
	r.HandleFunc("/Check_Unique_Email_Mobile", utils.AuthRequired(hndl.Check_Unique_Email_Mobile))
	r.HandleFunc("/Check_Unique_UserName/{UserName}", utils.AuthRequired(hndl.Check_Unique_UserName))
	r.HandleFunc("/GetEmployee_History_ByEmpID/{EmpID}", utils.AuthRequired(hndl.GetEmployee_History_ByEmpID))
	r.HandleFunc("/Get_UsersHistory_ByEmpID/{EmpID}", utils.AuthRequired(hndl.Get_UsersHistory_ByEmpID))
	r.HandleFunc("/EmployeesHistory", utils.AuthRequired(hndl.EmployeesHistory))
	r.HandleFunc("/UsersHistory", utils.AuthRequired(hndl.UsersHistory))
	r.HandleFunc("/GetUsersList", utils.AuthRequired(hndl.GetUsersList))
	r.HandleFunc("/GetApprovers/{LocID:[0-9]+}/{RoleID:[0-9]+}", utils.AuthRequired(hndl.GetApprovers))
	r.HandleFunc("/GetEmployeesList", utils.AuthRequired(hndl.GetEmployeesList))
	r.HandleFunc("/employee/{id:[0-9]+}/mode/{mode}", utils.AuthRequired(hndl.GetEmployeeByID))
	r.HandleFunc("/employee/create", utils.AuthRequired(hndl.CreateEmployee))
	//r.HandleFunc("/user/create/{empid:[0-9]+}",utils.AuthRequired(hndl.CreateUser)
	//r.HandleFunc("/employee/update",utils.AuthRequired(hndl.UpdateEmployee)
	r.HandleFunc("/employee/ReadExcel", utils.AuthRequired(hndl.EmployeeReadExcel))
	r.HandleFunc("/employeeDelete/{id}", utils.AuthRequired(hndl.DeleteEmployee))
	r.HandleFunc("/CreateOutWardCart", utils.AuthRequired(hndl.CreateOutWardCart))
	r.HandleFunc("/CreateInWardOutWard", utils.AuthRequired(hndl.CreateInWardOutWard))
	r.HandleFunc("/GetOutWardCart/{SenderEmpid:[0-9]+}", utils.AuthRequired(hndl.GetOutWardCart))
	r.HandleFunc("/Transfer", utils.AuthRequired(hndl.Transfer))
	r.HandleFunc("/DeleteOutWardCart", utils.AuthRequired(hndl.DeleteOutWardCart))
	r.HandleFunc("/OutWardDetails", utils.AuthRequired(hndl.OutWardDetails))
	r.HandleFunc("/InWardDetails", utils.AuthRequired(hndl.InWardDetails))
	r.HandleFunc("/OWApprovalDeatils", utils.AuthRequired(hndl.OWApprovalDeatils))
	r.HandleFunc("/GetOutWardDetailsByEmp/{SenderEmpid}", utils.AuthRequired(hndl.GetOutWardDetailsByEmp))
	r.HandleFunc("/GetInWardDetailsByEmp/{RcvrEmpID}", utils.AuthRequired(hndl.GetInWardDetailsByEmp))
	r.HandleFunc("/GetOutWardAssetDetailsByIwOwID/{IwOwID}", utils.AuthRequired(hndl.GetOutWardAssetDetailsByIwOwID))
	r.HandleFunc("/GetAssetdIDsNotEligbleForTransfer", utils.AuthRequired(hndl.GetAssetdIDsNotEligbleForTransfer))
	r.HandleFunc("/IWOWDetailsByAprvr/{AprvrEmpID}", utils.AuthRequired(hndl.IWOWDetailsByAprvr))
	r.HandleFunc("/OWApproval/{AprvrEmpID}/{IDInWardOutWard}", utils.AuthRequired(hndl.OWApproval))
	r.HandleFunc("/TransferApprovReject/{idInWardOutWard}/{status}", utils.AuthRequired(hndl.TransferApprovReject))
	r.HandleFunc("/IWReceive/{RcvrEmpID}/{IDInWardOutWard}", utils.AuthRequired(hndl.IWReceive))
	r.HandleFunc("/ReceiveIWAssets", utils.AuthRequired(hndl.ReceiveIWAssets))
	r.HandleFunc("/OwStatusChange/{OWid}/{Status}", utils.AuthRequired(hndl.OwStatusChange))
	r.HandleFunc("/UpdateIsMsngStcksRslvdMain/{IDInWardOutWard}", utils.AuthRequired(hndl.UpdateIsMsngStcksRslvdMain))
	r.HandleFunc("/LocationsDetails", utils.AuthRequired(hndl.LocationsDetails))
	r.HandleFunc("/VendorsDetails", utils.AuthRequired(hndl.VendorsDetails))
	r.HandleFunc("/VendorAssetDetails", utils.AuthRequired(hndl.VendorAssetDetails))
	r.HandleFunc("/GetVendorsAssetDetails/{vendorID}", utils.AuthRequired(hndl.GetVendorsAssetDetails))
	r.HandleFunc("/VednorsAssetMapInsert", utils.AuthRequired(hndl.VednorsAssetMapInsert))
	r.HandleFunc("/VednorsAssetMapUpdate", utils.AuthRequired(hndl.VednorsAssetMapUpdate))

}

//NonITAssetsRoutings ..
func NonITAssetsRoutings(r *mux.Router, hndl *nonitassetshndlr.INonITAsset) {
bi	
	r.HandleFunc("/NonITAssetDelete/{AssetID}", utils.AuthRequired(hndl.NonITAssetDelete))
	r.HandleFunc("/CheckDuplicateNonITAssetEntry/{MasterID:[0-9]+}/{LocID:[0-9]+}", utils.AuthRequired(hndl.CheckDuplicateNonITAssetEntry))
	r.HandleFunc("/Check_Unique_NonITAsset/{NonITAssetName}", utils.AuthRequired(hndl.Check_Unique_NonITAsset))
	r.HandleFunc("/GetNonITAssetReqListByEmp/{EmpID}", utils.AuthRequired(hndl.GetNonITAssetReqListByEmp))
	r.HandleFunc("/MyNonITAssetRequestList", utils.AuthRequired(hndl.MyNonITAssetRequestList))
	r.HandleFunc("/NonITAsset/Create", utils.AuthRequired(hndl.NonITAsset_Create))
	r.HandleFunc("/NonITAsset/Edit/{IDNonITAsset}", utils.AuthRequired(hndl.NonITAsset_Edit))
	r.HandleFunc("/GetNonITAssetMasterList", utils.AuthRequired(hndl.GetNonITAssetMasterList))
	r.HandleFunc("/GetNonITAssetLists/{LocID}", utils.AuthRequired(hndl.GetNonITAssetLists))
	r.HandleFunc("/NonITAsset/List", utils.AuthRequired(hndl.NonITAssetList))
	r.HandleFunc("/NonITAsset/AssetID/{AssetID}", utils.AuthRequired(hndl.NonITAsset_BY_AssetID))
	r.HandleFunc("/NonITAssets/Removestock", utils.AuthRequired(hndl.PostNonITAssets_oprtns_Removestock))
	r.HandleFunc("/NonITAssets/AddStock", utils.AuthRequired(hndl.PostNonITAssets_oprtns_AddStock))
	r.HandleFunc("/NonITAssets/Checkout/{AssetID}", utils.AuthRequired(hndl.NonITAssetsCheckout))
	r.HandleFunc("/NonITAssetAdd_Partial", utils.AuthRequired(hndl.NonITAssetAdd_Partial))
	r.HandleFunc("/NonITAssetemasterInsert", utils.AuthRequired(hndl.NonITAssetemasterInsert))
	r.HandleFunc("/GetNonITAssetGroups", utils.AuthRequired(hndl.GetNonITAssetGroups))
	r.HandleFunc("/NonITAssetMasterAdd", utils.AuthRequired(hndl.NonITAssetMasterAdd))
	r.HandleFunc("/GetNonITAssetOprtnListByID/{IDNonITAsset}", utils.AuthRequired(hndl.GetNonITAssetOprtnListByID))
	r.HandleFunc("/NonITAssetOprtnListByID/{IDNonITAsset}", utils.AuthRequired(hndl.NonITAssetOprtnListByID))
	r.HandleFunc("/NonITAssetRequest", utils.AuthRequired(hndl.NonITAssetRequest))
	r.HandleFunc("/PostNonITAssetRequest", utils.AuthRequired(hndl.PostNonITAssetRequest))
	r.HandleFunc("/GetNonITAssetReqList/{Apprvrid}", utils.AuthRequired(hndl.GetNonITAssetReqList))
	r.HandleFunc("/NonITAssetRequestList", utils.AuthRequired(hndl.NonITAssetRequestList))
	r.HandleFunc("/NonITAssetReq_ApprovalStatusList/{ReqGroupID}", utils.AuthRequired(hndl.NonITAssetReq_ApprovalStatusList))
	r.HandleFunc("/NonITAssetReqReject", utils.AuthRequired(hndl.NonITAssetReqReject))
	r.HandleFunc("/NonITAssetReqForward", utils.AuthRequired(hndl.NonITAssetReqForward))
	r.HandleFunc("/NonITAssetReqApprove", utils.AuthRequired(hndl.NonITAssetReqApprove))
	r.HandleFunc("/GetNonITAssetCheckinDetails/{LocID}", utils.AuthRequired(hndl.GetNonITAssetCheckinDetails))
	r.HandleFunc("/GetNonITAssetCheckinDetailsByAsset/{IDNonITAsset}", utils.AuthRequired(hndl.GetNonITAssetCheckinDetailsByAsset))
	r.HandleFunc("/GetNonITAssetCheckinDetailsByEmp/{EmpID}", utils.AuthRequired(hndl.GetNonITAssetCheckinDetailsByEmp))
	r.HandleFunc("/NonITAssetCheckinDetails", utils.AuthRequired(hndl.NonITAssetCheckinDetails))
	r.HandleFunc("/NonITAssetCheckin", utils.AuthRequired(hndl.NonITAssetCheckin))
	r.HandleFunc("/NonITAssetCheckinDetails_Partial", utils.AuthRequired(hndl.NonITAssetCheckinDetails_Partial))
	r.HandleFunc("/Getnonitassets_checkinByID/{CheckinID}", utils.AuthRequired(hndl.Getnonitassets_checkinByID))
	r.HandleFunc("/NonITAssetReqListByReqGroup/{ReqGroupID:[0-9]+}/{ApproverID:[0-9]+}", utils.AuthRequired(hndl.NonITAssetReqListByReqGroup))
}
