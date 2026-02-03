package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000008CreateTeamPlayersTable struct{}

func (r *M20260201000008CreateTeamPlayersTable) Signature() string {
	return "20260201000008_create_team_players_table"
}

func (r *M20260201000008CreateTeamPlayersTable) Up() error {
	return facades.Schema().Create("team_players", func(table schema.Blueprint) {
		table.ID("id")
		table.UnsignedBigInteger("team_id")
		table.UnsignedBigInteger("player_id")
		table.Integer("jersey_number").Default(0)
		table.Boolean("is_captain").Default(false)
		table.TimestampsTz()

		table.Foreign("team_id").References("id").On("teams").CascadeOnDelete()
		table.Foreign("player_id").References("id").On("players").CascadeOnDelete()
	})
}

func (r *M20260201000008CreateTeamPlayersTable) Down() error {
	return facades.Schema().DropIfExists("team_players")
}
