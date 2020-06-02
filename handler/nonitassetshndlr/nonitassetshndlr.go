package nonitassetshndlr

import (
	//"bytes"
	"encoding/json"
	//"fmt"
	//"image/png"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	//"github.com/boombuler/barcode"
	//"github.com/boombuler/barcode/qr"
	"github.com/gorilla/mux"
	"github.com/premsynfosys/AMS_WEB/models/CmnModel"
	"github.com/premsynfosys/AMS_WEB/models/NonITAssets_mdl"
	"github.com/premsynfosys/AMS_WEB/repository/CmnRep"
	"github.com/premsynfosys/AMS_WEB/repository/NonITAssetRep"
	"github.com/premsynfosys/AMS_WEB/utils"
)

// INonITAsset ..
type INonITAsset struct {
	Irepo    NonITAssetRep.NonITAssetIntrfc
	ICmnrepo CmnRep.CmnRepIntrfc
}

//NewITAssetHandler ..
func NewNonITAssetHandler(api string) *INonITAsset {
	return &INonITAsset{
		Irepo:    NonITAssetRep.NewAPIRepo(api),
		ICmnrepo: CmnRep.NewAPIRepo(api),
	}
}

//NonITAsset_Create ..
func (p *INonITAsset) NonITAsset_Create(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		data := new(NonITAssets_mdl.NonITAssets)
		// if r.URL.Query().Get("mode") == "clone" {
		// 	id := r.URL.Query().Get("id")
		// 	ids, _ := strconv.Atoi(id)
		// 	data, _ = p.Irepo.GetNonITAssetList_BY_AssetID(r.Context(), ids)
		// }
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Auth: auth,
			Data: data,
		}
		utils.ExecuteTemplate(w, r, "NonITAsset_Create", Mapdata)
	} else if r.Method == "POST" {
		mdl := NonITAssets_mdl.NonITAssets{}
		mdlOprtn := NonITAssets_mdl.NonITAssets_Oprtns{}

		mdl.NonITAssets_Master_ID, _ = strconv.Atoi(r.FormValue("IDNonITAssets_Master"))
		mdl.ModelNo = r.FormValue("ModelNo")
		//	mdl.IdentificationNo = r.FormValue("IdentificationNo")
		mdl.Description = r.FormValue("Description")
		mdl.ThresholdQnty, _ = strconv.Atoi(r.FormValue("ThresholdQnty"))
		mdl.StatusID = 25 //purchased
		LocationID, _ := strconv.Atoi(r.FormValue("LocationID"))
		mdl.LocationID = LocationID
		ReOrderStockPrice, err := strconv.ParseFloat(r.FormValue("ReOrderStockPrice"), 32)
		mdl.ReOrderStockPrice = ReOrderStockPrice
		ReOrderQuantity, _ := strconv.Atoi(r.FormValue("ReOrderQuantity"))
		mdl.ReOrderQuantity = ReOrderQuantity
		mdl.TotalQnty, _ = strconv.Atoi(r.FormValue("Quantity"))
		mdl.AvailableQnty, _ = strconv.Atoi(r.FormValue("Quantity"))

		mdl.Created_By = usr.IDEmployees
		img, handle, err := r.FormFile("Img")
		if err == nil {
			file, _ := os.Create("AppFiles/Images/NonITAsset/" + handle.Filename)
			io.Copy(file, img)
			defer file.Close()
			mdl.Img = handle.Filename
			defer img.Close()
		}

		VendorID, _ := strconv.Atoi(r.FormValue("VendorID"))
		mdlOprtn.VendorID = VendorID
		mdlOprtn.Warranty = r.FormValue("Warranty")
		mdlOprtn.Quantity, _ = strconv.Atoi(r.FormValue("Quantity"))
		mdlOprtn.UnitPrice, _ = strconv.ParseFloat(r.FormValue("UnitPrice"), 64)
		mdlOprtn.OrderedBy = usr.IDEmployees
		mdlOprtn.Created_By = usr.IDEmployees
		mdlOprtn.StatusID = 1

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

		mdl.NonITAssets_Oprtns = mdlOprtn

		err = p.Irepo.CreateNonITAsset(r.Context(), &mdl)
		if err == nil {
			utils.RespondwithJSON(w, r, http.StatusOK, nil)
		} else {
			utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
		}

	}

}

//GetNonITAssetMasterList ..
func (p *INonITAsset) GetNonITAssetMasterList(w http.ResponseWriter, r *http.Request) {
	data, err := p.Irepo.GetNonITAssetMasterList(r.Context())
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}

}

//GetNonITAssetLists ..
func (p *INonITAsset) GetNonITAssetLists(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	LocID := params["LocID"]
	id, _ := strconv.Atoi(LocID)
	data, err := p.Irepo.GetNonITAssetLists(r.Context(), id)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}

}

//NonITAssetList ..
func (p *INonITAsset) NonITAssetList(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Auth: auth,
		}
		utils.ExecuteTemplate(w, r, "NonITAssets_List", Mapdata)
	}

}

//NonITAsset_BY_AssetID ..
func (p *INonITAsset) NonITAsset_BY_AssetID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["AssetID"]
	id, _ := strconv.Atoi(ID)
	data, err := p.Irepo.GetNonITAssetList_BY_AssetID(r.Context(), id)
	if err != nil {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}
	usr, auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Auth: auth,
			Data: data,
		}
		utils.ExecuteTemplate(w, r, "NonITAssets_View", Mapdata)
	}

}

//NonITAsset_Edit ..
func (p *INonITAsset) NonITAsset_Edit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDNonITAsset := params["IDNonITAsset"]
	id, _ := strconv.Atoi(IDNonITAsset)
	usr, auth := utils.GetCookieUser(r)

	if r.Method == "GET" {
		data, _ := p.Irepo.GetNonITAssetList_BY_AssetID(r.Context(), id)
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Data: data,
			Auth: auth,
		}
		utils.ExecuteTemplate(w, r, "NonITAsset_Edit", Mapdata)
	}
	if r.Method == "POST" {
		mdl := NonITAssets_mdl.NonITAssets{}

		mdl.IDNonITAsset = id
		mdl.NonITAssets_Master_ID, _ = strconv.Atoi(r.FormValue("NonITAssets_Master_ID"))

		mdl.ModelNo = r.FormValue("ModelNo")
		mdl.Description = r.FormValue("Description")

		mdl.ThresholdQnty, _ = strconv.Atoi(r.FormValue("ThresholdQnty"))
		LocationID, _ := strconv.Atoi(r.FormValue("LocationID"))
		mdl.LocationID = LocationID
		ReOrderStockPrice, _ := strconv.ParseFloat(r.FormValue("ReOrderStockPrice"), 32)
		mdl.ReOrderStockPrice = ReOrderStockPrice
		ReOrderQuantity, _ := strconv.Atoi(r.FormValue("ReOrderQuantity"))
		mdl.ReOrderQuantity = ReOrderQuantity
		mdl.Modified_By = usr.EmployeeID
		mdl.Img = r.FormValue("Img1")
		img, handle, _ := r.FormFile("Img")
		if handle != nil {
			mdl.Img = handle.Filename
			file, _ := os.Create("AppFiles/Images/NonITAsset/" + handle.Filename)
			io.Copy(file, img)
			defer file.Close()
			defer img.Close()
		}

		_ = p.Irepo.PostNonITAssetEdit(r.Context(), &mdl)
		http.Redirect(w, r, "/NonITAsset/AssetID/"+IDNonITAsset, 301)
	}

}

//PostNonITAssets_oprtns_Removestock ..
func (p *INonITAsset) PostNonITAssets_oprtns_Removestock(w http.ResponseWriter, r *http.Request) {
	mdlOprtn := NonITAssets_mdl.NonITAssets_Oprtns{}

	mdlOprtn.Comments = r.FormValue("Comments")
	mdlOprtn.NonITAsset_ID, _ = strconv.Atoi(r.FormValue("NonITAsset_ID"))
	mdlOprtn.Quantity, _ = strconv.Atoi(r.FormValue("Quantity"))
	mdlOprtn.Created_By, _ = strconv.Atoi(r.FormValue("OrderedBy"))
	mdlOprtn.StatusID, _ = strconv.Atoi(r.FormValue("StatusID"))
	err := p.Irepo.PostNonITAssets_oprtns_Removestock(r.Context(), &mdlOprtn)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}

}

//PostNonITAssets_oprtns_AddStock ..
func (p *INonITAsset) PostNonITAssets_oprtns_AddStock(w http.ResponseWriter, r *http.Request) {
	mdlOprtn := NonITAssets_mdl.NonITAssets_Oprtns{}
	VendorID, _ := strconv.Atoi(r.FormValue("VendorID"))
	mdlOprtn.VendorID = VendorID
	mdlOprtn.Warranty = r.FormValue("Warranty")
	mdlOprtn.Comments = r.FormValue("Comments")
	mdlOprtn.NonITAsset_ID, _ = strconv.Atoi(r.FormValue("NonITAsset_ID"))
	mdlOprtn.Quantity, _ = strconv.Atoi(r.FormValue("Quantity"))
	mdlOprtn.OrderedBy, _ = strconv.Atoi(r.FormValue("OrderedBy"))
	mdlOprtn.Created_By, _ = strconv.Atoi(r.FormValue("OrderedBy"))
	mdlOprtn.UnitPrice, _ = strconv.ParseFloat(r.FormValue("UnitPrice"), 64)
	mdlOprtn.StatusID = 25

	err := p.Irepo.PostNonITAssets_oprtns_AddStock(r.Context(), &mdlOprtn)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}

}

//NonITAssetsCheckout ..
func (p *INonITAsset) NonITAssetsCheckout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["AssetID"]
	id, _ := strconv.Atoi(ID)
	usr, Auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		data, _ := p.Irepo.GetNonITAssetLists(r.Context(), usr.LocationID)
		for _, item := range data {
			if item.IDNonITAsset == id {
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
					Mdl    *NonITAssets_mdl.NonITAssets
					RoleID int
					Grade  int
				}{
					Mdl:    item,
					RoleID: RoleID,
					Grade:  Grade,
				}
				Mapdata := CmnModel.TemplateData{
					Data: strct,
					User: usr,
					Auth: Auth,
				}
				utils.ExecuteTemplate(w, r, "NonITAsset_CheckOut", Mapdata)
				break

			}
		}
	} else if r.Method == "POST" {
		mdl := NonITAssets_mdl.NonITAssets_checkout_checkin{}
		mdl.NonITAsset_ID = id
		mdl.CheckedOutTo = r.FormValue("AssignTo")
		mdl.CheckOut_Comments = r.FormValue("CheckOut_Comments")
		mdl.Created_By = usr.EmployeeID
		mdl.CheckOut_By = usr.EmployeeID
		if CheckedOutUserID, err := strconv.Atoi(r.FormValue("Users")); CheckedOutUserID != 0 && err == nil {
			mdl.CheckedOutUserID = CheckedOutUserID
			mdl.CheckedOutDate = r.FormValue("CheckOutDate")
			mdl.CheckOut_Qnty, _ = strconv.Atoi(r.FormValue("CheckOut_Qnty"))
		}
		if r.FormValue("CheckedOutPlace") != "" {
			mdl.CheckedOutPlace = r.FormValue("CheckedOutPlace")
			mdl.CheckedOutDate = r.FormValue("CheckOutDate")
			mdl.CheckOut_Qnty, _ = strconv.Atoi(r.FormValue("CheckOut_Qnty"))
		}
		var err error
		if mdl.CheckedOutTo == "Location" {
			IW := CmnModel.InWardOutWard{}
			//	IW.ApproverEmpID, _ = strconv.Atoi(r.FormValue("Approver"))
			IW.ReceiverEmpID, _ = strconv.Atoi(r.FormValue("Receiver"))
			IW.SenderEmpID, _ = strconv.Atoi(r.FormValue("EmpID"))
			IW.FromLocationID, _ = strconv.Atoi(r.FormValue("AssetLocation"))
			IW.ToLocationID, _ = strconv.Atoi(r.FormValue("Locations"))
			IW.Description = r.FormValue("CheckOut_Comments")
			IW.TransferStatus = 9
			transactionid := "TR" + time.Now().Format("2006-01-02 15:04:05")
			rpl := strings.NewReplacer(":", "",
				"-", "",
				" ", "")
			res := rpl.Replace(transactionid)
			IW.TransactionID = res

			ListIWA := []*CmnModel.InWardOutWardAsset{}
			IWA := CmnModel.InWardOutWardAsset{}
			IWA.AssetID = mdl.NonITAsset_ID
			IWA.AssetType = "nonitasset"
			IWA.TransferStatus = 9
			IWA.Quantity, _ = strconv.Atoi(r.FormValue("CheckOut_Qnty"))
			ListIWA = append(ListIWA, &IWA)
			IW.ListInWardOutWardAsset = ListIWA
			err = p.ICmnrepo.CreateInWardOutWard(r.Context(), &IW)
		} else {
			err = p.Irepo.PostNonITAssets_CheckOut(r.Context(), &mdl)
		}

		if err == nil {
			http.Redirect(w, r, "/NonITAsset/AssetID/"+params["AssetID"]+"", 301)
		}
	}
}

func (p *INonITAsset) CheckDuplicateNonITAssetEntry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	LocID := params["LocID"]
	locID, _ := strconv.Atoi(LocID)
	MasterID := params["MasterID"]
	MstrID, _ := strconv.Atoi(MasterID)
	data, err := p.Irepo.CheckDuplicateNonITAssetEntry(r.Context(), MstrID, locID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *INonITAsset) Check_Unique_NonITAsset(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ConsumableName := params["NonITAssetName"]
	data, err := p.Irepo.Check_Unique_NonITAsset(r.Context(), ConsumableName)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *INonITAsset) NonITAssetAdd_Partial(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "NonITAssetAdd_Partial", nil)
	}
}
func (p *INonITAsset) NonITAssetMasterAdd(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "NonITAssetMasterAdd", Mapdata)
	}
}

func (p *INonITAsset) NonITAssetemasterInsert(w http.ResponseWriter, r *http.Request) {
	mdl := NonITAssets_mdl.NonITAssets_Master{}
	mdl.NonITAssets_Name = r.FormValue("NonITAssets_Name")
	mdl.NonITAssets_GroupID, _ = strconv.Atoi(r.FormValue("NonITAssets_GroupID"))
	mdl.NonITAssets_GroupName = r.FormValue("NonITAssets_GroupName")
	err := p.Irepo.NonITAssetemasterInsert(r.Context(), &mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *INonITAsset) GetNonITAssetGroups(w http.ResponseWriter, r *http.Request) {

	data, err := p.Irepo.GetNonITAssetGroups(r.Context())
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *INonITAsset) GetNonITAssetOprtnListByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDNonITAsset := params["IDNonITAsset"]
	id, _ := strconv.Atoi(IDNonITAsset)
	data, err := p.Irepo.GetNonITAssetOprtnListByID(r.Context(), id)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *INonITAsset) NonITAssetOprtnListByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDConsumables := params["IDNonITAsset"]
	id, _ := strconv.Atoi(IDConsumables)
	if r.Method == "GET" {
		usr, auth := utils.GetCookieUser(r)
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Auth: auth,
			Data: id,
		}
		utils.ExecuteTemplate(w, r, "NonITAssetOprtnList", Mapdata)
	}

}

func (p *INonITAsset) NonITAssetRequest(w http.ResponseWriter, r *http.Request) {
	user, auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		NonITAssetGroups, _ := p.Irepo.GetNonITAssetGroups(r.Context())
		LocID, _ := strconv.Atoi(r.URL.Query().Get("LocID"))
		mapRoles, _ := p.ICmnrepo.GetMultiLevelApproval_Map(r.Context())
		var RoleID int
		var Grade int
		userroleid := user.RoleID
		for i := 0; i < len(mapRoles); i++ {
			if mapRoles[i].FeatureName == "NonITAsset Request" {
				if mapRoles[i].MultiLevelApproval_Map.RoleID != userroleid && RoleID == 0 {
					RoleID = mapRoles[i].MultiLevelApproval_Map.RoleID
					Grade = mapRoles[i].MultiLevelApproval_Map.Grade
				} else if userroleid == mapRoles[i].MultiLevelApproval_Map.RoleID {
					if len(mapRoles) > i+1 {
						if mapRoles[i+1].FeatureName == "NonITAsset Request" {
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
			NonITAssetGroups []*NonITAssets_mdl.NonITAssets_Groups
			UsersEmployees   []*CmnModel.User
			Grade            int
			RoleID           int
		}{
			NonITAssetGroups: NonITAssetGroups,
			UsersEmployees:   Approvers,
			Grade:            Grade,
			RoleID:           RoleID,
		}

		Mapdata := CmnModel.TemplateData{
			Data: mapData,
			Auth: auth,
			User: user,
		}

		utils.ExecuteTemplate(w, r, "NonITAssetRequest", Mapdata)
	}
}

func (p *INonITAsset) PostNonITAssetRequest(w http.ResponseWriter, r *http.Request) {
	ListData := []*NonITAssets_mdl.NonITAssetRequest{}
	json.NewDecoder(r.Body).Decode(&ListData)
	err := p.Irepo.CreateNonITAssetRequest(r.Context(), ListData)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}

}
func (p *INonITAsset) GetNonITAssetReqList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Apprvrid := params["Apprvrid"]
	id, _ := strconv.Atoi(Apprvrid)
	data, err := p.Irepo.GetNonITAssetReqList(r.Context(), id)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *INonITAsset) NonITAssetRequestList(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)
	mapRoles, _ := p.ICmnrepo.GetMultiLevelApproval_Map(r.Context())
	requstmap := []*CmnModel.MultiLevelApproval_Main{}
	for _, itm := range mapRoles {
		if itm.FeatureName == "NonITAsset Request" {
			requstmap = append(requstmap, itm)
		}
	}
	Mapdata := CmnModel.TemplateData{
		Auth: auth,
		User: usr,
		Data: requstmap,
	}

	utils.ExecuteTemplate(w, r, "NonITAssetRequestList", Mapdata)
}

func (p *INonITAsset) NonITAssetReq_ApprovalStatusList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["ReqGroupID"]
	ReqGroupID, _ := strconv.Atoi(ID)
	data, err := p.Irepo.NonITAssetReq_ApprovalStatusList(r.Context(), ReqGroupID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	}
}
func (p *INonITAsset) NonITAssetReqReject(w http.ResponseWriter, r *http.Request) {
	mdl := NonITAssets_mdl.NonITAssetRequestApproval{}
	json.NewDecoder(r.Body).Decode(&mdl)
	data, err := p.Irepo.NonITAssetReqReject(r.Context(), &mdl)
	if err == nil && data {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//NonITAssetReqForward //
func (p *INonITAsset) NonITAssetReqForward(w http.ResponseWriter, r *http.Request) {
	Data := NonITAssets_mdl.NonITAssetRequestApproval{}
	json.NewDecoder(r.Body).Decode(&Data)
	err := p.Irepo.NonITAssetReqForward(r.Context(), &Data)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}

}

func (p *INonITAsset) NonITAssetReqListByReqGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["ReqGroupID"]
	ReqGroupID, _ := strconv.Atoi(ID)
	ApproverID := params["ApproverID"]
	approverID, _ := strconv.Atoi(ApproverID)
	data, err := p.Irepo.NonITAssetReqListByReqGroup(r.Context(), ReqGroupID, approverID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}
}

func (p *INonITAsset) NonITAssetReqApprove(w http.ResponseWriter, r *http.Request) {
	mdl := []*NonITAssets_mdl.NonITAssetRequest{}
	json.NewDecoder(r.Body).Decode(&mdl)
	err := p.Irepo.NonITAssetReqApprove(r.Context(), mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}
}

func (p *INonITAsset) GetNonITAssetCheckinDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	LocID := params["LocID"]
	LocIDs, _ := strconv.Atoi(LocID)

	data, err := p.Irepo.GetNonITAssetCheckinDetails(r.Context(), LocIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}
}

func (p *INonITAsset) NonITAssetCheckinDetails(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		Auth: auth,
		User: usr,
	}
	utils.ExecuteTemplate(w, r, "NonITAssetCheckinDetails", Mapdata)
}

func (p *INonITAsset) NonITAssetCheckinDetails_Partial(w http.ResponseWriter, r *http.Request) {

	mode := r.URL.Query().Get("mode")
	ID, _ := strconv.Atoi(r.URL.Query().Get("ID"))

	utils.ExecuteTemplate(w, r, "NonITAssetCheckinDetails_Partial", struct {
		Mode string
		ID   int
	}{
		Mode: mode,
		ID:   ID,
	})
}

func (p *INonITAsset) NonITAssetCheckin(w http.ResponseWriter, r *http.Request) {

	mdl := NonITAssets_mdl.NonITAssets_checkin{}
	json.NewDecoder(r.Body).Decode(&mdl)
	err := p.Irepo.NonITAssetCheckin(r.Context(), &mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}
}

func (p *INonITAsset) Getnonitassets_checkinByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	CheckinID := params["CheckinID"]
	CheckinIDs, _ := strconv.Atoi(CheckinID)

	data, err := p.Irepo.Getnonitassets_checkinByID(r.Context(), CheckinIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}
}

func (p *INonITAsset) GetNonITAssetCheckinDetailsByAsset(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDNonITAsset := params["IDNonITAsset"]
	IDNonITAssets, _ := strconv.Atoi(IDNonITAsset)

	data, err := p.Irepo.GetNonITAssetCheckinDetailsByAsset(r.Context(), IDNonITAssets)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}
}

func (p *INonITAsset) GetNonITAssetCheckinDetailsByEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	EmpID := params["EmpID"]
	EmpIDs, _ := strconv.Atoi(EmpID)

	data, err := p.Irepo.GetNonITAssetCheckinDetailsByEmp(r.Context(), EmpIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
	}
}

func (p *INonITAsset) GetNonITAssetReqListByEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	EmpID := params["EmpID"]
	id, _ := strconv.Atoi(EmpID)
	data, err := p.Irepo.GetNonITAssetReqListByEmp(r.Context(), id)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *INonITAsset) MyNonITAssetRequestList(w http.ResponseWriter, r *http.Request) {
	usr, auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		Auth: auth,
		User: usr,
	}
	utils.ExecuteTemplate(w, r, "MyNonITAssetRequestList", Mapdata)
}


func (p *INonITAsset) NonITAssetDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	AssetID := params["AssetID"]
	AssetIDs, _ := strconv.Atoi(AssetID)
	_ = p.Irepo.NonITAssetDelete(r.Context(), AssetIDs)

	http.Redirect(w, r, "/NonITAsset/List", 301)

}