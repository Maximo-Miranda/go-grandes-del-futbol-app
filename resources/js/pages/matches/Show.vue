<script setup lang="ts">
import { Head, usePage } from "@inertiajs/vue3";

const page = usePage<{ match: any; events: any[] }>();
const match = page.props.match;
const events = page.props.events || [];
</script>

<template>
  <Head><title>Partido</title></Head>
  <div>
    <v-btn variant="text" href="/matches" prepend-icon="mdi-arrow-left" class="mb-4">Volver</v-btn>
    <v-card class="mb-6">
      <v-card-text class="text-center pa-8">
        <div class="text-caption text-medium-emphasis mb-2">{{ match.tournament?.name }}</div>
        <v-row align="center" justify="center">
          <v-col cols="4" class="text-center">
            <div class="text-h6 font-weight-bold">{{ match.home_team?.name || "TBD" }}</div>
          </v-col>
          <v-col cols="4" class="text-center">
            <div v-if="match.status === 'completed'" class="text-h3 font-weight-bold">
              {{ match.home_score }} - {{ match.away_score }}
            </div>
            <div v-else class="text-h5 text-medium-emphasis">vs</div>
            <v-chip size="small" :color="match.status === 'completed' ? 'success' : 'warning'" class="mt-2">
              {{ match.status }}
            </v-chip>
          </v-col>
          <v-col cols="4" class="text-center">
            <div class="text-h6 font-weight-bold">{{ match.away_team?.name || "TBD" }}</div>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

    <v-card v-if="events.length > 0">
      <v-card-title>Eventos</v-card-title>
      <v-list>
        <v-list-item v-for="e in events" :key="e.id">
          <v-list-item-title>{{ e.minute }}' — {{ e.event_type }} — {{ e.player?.name }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-card>
  </div>
</template>
