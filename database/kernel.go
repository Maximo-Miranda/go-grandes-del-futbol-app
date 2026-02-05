package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"grandesdelfutbol/database/migrations"
)

type Kernel struct{}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20260201000001CreateUsersTable{},
		&migrations.M20260201000002CreatePlayersTable{},
		&migrations.M20260201000003CreateVenuesTable{},
		&migrations.M20260201000004CreateTournamentsTable{},
		&migrations.M20260201000005CreateTeamsTable{},
		&migrations.M20260201000006CreateGroupsTable{},
		&migrations.M20260201000007CreateTournamentTeamsTable{},
		&migrations.M20260201000008CreateTeamPlayersTable{},
		&migrations.M20260201000009CreateMatchesTable{},
		&migrations.M20260201000010CreateMatchEventsTable{},
		&migrations.M20260201000011CreateMatchLineupsTable{},
		&migrations.M20260201000012CreateStandingsTable{},
		&migrations.M20260201000013CreateJobsTable{},
		&migrations.M20260204000001AddMinutesPlayedToMatchLineups{},
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{}
}
