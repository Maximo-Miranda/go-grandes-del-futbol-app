package migrations

import (
	"github.com/goravel/framework/facades"
)

type M20260212000001AddUuidToVenues struct{}

func (r *M20260212000001AddUuidToVenues) Signature() string {
	return "20260212000001_add_uuid_to_venues"
}

func (r *M20260212000001AddUuidToVenues) Up() error {
	_, err := facades.Orm().Query().Exec(`
		ALTER TABLE venues ADD COLUMN uuid UUID DEFAULT gen_random_uuid();
		UPDATE venues SET uuid = gen_random_uuid() WHERE uuid IS NULL;
		ALTER TABLE venues ALTER COLUMN uuid SET NOT NULL;
		CREATE UNIQUE INDEX idx_venues_uuid ON venues(uuid);
	`)
	return err
}

func (r *M20260212000001AddUuidToVenues) Down() error {
	_, err := facades.Orm().Query().Exec(`
		DROP INDEX IF EXISTS idx_venues_uuid;
		ALTER TABLE venues DROP COLUMN IF EXISTS uuid;
	`)
	return err
}
