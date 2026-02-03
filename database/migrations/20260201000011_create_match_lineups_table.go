package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000011CreateMatchLineupsTable struct{}

func (r *M20260201000011CreateMatchLineupsTable) Signature() string {
	return "20260201000011_create_match_lineups_table"
}

func (r *M20260201000011CreateMatchLineupsTable) Up() error {
	return facades.Schema().Create("match_lineups", func(table schema.Blueprint) {
		table.ID("id")
		table.UnsignedBigInteger("match_id")
		table.UnsignedBigInteger("team_id")
		table.UnsignedBigInteger("player_id")
		table.Boolean("is_starter").Default(true)
		table.TimestampsTz()

		table.Foreign("match_id").References("id").On("matches").CascadeOnDelete()
		table.Foreign("team_id").References("id").On("teams").CascadeOnDelete()
		table.Foreign("player_id").References("id").On("players").CascadeOnDelete()
	})
}

func (r *M20260201000011CreateMatchLineupsTable) Down() error {
	return facades.Schema().DropIfExists("match_lineups")
}
