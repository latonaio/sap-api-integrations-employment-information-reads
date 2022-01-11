package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PersonIDExternal             string      `json:"personIdExternal"`
			UserID                       string      `json:"userId"`
			LastModifiedDateTime         string      `json:"lastModifiedDateTime"`
			ServiceDate                  string      `json:"serviceDate"`
			EndDate                      string      `json:"endDate"`
			BonusPayExpirationDate       string      `json:"bonusPayExpirationDate"`
			CreatedDateTime              string      `json:"createdDateTime"`
			IsContingentWorker           bool        `json:"isContingentWorker"`
			InitialOptionGrant           string      `json:"initialOptionGrant"`
			InitialStockGrant            string      `json:"initialStockGrant"`
			EmployeeFirstEmployment      string      `json:"employeeFirstEmployment"`
			EligibleForStock             bool        `json:"eligibleForStock"`
			EligibleForSalContinuation   string      `json:"eligibleForSalContinuation"`
			OkToRehire                   string      `json:"okToRehire"`
			AssignmentIDExternal         string      `json:"assignmentIdExternal"`
			ProfessionalServiceDate      string      `json:"professionalServiceDate"`
			SalaryEndDate                string      `json:"salaryEndDate"`
			SeniorityDate                string      `json:"seniorityDate"`
			StartDate                    string      `json:"startDate"`
			HiringNotCompleted           bool        `json:"hiringNotCompleted"`
			IsECRecord                   bool        `json:"isECRecord"`
			PrevEmployeeID               string      `json:"prevEmployeeId"`
			RegretTermination            string      `json:"regretTermination"`
			CreatedOn                    string      `json:"createdOn"`
			LastDateWorked               string      `json:"lastDateWorked"`
			FirstDateWorked              string      `json:"firstDateWorked"`
			OriginalStartDate            string      `json:"originalStartDate"`
			PayrollEndDate               string      `json:"payrollEndDate"`
			BenefitsEligibilityStartDate string      `json:"benefitsEligibilityStartDate"`
			StockEndDate                 string      `json:"StockEndDate"`
			BenefitsEndDate              string      `json:"benefitsEndDate"`
			AssignmentClass              string      `json:"assignmentClass"`
			LastModifiedBy               string      `json:"lastModifiedBy"`
			LastModifiedOn               string      `json:"lastModifiedOn"`
			CreatedBy                    string      `json:"createdBy"`
			PaymentInformationNav        struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"paymentInformationNav"`
			EmpBeneficiaryNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"empBeneficiaryNav"`
			PhotoNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"photoNav"`
			CompInfoNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"compInfoNav"`
			EmpGlobalAssignmentNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"empGlobalAssignmentNav"`
			EmpJobRelationshipNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"empJobRelationshipNav"`
			PersonNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"personNav"`
			EmpWorkPermitNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"empWorkPermitNav"`
			WorkOrderNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"workOrderNav"`
			UserNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"userNav"`
			JobInfoNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobInfoNav"`
			EmpPensionPayoutNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"empPensionPayoutNav"`
			WfRequestNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"wfRequestNav"`
			CostDistributionNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"costDistributionNav"`
			EmpPayCompNonRecurringNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"empPayCompNonRecurringNav"`
		} `json:"results"`
	} `json:"d"`
}
