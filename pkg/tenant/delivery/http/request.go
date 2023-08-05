package http

type TenantCreateRequest struct {
	OrgName string `json:"org_name"`
}

type TenantCreateResponse struct {
	OrgName     string `json:"org_name"`
	OrgUUID     string `json:"org_uuid"`
	AdminApiKey string `json:"admin_api_key"`
}
