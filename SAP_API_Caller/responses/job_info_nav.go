package responses

type JobInfoNav struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SeqNumber                                 string      `json:"seqNumber"`
			UserID                                    string      `json:"userId"`
			StartDate                                 string      `json:"startDate"`
			OccupationalLevels                        interface{} `json:"occupationalLevels"`
			WorkscheduleCode                          string      `json:"workscheduleCode"`
			FgtsOptant                                bool        `json:"fgtsOptant"`
			CommitmentIndicator                       interface{} `json:"commitmentIndicator"`
			EffectiveLatestChange                     bool        `json:"effectiveLatestChange"`
			EndDate                                   string      `json:"endDate"`
			ContractType                              interface{} `json:"contractType"`
			MandatoryWorkBreakRecord                  interface{} `json:"mandatoryWorkBreakRecord"`
			CreatedDateTime                           string      `json:"createdDateTime"`
			JobCode                                   string      `json:"jobCode"`
			ValidFrom                                 interface{} `json:"validFrom"`
			FgtsDate                                  interface{} `json:"fgtsDate"`
			PayScaleLevel                             interface{} `json:"payScaleLevel"`
			Division                                  string      `json:"division"`
			TimeTypeProfileCode                       string      `json:"timeTypeProfileCode"`
			EeoClass                                  interface{} `json:"eeoClass"`
			Hazard                                    bool        `json:"hazard"`
			FromCurrency                              interface{} `json:"fromCurrency"`
			Eeo5JobCategory                           interface{} `json:"eeo5JobCategory"`
			FlsaStatus                                interface{} `json:"flsaStatus"`
			EducationalEntity                         interface{} `json:"educationalEntity"`
			CostCenter                                string      `json:"costCenter"`
			ResidentVote                              bool        `json:"residentVote"`
			TimeRecordingProfileCode                  interface{} `json:"timeRecordingProfileCode"`
			LaborProtection                           bool        `json:"laborProtection"`
			ExclExecutiveSector                       bool        `json:"exclExecutiveSector"`
			IsFulltimeEmployee                        bool        `json:"isFulltimeEmployee"`
			EmplStatus                                string      `json:"emplStatus"`
			PayScaleType                              string      `json:"payScaleType"`
			CountryOfCompany                          string      `json:"countryOfCompany"`
			Eeo6JobCategory                           interface{} `json:"eeo6JobCategory"`
			CreatedOn                                 string      `json:"createdOn"`
			MunicipalInseeCode                        interface{} `json:"municipalInseeCode"`
			HealthRisk                                bool        `json:"healthRisk"`
			TimeRecordingAdmissibilityCode            interface{} `json:"timeRecordingAdmissibilityCode"`
			Fte                                       string      `json:"fte"`
			PayGrade                                  string      `json:"payGrade"`
			TravelDistance                            interface{} `json:"travelDistance"`
			Event                                     string      `json:"event"`
			TimeRecordingVariant                      string      `json:"timeRecordingVariant"`
			AssedicCertObjectNum                      interface{} `json:"assedicCertObjectNum"`
			ManagerID                                 string      `json:"managerId"`
			LastModifiedOn                            string      `json:"lastModifiedOn"`
			MandatoryInternship                       interface{} `json:"mandatoryInternship"`
			DynamicBreakConfigCode                    interface{} `json:"dynamicBreakConfigCode"`
			WorkerCategory                            interface{} `json:"workerCategory"`
			DefaultOvertimeCompensationVariant        string      `json:"defaultOvertimeCompensationVariant"`
			BusinessUnit                              string      `json:"businessUnit"`
			LastModifiedDateTime                      string      `json:"lastModifiedDateTime"`
			Notes                                     interface{} `json:"notes"`
			HarmfulAgentExposure                      interface{} `json:"harmfulAgentExposure"`
			JobTitle                                  string      `json:"jobTitle"`
			SickPaySupplement                         interface{} `json:"sickPaySupplement"`
			ElectoralCollegeForWorksCouncil           interface{} `json:"electoralCollegeForWorksCouncil"`
			IntegrationAgent                          interface{} `json:"integrationAgent"`
			FamilyRelationshipWithEmployer            interface{} `json:"familyRelationshipWithEmployer"`
			ExchangeRate                              interface{} `json:"exchangeRate"`
			InternshipLevel                           interface{} `json:"internshipLevel"`
			HolidayCalendarCode                       string      `json:"holidayCalendarCode"`
			EmpRelationship                           interface{} `json:"empRelationship"`
			PermitIndicator                           bool        `json:"permitIndicator"`
			StandardHours                             string      `json:"standardHours"`
			ElectoralCollegeForWorkersRepresentatives interface{} `json:"electoralCollegeForWorkersRepresentatives"`
			Eeo4JobCategory                           interface{} `json:"eeo4JobCategory"`
			EventReason                               string      `json:"eventReason"`
			IsCompetitionClauseActive                 bool        `json:"isCompetitionClauseActive"`
			InternshipSchool                          interface{} `json:"internshipSchool"`
			AssedicCertInitialStateNum                interface{} `json:"assedicCertInitialStateNum"`
			IncludedChallengedPersonQuota             interface{} `json:"includedChallengedPersonQuota"`
			FgtsPercent                               interface{} `json:"fgtsPercent"`
			ToCurrency                                interface{} `json:"toCurrency"`
			Position                                  string      `json:"position"`
			PayScaleArea                              string      `json:"payScaleArea"`
			ProbationPeriodEndDate                    interface{} `json:"probationPeriodEndDate"`
			Timezone                                  string      `json:"timezone"`
			WorkingDaysPerWeek                        string      `json:"workingDaysPerWeek"`
			RegularTemp                               string      `json:"regularTemp"`
			Pcfm                                      bool        `json:"pcfm"`
			WorkLocation                              interface{} `json:"workLocation"`
			ContractReferenceForAed                   interface{} `json:"contractReferenceForAed"`
			IsSideLineJobAllowed                      bool        `json:"isSideLineJobAllowed"`
			Company                                   string      `json:"company"`
			Retired                                   bool        `json:"retired"`
			Department                                string      `json:"department"`
			Eeo1JobCategory                           interface{} `json:"eeo1JobCategory"`
			EmployeeClass                             string      `json:"employeeClass"`
			EmploymentType                            string      `json:"employmentType"`
			LastModifiedBy                            string      `json:"lastModifiedBy"`
			PayScaleGroup                             interface{} `json:"payScaleGroup"`
			EmployeeWorkgroupMembership               interface{} `json:"employeeWorkgroupMembership"`
			PositionEntryDate                         string      `json:"positionEntryDate"`
			CreatedBy                                 string      `json:"createdBy"`
			Location                                  string      `json:"location"`
			Eeo5JobCategoryNav                        struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"eeo5JobCategoryNav"`
			ToCurrencyNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"toCurrencyNav"`
			HolidayCalendarCodeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"holidayCalendarCodeNav"`
			OccupationalLevelsNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"occupationalLevelsNav"`
			CompanyNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"companyNav"`
			PayScaleAreaNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"payScaleAreaNav"`
			DepartmentNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"departmentNav"`
			BusinessUnitNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"businessUnitNav"`
			FromCurrencyNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"fromCurrencyNav"`
			ContractTypeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"contractTypeNav"`
			PayScaleTypeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"payScaleTypeNav"`
			ElectoralCollegeForWorkersRepresentativesNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"electoralCollegeForWorkersRepresentativesNav"`
			EmployeeWorkgroupMembershipNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"employeeWorkgroupMembershipNav"`
			CostCenterNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"costCenterNav"`
			HarmfulAgentExposureNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"harmfulAgentExposureNav"`
			FlsaStatusNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"flsaStatusNav"`
			EmploymentTypeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"employmentTypeNav"`
			UserNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"userNav"`
			PayScaleGroupNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"payScaleGroupNav"`
			TimeTypeProfileCodeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"timeTypeProfileCodeNav"`
			ElectoralCollegeForWorksCouncilNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"electoralCollegeForWorksCouncilNav"`
			DivisionNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"divisionNav"`
			InternshipLevelNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"internshipLevelNav"`
			Eeo6JobCategoryNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"eeo6JobCategoryNav"`
			EeoClassNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"eeoClassNav"`
			WorkscheduleCodeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"workscheduleCodeNav"`
			SickPaySupplementNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"sickPaySupplementNav"`
			EmploymentNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"employmentNav"`
			Eeo4JobCategoryNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"eeo4JobCategoryNav"`
			WorkerCategoryNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"workerCategoryNav"`
			CountryOfCompanyNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"countryOfCompanyNav"`
			CommitmentIndicatorNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"commitmentIndicatorNav"`
			ManagerEmploymentNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"managerEmploymentNav"`
			IncludedChallengedPersonQuotaNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"includedChallengedPersonQuotaNav"`
			FamilyRelationshipWithEmployerNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"familyRelationshipWithEmployerNav"`
			EventNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"eventNav"`
			ManagerUserNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"managerUserNav"`
			LocationNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"locationNav"`
			PayScaleLevelNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"payScaleLevelNav"`
			RegularTempNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"regularTempNav"`
			EventReasonNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"eventReasonNav"`
			EmployeeClassNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"employeeClassNav"`
			JobCodeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobCodeNav"`
			EmplStatusNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"emplStatusNav"`
			MandatoryInternshipNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"mandatoryInternshipNav"`
			InternshipSchoolNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"internshipSchoolNav"`
			MandatoryWorkBreakRecordNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"mandatoryWorkBreakRecordNav"`
			PositionNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"positionNav"`
			WfRequestNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"wfRequestNav"`
			Eeo1JobCategoryNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"eeo1JobCategoryNav"`
			EmpRelationshipNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"empRelationshipNav"`
			PayGradeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"payGradeNav"`
		} `json:"results"`
	} `json:"d"`
}
