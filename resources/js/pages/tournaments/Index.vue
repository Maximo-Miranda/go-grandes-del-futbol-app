<script setup lang="ts">
import { Head, usePage, router } from "@inertiajs/vue3";

const page = usePage<{ tournaments: any[] }>();
const tournaments = page.props.tournaments || [];

const statusColors: Record<string, string> = {
  draft: "grey", active: "success", completed: "primary", cancelled: "error",
};
const statusLabels: Record<string, string> = {
  draft: "Borrador", active: "Activo", completed: "Finalizado", cancelled: "Cancelado",
};

const deleteTournament = (id: number) => {
  if (confirm("¿Estás seguro de eliminar este torneo?")) {
    router.delete(`/tournaments/${id}`);
  }
};
</script>

<template>
  <Head><title>Torneos</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">Torneos</h1>
      <v-btn color="primary" href="/tournaments/create" prepend-icon="mdi-plus">Nuevo Torneo</v-btn>
    </div>
    <v-card>
      <v-table v-if="tournaments.length > 0">
        <thead>
          <tr>
            <th>Nombre</th><th>Formato</th><th>Tipo</th><th>Estado</th><th>Sede</th><th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="t in tournaments" :key="t.id">
            <td><a :href="`/tournaments/${t.id}`">{{ t.name }}</a></td>
            <td>{{ t.format }}</td>
            <td>{{ t.game_type }}</td>
            <td><v-chip :color="statusColors[t.status]" size="small">{{ statusLabels[t.status] || t.status }}</v-chip></td>
            <td>{{ t.venue?.name || "—" }}</td>
            <td>
              <v-btn icon="mdi-pencil" size="small" variant="text" :href="`/tournaments/${t.id}/edit`" />
              <v-btn icon="mdi-delete" size="small" variant="text" color="error" @click="deleteTournament(t.id)" />
            </td>
          </tr>
        </tbody>
      </v-table>
      <v-card-text v-else class="text-center text-medium-emphasis pa-8">
        No hay torneos creados aún. ¡Crea el primero!
      </v-card-text>
    </v-card>
  </div>
</template>
