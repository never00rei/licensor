package models

import "time"

type Tenants struct {
	OrgID     int       `db:"org_id"`
	OrgName   string    `db:"org_name"`
	OrgUUID   string    `db:"org_uuid"`
	TableName string    `db:"table_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
