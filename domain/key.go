package domain

// This represents a key value which will be used to validate and sign JWTs.
type Key struct {
	KeyID   int    `db:"key_id"`
	OrgUUID string `db:"org_uuid"`
	Key     []byte `db:"key"`
}
