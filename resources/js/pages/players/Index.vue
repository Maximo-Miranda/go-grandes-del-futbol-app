<script setup lang="ts">
import { Head, usePage, router } from "@inertiajs/vue3";

const page = usePage<{ players: any[] }>();
const players = page.props.players || [];

const deletePlayer = (id: number) => {
  if (confirm("¿Estás seguro de eliminar este jugador?")) router.delete(`/players/${id}`);
};
</script>

<template>
  <Head><title>Jugadores</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">Jugadores</h1>
      <v-btn color="primary" href="/players/create" prepend-icon="mdi-plus">Nuevo Jugador</v-btn>
    </div>
    <v-card>
      <v-table v-if="players.length > 0">
        <thead><tr><th>Nombre</th><th>Apodo</th><th>Posición</th><th>Teléfono</th><th>Acciones</th></tr></thead>
        <tbody>
          <tr v-for="p in players" :key="p.id">
            <td><a :href="`/players/${p.id}`">{{ p.name }}</a></td>
            <td>{{ p.nickname || "—" }}</td>
            <td>{{ p.position || "—" }}</td>
            <td>{{ p.phone || "—" }}</td>
            <td>
              <v-btn icon="mdi-pencil" size="small" variant="text" :href="`/players/${p.id}/edit`" />
              <v-btn icon="mdi-delete" size="small" variant="text" color="error" @click="deletePlayer(p.id)" />
            </td>
          </tr>
        </tbody>
      </v-table>
      <v-card-text v-else class="text-center text-medium-emphasis pa-8">
        No hay jugadores registrados aún. ¡Registra el primero!
      </v-card-text>
    </v-card>
  </div>
</template>
