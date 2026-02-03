<script setup lang="ts">
import { Head, usePage } from "@inertiajs/vue3";

const page = usePage<{ tournament: any; teams: any[]; matches: any[]; standings: any[] }>();
const tournament = page.props.tournament;
const teams = page.props.teams || [];
const matches = page.props.matches || [];
const standings = page.props.standings || [];
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
          <v-card-title>Equipos ({{ teams.length }})</v-card-title>
          <v-list density="compact">
            <v-list-item v-for="team in teams" :key="team.id" :href="`/teams/${team.id}`">
              <v-list-item-title>{{ team.name }}</v-list-item-title>
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
  </div>
</template>
