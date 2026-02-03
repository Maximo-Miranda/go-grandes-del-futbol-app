<script setup lang="ts">
import { Head, router, Link } from "@inertiajs/vue3";

const props = defineProps<{ standings: any[]; tournaments: any[]; tournamentId: string }>();

const goToTournament = (v: number | null) => {
  if (v) {
    router.get("/standings", { tournament_id: v }, { preserveState: true });
  } else {
    router.get("/standings", {}, { preserveState: true });
  }
};
</script>

<template>
  <Head><title>Tabla de Posiciones</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-2">Tabla de Posiciones</h1>
    <p class="text-medium-emphasis mb-6">Clasificación por torneo</p>

    <v-card class="mb-4">
      <v-card-text>
        <v-select
          label="Filtrar por torneo"
          :items="props.tournaments"
          item-title="name"
          item-value="id"
          :model-value="props.tournamentId ? Number(props.tournamentId) : null"
          clearable
          @update:model-value="goToTournament"
        />
      </v-card-text>
    </v-card>

    <v-card>
      <v-table v-if="props.standings.length > 0">
        <thead>
          <tr>
            <th>#</th>
            <th>Equipo</th>
            <th>Grupo</th>
            <th>PJ</th>
            <th>G</th>
            <th>E</th>
            <th>P</th>
            <th>GF</th>
            <th>GC</th>
            <th>DG</th>
            <th class="font-weight-bold">Pts</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(s, i) in props.standings" :key="s.id">
            <td>{{ i + 1 }}</td>
            <td class="font-weight-medium">
              <Link :href="`/teams/${s.team?.id}`" class="text-primary">
                {{ s.team?.name }}
              </Link>
            </td>
            <td>{{ s.group?.name || "—" }}</td>
            <td>{{ s.played }}</td>
            <td>{{ s.won }}</td>
            <td>{{ s.drawn }}</td>
            <td>{{ s.lost }}</td>
            <td>{{ s.goals_for }}</td>
            <td>{{ s.goals_against }}</td>
            <td>{{ s.goal_difference > 0 ? "+" : "" }}{{ s.goal_difference }}</td>
            <td class="font-weight-bold text-primary">{{ s.points }}</td>
          </tr>
        </tbody>
      </v-table>
      <v-card-text v-else class="text-center pa-8">
        <v-icon size="64" color="grey-lighten-1" class="mb-4">mdi-podium</v-icon>
        <p class="text-medium-emphasis">Selecciona un torneo para ver las posiciones.</p>
      </v-card-text>
    </v-card>
  </div>
</template>
