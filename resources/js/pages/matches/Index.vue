<script setup lang="ts">
import { Head, usePage } from "@inertiajs/vue3";

const page = usePage<{ matches: any[]; tournaments: any[] }>();
const matches = page.props.matches || [];
</script>

<template>
  <Head><title>Partidos</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Partidos</h1>
    <v-card>
      <v-table v-if="matches.length > 0">
        <thead><tr><th>Local</th><th>Resultado</th><th>Visitante</th><th>Torneo</th><th>Estado</th><th></th></tr></thead>
        <tbody>
          <tr v-for="m in matches" :key="m.id">
            <td>{{ m.home_team?.name || "TBD" }}</td>
            <td class="text-center font-weight-bold">
              <span v-if="m.status === 'completed'">{{ m.home_score }} - {{ m.away_score }}</span>
              <span v-else>vs</span>
            </td>
            <td>{{ m.away_team?.name || "TBD" }}</td>
            <td>{{ m.tournament?.name }}</td>
            <td><v-chip size="small" :color="m.status === 'completed' ? 'success' : 'grey'">{{ m.status }}</v-chip></td>
            <td><v-btn icon="mdi-eye" size="small" variant="text" :href="`/matches/${m.id}`" /></td>
          </tr>
        </tbody>
      </v-table>
      <v-card-text v-else class="text-center text-medium-emphasis pa-8">No hay partidos programados</v-card-text>
    </v-card>
  </div>
</template>
