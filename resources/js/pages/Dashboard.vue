<script setup lang="ts">
import { Head } from "@inertiajs/vue3";

defineProps<{
  league?: {
    teamA: string;
    teamB: string;
    winsA: number;
    winsB: number;
    draws: number;
    totalMatches: number;
  };
  topScorers?: Array<{ name: string; goals: number; photo?: string; team: string }>;
  topAssisters?: Array<{ name: string; assists: number; photo?: string }>;
  mvpOfWeek?: { name: string; photo?: string; goals: number; assists: number; rating: number };
  recentResults?: Array<{ home: string; away: string; homeScore: number; awayScore: number; date: string; mvp?: string }>;
  funFacts?: {
    biggestWin?: string;
    longestStreak?: string;
    mostCards?: string;
  };
}>();
</script>

<template>
  <Head><title>Dashboard - Grandes del Fútbol</title></Head>

  <v-container fluid class="pa-4 pa-sm-6">
    <div class="text-h4 font-weight-black mb-1">🏟️ La Cancha</div>
    <div class="text-body-2 text-grey mb-6">¿Quién manda hoy?</div>

    <v-row>
      <!-- League Standing -->
      <v-col cols="12" md="6">
        <v-card rounded="xl" class="pa-6" color="green-darken-4" elevation="8">
          <div class="text-caption text-amber font-weight-bold mb-4">🏆 LA LIGA</div>
          
          <div v-if="league" class="d-flex align-center justify-space-around">
            <div class="text-center">
              <div class="text-h2 font-weight-black text-white">{{ league.winsA }}</div>
              <div class="text-h6 text-white font-weight-bold">{{ league.teamA }}</div>
              <div class="text-caption text-grey-lighten-1">victorias</div>
            </div>
            
            <div class="text-center">
              <div class="text-h4 text-amber font-weight-bold">VS</div>
              <div class="text-caption text-grey-lighten-1 mt-1">{{ league.draws }} empates</div>
              <div class="text-caption text-grey-lighten-1">{{ league.totalMatches }} partidos</div>
            </div>
            
            <div class="text-center">
              <div class="text-h2 font-weight-black text-white">{{ league.winsB }}</div>
              <div class="text-h6 text-white font-weight-bold">{{ league.teamB }}</div>
              <div class="text-caption text-grey-lighten-1">victorias</div>
            </div>
          </div>

          <div v-else class="text-center text-grey-lighten-1 py-4">
            <v-icon size="48">mdi-trophy-outline</v-icon>
            <div class="mt-2">La liga aún no ha comenzado</div>
          </div>
        </v-card>
      </v-col>

      <!-- MVP de la Jornada -->
      <v-col cols="12" md="6">
        <v-card rounded="xl" class="pa-6" style="background: linear-gradient(145deg, #1a1a2e, #16213e);" elevation="8">
          <div class="text-caption text-amber font-weight-bold mb-4">⭐ JUGADOR DE LA JORNADA</div>
          
          <div v-if="mvpOfWeek" class="d-flex align-center ga-4">
            <v-avatar size="80" color="amber">
              <v-img v-if="mvpOfWeek.photo" :src="mvpOfWeek.photo" cover />
              <v-icon v-else size="40">mdi-account</v-icon>
            </v-avatar>
            <div>
              <div class="text-h5 font-weight-bold text-white">{{ mvpOfWeek.name }}</div>
              <div class="d-flex ga-3 mt-1">
                <span class="text-body-2 text-grey-lighten-1">⚽ {{ mvpOfWeek.goals }} goles</span>
                <span class="text-body-2 text-grey-lighten-1">🅰️ {{ mvpOfWeek.assists }} asist.</span>
              </div>
              <v-chip color="amber" size="small" variant="flat" class="font-weight-bold mt-2">
                Rating: {{ mvpOfWeek.rating }}
              </v-chip>
            </div>
          </div>

          <div v-else class="text-center text-grey-lighten-1 py-4">
            <v-icon size="48">mdi-star-outline</v-icon>
            <div class="mt-2">Aún no hay MVP</div>
          </div>
        </v-card>
      </v-col>

      <!-- Top Goleadores -->
      <v-col cols="12" md="6">
        <v-card rounded="xl" class="pa-5" elevation="4">
          <div class="text-caption text-amber font-weight-bold mb-3">⚽ GOLEADORES</div>
          <div v-if="topScorers && topScorers.length">
            <div v-for="(s, i) in topScorers.slice(0, 7)" :key="i" class="d-flex align-center py-2" :class="{ 'border-b': i < topScorers.length - 1 }">
              <span class="text-h6 font-weight-bold mr-3" :class="i === 0 ? 'text-amber' : i === 1 ? 'text-grey-lighten-1' : i === 2 ? 'text-brown' : 'text-grey'" style="min-width: 28px;">
                {{ i === 0 ? '🥇' : i === 1 ? '🥈' : i === 2 ? '🥉' : (i + 1) }}
              </span>
              <v-avatar size="32" class="mr-2" color="green-darken-3">
                <v-img v-if="s.photo" :src="s.photo" />
                <v-icon v-else size="18">mdi-account</v-icon>
              </v-avatar>
              <div class="flex-grow-1">
                <div class="text-body-2 font-weight-medium">{{ s.name }}</div>
                <div class="text-caption text-grey">{{ s.team }}</div>
              </div>
              <v-chip color="amber" variant="flat" size="small" class="font-weight-bold">{{ s.goals }}</v-chip>
            </div>
          </div>
          <div v-else class="text-center text-grey py-4">Sin goles aún 😢</div>
        </v-card>
      </v-col>

      <!-- Top Asistidores -->
      <v-col cols="12" md="6">
        <v-card rounded="xl" class="pa-5" elevation="4">
          <div class="text-caption text-amber font-weight-bold mb-3">🅰️ ASISTIDORES</div>
          <div v-if="topAssisters && topAssisters.length">
            <div v-for="(a, i) in topAssisters.slice(0, 7)" :key="i" class="d-flex align-center py-2">
              <span class="text-h6 font-weight-bold mr-3" :class="i === 0 ? 'text-amber' : 'text-grey'" style="min-width: 28px;">
                {{ i === 0 ? '🥇' : i === 1 ? '🥈' : i === 2 ? '🥉' : (i + 1) }}
              </span>
              <v-avatar size="32" class="mr-2" color="green-darken-3">
                <v-img v-if="a.photo" :src="a.photo" />
                <v-icon v-else size="18">mdi-account</v-icon>
              </v-avatar>
              <span class="text-body-2 font-weight-medium flex-grow-1">{{ a.name }}</span>
              <v-chip color="blue" variant="flat" size="small" class="font-weight-bold">{{ a.assists }}</v-chip>
            </div>
          </div>
          <div v-else class="text-center text-grey py-4">Sin asistencias aún</div>
        </v-card>
      </v-col>

      <!-- Últimos Resultados -->
      <v-col cols="12" md="8">
        <v-card rounded="xl" class="pa-5" elevation="4">
          <div class="text-caption text-amber font-weight-bold mb-3">📋 ÚLTIMOS PARTIDOS</div>
          <div v-if="recentResults && recentResults.length">
            <v-card v-for="(r, i) in recentResults.slice(0, 5)" :key="i" class="mb-2 pa-3" color="grey-lighten-4" variant="flat" rounded="lg">
              <div class="d-flex align-center justify-space-between">
                <span class="text-body-2 font-weight-medium flex-grow-1 text-right">{{ r.home }}</span>
                <v-chip class="mx-3 font-weight-bold" :color="r.homeScore > r.awayScore ? 'green' : r.homeScore < r.awayScore ? 'red' : 'grey'" variant="flat" size="small">
                  {{ r.homeScore }} - {{ r.awayScore }}
                </v-chip>
                <span class="text-body-2 font-weight-medium flex-grow-1">{{ r.away }}</span>
                <div class="ml-3 text-right" style="min-width: 80px;">
                  <div class="text-caption text-grey">{{ r.date }}</div>
                  <div v-if="r.mvp" class="text-caption text-amber">⭐ {{ r.mvp }}</div>
                </div>
              </div>
            </v-card>
          </div>
          <div v-else class="text-center text-grey py-4">
            <v-icon size="48" class="mb-2">mdi-soccer-field</v-icon>
            <div>No hay partidos registrados</div>
          </div>
        </v-card>
      </v-col>

      <!-- Fun Facts -->
      <v-col cols="12" md="4">
        <v-card rounded="xl" class="pa-5" elevation="4">
          <div class="text-caption text-amber font-weight-bold mb-3">🎯 DATOS RANDOM</div>
          <div v-if="funFacts" class="d-flex flex-column ga-3">
            <div v-if="funFacts.biggestWin">
              <div class="text-caption text-grey">Goleada más grande</div>
              <div class="text-body-2 font-weight-medium">{{ funFacts.biggestWin }}</div>
            </div>
            <div v-if="funFacts.longestStreak">
              <div class="text-caption text-grey">Racha más larga</div>
              <div class="text-body-2 font-weight-medium">{{ funFacts.longestStreak }}</div>
            </div>
            <div v-if="funFacts.mostCards">
              <div class="text-caption text-grey">Más tarjetas 😂</div>
              <div class="text-body-2 font-weight-medium">{{ funFacts.mostCards }}</div>
            </div>
          </div>
          <div v-else class="text-center text-grey py-4">
            <div>🔜 Próximamente</div>
            <div class="text-caption">A jugar para llenar esto</div>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
