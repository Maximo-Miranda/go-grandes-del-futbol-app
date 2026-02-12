package migrations

import (
	"github.com/goravel/framework/facades"
)

type M20260213000001MakePlayerUserIdRequiredAndDropName struct{}

func (r *M20260213000001MakePlayerUserIdRequiredAndDropName) Signature() string {
	return "20260213000001_make_player_user_id_required_and_drop_name"
}

func (r *M20260213000001MakePlayerUserIdRequiredAndDropName) Up() error {
	// Clean existing data — production users will recreate their profiles on next login
	if _, err := facades.Orm().Query().Exec(`DELETE FROM match_lineups WHERE player_id IN (SELECT id FROM players)`); err != nil {
		return err
	}
	if _, err := facades.Orm().Query().Exec(`DELETE FROM match_events WHERE player_id IN (SELECT id FROM players)`); err != nil {
		return err
	}
	if _, err := facades.Orm().Query().Exec(`DELETE FROM team_players WHERE player_id IN (SELECT id FROM players)`); err != nil {
		return err
	}
	if _, err := facades.Orm().Query().Exec(`TRUNCATE TABLE players RESTART IDENTITY CASCADE`); err != nil {
		return err
	}

	// Make user_id NOT NULL
	if _, err := facades.Orm().Query().Exec(`ALTER TABLE players ALTER COLUMN user_id SET NOT NULL`); err != nil {
		return err
	}

	// Add unique constraint to enforce 1:1 relationship
	if _, err := facades.Orm().Query().Exec(`ALTER TABLE players ADD CONSTRAINT players_user_id_unique UNIQUE (user_id)`); err != nil {
		return err
	}

	// Drop the name column (will use user.name instead)
	_, err := facades.Orm().Query().Exec(`ALTER TABLE players DROP COLUMN name`)
	return err
}

func (r *M20260213000001MakePlayerUserIdRequiredAndDropName) Down() error {
	if _, err := facades.Orm().Query().Exec(`ALTER TABLE players ADD COLUMN name VARCHAR`); err != nil {
		return err
	}
	if _, err := facades.Orm().Query().Exec(`ALTER TABLE players DROP CONSTRAINT IF EXISTS players_user_id_unique`); err != nil {
		return err
	}
	_, err := facades.Orm().Query().Exec(`ALTER TABLE players ALTER COLUMN user_id DROP NOT NULL`)
	return err
}
