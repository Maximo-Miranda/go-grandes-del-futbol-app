<script setup lang="ts">
import { Head, usePage, router } from "@inertiajs/vue3";

const page = usePage<{ teams: any[] }>();
const teams = page.props.teams || [];

const deleteTeam = (id: number) => {
  if (confirm("¿Estás seguro de eliminar este equipo?")) router.delete(`/teams/${id}`);
};
</script>

<template>
  <Head><title>Equipos</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">Equipos</h1>
      <v-btn color="primary" href="/teams/create" prepend-icon="mdi-plus">Nuevo Equipo</v-btn>
    </div>
    <v-row>
      <v-col v-for="team in teams" :key="team.id" cols="12" sm="6" md="4" lg="3">
        <v-card :href="`/teams/${team.id}`" hover>
          <v-card-text class="text-center pa-6">
            <v-avatar size="64" :color="team.color || 'primary'" class="mb-3">
              <v-icon size="32" color="white">mdi-shield-half-full</v-icon>
            </v-avatar>
            <div class="text-h6 font-weight-bold">{{ team.name }}</div>
            <div class="text-caption text-medium-emphasis">{{ team.players?.length || 0 }} jugadores</div>
          </v-card-text>
          <v-card-actions class="justify-center">
            <v-btn icon="mdi-pencil" size="small" variant="text" :href="`/teams/${team.id}/edit`" />
            <v-btn icon="mdi-delete" size="small" variant="text" color="error" @click.prevent="deleteTeam(team.id)" />
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-card v-if="teams.length === 0" class="text-center pa-8 text-medium-emphasis">
      No hay equipos registrados aún. ¡Crea el primero!
    </v-card>
  </div>
</template>
