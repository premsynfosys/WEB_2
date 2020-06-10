package itassethndlr

import (
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gorilla/mux"
	"github.com/premsynfosys/AMS_WEB/models/CmnModel"
	"github.com/premsynfosys/AMS_WEB/models/ITAssetsmodel"
	"github.com/premsynfosys/AMS_WEB/repository/CmnRep"
	"github.com/premsynfosys/AMS_WEB/repository/ITAssetRep"
	"github.com/premsynfosys/AMS_WEB/utils"
)

// IITAsset ..
type IITAsset struct {
	Irepo    ITAssetRep.ITAssetIntrfc
	ICmnrepo CmnRep.CmnRepIntrfc
}

//NewITAssetHandler ..
func NewITAssetHandler(api string) *IITAsset {
	return &IITAsset{
		Irepo:    ITAssetRep.NewAPIRepo(api),
		ICmnrepo: CmnRep.NewAPIRepo(api),
	}
}

//GetEmployeeITAssetsByID ..
func (p *IITAsset) GetEmployeeITAssetsByID(w http.ResponseWriter, r *http.Request) {
	isCheckin, _ := strconv.Atoi(r.URL.Query().Get("isCheckin"))
	EmpID, _ := strconv.Atoi(r.URL.Query().Get("EmpID"))
	body, _ := p.Irepo.GetEmployeeITAssetsByID(r.Context(), EmpID, isCheckin)
	utils.RespondwithJSON(w, r, http.StatusOK, body)

}

//EmployeeITAssetsByID ..
func (p *IITAsset) EmployeeITAssetsByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	id, _ := strconv.Atoi(ID)

	utils.ExecuteTemplate(w, r, "EmployeeItassets", id)
}

//______________________________________________________________________________________________
//CreateITAsset a new post
func (p *IITAsset) CreateITAsset(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		data := &ITAssetsmodel.ITAssetModel{}

		if r.URL.Query().Get("mode") == "clone" {
			id := r.URL.Query().Get("id")
			_id, _ := strconv.Atoi(id)
			data = p.returnITAssetsByID(r, _id)
			//data.ITAssetIdentificationNo = ""
			data.ITAssetStatus = 1

		} else {
			status, _ := p.ICmnrepo.GetStatus(r.Context(), "itasset")
			data.ITAssetStatusList = status
			data.ITAssetGroupList, _ = p.Irepo.GetITAssetGroups(r.Context())

		}

		Mapdata := CmnModel.TemplateData{
			User: usr,
			Data: data,
			Auth: auth,
		}

		utils.ExecuteTemplate(w, r, "ITAssetTmpl", Mapdata)
	}

	if r.Method == "POST" {
		//r.ParseForm()
		mdl := ITAssetsmodel.ITAssetModel{}
		mdl.ITAssetName = r.FormValue("ITAssetName")
		ITAssetGroup, _ := strconv.Atoi(r.FormValue("ITAssetGroup"))
		mdl.ITAssetGroup = ITAssetGroup
		mdl.ITAssetModel = r.FormValue("ITAssetModel")
		mdl.ITAssetSerialNo = r.FormValue("ITAssetSerialNo")
		//mdl.ITAssetIdentificationNo = r.FormValue("ITAssetIdentificationNo")
		mdl.ITAssetDescription = r.FormValue("ITAssetDescription")

		mdl.ITAssetAssignTo = r.FormValue("CanITAssetAssignTo")
		ITAssetPrice, _ := strconv.ParseFloat(r.FormValue("ITAssetPrice"), 32)
		mdl.ITAssetPrice = ITAssetPrice
		mdl.ITAssetWarranty = r.FormValue("ITAssetWarranty")
		VendorID, _ := strconv.Atoi(r.FormValue("Vendor"))
		mdl.Vendor = VendorID
		LocationID, _ := strconv.Atoi(r.FormValue("Location"))
		mdl.Location = LocationID

		ITAssetStatus, _ := strconv.Atoi(r.FormValue("ITAssetStatus"))
		mdl.ITAssetStatus = ITAssetStatus
		mdl.ITAssetFileUpld = r.FormValue("ITAssetFileUpld")

		img, handle, err := r.FormFile("ITAssetImg")
		if err == nil {
			file, _ := os.Create("AppFiles/Images/ITAssets/" + handle.Filename)
			io.Copy(file, img)
			defer file.Close()
			mdl.ITAssetImg = handle.Filename
			defer img.Close()
		}

		customfields := make(map[string][]string)
		customfieldsRaw := []byte(r.FormValue("customfields"))
		json.Unmarshal(customfieldsRaw, &customfields)

		if r.FormValue("CustomFields1") != "" {
			mdl.CustomFields1Value = r.FormValue("CustomFields1")
			mdl.CustomFields1 = customfields["CustomFields1"][0]
			mdl.CustomFields1Type = customfields["CustomFields1"][1]
		}
		if r.FormValue("CustomFields2") != "" {
			mdl.CustomFields2Value = r.FormValue("CustomFields2")
			mdl.CustomFields2 = customfields["CustomFields2"][0]
			mdl.CustomFields2Type = customfields["CustomFields2"][1]
		}

		if r.FormValue("CustomFields3") != "" {
			mdl.CustomFields3Value = r.FormValue("CustomFields3")
			mdl.CustomFields3 = customfields["CustomFields3"][0]
			mdl.CustomFields3Type = customfields["CustomFields3"][1]
		}

		if r.FormValue("CustomFields4") != "" {
			mdl.CustomFields4Value = r.FormValue("CustomFields4")
			mdl.CustomFields4 = customfields["CustomFields4"][0]
			mdl.CustomFields4Type = customfields["CustomFields4"][1]
		}
		mdl.CreatedBy = usr.EmployeeID
		_, _ = p.Irepo.CreateITAsset(r.Context(), &mdl)
		//http.Redirect(w, r, "/ITAssets", 301)
	}
}

// ITAssets ..
func (p *IITAsset) ITAssets(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{

		User: usr,
		Auth: Auth,
	}
	utils.ExecuteTemplate(w, r, "ITAssetDetailsTmpl", Mapdata)
}

// ITAssetRequest ..
func (p *IITAsset) ITAssetRequest(w http.ResponseWriter, r *http.Request) {
	user, auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		ITAssetGroups := []*ITAssetsmodel.ITAssetGroup{}
		ITAssetGroups, _ = p.Irepo.GetITAssetGroups(r.Context())
		LocID, _ := strconv.Atoi(r.URL.Query().Get("LocID"))
		mapRoles, _ := p.ICmnrepo.GetMultiLevelApproval_Map(r.Context())
		var RoleID int
		var Grade int
		userroleid := user.RoleID
		for i := 0; i < len(mapRoles); i++ {
			if mapRoles[i].FeatureName == "ITAsset Request" {
				if mapRoles[i].MultiLevelApproval_Map.RoleID != userroleid && RoleID == 0 {
					RoleID = mapRoles[i].MultiLevelApproval_Map.RoleID
					Grade = mapRoles[i].MultiLevelApproval_Map.Grade
				} else if userroleid == mapRoles[i].MultiLevelApproval_Map.RoleID {
					if len(mapRoles) > i+1 {
						if mapRoles[i+1].FeatureName == "ITAsset Request" {
							RoleID = mapRoles[i+1].MultiLevelApproval_Map.RoleID
							Grade = mapRoles[i+1].MultiLevelApproval_Map.Grade
						}
					} else {
						RoleID = 0
						Grade = 1
					}
				}

			}
		}

		Approvers, _ := p.ICmnrepo.GetApprovers(r.Context(), LocID, RoleID)
		//mapData := make(map[string]interface{})
		mapData := struct {
			ITAssetGroups  []*ITAssetsmodel.ITAssetGroup
			UsersEmployees []*CmnModel.User
			Grade          int
			RoleID         int
		}{
			ITAssetGroups:  ITAssetGroups,
			UsersEmployees: Approvers,
			Grade:          Grade,
			RoleID:         RoleID,
		}

		Mapdata := CmnModel.TemplateData{
			Data: mapData,
			Auth: auth,
			User: user,
		}

		utils.ExecuteTemplate(w, r, "ITAssetRequest", Mapdata)
	}
}

//PostITAssetRequest ..
func (p *IITAsset) PostITAssetRequest(w http.ResponseWriter, r *http.Request) {
	ListData := []*ITAssetsmodel.ITAssetRequest{}
	json.NewDecoder(r.Body).Decode(&ListData)
	err := p.Irepo.CreateITAssetRequest(r.Context(), ListData)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}

}

//PostITAssetRequest ..
func (p *IITAsset) ITAssetGroups_Create(w http.ResponseWriter, r *http.Request) {
	ITAssetGroup := ITAssetsmodel.ITAssetGroup{}
	ITAssetGroup.ITAssetgroupName = r.FormValue("ITAssetgroupName")
	ITAssetGroup.CreatedBy, _ = strconv.Atoi(r.FormValue("CreatedBy"))
	err := p.Irepo.ITAssetGroups_Create(r.Context(), &ITAssetGroup)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}

}

//ITAssetRequestList ..
func (p *IITAsset) ITAssetRequestList(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)
	mapRoles, _ := p.ICmnrepo.GetMultiLevelApproval_Map(r.Context())
	requstmap := []*CmnModel.MultiLevelApproval_Main{}
	for _, itm := range mapRoles {
		if itm.FeatureName == "ITAsset Request" {
			requstmap = append(requstmap, itm)
		}
	}
	Mapdata := CmnModel.TemplateData{
		Auth: auth,
		User: usr,
		Data: requstmap,
	}

	utils.ExecuteTemplate(w, r, "ITAssetRequestList", Mapdata)
}

//GetITAssetReqListByApprover ..
func (p *IITAsset) GetITAssetReqListByApprover(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["Apprvrid"]
	Apprvrid, _ := strconv.Atoi(ID)
	data, err := p.Irepo.GetITAssetReqList(r.Context(), Apprvrid)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	}
}

//ITAssetReqReject ..
func (p *IITAsset) ITAssetReqReject(w http.ResponseWriter, r *http.Request) {
	mdl := ITAssetsmodel.ITAssetRequestApproval{}
	json.NewDecoder(r.Body).Decode(&mdl)
	data, err := p.Irepo.ITAssetReqReject(r.Context(), &mdl)
	if err == nil && data {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	}
}

//ITAssetReqApprove ..
func (p *IITAsset) ITAssetReqApprove(w http.ResponseWriter, r *http.Request) {
	mdl := []*ITAssetsmodel.ITAssetRequest{}
	json.NewDecoder(r.Body).Decode(&mdl)
	data, err := p.Irepo.ITAssetReqApprove(r.Context(), mdl)

	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	}
}

//GetITAssetReqListByReqGroup ..
func (p *IITAsset) GetITAssetReqListByReqGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["ReqGroupID"]
	ReqGroupID, _ := strconv.Atoi(ID)
	ApproverID := params["ApproverID"]
	approverID, _ := strconv.Atoi(ApproverID)
	data, err := p.Irepo.GetITAssetReqListByReqGroup(r.Context(), ReqGroupID, approverID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	}
}

//GetITAssetReqListByReqGroup ..
func (p *IITAsset) ITAssetReq_ApprovalStatusList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["ReqGroupID"]
	ReqGroupID, _ := strconv.Atoi(ID)
	data, err := p.Irepo.ITAssetReq_ApprovalStatusList(r.Context(), ReqGroupID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	}
}

// GetITAssets ..
func (p *IITAsset) GetITAssets(w http.ResponseWriter, r *http.Request) {
	LocID, _ := strconv.Atoi(r.URL.Query().Get("LocID"))
	data, _ := p.Irepo.GetITAssets(r.Context(), LocID)
	utils.RespondwithJSON(w, r, http.StatusOK, data)
}
func (p *IITAsset) returnITAssetsByID(r *http.Request, id int) *ITAssetsmodel.ITAssetModel {
	LocID, _ := strconv.Atoi(r.URL.Query().Get("LocID"))
	data, _ := p.Irepo.GetITAssets(r.Context(), LocID)
	status, _ := p.ICmnrepo.GetStatus(r.Context(), "itasset")
	mdl := &ITAssetsmodel.ITAssetModel{}
	for _, item := range data {
		if item.IDITAssets == id {
			mdl = item
		}
	}
	mdl.ITAssetStatusList = status
	return mdl
}

//GetITAssetGroups ..
func (p *IITAsset) GetITAssetGroups(w http.ResponseWriter, r *http.Request) {
	body, _ := p.Irepo.GetITAssetGroups(r.Context())
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

// GetITAssetsByID ..
func (p *IITAsset) GetITAssetsByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	id, _ := strconv.Atoi(ID)
	item := p.returnITAssetsByID(r, id)

	lst, _ := p.ICmnrepo.GetLocations(r.Context())
	for _, itm := range lst {
		if item.Location == itm.IDLocations {
			item.LocationData = *itm
		}
	}
	lst2, _ := p.ICmnrepo.GetVendors(r.Context())
	for _, itm := range lst2 {
		if item.Vendor == itm.Idvendors {
			item.VendorData = *itm
		}
	}
	usr, auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		Data: item,
		Auth: auth,
		User: usr,
	}
	utils.ExecuteTemplate(w, r, "ITAssetViewTmpl", Mapdata)
}

// GetITAssetsEditByID ..
func (p *IITAsset) GetITAssetsEditByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	id, _ := strconv.Atoi(ID)
	usr, auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		item := p.returnITAssetsByID(r, id)
		status, _ := p.ICmnrepo.GetStatus(r.Context(), "itasset")
		item.ITAssetStatusList = status
		item.ITAssetGroupList, _ = p.Irepo.GetITAssetGroups(r.Context())

		Mapdata := CmnModel.TemplateData{
			Data: item,
			Auth: auth,
			User: usr,
		}
		utils.ExecuteTemplate(w, r, "ITAssetEditTmpl", Mapdata)

	}
	if r.Method == "POST" {
		mdl := ITAssetsmodel.ITAssetModel{}
		mdl.IDITAssets = id
		mdl.ITAssetName = r.FormValue("ITAssetName")
		ITAssetGroup, _ := strconv.Atoi(r.FormValue("ITAssetGroup"))
		mdl.ITAssetGroup = ITAssetGroup
		mdl.ITAssetModel = r.FormValue("ITAssetModel")
		mdl.ITAssetSerialNo = r.FormValue("ITAssetSerialNo")
		//mdl.ITAssetIdentificationNo = r.FormValue("ITAssetIdentificationNo")
		mdl.ITAssetDescription = r.FormValue("ITAssetDescription")

		//mdl.ITAssetAssignTo = r.FormValue("CanITAssetAssignTo")
		ITAssetPrice, _ := strconv.ParseFloat(r.FormValue("ITAssetPrice"), 32)
		mdl.ITAssetPrice = ITAssetPrice
		mdl.ITAssetWarranty = r.FormValue("ITAssetWarranty")
		VendorID, _ := strconv.Atoi(r.FormValue("Vendor"))
		mdl.Vendor = VendorID
		LocationID, _ := strconv.Atoi(r.FormValue("Location"))
		mdl.Location = LocationID

		ITAssetStatus, _ := strconv.Atoi(r.FormValue("ITAssetStatus"))
		mdl.ITAssetStatus = ITAssetStatus
		mdl.ITAssetFileUpld = r.FormValue("ITAssetFileUpld")
		mdl.ITAssetImg = r.FormValue("ITAssetImg1")
		img, handle, _ := r.FormFile("ITAssetImg")
		if handle != nil {
			mdl.ITAssetImg = handle.Filename
			file, _ := os.Create("AppFiles/Images/ITAssets/" + handle.Filename)
			io.Copy(file, img)
			defer file.Close()
			defer img.Close()
		}
		customfields := make(map[string][]string)
		customfieldsRaw := []byte(r.FormValue("customfields"))
		json.Unmarshal(customfieldsRaw, &customfields)

		if r.FormValue("CustomFields1") != "" {
			mdl.CustomFields1Value = r.FormValue("CustomFields1")
			mdl.CustomFields1 = customfields["CustomFields1"][0]
			mdl.CustomFields1Type = customfields["CustomFields1"][1]
		}
		if r.FormValue("CustomFields2") != "" {
			mdl.CustomFields2Value = r.FormValue("CustomFields2")
			mdl.CustomFields2 = customfields["CustomFields2"][0]
			mdl.CustomFields2Type = customfields["CustomFields2"][1]
		}

		if r.FormValue("CustomFields3") != "" {
			mdl.CustomFields3Value = r.FormValue("CustomFields3")
			mdl.CustomFields3 = customfields["CustomFields3"][0]
			mdl.CustomFields3Type = customfields["CustomFields3"][1]
		}

		if r.FormValue("CustomFields4") != "" {
			mdl.CustomFields4Value = r.FormValue("CustomFields4")
			mdl.CustomFields4 = customfields["CustomFields4"][0]
			mdl.CustomFields4Type = customfields["CustomFields4"][1]
		}
		mdl.ModifiedBy = usr.EmployeeID
		_ = p.Irepo.UpadteITAsset(r.Context(), &mdl)
		
	}
}

//ITAssetsBulkEdit cc
func (p *IITAsset) ITAssetsBulkEdit(w http.ResponseWriter, r *http.Request) {
	//	params := mux.Vars(r)
	//	ID := params["id"]
	//	id, _ := strconv.Atoi(ID)
	if r.Method == "GET" {
		SelectedIds := make(map[string]string)
		idbytes := []byte(`` + r.FormValue("SelectedIds") + ``)
		json.Unmarshal(idbytes, &SelectedIds)

		utils.ExecuteTemplate(w, r, "ITAssetEditTmplBulk", SelectedIds)

	} else if r.Method == "POST" {

		mdl := ITAssetsmodel.ITAssetModel{}

		data := make(map[string]string)
		//idsMap := make(map[string]string)
		dataByte := []byte(`` + r.FormValue("data") + ``)
		json.Unmarshal(dataByte, &data)

		mdl.ITAssetDescription = data["ITAssetDescription"]

		if data["ITAssetWarranty"] != "" {
			mdl.ITAssetWarranty = data["ITAssetWarranty"]
		}
		VendorID, _ := strconv.Atoi(data["Vendor"])
		mdl.Vendor = VendorID
		LocationID, _ := strconv.Atoi(data["Location"])
		mdl.Location = LocationID

		ids := data["IDS"]
		var s []string
		//	json.Unmarshal(idsByte, &idsMap)
		if len(ids) > 0 {
			ids = ids[0 : len(ids)-1]
			s = strings.Split(ids, ",")
		}

		_ = p.Irepo.ITAssetsBulkEdit(r.Context(), &mdl, &s)
		http.Redirect(w, r, "/ITAssets", 301)
	}
}

// ITAssetsCheckout ..
func (p *IITAsset) ITAssetsCheckout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	id, _ := strconv.Atoi(ID)
	usr, Auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		//LocID, _ := strconv.Atoi(r.URL.Query().Get("LocID"))
		data := p.returnITAssetsByID(r, id)

		mapRoles, _ := p.ICmnrepo.GetMultiLevelApproval_Map(r.Context())
		var RoleID int
		var Grade int
		for i := 0; i < len(mapRoles); i++ {
			if mapRoles[i].FeatureName == "OutWard Request" {
				RoleID = mapRoles[i].MultiLevelApproval_Map.RoleID
				Grade = mapRoles[i].MultiLevelApproval_Map.Grade
				break
			}
		}
		strct := struct {
			Mdl    *ITAssetsmodel.ITAssetModel
			RoleID int
			Grade  int
		}{
			Mdl:    data,
			RoleID: RoleID,
			Grade:  Grade,
		}
		Mapdata := CmnModel.TemplateData{
			Data: strct,
			User: usr,
			Auth: Auth,
		}
		utils.ExecuteTemplate(w, r, "ITAssetCheckoutTmpl", Mapdata)

	} else if r.Method == "POST" {
		mdl := ITAssetsmodel.ITassetCheckout{}
		mdl.AssetID = id
		mdl.CheckedOutTo = r.FormValue("AssignTo")
		CheckedOutUserID, err := strconv.Atoi(r.FormValue("Users"))
		if CheckedOutUserID != 0 {
			mdl.CheckedOutUserID = CheckedOutUserID
			mdl.CheckedOutDate = r.FormValue("CheckOutDate")
			mdl.ExpectedCheckInDate = r.FormValue("ExpectedCheckInDate")
		}
		CheckedOutAssetID, err := strconv.Atoi(r.FormValue("Assets"))
		if CheckedOutAssetID != 0 {
			mdl.CheckedOutAssetID = CheckedOutAssetID
			mdl.CheckedOutDate = r.FormValue("CheckOutDate")
			mdl.ExpectedCheckInDate = r.FormValue("ExpectedCheckInDate")
		}

		mdl.Comments = r.FormValue("Description")
		if mdl.CheckedOutTo == "Location" {
			IW := CmnModel.InWardOutWard{}
			//IW.ApproverEmpID, _ = strconv.Atoi(r.FormValue("Approver"))
			IW.ReceiverEmpID, _ = strconv.Atoi(r.FormValue("Receiver"))
			IW.SenderEmpID, _ = strconv.Atoi(r.FormValue("EmpID"))
			IW.FromLocationID, _ = strconv.Atoi(r.FormValue("AssetLocation"))
			IW.ToLocationID, _ = strconv.Atoi(r.FormValue("Locations"))
			IW.Description = r.FormValue("Description")
			IW.TransferStatus = 9
			transactionid := "TR" + time.Now().Format("2006-01-02 15:04:05")

			rpl := strings.NewReplacer(":", "",
				"-", "",
				" ", "")
			res := rpl.Replace(transactionid)
			IW.TransactionID = res
			//

			ListIWA := []*CmnModel.InWardOutWardAsset{}
			IWA := CmnModel.InWardOutWardAsset{}
			IWA.AssetID = mdl.AssetID
			IWA.AssetType = "itasset"
			IWA.Quantity = 1
			IWA.TransferStatus = 9
			ListIWA = append(ListIWA, &IWA)
			IW.ListInWardOutWardAsset = ListIWA
			IW.CreatedBy = usr.EmployeeID

			aprvl := CmnModel.InWardOutWardApproval{}
			aprvl.ApproverID, _ = strconv.Atoi(r.FormValue("Approver"))
			aprvl.Grade, _ = strconv.Atoi(r.FormValue("Grade"))
			aprvl.RoleID, _ = strconv.Atoi(r.FormValue("ApprvrRoleID"))

			aprvl.Status = 9
			IW.InWardOutWardApproval = aprvl
			err = p.ICmnrepo.CreateInWardOutWard(r.Context(), &IW)
		} else {
			mdl.CheckOut_By = usr.EmployeeID
			mdl.CreatedBy = usr.EmployeeID
			err = p.Irepo.CreateITAssetsCheckout(r.Context(), &mdl)
		}

		if err == nil {
			http.Redirect(w, r, "/ITAssets/view/"+params["id"]+"", 301)
		}
	}
}

// ITAssetViewFiles ..
func (p *IITAsset) ITAssetViewFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		params := mux.Vars(r)
		ID := params["id"]
		id, _ := strconv.Atoi(ID)
		utils.ExecuteTemplate(w, r, "ITAssetView_FilesTmpl", id)
	}
}

//GetITassetsFilesByID ..
func (p *IITAsset) GetITassetsFilesByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	id, _ := strconv.Atoi(ID)
	body, _ := p.Irepo.GetITassetsFilesByID(r.Context(), id)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

//ITAssetRetire ..
func (p *IITAsset) ITAssetRetire(w http.ResponseWriter, r *http.Request) {
	mdl := CmnModel.Retire{}
	mdl.AssetID, _ = strconv.Atoi(r.FormValue("AssetID"))
	mdl.Reason, _ = strconv.Atoi(r.FormValue("Reason"))
	mdl.RetireDate = r.FormValue("RetireDate")
	mdl.Commnets = r.FormValue("Commnets")
	mdl.RetiredBy, _ = strconv.Atoi(r.FormValue("RetiredBy"))
	body := p.Irepo.ITAssetRetire(r.Context(), &mdl)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

//GetCustomFields ..
func (p *IITAsset) GetCustomFields(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	id, _ := strconv.Atoi(ID)
	Mod := params["Mod"]
	body, _ := p.Irepo.GetCustomFields(r.Context(), id, Mod)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

// CreateITAssetsCheckIn ..
func (p *IITAsset) CreateITAssetsCheckIn(w http.ResponseWriter, r *http.Request) {

	mdl := ITAssetsmodel.ITassetCheckout{}
	mdl.CheckinDate = r.FormValue("CheckInDate")
	mdl.CheckInComments = r.FormValue("CheckInComments")
	IDITAssetCheckOutCheckIN, err := strconv.Atoi(r.FormValue("IDITAssetCheckOutCheckIN"))
	mdl.IDITAssetCheckOutCheckIN = IDITAssetCheckOutCheckIN
	err = p.Irepo.CreateITAssetsCheckIn(r.Context(), &mdl)
	if err != nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	}
}

// UploadFiles ..
func (p *IITAsset) UploadFiles(w http.ResponseWriter, r *http.Request) {
	mdl := ITAssetsmodel.ITassetsFiles{}

	file, handle, err := r.FormFile("ITAssetFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	_file, _ := os.Create("AppFiles/Files/ITAssets/" + handle.Filename)
	mdl.Descrption = r.FormValue("Description")
	size := handle.Size / 1000
	mdl.Size = strconv.Itoa(int(size)) + " Kb"
	//mdl.Path = "AppFiles/Files/ITAssets/" + handle.Filename
	mdl.FileType = strings.Split(handle.Filename, ".")[1]
	mdl.AssetID, _ = strconv.Atoi(r.FormValue("AssetID"))

	io.Copy(_file, file)
	defer file.Close()

	mdl.Name = handle.Filename

	err = p.Irepo.CreateITAssetFiles(r.Context(), &mdl)
	url := "/ITAssets/view/" + r.FormValue("AssetID") + ""
	http.Redirect(w, r, url, 301)

}

//Customfields ..
func (p *IITAsset) Customfields(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	Mod := params["Mod"]
	id, _ := strconv.Atoi(ID)
	utils.ExecuteTemplate(w, r, "Customfields", struct {
		ID  int
		Mod string
	}{
		ID:  id,
		Mod: Mod,
	})
}

//QrGenerator  ..
func (p *IITAsset) QrGenerator(w http.ResponseWriter, r *http.Request) {
	//deleteing all generated qr codes
	dirRead, _ := os.Open("AppFiles/Images/QR/")
	dirFiles, _ := dirRead.Readdir(0)
	// Loop over the directory's files.
	for index := range dirFiles {
		fileHere := dirFiles[index]
		nameHere := fileHere.Name()
		fullPath := "AppFiles/Images/QR/" + nameHere
		os.Remove(fullPath)
	}

	name := r.URL.Query().Get("name")
	Idntfctn := r.URL.Query().Get("Idntfctn")
	//data, _ := p.Irepo.GetITAssets(r.Context(), LocID)
	// for _, item := range data {
	// 	if item.IDITAssets == id {
	// 		utils.ExecuteTemplate(w, r, "QrGenerator", item)
	// 		break
	// 	}
	// }
	utils.ExecuteTemplate(w, r, "QrGenerator", struct {
		Name     string
		Idntfctn string
	}{
		Name:     name,
		Idntfctn: Idntfctn,
	})
}

// PrintQr ..
func (p *IITAsset) PrintQr(w http.ResponseWriter, r *http.Request) {
	size := r.URL.Query().Get("size")
	url := r.URL.Query().Get("url")
	IdntficationNumber := r.URL.Query().Get("IdntficationNumber")
	sizes, err := strconv.Atoi(size)
	if err != nil {
		sizes = 256
	}
	qrCode, err := qr.Encode(url, qr.L, qr.Auto)
	if err != nil {
		fmt.Println(err.Error())
	}
	qrCode, err = barcode.Scale(qrCode, sizes, sizes)
	if err != nil {
		fmt.Println(err.Error())
	}
	// create the output file
	result := strings.Replace(IdntficationNumber, "/", "", -1)

	file, err := os.Create("AppFiles/Images/QR/" + strings.TrimSpace(result) + ".png")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	// encode the barcode as png
	png.Encode(file, qrCode)
}

//ITAssetReadExcel ..
func (p *IITAsset) ITAssetReadExcel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "ITAssetReadExcelTmpl", nil)
	}
	if r.Method == "POST" {
		r.ParseForm()
		resexceldata := make([]map[string]string, 1, 1)
		resmaps := make(map[string]string)
		exceldata := []byte(` ` + r.FormValue("exceldata") + ``)
		maps := []byte(`` + r.FormValue("map") + ``)
		json.Unmarshal(exceldata, &resexceldata)
		json.Unmarshal(maps, &resmaps)

		ListITassetmodel := []*ITAssetsmodel.ITAssetModel{}

		for _, item := range resexceldata {
			ITassetmodel := ITAssetsmodel.ITAssetModel{}
			ITassetmodel.ITAssetName = item[resmaps["ITAssetName"]]
			ITassetmodel.ITAssetModel = item[resmaps["ITAssetModel"]]
			ITassetmodel.ITAssetSerialNo = item[resmaps["ITAssetSerialNo"]]
			ITassetmodel.ITAssetDescription = item[resmaps["ITAssetDescription"]]

			price, err := strconv.ParseFloat(item[resmaps["ITAssetPrice"]], 32)
			if err == nil {
				ITassetmodel.ITAssetPrice = price
			}

			ListITassetmodel = append(ListITassetmodel, &ITassetmodel)
		}
		err := p.Irepo.BulkCreateITAsset(r.Context(), ListITassetmodel)
		if err != nil {
			utils.RespondwithJSON(w, r, http.StatusOK, nil)
		}

	}
}

//ITAsset_Service  ..
func (p *IITAsset) ITAsset_Service(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ITAssetID := params["ITAssetID"]

	mode := r.URL.Query().Get("mode")
	IDITAsset_Services := r.URL.Query().Get("IDITAsset_Services")
	data := map[string]interface{}{
		"mode":               mode,
		"ITAssetID":          ITAssetID,
		"IDITAsset_Services": IDITAsset_Services,
	}
	utils.ExecuteTemplate(w, r, "ITAsset_Service", data)

}

//ITAssetRetire ..
func (p *IITAsset) ITasset_services_Insert(w http.ResponseWriter, r *http.Request) {
	mdl := ITAssetsmodel.ITasset_services{}
	json.NewDecoder(r.Body).Decode(&mdl)
	err := p.Irepo.ITasset_services_Insert(r.Context(), &mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}

}

//ITasset_services_start_Update ..
func (p *IITAsset) ITasset_services_start_Update(w http.ResponseWriter, r *http.Request) {
	mdl := ITAssetsmodel.ITasset_services{}
	json.NewDecoder(r.Body).Decode(&mdl)
	err := p.Irepo.ITasset_services_start_Update(r.Context(), &mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//ITasset_services_Complete_Update ..
func (p *IITAsset) ITasset_services_Complete_Update(w http.ResponseWriter, r *http.Request) {
	mdl := ITAssetsmodel.ITasset_services{}
	json.NewDecoder(r.Body).Decode(&mdl)
	err := p.Irepo.ITasset_services_Complete_Update(r.Context(), &mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//ITasset_services_Extend_Update ..
func (p *IITAsset) ITasset_services_Extend_Update(w http.ResponseWriter, r *http.Request) {
	mdl := ITAssetsmodel.ITasset_services{}
	json.NewDecoder(r.Body).Decode(&mdl)
	err := p.Irepo.ITasset_services_Extend_Update(r.Context(), &mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

// ITAssetView_Maintenance ..
func (p *IITAsset) ITAssetView_Maintenance(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		params := mux.Vars(r)
		ITAssetID := params["ITAssetID"]
		utils.ExecuteTemplate(w, r, "ITAssetView_Maintenance", ITAssetID)
	}
}

//GetITAssetservices_List ..
func (p *IITAsset) GetITAssetservices_List(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ITAssetID := params["ITAssetID"]
	id, _ := strconv.Atoi(ITAssetID)

	body, err := p.Irepo.GetITAssetservices_List(r.Context(), id)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, body)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}

}

func (p *IITAsset) GetITAssetToCheckoutToITasset(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	LocID := params["LocID"]
	LocIDs, _ := strconv.Atoi(LocID)
	AssetID := params["AssetID"]
	AssetIDs, _ := strconv.Atoi(AssetID)
	body, err := p.Irepo.GetITAssetToCheckoutToITasset(r.Context(), LocIDs, AssetIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, body)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}

}

func (p *IITAsset) ITAsset_Service_Request_Resolve(w http.ResponseWriter, r *http.Request) {
	mdl := ITAssetsmodel.ITAsset_service_request{}
	json.NewDecoder(r.Body).Decode(&mdl)
	err := p.Irepo.ITAsset_Service_Request_Resolve(r.Context(), &mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

// ITAsset_Service_Request ..
func (p *IITAsset) ITAsset_Service_Request(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)

	if r.Method == "GET" {
		params := mux.Vars(r)
		ITAssetID := params["ITAssetID"]
		Mapdata := CmnModel.TemplateData{
			Auth: auth,
			User: usr,
			Data: ITAssetID,
		}
		utils.ExecuteTemplate(w, r, "ITAsset_Service_Request", Mapdata)
	} else if r.Method == "POST" {
		mdl := ITAssetsmodel.ITAsset_service_request{}
		json.NewDecoder(r.Body).Decode(&mdl)
		err := p.Irepo.ITAsset_Service_Request(r.Context(), &mdl)

		if err == nil {
			utils.RespondwithJSON(w, r, http.StatusOK, nil)
		} else {
			utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
		}
	}

}

//GetITAsset_service_request_List ..
func (p *IITAsset) GetITAsset_service_request_List(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	EmpID := params["EmpID"]
	id, _ := strconv.Atoi(EmpID)

	body, _ := p.Irepo.GetITAsset_service_request_List(r.Context(), id)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

//ITAsset_Service_Request_List  ..
func (p *IITAsset) ITAsset_Service_Request_List(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		Auth: auth,
		User: usr,
	}
	utils.ExecuteTemplate(w, r, "ITAsset_Service_Request_List", Mapdata)

}

// GetITAssets ..
func (p *IITAsset) GetITAsset_Retired(w http.ResponseWriter, r *http.Request) {
	LocID, _ := strconv.Atoi(r.URL.Query().Get("LocID"))
	EmpID, _ := strconv.Atoi(r.URL.Query().Get("EmpID"))
	data, _ := p.Irepo.GetITAsset_Retired(r.Context(), LocID, EmpID)
	utils.RespondwithJSON(w, r, http.StatusOK, data)
}

//ITAssetRetiredList  ..
func (p *IITAsset) ITAssetRetiredList(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		Auth: auth,
		User: usr,
	}
	utils.ExecuteTemplate(w, r, "ITAssetRetiredList", Mapdata)

}

//EmployeesHistory ..
func (p *IITAsset) ITAssetHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	AssetID := params["AssetID"]
	id := AssetID
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: id,
	}
	utils.ExecuteTemplate(w, r, "ITAsset_History", Mapdata)
}

func (p *IITAsset) GetITAsset_History_ByAssetID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	AssetID := params["AssetID"]
	id, _ := strconv.Atoi(AssetID)
	data, err := p.Irepo.Get_ITAssetsHistory_ByAsset(r.Context(), id)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *IITAsset) ITAssetReqForward(w http.ResponseWriter, r *http.Request) {
	Data := ITAssetsmodel.ITAssetRequestApproval{}
	json.NewDecoder(r.Body).Decode(&Data)
	err := p.Irepo.ITAssetReqForward(r.Context(), &Data)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}

}

func (p *IITAsset) GetITAssetReqListByEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["EmpID"]
	EmpID, _ := strconv.Atoi(ID)
	data, err := p.Irepo.GetITAssetReqListByEmp(r.Context(), EmpID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	}
}

func (p *IITAsset) MyITAssetRequestList(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	utils.ExecuteTemplate(w, r, "MyITAssetRequestList", Mapdata)
}

func (p *IITAsset) ITAssetDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	AssetID := params["AssetID"]
	AssetIDs, _ := strconv.Atoi(AssetID)
	_ = p.Irepo.ITAssetDelete(r.Context(), AssetIDs)

	http.Redirect(w, r, "/ITAssets", 301)

}
