package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000002CreatePlayersTable struct{}

func (r *M20260201000002CreatePlayersTable) Signature() string {
	return "20260201000002_create_players_table"
}

func (r *M20260201000002CreatePlayersTable) Up() error {
	return facades.Schema().Create("players", func(table schema.Blueprint) {
		table.ID("id")
		table.UnsignedBigInteger("user_id").Nullable()
		table.String("name")
		table.String("nickname").Nullable()
		table.String("document_id").Nullable()
		table.String("phone").Nullable()
		table.String("position").Nullable()
		table.String("photo").Nullable()
		table.TimestampsTz()

		table.Foreign("user_id").References("id").On("users").CascadeOnDelete()
	})
}

func (r *M20260201000002CreatePlayersTable) Down() error {
	return facades.Schema().DropIfExists("players")
}
