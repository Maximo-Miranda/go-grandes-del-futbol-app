package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000013CreateJobsTable struct{}

func (r *M20260201000013CreateJobsTable) Signature() string {
	return "20260201000013_create_jobs_table"
}

func (r *M20260201000013CreateJobsTable) Up() error {
	return facades.Schema().Create("jobs", func(table schema.Blueprint) {
		table.ID("id")
		table.String("queue")
		table.Text("payload")
		table.Integer("attempts").Default(0)
		table.TimestampTz("reserved_at").Nullable()
		table.TimestampTz("available_at")
		table.TimestampsTz()

		table.Index("queue")
	})
}

func (r *M20260201000013CreateJobsTable) Down() error {
	return facades.Schema().DropIfExists("jobs")
}
