package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260213000003MakeMatchesFieldsRequired struct{}

func (r *M20260213000003MakeMatchesFieldsRequired) Signature() string {
	return "20260213000003_make_matches_fields_required"
}

func (r *M20260213000003MakeMatchesFieldsRequired) Up() error {
	err := facades.Schema().Table("matches", func(table schema.Blueprint) {
		table.TimestampTz("match_date").Change()
		table.Integer("round").Change()
	})
	if err != nil {
		return err
	}

	_, err = facades.Orm().Query().Exec(`
		ALTER TABLE matches ALTER COLUMN venue_id SET NOT NULL;
	`)
	return err
}

func (r *M20260213000003MakeMatchesFieldsRequired) Down() error {
	err := facades.Schema().Table("matches", func(table schema.Blueprint) {
		table.TimestampTz("match_date").Nullable().Change()
		table.Integer("round").Default(1).Change()
	})
	if err != nil {
		return err
	}

	_, err = facades.Orm().Query().Exec(`
		ALTER TABLE matches ALTER COLUMN venue_id DROP NOT NULL;
	`)
	return err
}
