package cnsmblhndlr

import (
	"encoding/json"

	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	CmnModel "github.com/premsynfosys/AMS_WEB/models/CmnModel"
	ConsumableModel "github.com/premsynfosys/AMS_WEB/models/ConsumableModel"
	CmnRep "github.com/premsynfosys/AMS_WEB/repository/CmnRep"
	ConsumableRep "github.com/premsynfosys/AMS_WEB/repository/ConsumableRep"
	utils "github.com/premsynfosys/AMS_WEB/utils"
)

//IConsumables ..
type IConsumables struct {
	Irepo    ConsumableRep.ConsumableIntrfc
	ICmnrepo CmnRep.CmnRepIntrfc
}

//NewConsumablesHandler ..
func NewConsumablesHandler(api string) *IConsumables {
	return &IConsumables{
		Irepo:    ConsumableRep.NewAPIRepo(api),
		ICmnrepo: CmnRep.NewAPIRepo(api),
	}
}

//Consumable ..
func (p *IConsumables) Consumable(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := &ConsumableModel.Consumables{}
		// if r.URL.Query().Get("mode") == "clone" {
		// 	id := r.URL.Query().Get("id")
		// 	_id, _ := strconv.Atoi(id)
		// 	data, _ = p.Irepo.GetConsumableListByID(r.Context(), _id)
		// }
		data.Consumablemaster.ConsumableGroups, _ = p.Irepo.GetConsumableGroups(r.Context())
		usr, auth := utils.GetCookieUser(r)
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Data: data,
			Auth: auth,
		}

		utils.ExecuteTemplate(w, r, "Consumables", Mapdata)
	}

}

//ConsumableEdit ..
func (p *IConsumables) ConsumableEdit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDConsumables := params["IDConsumables"]
	id, _ := strconv.Atoi(IDConsumables)
	usr, auth := utils.GetCookieUser(r)
	if r.Method == "GET" {

		data, _ := p.Irepo.GetConsumableListByID(r.Context(), id)
		data.Consumablemaster.ConsumableGroups, _ = p.Irepo.GetConsumableGroups(r.Context())

		Mapdata := CmnModel.TemplateData{
			User: usr,
			Data: data,
			Auth: auth,
		}
		utils.ExecuteTemplate(w, r, "ConsumableEdit", Mapdata)
	}
	if r.Method == "POST" {
		mdl := ConsumableModel.Consumables{}
		// mdl.IDConsumable = id

		// mdl.ConsumableName = r.FormValue("ConsumableName")
		// GroupID, _ := strconv.Atoi(r.FormValue("GroupID"))
		// mdl.GroupID = GroupID
		mdl.IDConsumables = id
		mdl.IDconsumableMaster, _ = strconv.Atoi(r.FormValue("IDconsumableMaster"))
		mdl.Description = r.FormValue("Description")
		mdl.ThresholdQnty, _ = strconv.Atoi(r.FormValue("ThresholdQnty"))
		LocationID, _ := strconv.Atoi(r.FormValue("LocationID"))
		mdl.LocationID = LocationID
		ReOrderStockPrice, _ := strconv.ParseFloat(r.FormValue("ReOrderStockPrice"), 32)
		mdl.ReOrderStockPrice = ReOrderStockPrice
		ReOrderQuantity, _ := strconv.Atoi(r.FormValue("ReOrderQuantity"))
		mdl.ReOrderQuantity = ReOrderQuantity
		mdl.ModifiedBy = usr.EmployeeID
		mdl.Img = r.FormValue("Img1")
		img, handle, _ := r.FormFile("Img")
		if handle != nil {
			mdl.Img = handle.Filename
			file, _ := os.Create("AppFiles/Images/consumables/" + handle.Filename)
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
		_ = p.Irepo.UpdateConsumable(r.Context(), &mdl)
	}

}

//PostCreateConsumable ..
func (p *IConsumables) PostCreateConsumable(w http.ResponseWriter, r *http.Request) {
	usr, _ := utils.GetCookieUser(r)
	mdl := ConsumableModel.Consumables{}
	// mdl.ConsumableName = r.FormValue("ConsumableName")
	// GroupID, _ := strconv.Atoi(r.FormValue("GroupID"))
	// mdl.GroupID = GroupID
	//	SubGroupID, _ := strconv.Atoi(r.FormValue("SubGroupID"))
	//mdl.SubGroupID = SubGroupID

	mdl.IDconsumableMaster, _ = strconv.Atoi(r.FormValue("IDconsumableMaster"))
	//mdl.IdentificationNo = r.FormValue("IdentificationNo")
	mdl.Description = r.FormValue("Description")
	mdl.ThresholdQnty, _ = strconv.Atoi(r.FormValue("ThresholdQnty"))
	//mdl.Warranty = r.FormValue("Warranty")
	mdl.StatusID = 16 //purchased
	LocationID, _ := strconv.Atoi(r.FormValue("LocationID"))
	mdl.LocationID = LocationID
	ReOrderStockPrice, err := strconv.ParseFloat(r.FormValue("ReOrderStockPrice"), 32)
	mdl.ReOrderStockPrice = ReOrderStockPrice
	ReOrderQuantity, _ := strconv.Atoi(r.FormValue("ReOrderQuantity"))
	mdl.ReOrderQuantity = ReOrderQuantity
	mdl.CreatedBy = usr.EmployeeID
	img, handle, err := r.FormFile("Img")
	if err == nil {
		file, _ := os.Create("AppFiles/Images/consumables/" + handle.Filename)
		io.Copy(file, img)
		defer file.Close()
		mdl.Img = handle.Filename
		defer img.Close()
	}

	customfields := make(map[string][]string)
	customfieldsRaw := []byte(r.FormValue("customfields"))
	json.Unmarshal(customfieldsRaw, &customfields)
	Oprtns := ConsumableModel.ConsumableOprtns{}
	Quantity, err := strconv.Atoi(r.FormValue("AddStockQnty"))
	if err == nil {
		Oprtns.Quantity = Quantity
	}
	UnitPrice, err := strconv.ParseFloat(r.FormValue("UnitPrice"), 32)
	if err == nil {
		Oprtns.UnitPrice = UnitPrice
	}
	VendorID, err := strconv.Atoi(r.FormValue("VendorID"))
	if err == nil {
		Oprtns.VendorID = VendorID
	}
	Oprtns.OrderedBy = usr.IDEmployees
	Oprtns.StatusID = 16
	Oprtns.Comments = r.FormValue("Description")
	mdl.ConsumableOprtns = Oprtns

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
	err = p.Irepo.CreateConsumable(r.Context(), &mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//PostConsumableOprtnsAddStock ..
func (p *IConsumables) PostConsumableOprtnsAddStock(w http.ResponseWriter, r *http.Request) {
	mdl := ConsumableModel.ConsumableOprtns{}
	mdl.ConsumableID, _ = strconv.Atoi(r.FormValue("IDConsumables"))
	AddStockQnty, _ := strconv.Atoi(r.FormValue("Quantity"))
	mdl.Quantity = AddStockQnty
	mdl.UnitPrice, _ = strconv.ParseFloat(r.FormValue("UnitPrice"), 32)
	mdl.VendorID, _ = strconv.Atoi(r.FormValue("VendorID"))
	mdl.OrderedBy, _ = strconv.Atoi(r.FormValue("OrderedBy"))
	mdl.Comments = r.FormValue("Comments")
	mdl.StatusID = 16 //purchased
	err := p.Irepo.ConsumableOprtnsAddStock(r.Context(), &mdl)
	if err == nil {
		http.Redirect(w, r, "/Consumables/ConsumableView/"+r.FormValue("IDConsumables"), 301)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//PostConsumableOprtnsRemovestock ..
func (p *IConsumables) PostConsumableOprtnsRemovestock(w http.ResponseWriter, r *http.Request) {
	mdl := ConsumableModel.ConsumableOprtns{}
	mdl.ConsumableID, _ = strconv.Atoi(r.FormValue("IDConsumables"))
	RemoveStockQnty, _ := strconv.Atoi(r.FormValue("Quantity"))
	mdl.Quantity = RemoveStockQnty
	//mdl.UnitPrice, _ = strconv.ParseFloat(r.FormValue("UnitPrice"), 32)
	mdl.OrderedBy, _ = strconv.Atoi(r.FormValue("OrderedBy"))
	mdl.Comments = r.FormValue("Comments")
	mdl.StatusID, _ = strconv.Atoi(r.FormValue("StatusID"))
	err := p.Irepo.PostConsumableOprtnsRemovestock(r.Context(), &mdl)
	if err == nil {
		http.Redirect(w, r, "/Consumables/ConsumableView/"+r.FormValue("IDConsumables"), 301)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//GetConsumableList ..
func (p *IConsumables) GetConsumableList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	LocID := params["LocID"]
	locID, _ := strconv.Atoi(LocID)
	data, err := p.Irepo.GetConsumableList(r.Context(), locID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//ConsumableList ..
func (p *IConsumables) ConsumableList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		usr, auth := utils.GetCookieUser(r)
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Auth: auth,
		}
		utils.ExecuteTemplate(w, r, "ConsumableList", Mapdata)
	}

}

//ConsumableView ..
func (p *IConsumables) ConsumableView(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		params := mux.Vars(r)
		IDConsumables := params["IDConsumables"]
		id, _ := strconv.Atoi(IDConsumables)
		usr, auth := utils.GetCookieUser(r)

		data, err := p.Irepo.GetConsumableListByID(r.Context(), id)
		data.ListStatus, _ = p.ICmnrepo.GetStatus(r.Context(), "consumable")
		if err != nil {
			utils.ExecuteTemplate(w, r, "0", nil)
			return
		}
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Auth: auth,
			Data: data,
		}
		utils.ExecuteTemplate(w, r, "ConsumableView", Mapdata)
	}

}

//GetConsumableMasterList ..
func (p *IConsumables) GetConsumableMasterList(w http.ResponseWriter, r *http.Request) {
	data, err := p.Irepo.GetConsumableMasterList(r.Context())
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//GetConsumableListByID ..
func (p *IConsumables) GetConsumableListByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDConsumables := params["IDConsumables"]
	id, _ := strconv.Atoi(IDConsumables)
	if r.Method == "GET" {
		data, err := p.Irepo.GetConsumableListByID(r.Context(), id)
		if err == nil {
			utils.RespondwithJSON(w, r, http.StatusOK, data)
		} else {
			utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
		}
	}

}

//GetConsumableOprtnListByID ..
func (p *IConsumables) GetConsumableOprtnListByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDConsumables := params["IDConsumables"]
	id, _ := strconv.Atoi(IDConsumables)
	if r.Method == "GET" {
		data, err := p.Irepo.GetConsumableOprtnListByID(r.Context(), id)
		if err == nil {
			utils.RespondwithJSON(w, r, http.StatusOK, data)
		} else {
			utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
		}
	}

}

//GetConsumableOprtnList ..
func (p *IConsumables) GetConsumableOprtnList(w http.ResponseWriter, r *http.Request) {
	data, err := p.Irepo.GetConsumableOprtnList(r.Context())
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}

}

//ConsumableHistory ..
func (p *IConsumables) ConsumableHistory(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		usr, auth := utils.GetCookieUser(r)
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Auth: auth,
		}
		utils.ExecuteTemplate(w, r, "ConsumableHistory", Mapdata)
	}

}

//ConsumableOprtnListByID ..
func (p *IConsumables) ConsumableOprtnListByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDConsumables := params["IDConsumables"]
	id, _ := strconv.Atoi(IDConsumables)
	if r.Method == "GET" {
		usr, auth := utils.GetCookieUser(r)
		Mapdata := CmnModel.TemplateData{
			User: usr,
			Auth: auth,
			Data: id,
		}
		utils.ExecuteTemplate(w, r, "ConsumablesOprtnList", Mapdata)
	}

}

//ConsumableBulkEdit cc
func (p *IConsumables) ConsumableBulkEdit(w http.ResponseWriter, r *http.Request) {
	//	params := mux.Vars(r)
	//	ID := params["id"]
	//	id, _ := strconv.Atoi(ID)
	if r.Method == "GET" {
		SelectedIds := make(map[string]string)
		idbytes := []byte(`` + r.FormValue("SelectedIds") + ``)
		json.Unmarshal(idbytes, &SelectedIds)
		data, _ := p.Irepo.GetConsumableGroups(r.Context())
		Mapdata := CmnModel.TemplateData{
			Data: struct {
				ConsumableGroup []*ConsumableModel.ConsumableGroup
				SelectedIds     map[string]string
			}{
				ConsumableGroup: data,
				SelectedIds:     SelectedIds,
			},
		}
		utils.ExecuteTemplate(w, r, "ConsumablesBulkEdit", Mapdata)

	} else if r.Method == "POST" {
		mdl := ConsumableModel.Consumables{}
		data := make(map[string]string)
		//idsMap := make(map[string]string)
		dataByte := []byte(`` + r.FormValue("data") + ``)
		json.Unmarshal(dataByte, &data)
		mdl.Description = data["Description"]
		// GroupID, _ := strconv.Atoi(data["GroupID"])
		// mdl.GroupID = GroupID
		// SubGroupID, _ := strconv.Atoi(data["SubGroupID"])
		// mdl.SubGroupID = SubGroupID

		ThresholdQnty, _ := strconv.Atoi(data["ThresholdQnty"])
		mdl.ThresholdQnty = ThresholdQnty
		ReOrderStockPrice, _ := strconv.ParseFloat(data["ReOrderStockPrice"], 32)
		mdl.ReOrderStockPrice = ReOrderStockPrice
		ReOrderQuantity, _ := strconv.Atoi(data["ReOrderQuantity"])
		mdl.ReOrderQuantity = ReOrderQuantity

		LocationID, _ := strconv.Atoi(data["LocationID"])
		mdl.LocationID = LocationID

		ids := data["IDS"]
		var s []string
		//	json.Unmarshal(idsByte, &idsMap)
		if len(ids) > 0 {
			ids = ids[0 : len(ids)-1]
			s = strings.Split(ids, ",")
		}

		_ = p.Irepo.ConsumableBulkEdit(r.Context(), &mdl, &s)
		http.Redirect(w, r, "/ITAssets", 301)
	}
}

//ITAssetRetire ..
func (p *IConsumables) CnsmblRetire(w http.ResponseWriter, r *http.Request) {
	mdl := []*CmnModel.InWardOutWardAsset{}
	json.NewDecoder(r.Body).Decode(&mdl)
	body := p.Irepo.CnsmblBulkRetire(r.Context(), mdl)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

func (p *IConsumables) CheckDuplicateAssetEntry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	LocID := params["LocID"]
	locID, _ := strconv.Atoi(LocID)
	MasterID := params["MasterID"]
	MstrID, _ := strconv.Atoi(MasterID)
	data, err := p.Irepo.CheckDuplicateAssetEntry(r.Context(), MstrID, locID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *IConsumables) ConsumableAdd_Partial(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "ConsumableAdd_Partial", nil)
	}
}
func (p *IConsumables) ConsumablemasterInsert(w http.ResponseWriter, r *http.Request) {
	mdl := ConsumableModel.Consumablemaster{}
	mdl.ConsumableName = r.FormValue("consumableName")
	mdl.GroupID, _ = strconv.Atoi(r.FormValue("GroupID"))
	mdl.GroupName = r.FormValue("GroupName")
	err := p.Irepo.ConsumablemasterInsert(r.Context(), &mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *IConsumables) GetConsumableGroups(w http.ResponseWriter, r *http.Request) {

	data, err := p.Irepo.GetConsumableGroups(r.Context())
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *IConsumables) Check_Unique_Consumable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ConsumableName := params["ConsumableName"]
	data, err := p.Irepo.Check_Unique_Consumable(r.Context(), ConsumableName)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *IConsumables) GetVendorsByConsumable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ConsumableID := params["ConsumableID"]
	id, _ := strconv.Atoi(ConsumableID)
	data, err := p.Irepo.GetVendorsByConsumable(r.Context(), id)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}

}

func (p *IConsumables) ConsumableBulkDelete(w http.ResponseWriter, r *http.Request) {
	var ids []string
	json.NewDecoder(r.Body).Decode(&ids)
	err := p.Irepo.ConsumableBulkDelete(r.Context(), ids)

	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *IConsumables) ConsumableDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	AssetID := params["AssetID"]
	AssetIDs, _ := strconv.Atoi(AssetID)
	_ = p.Irepo.ConsumableDelete(r.Context(), AssetIDs)

	http.Redirect(w, r, "/Consumables/ConsumableList", 301)

}

func (p *IConsumables) GetConsumableMastersByVendors(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	VendorID := params["VendorID"]
	id, _ := strconv.Atoi(VendorID)
	data, err := p.Irepo.GetConsumableMastersByVendors(r.Context(), id)

	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *IConsumables) ConsumablesReadExcel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "ConsumableReadExcel", nil)
	}
	if r.Method == "POST" {
		r.ParseForm()
		resexceldata := make([]map[string]string, 1, 1)
		resmaps := make(map[string]string)
		exceldata := []byte(` ` + r.FormValue("exceldata") + ``)
		maps := []byte(`` + r.FormValue("map") + ``)
		json.Unmarshal(exceldata, &resexceldata)
		json.Unmarshal(maps, &resmaps)

		ListITassetmodel := []*ConsumableModel.Consumables{}
		if len(resexceldata) == 0 || len(resmaps) == 0 {
			utils.RespondwithJSON(w, r, http.StatusBadRequest, "No data available in sheet")
			return
		}
		var err error
		for _, item := range resexceldata {
			asset := ConsumableModel.Consumables{}
			asset.IDconsumableMaster, err = strconv.Atoi(item[resmaps["IDconsumableMaster"]])
			if err != nil {
				utils.RespondwithJSON(w, r, http.StatusBadRequest, "Invalid Asset ID  in sheet")
				return
			}
			asset.Description = item[resmaps["Description"]]
			asset.TotalQnty = 0
			asset.ThresholdQnty, err = strconv.Atoi(item[resmaps["ThresholdQnty"]])
			if err != nil {
				utils.RespondwithJSON(w, r, http.StatusBadRequest, "Invalid Threshold Quantity "+item[resmaps["ThresholdQnty"]]+" in sheet")
				return
			}
			asset.ReOrderQuantity, err = strconv.Atoi(item[resmaps["ReOrderQuantity"]])
			if err != nil {
				utils.RespondwithJSON(w, r, http.StatusBadRequest, "Invalid ReOrder Quantity "+item[resmaps["ReOrderQuantity"]]+" in sheet")
				return
			}
			asset.StatusID = 16 //purchased
			asset.LocationID, err = strconv.Atoi(r.URL.Query().Get("locationid"))
			asset.CreatedBy, err = strconv.Atoi(r.URL.Query().Get("createdby"))
			ListITassetmodel = append(ListITassetmodel, &asset)
		}
		err = p.Irepo.BulkCreateConsumables(r.Context(), ListITassetmodel)
		if err == nil {
			utils.RespondwithJSON(w, r, http.StatusOK, nil)
		} else {
			utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
		}

	}
}
