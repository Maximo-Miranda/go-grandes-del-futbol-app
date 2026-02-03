package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000005CreateTeamsTable struct{}

func (r *M20260201000005CreateTeamsTable) Signature() string {
	return "20260201000005_create_teams_table"
}

func (r *M20260201000005CreateTeamsTable) Up() error {
	return facades.Schema().Create("teams", func(table schema.Blueprint) {
		table.ID("id")
		table.String("name")
		table.String("logo").Nullable()
		table.String("color").Nullable()
		table.String("contact_phone").Nullable()
		table.TimestampsTz()
	})
}

func (r *M20260201000005CreateTeamsTable) Down() error {
	return facades.Schema().DropIfExists("teams")
}
