package ilcd

type ProcessModelling struct {
	InventoryMethod        *ProcessInventoryMethod `xml:"LCIMethodAndAllocation"`
	DataSources            *ProcessSources         `xml:"dataSourcesTreatmentAndRepresentativeness"`
	Completeness           *Completeness           `xml:"completeness"`
	Reviews                []Review                `xml:"validation>review"`
	ComplianceDeclarations []ComplianceDeclaration `xml:"complianceDeclarations>compliance"`
}

type ProcessInventoryMethod struct {
	Type          string                 `xml:"typeOfDataSet"`
	MethodReports []Ref                  `xml:"referenceToLCAMethodDetails"`
	EpdExt        *EpdInventoryMethodExt `xml:"http://lca.jrc.it/ILCD/Common other"`
}

type ProcessSources struct {
	DataHandling []Ref          `xml:"referenceToDataHandlingPrinciples"`
	DataSources  []Ref          `xml:"referenceToDataSource"`
	UseAdvice    LangString     `xml:"useAdviceForDataSet"`
	EpdExt       *EpdSourcesExt `xml:"http://lca.jrc.it/ILCD/Common other"`
}

type Completeness struct {
	ProductModel           *string                      `xml:"completenessProductModel"`
	SupportedImpactMethods []Ref                        `xml:"referenceToSupportedImpactAssessmentMethods"`
	ElementaryFlows        []ElementaryFlowCompleteness `xml:"completenessElementaryFlows"`
	OtherProblemField      *LangString                  `xml:"completenessOtherProblemField"`
}

type ElementaryFlowCompleteness struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

type Review struct {
	Type                  string                 `xml:"type,attr"`
	Scopes                []ReviewScope          `xml:"scope"`
	DataQualityIndicators []DataQualityIndicator `xml:"dataQualityIndicator"`
	ReviewDetails         LangString             `xml:"reviewDetails"`
	Reviewers             []Ref                  `xml:"http://lca.jrc.it/ILCD/Common referenceToNameOfReviewerAndInstitution"`
	OtherDetails          LangString             `xml:"otherReviewDetails"`
	ReviewReport          *Ref                   `xml:"http://lca.jrc.it/ILCD/Common referenceToCompleteReviewReport"`
}

type ReviewScope struct {
	Name    string         `xml:"name,attr"`
	Methods []ReviewMethod `xml:"method"`
}

type ReviewMethod struct {
	Name string `xml:"name,attr"`
}

type DataQualityIndicator struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type ComplianceDeclaration struct {
	ComplianceSystem         *Ref    `xml:"referenceToComplianceSystem"`
	OverallCompliance        *string `xml:"approvalOfOverallCompliance"`
	NomenclatureCompliance   *string `xml:"nomenclatureCompliance"`
	MethodologicalCompliance *string `xml:"methodologicalCompliance"`
	ReviewCompliance         *string `xml:"reviewCompliance"`
	DocumentationCompliance  *string `xml:"documentationCompliance"`
	QualityCompliance        *string `xml:"qualityCompliance"`
}
