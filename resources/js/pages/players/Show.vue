<script setup lang="ts">
import { Head, Link } from "@inertiajs/vue3";
import { computed } from "vue";

interface PlayerStats {
  matchesPlayed: number;
  goals: number;
  assists: number;
  yellow_cards: number;
  red_cards: number;
}

interface RecentMatch {
  date: string;
  opponent: string;
  result: string;
  goals: number;
  assists: number;
  cards: string;
  isMvp: boolean;
}

const props = defineProps<{
  player: {
    id: number;
    name: string;
    nickname?: string;
    position?: string;
    photo?: string;
    document_id?: string;
    phone?: string;
  };
  stats?: PlayerStats;
  recentMatches?: RecentMatch[];
}>();

const displayName = computed(() => props.player.nickname || props.player.name);

const photoUrl = computed(() => {
  if (props.player.photo) {
    return `/players/photo/${props.player.photo.split('/').pop()}`;
  }
  return null;
});

const positionLabel = computed(() => {
  const map: Record<string, string> = {
    goalkeeper: "POR",
    defender: "DEF",
    midfielder: "MED",
    forward: "DEL",
  };
  return map[props.player.position || ""] || "JUG";
});

const positionFullLabel = computed(() => {
  const map: Record<string, string> = {
    goalkeeper: "Portero",
    defender: "Defensa",
    midfielder: "Mediocampista",
    forward: "Delantero",
  };
  return map[props.player.position || ""] || "Jugador";
});

const positionColor = computed(() => {
  const map: Record<string, string> = {
    goalkeeper: "#FFC107",
    defender: "#2196F3",
    midfielder: "#4CAF50",
    forward: "#F44336",
  };
  return map[props.player.position || ""] || "#9E9E9E";
});

// Calculate overall rating based on stats
const overall = computed(() => {
  const s = props.stats;
  if (!s || s.matchesPlayed === 0) return 50;
  
  // Simple formula based on stats
  let rating = 50;
  rating += Math.min(s.goals * 3, 20);
  rating += Math.min(s.assists * 2, 15);
  rating += Math.min(s.matchesPlayed * 1, 10);
  rating -= s.yellow_cards * 2;
  rating -= s.red_cards * 5;
  
  return Math.max(30, Math.min(99, Math.round(rating)));
});

// Rating color function (reserved for future use)
const _ratingColor = (val: number) => {
  if (val >= 80) return "#4CAF50";
  if (val >= 60) return "#8BC34A";
  if (val >= 40) return "#FFC107";
  return "#F44336";
};
void _ratingColor; // Suppress unused warning
</script>

<template>
  <Head><title>{{ displayName }} - Jugador</title></Head>

  <v-container fluid class="pa-4">
    <Link href="/players" class="text-decoration-none">
      <v-btn variant="text" prepend-icon="mdi-arrow-left" class="mb-4">Volver a Jugadores</v-btn>
    </Link>

    <v-row justify="center">
      <!-- FIFA-style Card -->
      <v-col cols="12" sm="6" md="4" lg="3">
        <v-card
          class="fifa-card mx-auto"
          rounded="xl"
          elevation="12"
          style="max-width: 320px; background: linear-gradient(145deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%); overflow: hidden;"
        >
          <!-- Card Header -->
          <div class="d-flex align-center justify-space-between px-4 pt-4">
            <div class="text-center">
              <div class="text-h2 font-weight-black text-amber" style="line-height: 1;">{{ overall }}</div>
              <v-chip :color="positionColor" size="small" variant="flat" class="font-weight-bold mt-1">
                {{ positionLabel }}
              </v-chip>
            </div>
            <v-avatar size="120" :color="positionColor" class="elevation-8">
              <v-img v-if="photoUrl" :src="photoUrl" cover />
              <v-icon v-else size="60" color="white">mdi-account</v-icon>
            </v-avatar>
          </div>

          <!-- Player Name -->
          <div class="text-center mt-3 px-4">
            <div class="text-h5 font-weight-black text-white text-uppercase" style="letter-spacing: 2px;">
              {{ displayName }}
            </div>
            <div v-if="player.nickname && player.nickname !== player.name" class="text-caption text-grey">
              {{ player.name }}
            </div>
            <div class="text-caption text-grey-lighten-1 mt-1">{{ positionFullLabel }}</div>
          </div>

          <v-divider class="mx-4 mt-3" color="amber" opacity="0.3" />

          <!-- Season Stats -->
          <div class="px-4 py-3">
            <div class="text-caption text-amber font-weight-bold mb-2">📊 ESTADÍSTICAS</div>
            <v-row no-gutters class="text-center">
              <v-col
                v-for="s in [
                  { icon: '⚽', value: stats?.goals || 0, label: 'Goles' },
                  { icon: '🅰️', value: stats?.assists || 0, label: 'Asist.' },
                  { icon: '🏟️', value: stats?.matchesPlayed || 0, label: 'Partidos' },
                  { icon: '🟨', value: stats?.yellow_cards || 0, label: 'Amarillas' },
                  { icon: '🟥', value: stats?.red_cards || 0, label: 'Rojas' },
                ]"
                :key="s.label"
                cols="4"
                class="py-1"
              >
                <div class="text-body-1 font-weight-bold text-white">{{ s.icon }} {{ s.value }}</div>
                <div class="text-caption text-grey">{{ s.label }}</div>
              </v-col>
            </v-row>
          </div>

          <!-- Player Info -->
          <div class="px-4 pb-3">
            <div class="d-flex ga-2 flex-wrap justify-center">
              <v-chip v-if="player.document_id" size="small" color="blue-grey" variant="outlined">
                📋 {{ player.document_id }}
              </v-chip>
              <v-chip v-if="player.phone" size="small" color="blue-grey" variant="outlined">
                📱 {{ player.phone }}
              </v-chip>
            </div>
          </div>

          <!-- Actions -->
          <div class="px-4 pb-4 text-center">
            <Link :href="`/players/${player.id}/edit`" class="text-decoration-none">
              <v-btn color="amber" variant="tonal" prepend-icon="mdi-pencil" size="small">
                Editar
              </v-btn>
            </Link>
          </div>
        </v-card>
      </v-col>

      <!-- Match History -->
      <v-col cols="12" sm="6" md="5" lg="4">
        <v-card rounded="xl" color="grey-darken-4" class="pa-4">
          <div class="text-h6 text-amber font-weight-bold mb-3">📋 Últimos Partidos</div>

          <div v-if="recentMatches && recentMatches.length > 0">
            <v-card
              v-for="(m, i) in recentMatches"
              :key="i"
              class="mb-2 pa-3"
              :color="m.result === 'W' ? 'green-darken-4' : m.result === 'L' ? 'red-darken-4' : 'grey-darken-3'"
              variant="flat"
              rounded="lg"
            >
              <div class="d-flex align-center justify-space-between">
                <div>
                  <div class="text-body-2 text-white font-weight-medium">vs {{ m.opponent || 'TBD' }}</div>
                  <div class="text-caption text-grey-lighten-1">{{ m.date || 'Fecha TBD' }}</div>
                </div>
                <div class="text-right">
                  <div class="d-flex ga-1 align-center justify-end">
                    <span v-if="m.goals > 0" class="text-caption">⚽×{{ m.goals }}</span>
                    <span v-if="m.assists > 0" class="text-caption">🅰️×{{ m.assists }}</span>
                    <span v-if="m.cards === 'yellow'" class="text-caption">🟨</span>
                    <span v-if="m.cards === 'red'" class="text-caption">🟥</span>
                    <span v-if="m.isMvp" class="text-caption">⭐</span>
                  </div>
                  <v-chip
                    v-if="m.result"
                    :color="m.result === 'W' ? 'green' : m.result === 'L' ? 'red' : 'grey'"
                    size="x-small"
                    variant="flat"
                    class="font-weight-bold mt-1"
                  >
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
