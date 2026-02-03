<script setup lang="ts">
import { Head } from "@inertiajs/vue3";
import { computed } from "vue";

interface PlayerStats {
  matchesPlayed: number;
  goals: number;
  assists: number;
  yellowCards: number;
  redCards: number;
  mvpAwards: number;
  cleanSheets: number;
}

interface PlayerRatings {
  pace: number;
  shooting: number;
  passing: number;
  dribbling: number;
  defense: number;
  physical: number;
}

const props = defineProps<{
  player: {
    id: number;
    firstName: string;
    lastName: string;
    nickname?: string;
    position?: string;
    photoUrl?: string;
    jerseyNumber?: number;
    dominantFoot?: string;
    heightCm?: number;
    weightKg?: number;
  };
  stats?: PlayerStats;
  ratings?: PlayerRatings;
  recentMatches?: Array<{
    date: string;
    opponent: string;
    result: string;
    goals: number;
    assists: number;
    cards: string;
    isMvp: boolean;
  }>;
}>();

const displayName = computed(() => props.player.nickname || props.player.firstName);
const overall = computed(() => {
  if (!props.ratings) return 50;
  const r = props.ratings;
  const weights: Record<string, number[]> = {
    portero: [0.05, 0.05, 0.1, 0.05, 0.5, 0.25],
    defensa: [0.1, 0.05, 0.15, 0.05, 0.4, 0.25],
    mediocampista: [0.15, 0.15, 0.3, 0.2, 0.1, 0.1],
    delantero: [0.2, 0.35, 0.15, 0.2, 0.02, 0.08],
  };
  const w = weights[props.player.position || "mediocampista"] || weights.mediocampista;
  return Math.round(r.pace * w[0] + r.shooting * w[1] + r.passing * w[2] + r.dribbling * w[3] + r.defense * w[4] + r.physical * w[5]);
});

const positionLabel = computed(() => {
  const map: Record<string, string> = { portero: "POR", defensa: "DEF", mediocampista: "MED", delantero: "DEL" };
  return map[props.player.position || ""] || "JUG";
});

const positionColor = computed(() => {
  const map: Record<string, string> = { portero: "#FFC107", defensa: "#2196F3", mediocampista: "#4CAF50", delantero: "#F44336" };
  return map[props.player.position || ""] || "#9E9E9E";
});

const ratingColor = (val: number) => {
  if (val >= 80) return "#4CAF50";
  if (val >= 60) return "#8BC34A";
  if (val >= 40) return "#FFC107";
  return "#F44336";
};

const statItems = computed(() => [
  { label: "PAC", value: props.ratings?.pace || 50, fullLabel: "Velocidad" },
  { label: "SHO", value: props.ratings?.shooting || 50, fullLabel: "Disparo" },
  { label: "PAS", value: props.ratings?.passing || 50, fullLabel: "Pase" },
  { label: "DRI", value: props.ratings?.dribbling || 50, fullLabel: "Regate" },
  { label: "DEF", value: props.ratings?.defense || 50, fullLabel: "Defensa" },
  { label: "PHY", value: props.ratings?.physical || 50, fullLabel: "Físico" },
]);
</script>

<template>
  <Head><title>{{ displayName }} - Carta FIFA</title></Head>

  <v-container fluid class="pa-4">
    <v-row justify="center">
      <!-- FIFA Card -->
      <v-col cols="12" sm="6" md="4" lg="3">
        <v-card class="fifa-card mx-auto" rounded="xl" elevation="12" style="max-width: 320px; background: linear-gradient(145deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%); overflow: hidden;">
          <!-- Card Header -->
          <div class="d-flex align-center justify-space-between px-4 pt-4">
            <div class="text-center">
              <div class="text-h2 font-weight-black text-amber" style="line-height: 1;">{{ overall }}</div>
              <v-chip :color="positionColor" size="small" variant="flat" class="font-weight-bold mt-1">
                {{ positionLabel }}
              </v-chip>
            </div>
            <v-avatar size="120" :color="positionColor" class="elevation-8">
              <v-img v-if="player.photoUrl" :src="player.photoUrl" cover />
              <v-icon v-else size="60" color="white">mdi-account</v-icon>
            </v-avatar>
          </div>

          <!-- Player Name -->
          <div class="text-center mt-3 px-4">
            <div class="text-h5 font-weight-black text-white text-uppercase" style="letter-spacing: 2px;">
              {{ displayName }}
            </div>
            <div v-if="player.nickname && player.nickname !== player.firstName" class="text-caption text-grey">
              {{ player.firstName }} {{ player.lastName }}
            </div>
          </div>

          <v-divider class="mx-4 mt-3" color="amber" opacity="0.3" />

          <!-- Stats Grid -->
          <div class="px-4 py-3">
            <v-row no-gutters class="text-center">
              <v-col v-for="stat in statItems" :key="stat.label" cols="4" class="py-1">
                <div class="text-h5 font-weight-black" :style="{ color: ratingColor(stat.value) }">
                  {{ stat.value }}
                </div>
                <div class="text-caption font-weight-bold text-grey-lighten-1">{{ stat.label }}</div>
              </v-col>
            </v-row>
          </div>

          <v-divider class="mx-4" color="amber" opacity="0.3" />

          <!-- Season Stats -->
          <div class="px-4 py-3">
            <div class="text-caption text-amber font-weight-bold mb-2">📊 TEMPORADA</div>
            <v-row no-gutters class="text-center">
              <v-col v-for="s in [
                { icon: '⚽', value: stats?.goals || 0, label: 'Goles' },
                { icon: '🅰️', value: stats?.assists || 0, label: 'Asist.' },
                { icon: '🏟️', value: stats?.matchesPlayed || 0, label: 'Partidos' },
                { icon: '🟨', value: stats?.yellowCards || 0, label: 'Amarillas' },
                { icon: '🟥', value: stats?.redCards || 0, label: 'Rojas' },
                { icon: '⭐', value: stats?.mvpAwards || 0, label: 'MVP' },
              ]" :key="s.label" cols="4" class="py-1">
                <div class="text-body-1 font-weight-bold text-white">{{ s.icon }} {{ s.value }}</div>
                <div class="text-caption text-grey">{{ s.label }}</div>
              </v-col>
            </v-row>
          </div>

          <!-- Player Info -->
          <div class="px-4 pb-3">
            <div class="d-flex ga-2 flex-wrap justify-center">
              <v-chip v-if="player.jerseyNumber" size="small" color="amber" variant="outlined">
                #{{ player.jerseyNumber }}
              </v-chip>
              <v-chip v-if="player.dominantFoot" size="small" color="blue-grey" variant="outlined">
                {{ player.dominantFoot === 'derecho' ? '🦶 Derecho' : player.dominantFoot === 'izquierdo' ? '🦶 Izquierdo' : '🦶 Ambidiestro' }}
              </v-chip>
              <v-chip v-if="player.heightCm" size="small" color="blue-grey" variant="outlined">
                📏 {{ player.heightCm }}cm
              </v-chip>
            </div>
          </div>
        </v-card>
      </v-col>

      <!-- Match History -->
      <v-col cols="12" sm="6" md="5" lg="4">
        <v-card rounded="xl" color="grey-darken-4" class="pa-4">
          <div class="text-h6 text-amber font-weight-bold mb-3">📋 Últimos Partidos</div>
          
          <div v-if="recentMatches && recentMatches.length">
            <v-card v-for="(m, i) in recentMatches" :key="i" class="mb-2 pa-3" :color="m.result === 'W' ? 'green-darken-4' : m.result === 'L' ? 'red-darken-4' : 'grey-darken-3'" variant="flat" rounded="lg">
              <div class="d-flex align-center justify-space-between">
                <div>
                  <div class="text-body-2 text-white font-weight-medium">vs {{ m.opponent }}</div>
                  <div class="text-caption text-grey-lighten-1">{{ m.date }}</div>
                </div>
                <div class="text-right">
                  <div class="d-flex ga-1 align-center justify-end">
                    <span v-if="m.goals > 0" class="text-caption">⚽×{{ m.goals }}</span>
                    <span v-if="m.assists > 0" class="text-caption">🅰️×{{ m.assists }}</span>
                    <span v-if="m.cards === 'yellow'" class="text-caption">🟨</span>
                    <span v-if="m.cards === 'red'" class="text-caption">🟥</span>
                    <span v-if="m.isMvp" class="text-caption">⭐</span>
                  </div>
                  <v-chip :color="m.result === 'W' ? 'green' : m.result === 'L' ? 'red' : 'grey'" size="x-small" variant="flat" class="font-weight-bold mt-1">
                    {{ m.result === 'W' ? 'VICTORIA' : m.result === 'L' ? 'DERROTA' : 'EMPATE' }}
                  </v-chip>
                </div>
              </div>
            </v-card>
          </div>
          
          <div v-else class="text-center text-grey py-8">
            <v-icon size="48" class="mb-2">mdi-soccer-field</v-icon>
            <div>Sin partidos aún</div>
            <div class="text-caption">¡A la cancha parcero! 🔥</div>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.fifa-card {
  transition: transform 0.3s ease;
}
.fifa-card:hover {
  transform: scale(1.02);
}
</style>
