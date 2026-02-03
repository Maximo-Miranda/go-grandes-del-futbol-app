package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000010CreateMatchEventsTable struct{}

func (r *M20260201000010CreateMatchEventsTable) Signature() string {
	return "20260201000010_create_match_events_table"
}

func (r *M20260201000010CreateMatchEventsTable) Up() error {
	return facades.Schema().Create("match_events", func(table schema.Blueprint) {
		table.ID("id")
		table.UnsignedBigInteger("match_id")
		table.UnsignedBigInteger("player_id")
		table.UnsignedBigInteger("team_id")
		table.String("event_type")
		table.Integer("minute").Default(0)
		table.TimestampsTz()

		table.Foreign("match_id").References("id").On("matches").CascadeOnDelete()
		table.Foreign("player_id").References("id").On("players").CascadeOnDelete()
		table.Foreign("team_id").References("id").On("teams").CascadeOnDelete()
	})
}

func (r *M20260201000010CreateMatchEventsTable) Down() error {
	return facades.Schema().DropIfExists("match_events")
}
