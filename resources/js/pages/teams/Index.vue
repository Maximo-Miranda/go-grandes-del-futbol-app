<script setup lang="ts">
import { Head, usePage, router, Link } from "@inertiajs/vue3";

const page = usePage<{ teams: any[] }>();
const teams = page.props.teams || [];

const deleteTeam = (id: number, event: Event) => {
  event.preventDefault();
  event.stopPropagation();
  if (confirm("¿Estás seguro de eliminar este equipo?")) {
    router.delete(`/teams/${id}`, {
      onSuccess: () => {
        router.visit(window.location.href, { preserveScroll: true });
      },
    });
  }
};
</script>

<template>
  <Head><title>Equipos</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">Equipos</h1>
      <Link href="/teams/create" class="text-decoration-none">
        <v-btn color="primary" prepend-icon="mdi-plus">Nuevo Equipo</v-btn>
      </Link>
    </div>
    <v-row>
      <v-col v-for="team in teams" :key="team.id" cols="12" sm="6" md="4" lg="3">
        <Link :href="`/teams/${team.id}`" class="text-decoration-none">
          <v-card hover>
            <v-card-text class="text-center pa-6">
              <v-avatar size="64" :color="team.color || 'primary'" class="mb-3">
                <v-icon size="32" color="white">mdi-shield-half-full</v-icon>
              </v-avatar>
              <div class="text-h6 font-weight-bold">{{ team.name }}</div>
              <div class="text-caption text-medium-emphasis">{{ team.players?.length || 0 }} jugadores</div>
            </v-card-text>
            <v-card-actions class="justify-center">
              <Link :href="`/teams/${team.id}/edit`" class="text-decoration-none" @click.stop>
                <v-btn icon="mdi-pencil" size="small" variant="text" />
              </Link>
              <v-btn icon="mdi-delete" size="small" variant="text" color="error" @click="deleteTeam(team.id, $event)" />
            </v-card-actions>
          </v-card>
        </Link>
      </v-col>
    </v-row>
    <v-card v-if="teams.length === 0" class="text-center pa-8">
      <v-icon size="64" color="grey-lighten-1" class="mb-4">mdi-shield-outline</v-icon>
      <p class="text-medium-emphasis">No hay equipos registrados aún. ¡Crea el primero!</p>
      <Link href="/teams/create" class="text-decoration-none">
        <v-btn color="primary" variant="tonal" class="mt-4">Crear Equipo</v-btn>
      </Link>
    </v-card>
  </div>
</template>
