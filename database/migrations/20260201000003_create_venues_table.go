package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000003CreateVenuesTable struct{}

func (r *M20260201000003CreateVenuesTable) Signature() string {
	return "20260201000003_create_venues_table"
}

func (r *M20260201000003CreateVenuesTable) Up() error {
	return facades.Schema().Create("venues", func(table schema.Blueprint) {
		table.ID("id")
		table.String("name")
		table.String("address").Nullable()
		table.String("surface_type").Nullable()
		table.Integer("capacity").Default(0)
		table.TimestampsTz()
	})
}

func (r *M20260201000003CreateVenuesTable) Down() error {
	return facades.Schema().DropIfExists("venues")
}
