package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260211000002CreatePitchesTable struct{}

func (r *M20260211000002CreatePitchesTable) Signature() string {
	return "20260211000002_create_pitches_table"
}

func (r *M20260211000002CreatePitchesTable) Up() error {
	return facades.Schema().Create("pitches", func(table schema.Blueprint) {
		table.ID("id")
		table.UnsignedBigInteger("venue_id")
		table.String("name")
		table.String("surface_type")
		table.String("pitch_type")
		table.TimestampsTz()
		table.Foreign("venue_id").References("id").On("venues").CascadeOnDelete()
	})
}

func (r *M20260211000002CreatePitchesTable) Down() error {
	return facades.Schema().DropIfExists("pitches")
}
