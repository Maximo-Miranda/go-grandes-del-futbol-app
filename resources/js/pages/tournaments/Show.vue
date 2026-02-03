<script setup lang="ts">
import { ref } from "vue";
import { Head, usePage, router } from "@inertiajs/vue3";

const page = usePage<{ tournament: any; teams: any[]; matches: any[]; standings: any[]; availableTeams: any[] }>();
const tournament = page.props.tournament;
const teams = page.props.teams || [];
const matches = page.props.matches || [];
const standings = page.props.standings || [];
const availableTeams = ref(page.props.availableTeams || []);

const showAddTeamDialog = ref(false);
const selectedTeamId = ref<number | null>(null);
const loadingTeams = ref(false);

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
    router.post(`/tournaments/${tournament.id}/teams`, { team_id: selectedTeamId.value }, {
      onSuccess: () => {
        showAddTeamDialog.value = false;
        selectedTeamId.value = null;
      }
    });
  }
};

const removeTeam = (teamId: number) => {
  if (confirm("¿Quitar este equipo del torneo?")) {
    router.delete(`/tournaments/${tournament.id}/teams/${teamId}`);
  }
};
</script>

<template>
  <Head><title>{{ tournament.name }}</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">{{ tournament.name }}</h1>
      <v-btn color="primary" :href="`/tournaments/${tournament.id}/edit`" prepend-icon="mdi-pencil">Editar</v-btn>
    </div>

    <v-row>
      <v-col cols="12" md="8">
        <v-card class="mb-4">
          <v-card-title>Información</v-card-title>
          <v-card-text>
            <p v-if="tournament.description">{{ tournament.description }}</p>
            <v-chip class="mr-2 mb-2">{{ tournament.format }}</v-chip>
            <v-chip class="mr-2 mb-2">{{ tournament.game_type }}</v-chip>
            <v-chip class="mr-2 mb-2" :color="tournament.status === 'active' ? 'success' : 'grey'">{{ tournament.status }}</v-chip>
            <v-chip v-if="tournament.venue" class="mr-2 mb-2" color="info" prepend-icon="mdi-map-marker">{{ tournament.venue.name }}</v-chip>
          </v-card-text>
        </v-card>

        <v-card v-if="matches.length > 0" class="mb-4">
          <v-card-title>Partidos</v-card-title>
          <v-list>
            <v-list-item v-for="m in matches" :key="m.id" :href="`/matches/${m.id}`">
              <v-list-item-title>{{ m.home_team?.name || "TBD" }} vs {{ m.away_team?.name || "TBD" }}</v-list-item-title>
              <v-list-item-subtitle>{{ m.match_date }} — {{ m.status }}</v-list-item-subtitle>
            </v-list-item>
          </v-list>
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
              <v-list-item-title>{{ t.team?.name || t.name }}</v-list-item-title>
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
            <thead><tr><th>#</th><th>Equipo</th><th>Pts</th><th>PJ</th></tr></thead>
            <tbody>
              <tr v-for="(s, i) in standings" :key="s.id">
                <td>{{ i + 1 }}</td>
                <td>{{ s.team?.name }}</td>
                <td class="font-weight-bold">{{ s.points }}</td>
                <td>{{ s.played }}</td>
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
            No hay equipos disponibles. Primero crea equipos.
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
