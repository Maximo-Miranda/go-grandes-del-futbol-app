<script setup lang="ts">
import { ref } from "vue";
import { Head, usePage, router, Link } from "@inertiajs/vue3";

const page = usePage<{ tournament: any; teams: any[]; matches: any[]; standings: any[]; availableTeams: any[] }>();
const tournament = page.props.tournament;
const teams = page.props.teams || [];
const matches = page.props.matches || [];
const standings = page.props.standings || [];
const availableTeams = ref(page.props.availableTeams || []);

const showAddTeamDialog = ref(false);
const selectedTeamId = ref<number | null>(null);
const loadingTeams = ref(false);

const formatDate = (dateStr: string | null) => {
  if (!dateStr) return "Sin fecha";
  return new Date(dateStr).toLocaleDateString("es-CO", { weekday: "short", day: "numeric", month: "short", hour: "2-digit", minute: "2-digit" });
};

const openAddTeamDialog = async () => {
  loadingTeams.value = true;
  try {
    const response = await fetch(`/tournaments/${tournament.id}/available-teams`);
    const data = await response.json();
    availableTeams.value = data.teams || [];
  } catch (e) {
    console.error(e);
  }
  loadingTeams.value = false;
  showAddTeamDialog.value = true;
};

const addTeam = () => {
  if (selectedTeamId.value) {
    router.post(
      `/tournaments/${tournament.id}/teams`,
      { team_id: selectedTeamId.value },
      {
        preserveScroll: true,
        onSuccess: () => {
          showAddTeamDialog.value = false;
          selectedTeamId.value = null;
          router.reload({ only: ["teams"] });
        },
      }
    );
  }
};

const removeTeam = (teamId: number) => {
  if (confirm("¿Quitar este equipo del torneo?")) {
    router.delete(`/tournaments/${tournament.id}/teams/${teamId}`, {
      preserveScroll: true,
    });
  }
};
</script>

<template>
  <Head><title>{{ tournament.name }}</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">{{ tournament.name }}</h1>
      <div class="d-flex ga-2">
        <Link :href="`/matches/create?tournament_id=${tournament.id}`" class="text-decoration-none">
          <v-btn color="success" variant="tonal" prepend-icon="mdi-soccer">Nuevo Partido</v-btn>
        </Link>
        <Link :href="`/tournaments/${tournament.id}/edit`" class="text-decoration-none">
          <v-btn color="primary" prepend-icon="mdi-pencil">Editar</v-btn>
        </Link>
      </div>
    </div>

    <v-row>
      <v-col cols="12" md="8">
        <v-card class="mb-4">
          <v-card-title>Información</v-card-title>
          <v-card-text>
            <p v-if="tournament.description" class="mb-4">{{ tournament.description }}</p>
            <div class="d-flex flex-wrap ga-2">
              <v-chip>{{ tournament.format }}</v-chip>
              <v-chip>{{ tournament.game_type }}</v-chip>
              <v-chip :color="tournament.status === 'active' ? 'success' : 'grey'">{{ tournament.status }}</v-chip>
              <v-chip v-if="tournament.venue" color="info" prepend-icon="mdi-map-marker">{{ tournament.venue.name }}</v-chip>
              <v-chip v-if="tournament.start_date" prepend-icon="mdi-calendar">
                {{ new Date(tournament.start_date).toLocaleDateString("es-CO") }}
                <span v-if="tournament.end_date"> — {{ new Date(tournament.end_date).toLocaleDateString("es-CO") }}</span>
              </v-chip>
            </div>
          </v-card-text>
        </v-card>

        <v-card v-if="matches.length > 0" class="mb-4">
          <v-card-title class="d-flex justify-space-between align-center">
            <span>Partidos ({{ matches.length }})</span>
            <Link :href="`/matches?tournament_id=${tournament.id}`" class="text-decoration-none">
              <v-btn size="small" variant="text">Ver todos</v-btn>
            </Link>
          </v-card-title>
          <v-list>
            <Link v-for="m in matches" :key="m.id" :href="`/matches/${m.id}`" class="text-decoration-none">
              <v-list-item>
                <v-list-item-title class="d-flex justify-space-between">
                  <span>{{ m.home_team?.name || "TBD" }} vs {{ m.away_team?.name || "TBD" }}</span>
                  <v-chip v-if="m.status === 'completed'" size="x-small" color="success">
                    {{ m.home_score }} - {{ m.away_score }}
                  </v-chip>
                </v-list-item-title>
                <v-list-item-subtitle>{{ formatDate(m.match_date) }} — {{ m.status }}</v-list-item-subtitle>
              </v-list-item>
            </Link>
          </v-list>
        </v-card>
        <v-card v-else class="mb-4">
          <v-card-text class="text-center pa-8 text-medium-emphasis">
            <v-icon size="48" color="grey-lighten-1" class="mb-2">mdi-soccer-field</v-icon>
            <p>No hay partidos programados</p>
            <Link :href="`/matches/create?tournament_id=${tournament.id}`" class="text-decoration-none">
              <v-btn color="primary" variant="tonal" size="small" class="mt-2">Crear Partido</v-btn>
            </Link>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" md="4">
        <v-card class="mb-4">
          <v-card-title class="d-flex justify-space-between align-center">
            <span>Equipos ({{ teams.length }})</span>
            <v-btn size="small" color="primary" variant="tonal" prepend-icon="mdi-plus" @click="openAddTeamDialog">
              Agregar
            </v-btn>
          </v-card-title>
          <v-list density="compact">
            <v-list-item v-for="t in teams" :key="t.team?.id || t.id">
              <template #prepend>
                <v-avatar size="32" color="primary">
                  <v-icon>mdi-shield</v-icon>
                </v-avatar>
              </template>
              <Link :href="`/teams/${t.team?.id || t.team_id}`" class="text-decoration-none">
                <v-list-item-title class="text-primary">{{ t.team?.name || t.name }}</v-list-item-title>
              </Link>
              <template #append>
                <v-btn icon="mdi-close" size="x-small" variant="text" color="error" @click="removeTeam(t.team?.id || t.team_id)" />
              </template>
            </v-list-item>
            <v-list-item v-if="teams.length === 0">
              <v-list-item-title class="text-medium-emphasis">Sin equipos aún</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-card>

        <v-card v-if="standings.length > 0">
          <v-card-title>Clasificación</v-card-title>
          <v-table density="compact">
            <thead><tr><th>#</th><th>Equipo</th><th>Pts</th><th>PJ</th><th>DG</th></tr></thead>
            <tbody>
              <tr v-for="(s, i) in standings" :key="s.id">
                <td>{{ i + 1 }}</td>
                <td>
                  <Link :href="`/teams/${s.team?.id}`" class="text-primary">{{ s.team?.name }}</Link>
                </td>
                <td class="font-weight-bold">{{ s.points }}</td>
                <td>{{ s.played }}</td>
                <td>{{ s.goal_difference > 0 ? '+' : '' }}{{ s.goal_difference }}</td>
              </tr>
            </tbody>
          </v-table>
        </v-card>
      </v-col>
    </v-row>

    <!-- Dialog para agregar equipo -->
    <v-dialog v-model="showAddTeamDialog" max-width="400">
      <v-card>
        <v-card-title>Agregar Equipo al Torneo</v-card-title>
        <v-card-text>
          <v-select
            v-model="selectedTeamId"
            :items="availableTeams"
            item-title="name"
            item-value="id"
            label="Seleccionar Equipo"
            :loading="loadingTeams"
            :disabled="availableTeams.length === 0"
          />
          <p v-if="availableTeams.length === 0 && !loadingTeams" class="text-medium-emphasis">
            No hay equipos disponibles. <Link href="/teams/create" class="text-primary">Crea uno primero</Link>.
          </p>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="showAddTeamDialog = false">Cancelar</v-btn>
          <v-btn color="primary" :disabled="!selectedTeamId" @click="addTeam">Agregar</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
