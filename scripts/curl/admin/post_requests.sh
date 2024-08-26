# POST REQUESTS FOR admin/user PATH.

func post_admin_user() {
  curl -X POST http://localhost:8080/admin/user \
    -H "Authorization: Bearer ${LICENSOR_API_KEY}" \
    -H "Content-Type: application/json" \
    --data "{\"Username\": \"${1}\",\"Email\":\"${2}\",\"IsAdmin\":\"${3}\"}" \
    | jq
}

func post_admin_tenant() {
  curl -X POST http://localhost:8080/admin/tenant \
    -H "Authorization: Bearer ${LICENSOR_API_KEY}" \
    -H "Content-Type: application/json" \
    --data "{\"org_name\":\"${1}\"}" \
    | jq
}
