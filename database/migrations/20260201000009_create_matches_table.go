package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000009CreateMatchesTable struct{}

func (r *M20260201000009CreateMatchesTable) Signature() string {
	return "20260201000009_create_matches_table"
}

func (r *M20260201000009CreateMatchesTable) Up() error {
	return facades.Schema().Create("matches", func(table schema.Blueprint) {
		table.ID("id")
		table.UnsignedBigInteger("tournament_id")
		table.UnsignedBigInteger("group_id").Nullable()
		table.UnsignedBigInteger("home_team_id")
		table.UnsignedBigInteger("away_team_id")
		table.UnsignedBigInteger("venue_id").Nullable()
		table.TimestampTz("match_date").Nullable()
		table.String("status").Default("scheduled")
		table.Integer("home_score").Default(0)
		table.Integer("away_score").Default(0)
		table.Integer("round").Default(1)
		table.TimestampsTz()

		table.Foreign("tournament_id").References("id").On("tournaments").CascadeOnDelete()
		table.Foreign("group_id").References("id").On("groups").NullOnDelete()
		table.Foreign("home_team_id").References("id").On("teams").CascadeOnDelete()
		table.Foreign("away_team_id").References("id").On("teams").CascadeOnDelete()
		table.Foreign("venue_id").References("id").On("venues").NullOnDelete()
	})
}

func (r *M20260201000009CreateMatchesTable) Down() error {
	return facades.Schema().DropIfExists("matches")
}
