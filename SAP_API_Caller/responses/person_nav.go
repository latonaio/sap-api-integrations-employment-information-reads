package responses

type PersonNav struct {
	D struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		PersonIDExternal     string `json:"personIdExternal"`
		LastModifiedDateTime string `json:"lastModifiedDateTime"`
		LastModifiedBy       string `json:"lastModifiedBy"`
		CreatedDateTime      string `json:"createdDateTime"`
		DateOfBirth          string `json:"dateOfBirth"`
		PerPersonUUID        string `json:"perPersonUuid"`
		CreatedOn            string `json:"createdOn"`
		LastModifiedOn       string `json:"lastModifiedOn"`
		CountryOfBirth       string `json:"countryOfBirth"`
		CreatedBy            string `json:"createdBy"`
		PersonID             string `json:"personId"`
		PersonalInfoNav      struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"personalInfoNav"`
		EmergencyContactNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"emergencyContactNav"`
		SecondaryAssignmentsNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"secondaryAssignmentsNav"`
		PersonEmpTerminationInfoNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"personEmpTerminationInfoNav"`
		PhoneNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"phoneNav"`
		EmploymentNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"employmentNav"`
		CountryOfBirthNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"countryOfBirthNav"`
		PersonRerlationshipNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"personRerlationshipNav"`
		NationalIDNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"nationalIdNav"`
		UserAccountNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"userAccountNav"`
		PersonTypeUsageNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"personTypeUsageNav"`
		EmailNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"emailNav"`
		HomeAddressNavDEFLT struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"homeAddressNavDEFLT"`
	} `json:"d"`
}