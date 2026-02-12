package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20260211000001ModifyVenuesAddPhotoRemoveFields struct{}

func (r *M20260211000001ModifyVenuesAddPhotoRemoveFields) Signature() string {
	return "20260211000001_modify_venues_add_photo_remove_fields"
}

func (r *M20260211000001ModifyVenuesAddPhotoRemoveFields) Up() error {
	return facades.Schema().Table("venues", func(table schema.Blueprint) {
		table.String("photo").Nullable()
		table.DropColumn("surface_type")
		table.DropColumn("capacity")
	})
}

func (r *M20260211000001ModifyVenuesAddPhotoRemoveFields) Down() error {
	return facades.Schema().Table("venues", func(table schema.Blueprint) {
		table.DropColumn("photo")
		table.String("surface_type")
		table.Integer("capacity").Default(0)
	})
}
