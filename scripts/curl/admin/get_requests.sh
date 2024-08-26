# GET REQUESTS FOR admin/user PATH.

func get_admin_users() {
  curl -X GET http://localhost:8080/admin/user/ \
    -H "Authorization: Bearer ${LICENSOR_API_KEY}" \
    -H "Content-Type: application/json" \
    | jq
}

func get_admin_user() {
  curl -X GET http://localhost:8080/admin/user/${1} \
    -H "Authorization: Bearer ${LICENSOR_API_KEY}" \
    -H "Content-Type: application/json" \
    | jq
}

# GET REQUESTS FOR admin/tenant PATH.

func get_admin_tenants() {
  curl -X GET http://localhost:8080/admin/tenant/ \
    -H "Authorization: Bearer ${LICENSOR_API_KEY}" \
    -H "Content-Type: application/json" \
    | jq
}

func get_admin_tenant() {
  curl -X GET http://localhost:8080/admin/tenant/${1} \
    -H "Authorization: Bearer ${LICENSOR_API_KEY}" \
    -H "Content-Type: application/json" \
    | jq
}
