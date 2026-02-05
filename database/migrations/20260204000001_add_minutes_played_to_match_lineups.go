package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260204000001AddMinutesPlayedToMatchLineups struct{}

func (r *M20260204000001AddMinutesPlayedToMatchLineups) Signature() string {
	return "20260204000001_add_minutes_played_to_match_lineups"
}

func (r *M20260204000001AddMinutesPlayedToMatchLineups) Up() error {
	return facades.Schema().Table("match_lineups", func(table schema.Blueprint) {
		table.Integer("minutes_played").Default(0)
	})
}

func (r *M20260204000001AddMinutesPlayedToMatchLineups) Down() error {
	return facades.Schema().Table("match_lineups", func(table schema.Blueprint) {
		table.DropColumn("minutes_played")
	})
}
