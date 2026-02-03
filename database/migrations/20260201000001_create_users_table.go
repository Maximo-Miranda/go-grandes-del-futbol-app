package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260201000001CreateUsersTable struct{}

func (r *M20260201000001CreateUsersTable) Signature() string {
	return "20260201000001_create_users_table"
}

func (r *M20260201000001CreateUsersTable) Up() error {
	return facades.Schema().Create("users", func(table schema.Blueprint) {
		table.ID("id")
		table.String("name")
		table.String("email")
		table.String("password")
		table.String("role").Default("player")
		table.Integer("sign_in_count").Default(0)
		table.TimestampTz("current_sign_in_at").Nullable()
		table.TimestampTz("last_sign_in_at").Nullable()
		table.String("current_sign_in_ip").Nullable()
		table.String("last_sign_in_ip").Nullable()
		table.Integer("failed_attempts").Default(0)
		table.TimestampTz("locked_at").Nullable()
		table.TimestampsTz()

		table.Index("email")
	})
}

func (r *M20260201000001CreateUsersTable) Down() error {
	return facades.Schema().DropIfExists("users")
}
