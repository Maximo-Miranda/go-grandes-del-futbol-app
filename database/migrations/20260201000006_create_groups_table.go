package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000006CreateGroupsTable struct{}

func (r *M20260201000006CreateGroupsTable) Signature() string {
	return "20260201000006_create_groups_table"
}

func (r *M20260201000006CreateGroupsTable) Up() error {
	return facades.Schema().Create("groups", func(table schema.Blueprint) {
		table.ID("id")
		table.UnsignedBigInteger("tournament_id")
		table.String("name")
		table.TimestampsTz()

		table.Foreign("tournament_id").References("id").On("tournaments").CascadeOnDelete()
	})
}

func (r *M20260201000006CreateGroupsTable) Down() error {
	return facades.Schema().DropIfExists("groups")
}
