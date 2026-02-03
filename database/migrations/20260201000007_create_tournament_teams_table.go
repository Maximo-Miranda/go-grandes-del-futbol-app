package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000007CreateTournamentTeamsTable struct{}

func (r *M20260201000007CreateTournamentTeamsTable) Signature() string {
	return "20260201000007_create_tournament_teams_table"
}

func (r *M20260201000007CreateTournamentTeamsTable) Up() error {
	return facades.Schema().Create("tournament_teams", func(table schema.Blueprint) {
		table.ID("id")
		table.UnsignedBigInteger("tournament_id")
		table.UnsignedBigInteger("team_id")
		table.UnsignedBigInteger("group_id").Nullable()
		table.TimestampsTz()

		table.Foreign("tournament_id").References("id").On("tournaments").CascadeOnDelete()
		table.Foreign("team_id").References("id").On("teams").CascadeOnDelete()
		table.Foreign("group_id").References("id").On("groups").NullOnDelete()
	})
}

func (r *M20260201000007CreateTournamentTeamsTable) Down() error {
	return facades.Schema().DropIfExists("tournament_teams")
}
