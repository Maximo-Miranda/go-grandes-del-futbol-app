<script setup lang="ts">
import { Head, usePage } from "@inertiajs/vue3";

interface Props extends Record<string, unknown> {
  stats: { tournaments: number; teams: number; matches: number; players: number };
  upcomingMatches: any[];
}

const page = usePage<Props>();
const stats = page.props.stats;
const upcomingMatches = page.props.upcomingMatches || [];

const cards = [
  { title: "Torneos", value: stats.tournaments, icon: "mdi-trophy", color: "primary" },
  { title: "Equipos", value: stats.teams, icon: "mdi-shield-half-full", color: "secondary" },
  { title: "Partidos", value: stats.matches, icon: "mdi-soccer", color: "accent" },
  { title: "Jugadores", value: stats.players, icon: "mdi-account-group", color: "info" },
];
</script>

<template>
  <Head><title>Dashboard</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Dashboard</h1>
    <v-row>
      <v-col v-for="card in cards" :key="card.title" cols="12" sm="6" md="3">
        <v-card :color="card.color" variant="tonal">
          <v-card-text class="d-flex align-center pa-6">
            <v-icon :icon="card.icon" size="48" class="mr-4" />
            <div>
              <div class="text-h4 font-weight-bold">{{ card.value }}</div>
              <div class="text-body-1">{{ card.title }}</div>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <v-card class="mt-6" v-if="upcomingMatches.length > 0">
      <v-card-title>Próximos Partidos</v-card-title>
      <v-list>
        <v-list-item v-for="match in upcomingMatches" :key="match.id" :href="`/matches/${match.id}`">
          <v-list-item-title>
            {{ match.home_team?.name || "TBD" }} vs {{ match.away_team?.name || "TBD" }}
          </v-list-item-title>
          <v-list-item-subtitle>
            {{ match.tournament?.name }} — {{ match.match_date }}
          </v-list-item-subtitle>
        </v-list-item>
      </v-list>
    </v-card>
  </div>
</template>
