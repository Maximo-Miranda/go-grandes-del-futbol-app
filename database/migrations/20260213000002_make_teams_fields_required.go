package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260213000002MakeTeamsFieldsRequired struct{}

func (r *M20260213000002MakeTeamsFieldsRequired) Signature() string {
	return "20260213000002_make_teams_fields_required"
}

func (r *M20260213000002MakeTeamsFieldsRequired) Up() error {
	return facades.Schema().Table("teams", func(table schema.Blueprint) {
		table.String("color").Change()
		table.String("contact_phone").Change()
	})
}

func (r *M20260213000002MakeTeamsFieldsRequired) Down() error {
	return facades.Schema().Table("teams", func(table schema.Blueprint) {
		table.String("color").Nullable().Change()
		table.String("contact_phone").Nullable().Change()
	})
}
