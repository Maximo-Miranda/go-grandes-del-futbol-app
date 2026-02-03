<script setup lang="ts">
import { Head, usePage } from "@inertiajs/vue3";

const page = usePage<{ team: any }>();
const team = page.props.team;
const players = team.players || [];
</script>

<template>
  <Head><title>{{ team.name }}</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <div class="d-flex align-center ga-4">
        <v-avatar size="56" :color="team.color || 'primary'">
          <v-icon size="28" color="white">mdi-shield-half-full</v-icon>
        </v-avatar>
        <h1 class="text-h4 font-weight-bold">{{ team.name }}</h1>
      </div>
      <v-btn color="primary" :href="`/teams/${team.id}/edit`" prepend-icon="mdi-pencil">Editar</v-btn>
    </div>

    <v-row>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title>Información</v-card-title>
          <v-card-text>
            <div v-if="team.contact_phone"><strong>Teléfono:</strong> {{ team.contact_phone }}</div>
            <div><strong>Jugadores:</strong> {{ players.length }}</div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title>Plantel</v-card-title>
          <v-list v-if="players.length > 0" density="compact">
            <v-list-item v-for="p in players" :key="p.id" :href="`/players/${p.id}`">
              <v-list-item-title>{{ p.name }}</v-list-item-title>
              <v-list-item-subtitle>{{ p.position || "Sin posición" }}</v-list-item-subtitle>
            </v-list-item>
          </v-list>
          <v-card-text v-else class="text-medium-emphasis">Sin jugadores asignados</v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>
