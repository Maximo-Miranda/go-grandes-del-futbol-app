<script setup lang="ts">
import { Head, usePage, Link, router } from "@inertiajs/vue3";
import { ref } from "vue";

const page = usePage<{ matches: any[]; tournaments: any[]; tournamentId?: string }>();
const matches = page.props.matches || [];
const tournaments = page.props.tournaments || [];
const selectedTournament = ref(page.props.tournamentId || "");

const filterByTournament = () => {
  if (selectedTournament.value) {
    router.get("/matches", { tournament_id: selectedTournament.value }, { preserveState: true });
  } else {
    router.get("/matches", {}, { preserveState: true });
  }
};

const formatDate = (dateStr: string | null) => {
  if (!dateStr) return "Sin fecha";
  const date = new Date(dateStr);
  return date.toLocaleDateString("es-CO", { weekday: "short", day: "numeric", month: "short", hour: "2-digit", minute: "2-digit" });
};

const statusColors: Record<string, string> = {
  scheduled: "info",
  live: "warning",
  completed: "success",
  cancelled: "error",
};
</script>

<template>
  <Head><title>Partidos</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">Partidos</h1>
      <Link href="/matches/create" class="text-decoration-none">
        <v-btn color="primary" prepend-icon="mdi-plus">Nuevo Partido</v-btn>
      </Link>
    </div>

    <v-card class="mb-4" v-if="tournaments.length > 0">
      <v-card-text>
        <v-row>
          <v-col cols="12" md="4">
            <v-select
              v-model="selectedTournament"
              :items="tournaments"
              item-title="name"
              item-value="id"
              label="Filtrar por Torneo"
              clearable
              density="compact"
              @update:model-value="filterByTournament"
            />
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

    <v-card>
      <v-table v-if="matches.length > 0">
        <thead>
          <tr>
            <th>Fecha</th>
            <th>Local</th>
            <th class="text-center">Resultado</th>
            <th>Visitante</th>
            <th>Torneo</th>
            <th>Estado</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="m in matches" :key="m.id">
            <td class="text-medium-emphasis">{{ formatDate(m.match_date) }}</td>
            <td class="font-weight-medium">{{ m.home_team?.name || "TBD" }}</td>
            <td class="text-center font-weight-bold">
              <span v-if="m.status === 'completed'" class="text-h6">{{ m.home_score }} - {{ m.away_score }}</span>
              <span v-else class="text-medium-emphasis">vs</span>
            </td>
            <td class="font-weight-medium">{{ m.away_team?.name || "TBD" }}</td>
            <td>
              <Link :href="`/tournaments/${m.tournament?.id}`" class="text-primary">
                {{ m.tournament?.name }}
              </Link>
            </td>
            <td>
              <v-chip size="small" :color="statusColors[m.status] || 'grey'">
                {{ m.status }}
              </v-chip>
            </td>
            <td>
              <Link :href="`/matches/${m.id}`" class="text-decoration-none">
                <v-btn icon="mdi-eye" size="small" variant="text" />
              </Link>
              <Link :href="`/matches/${m.id}/edit`" class="text-decoration-none">
                <v-btn icon="mdi-pencil" size="small" variant="text" />
              </Link>
            </td>
          </tr>
        </tbody>
      </v-table>
      <v-card-text v-else class="text-center text-medium-emphasis pa-8">
        <v-icon size="64" color="grey-lighten-1" class="mb-4">mdi-soccer-field</v-icon>
        <p>No hay partidos programados</p>
        <Link href="/matches/create" class="text-decoration-none">
          <v-btn color="primary" variant="tonal" class="mt-4">Crear Primer Partido</v-btn>
        </Link>
      </v-card-text>
    </v-card>
  </div>
</template>
