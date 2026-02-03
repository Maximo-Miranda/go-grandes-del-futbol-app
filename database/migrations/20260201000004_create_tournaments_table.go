package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000004CreateTournamentsTable struct{}

func (r *M20260201000004CreateTournamentsTable) Signature() string {
	return "20260201000004_create_tournaments_table"
}

func (r *M20260201000004CreateTournamentsTable) Up() error {
	return facades.Schema().Create("tournaments", func(table schema.Blueprint) {
		table.ID("id")
		table.String("name")
		table.Text("description").Nullable()
		table.String("format").Default("round_robin")
		table.String("game_type").Default("5v5")
		table.String("status").Default("draft")
		table.TimestampTz("start_date").Nullable()
		table.TimestampTz("end_date").Nullable()
		table.UnsignedBigInteger("venue_id").Nullable()
		table.TimestampsTz()

		table.Foreign("venue_id").References("id").On("venues").NullOnDelete()
	})
}

func (r *M20260201000004CreateTournamentsTable) Down() error {
	return facades.Schema().DropIfExists("tournaments")
}
