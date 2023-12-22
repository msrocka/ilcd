package ilcd

type ProcessModelling struct {
	Completeness           *Completeness           `xml:"completeness"`
	Reviews                []Review                `xml:"validation>review"`
	ComplianceDeclarations []ComplianceDeclaration `xml:"complianceDeclarations>compliance"`
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
	ReviewDetails         *LangString            `xml:"reviewDetails"`
	Reviewers             []Ref                  `xml:"referenceToNameOfReviewerAndInstitution"`
	OtherDetails          *LangString            `xml:"otherReviewDetails"`
	ReviewReport          *Ref                   `xml:"referenceToCompleteReviewReport"`
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
