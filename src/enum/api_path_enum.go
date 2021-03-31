package enum

const (
	GroupApiV1 = "/api/v1"
)

const (
	AgentRoot          = "/agent"
	GetAgentPolicyPath = AgentRoot
	UploadPolicyPath   = AgentRoot + "/template"

	PolicyRoot              = "/policy"
	GetAllPolicyPath        = PolicyRoot + "/:type"
	CreatePolicyPath        = PolicyRoot + "/:type"
	UpdatePolicyContentPath = PolicyRoot
	GetOnePolicyPath        = PolicyRoot + "/:type/:policyId"
	UpdatePolicyNamePath    = PolicyRoot + "/:type/:policyId"
	DeletePolicyPath        = PolicyRoot + "/:policyId"

	LicenseRoot          = "license"
	GetPublicKeyPath     = LicenseRoot + "/key"
	GetLicenseEnablePath = LicenseRoot + "/enable"
	UploadLicensePath    = LicenseRoot
	GetLicensePath       = LicenseRoot

	TemplateRoot               = "/template"
	DownloadPolicyTemplatePath = TemplateRoot
)
