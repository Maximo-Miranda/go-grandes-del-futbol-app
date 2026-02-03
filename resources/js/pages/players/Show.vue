<script setup lang="ts">
import { Head, usePage } from "@inertiajs/vue3";

const page = usePage<{ player: any; stats: { goals: number; yellow_cards: number; red_cards: number } }>();
const player = page.props.player;
const stats = page.props.stats;
</script>

<template>
  <Head><title>{{ player.name }}</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">{{ player.name }}</h1>
      <v-btn color="primary" :href="`/players/${player.id}/edit`" prepend-icon="mdi-pencil">Editar</v-btn>
    </div>
    <v-row>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title>Información</v-card-title>
          <v-card-text>
            <div v-if="player.nickname"><strong>Apodo:</strong> {{ player.nickname }}</div>
            <div v-if="player.position"><strong>Posición:</strong> {{ player.position }}</div>
            <div v-if="player.document_id"><strong>Documento:</strong> {{ player.document_id }}</div>
            <div v-if="player.phone"><strong>Teléfono:</strong> {{ player.phone }}</div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title>Estadísticas</v-card-title>
          <v-card-text>
            <v-row>
              <v-col cols="4" class="text-center">
                <div class="text-h4 font-weight-bold text-primary">{{ stats?.goals || 0 }}</div>
                <div class="text-caption">Goles</div>
              </v-col>
              <v-col cols="4" class="text-center">
                <div class="text-h4 font-weight-bold text-warning">{{ stats?.yellow_cards || 0 }}</div>
                <div class="text-caption">Amarillas</div>
              </v-col>
              <v-col cols="4" class="text-center">
                <div class="text-h4 font-weight-bold text-error">{{ stats?.red_cards || 0 }}</div>
                <div class="text-caption">Rojas</div>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>
