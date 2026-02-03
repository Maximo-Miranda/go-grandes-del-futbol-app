package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000012CreateStandingsTable struct{}

func (r *M20260201000012CreateStandingsTable) Signature() string {
	return "20260201000012_create_standings_table"
}

func (r *M20260201000012CreateStandingsTable) Up() error {
	return facades.Schema().Create("standings", func(table schema.Blueprint) {
		table.ID("id")
		table.UnsignedBigInteger("tournament_id")
		table.UnsignedBigInteger("group_id").Nullable()
		table.UnsignedBigInteger("team_id")
		table.Integer("played").Default(0)
		table.Integer("won").Default(0)
		table.Integer("drawn").Default(0)
		table.Integer("lost").Default(0)
		table.Integer("goals_for").Default(0)
		table.Integer("goals_against").Default(0)
		table.Integer("goal_difference").Default(0)
		table.Integer("points").Default(0)
		table.TimestampsTz()

		table.Foreign("tournament_id").References("id").On("tournaments").CascadeOnDelete()
		table.Foreign("group_id").References("id").On("groups").NullOnDelete()
		table.Foreign("team_id").References("id").On("teams").CascadeOnDelete()
	})
}

func (r *M20260201000012CreateStandingsTable) Down() error {
	return facades.Schema().DropIfExists("standings")
}
