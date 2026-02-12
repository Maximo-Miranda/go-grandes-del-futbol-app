package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260212000002MakeTournamentsFieldsRequired struct{}

func (r *M20260212000002MakeTournamentsFieldsRequired) Signature() string {
	return "20260212000002_make_tournaments_fields_required"
}

func (r *M20260212000002MakeTournamentsFieldsRequired) Up() error {
	err := facades.Schema().Table("tournaments", func(table schema.Blueprint) {
		table.String("format").Change()
		table.String("game_type").Change()
		table.String("status").Change()
		table.TimestampTz("start_date").Change()
		table.TimestampTz("end_date").Change()
	})
	if err != nil {
		return err
	}

	_, err = facades.Orm().Query().Exec(`
		ALTER TABLE tournaments ALTER COLUMN venue_id SET NOT NULL;
	`)
	return err
}

func (r *M20260212000002MakeTournamentsFieldsRequired) Down() error {
	err := facades.Schema().Table("tournaments", func(table schema.Blueprint) {
		table.String("format").Default("round_robin").Change()
		table.String("game_type").Default("5v5").Change()
		table.String("status").Default("draft").Change()
		table.TimestampTz("start_date").Nullable().Change()
		table.TimestampTz("end_date").Nullable().Change()
	})
	if err != nil {
		return err
	}

	_, err = facades.Orm().Query().Exec(`
		ALTER TABLE tournaments ALTER COLUMN venue_id DROP NOT NULL;
	`)
	return err
}
