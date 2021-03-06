package sap_api_input_reader

type EC_MC struct {
	ConnectionKey    string `json:"connection_key"`
	Result           bool   `json:"result"`
	RedisKey         string `json:"redis_key"`
	Filepath         string `json:"filepath"`
	PersonIDExternal struct {
		PersonIDExternal         string `json:"candidate_id"`
		ShipToParty              string `json:"deliver_to"`
		OriginalDeliveryQuantity string `json:"quantity"`
		ActualDeliveryQuantity   string `json:"picked_quantity"`
		Price                    string `json:"price"`
		Batch                    string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string `json:"api_schema"`
	MaterialCode            string `json:"material_code"`
	ShippingPoint           string `json:"plant/supplier"`
	Stock                   string `json:"stock"`
	DeliveryDocumentType    string `json:"document_type"`
	DeliveryDocument        string `json:"document_no"`
	PlannedGoodsIssueDate   string `json:"planned_date"`
	ActualGoodsMovementDate string `json:"validated_date"`
	Deleted                 bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey         string `json:"connection_key"`
	Result                bool   `json:"result"`
	RedisKey              string `json:"redis_key"`
	Filepath              string `json:"filepath"`
	EmploymentInformation struct {
		PersonIDExternal             string      `json:"PersonIDExternal"`
		UserID                       string      `json:"UserID"`
		LastModifiedDateTime         string      `json:"LastModifiedDateTime"`
		ServiceDate                  string      `json:"ServiceDate"`
		EndDate                      string      `json:"EndDate"`
		BonusPayExpirationDate       string      `json:"BonusPayExpirationDate"`
		CreatedDateTime              string      `json:"CreatedDateTime"`
		IsContingentWorker           bool        `json:"IsContingentWorker"`
		InitialOptionGrant           string      `json:"InitialOptionGrant"`
		InitialStockGrant            string      `json:"InitialStockGrant"`
		EmployeeFirstEmployment      string      `json:"EmployeeFirstEmployment"`
		EligibleForStock             bool        `json:"EligibleForStock"`
		EligibleForSalContinuation   string      `json:"EligibleForSalContinuation"`
		OkToRehire                   string      `json:"OkToRehire"`
		AssignmentIDExternal         string      `json:"AssignmentIDExternal"`
		ProfessionalServiceDate      string      `json:"ProfessionalServiceDate"`
		SalaryEndDate                string      `json:"SalaryEndDate"`
		SeniorityDate                string      `json:"SeniorityDate"`
		StartDate                    string      `json:"StartDate"`
		HiringNotCompleted           bool        `json:"HiringNotCompleted"`
		IsECRecord                   bool        `json:"IsECRecord"`
		PrevEmployeeID               string      `json:"PrevEmployeeID"`
		RegretTermination            string      `json:"RegretTermination"`
		CreatedOn                    string      `json:"CreatedOn"`
		LastDateWorked               string      `json:"LastDateWorked"`
		FirstDateWorked              string      `json:"FirstDateWorked"`
		OriginalStartDate            string      `json:"OriginalStartDate"`
		PayrollEndDate               string      `json:"PayrollEndDate"`
		BenefitsEligibilityStartDate string      `json:"BenefitsEligibilityStartDate"`
		StockEndDate                 string      `json:"StockEndDate"`
		BenefitsEndDate              string      `json:"BenefitsEndDate"`
		AssignmentClass              string      `json:"AssignmentClass"`
		LastModifiedBy               string      `json:"LastModifiedBy"`
		LastModifiedOn               string      `json:"LastModifiedOn"`
		CreatedBy                    string      `json:"CreatedBy"`
		PersonNav                    struct {
			LastModifiedDateTime string `json:"LastModifiedDateTime"`
			LastModifiedBy       string `json:"LastModifiedBy"`
			CreatedDateTime      string `json:"CreatedDateTime"`
			DateOfBirth          string `json:"DateOfBirth"`
			PerPersonUUID        string `json:"PerPersonUUID"`
			CreatedOn            string `json:"CreatedOn"`
			LastModifiedOn       string `json:"LastModifiedOn"`
			CountryOfBirth       string `json:"CountryOfBirth"`
			CreatedBy            string `json:"CreatedBy"`
			PersonID             string `json:"PersonID"`
			PersonalInfoNav      struct {
				StartDate            string `json:"StartDate"`
				LastName             string `json:"LastName"`
				BusinessLastName     string `json:"BusinessLastName"`
				LastModifiedDateTime string `json:"LastModifiedDateTime"`
				Gender               string `json:"Gender"`
				EndDate              string `json:"EndDate"`
				DisplayName          string `json:"DisplayName"`
				CreatedDateTime      string `json:"CreatedDateTime"`
				Title                string `json:"Title"`
				CreatedOn            string `json:"CreatedOn"`
				BusinessFirstName    string `json:"BusinessFirstName"`
				AttachmentID         string `json:"AttachmentID"`
				PreferredName        string `json:"PreferredName"`
				Initials             string `json:"Initials"`
				LastModifiedBy       string `json:"LastModifiedBy"`
				Script               string `json:"Script"`
				LastModifiedOn       string `json:"LastModifiedOn"`
				FirstName            string `json:"FirstName"`
				Nationality          string `json:"Nationality"`
				CreatedBy            string `json:"CreatedBy"`
				MiddleName           string `json:"MiddleName"`
				NativePreferredLang  string `json:"NativePreferredLang"`
				Salutation           string `json:"Salutation"`
				MaritalStatus        string `json:"MaritalStatus"`
			} `json:"PersonalInfoNav"`
		} `json:"PersonNav"`
	} `json:"EmploymentInformation"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	PersonIDExternal string   `json:"person_id_external"`
	Deleted          bool     `json:"deleted"`
}
