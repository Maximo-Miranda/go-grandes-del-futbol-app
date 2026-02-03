<script setup lang="ts">
import { Head, usePage, router, Link } from "@inertiajs/vue3";

const page = usePage<{ players: any[] }>();
const players = page.props.players || [];

const deletePlayer = (id: number) => {
  if (confirm("¿Estás seguro de eliminar este jugador?")) {
    router.delete(`/players/${id}`);
  }
};
</script>

<template>
  <Head><title>Jugadores</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">Jugadores</h1>
      <Link href="/players/create" class="text-decoration-none">
        <v-btn color="primary" prepend-icon="mdi-plus">Nuevo Jugador</v-btn>
      </Link>
    </div>
    <v-card>
      <v-table v-if="players.length > 0">
        <thead>
          <tr>
            <th>Nombre</th>
            <th>Apodo</th>
            <th>Posición</th>
            <th>Teléfono</th>
            <th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="p in players" :key="p.id">
            <td>
              <Link :href="`/players/${p.id}`" class="text-primary font-weight-medium">
                {{ p.name }}
              </Link>
            </td>
            <td>{{ p.nickname || "—" }}</td>
            <td>{{ p.position || "—" }}</td>
            <td>{{ p.phone || "—" }}</td>
            <td>
              <Link :href="`/players/${p.id}/edit`" class="text-decoration-none">
                <v-btn icon="mdi-pencil" size="small" variant="text" />
              </Link>
              <v-btn icon="mdi-delete" size="small" variant="text" color="error" @click="deletePlayer(p.id)" />
            </td>
          </tr>
        </tbody>
      </v-table>
      <v-card-text v-else class="text-center pa-8">
        <v-icon size="64" color="grey-lighten-1" class="mb-4">mdi-account-group</v-icon>
        <p class="text-medium-emphasis">No hay jugadores registrados aún. ¡Registra el primero!</p>
        <Link href="/players/create" class="text-decoration-none">
          <v-btn color="primary" variant="tonal" class="mt-4">Registrar Jugador</v-btn>
        </Link>
      </v-card-text>
    </v-card>
  </div>
</template>
