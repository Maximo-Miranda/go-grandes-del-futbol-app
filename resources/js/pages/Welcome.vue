<script setup lang="ts">
import { Head } from "@inertiajs/vue3";

defineProps<{
  version?: string;
  topScorers?: Array<{ name: string; goals: number; photo?: string }>;
  recentResults?: Array<{ home: string; away: string; homeScore: number; awayScore: number; date: string }>;
  leagueStanding?: { teamA: string; teamB: string; winsA: number; winsB: number; draws: number };
  totalGoals?: number;
  totalMatches?: number;
  totalPlayers?: number;
}>();
</script>

<template>
  <Head><title>Grandes del Fútbol ⚽</title></Head>
  <v-app>
    <v-main>
      <!-- Hero -->
      <div class="d-flex flex-column align-center justify-center text-center px-4" style="min-height: 100vh; background: linear-gradient(160deg, #0D3B0D 0%, #1B5E20 30%, #2E7D32 60%, #1B5E20 100%);">
        
        <div class="mb-4" style="font-size: 80px; line-height: 1;">⚽</div>
        
        <h1 class="text-white" style="font-size: clamp(2rem, 6vw, 3.5rem); font-weight: 800; letter-spacing: -1px;">
          Grandes del Fútbol
        </h1>
        
        <p class="text-h6 font-weight-light mt-2 mb-2" style="color: #A5D6A7; max-width: 500px;">
          Donde cada gol cuenta y cada jugada se convierte en leyenda
        </p>
        
        <p class="text-body-2 mb-8" style="color: #81C784;">
          La liga de los parceros 🤝
        </p>

        <!-- Stats rápidas -->
        <v-row v-if="totalGoals || totalMatches || totalPlayers" class="justify-center ga-6 mb-8" no-gutters>
          <v-col cols="auto" class="text-center">
            <div class="text-h3 font-weight-bold text-amber">{{ totalGoals || 0 }}</div>
            <div class="text-caption" style="color: #A5D6A7;">⚽ Goles</div>
          </v-col>
          <v-col cols="auto" class="text-center">
            <div class="text-h3 font-weight-bold text-amber">{{ totalMatches || 0 }}</div>
            <div class="text-caption" style="color: #A5D6A7;">🏟️ Partidos</div>
          </v-col>
          <v-col cols="auto" class="text-center">
            <div class="text-h3 font-weight-bold text-amber">{{ totalPlayers || 0 }}</div>
            <div class="text-caption" style="color: #A5D6A7;">👥 Parceros</div>
          </v-col>
        </v-row>

        <!-- Liga actual -->
        <v-card v-if="leagueStanding" class="mb-8 pa-4" color="rgba(255,255,255,0.1)" variant="flat" rounded="xl" style="backdrop-filter: blur(10px); max-width: 400px; width: 100%;">
          <div class="text-caption text-amber font-weight-bold mb-2">🏆 LA LIGA</div>
          <div class="d-flex align-center justify-space-between">
            <div class="text-center">
              <div class="text-h4 font-weight-bold text-white">{{ leagueStanding.winsA }}</div>
              <div class="text-body-2 text-white">{{ leagueStanding.teamA }}</div>
            </div>
            <div class="text-center px-4">
              <div class="text-body-2" style="color: #81C784;">{{ leagueStanding.draws }} empates</div>
              <div class="text-h6 text-amber">VS</div>
            </div>
            <div class="text-center">
              <div class="text-h4 font-weight-bold text-white">{{ leagueStanding.winsB }}</div>
              <div class="text-body-2 text-white">{{ leagueStanding.teamB }}</div>
            </div>
          </div>
        </v-card>

        <!-- Botones -->
        <div class="d-flex ga-4 justify-center flex-wrap mb-12">
          <v-btn size="x-large" color="amber" class="text-black font-weight-bold" href="/auth/login" prepend-icon="mdi-login" rounded="pill" elevation="8">
            Entrar a la Cancha
          </v-btn>
          <v-btn size="x-large" variant="outlined" color="amber" href="/auth/register" prepend-icon="mdi-account-plus" rounded="pill">
            Unirme a la Liga
          </v-btn>
        </div>

        <!-- Top goleadores -->
        <v-card v-if="topScorers && topScorers.length" class="mb-8 pa-4" color="rgba(255,255,255,0.08)" variant="flat" rounded="xl" style="backdrop-filter: blur(10px); max-width: 500px; width: 100%;">
          <div class="text-caption text-amber font-weight-bold mb-3">⚽ TOP GOLEADORES</div>
          <div v-for="(scorer, i) in topScorers.slice(0, 5)" :key="i" class="d-flex align-center mb-2">
            <span class="text-h6 font-weight-bold text-amber mr-3" style="min-width: 24px;">{{ i + 1 }}</span>
            <v-avatar size="32" class="mr-2" color="green-darken-3">
              <v-img v-if="scorer.photo" :src="scorer.photo" />
              <v-icon v-else size="20">mdi-account</v-icon>
            </v-avatar>
            <span class="text-body-1 text-white flex-grow-1">{{ scorer.name }}</span>
            <v-chip size="small" color="amber" variant="flat" class="font-weight-bold">{{ scorer.goals }} ⚽</v-chip>
          </div>
        </v-card>

        <!-- Features -->
        <v-row class="justify-center ga-4 mt-4 mb-8" style="max-width: 700px;">
          <v-col v-for="feat in [
            { icon: 'mdi-card-account-details', title: 'Tu Carta FIFA', desc: 'Stats personales como en el videojuego' },
            { icon: 'mdi-trophy', title: 'La Liga', desc: '¿Quién manda en la cancha?' },
            { icon: 'mdi-chart-line', title: 'Estadísticas', desc: 'Goles, asistencias, tarjetas y más' },
            { icon: 'mdi-fire', title: 'MVP', desc: 'El crack de cada partido' },
          ]" :key="feat.title" cols="6" sm="3">
            <div class="text-center">
              <v-icon size="36" color="amber" class="mb-1">{{ feat.icon }}</v-icon>
              <div class="text-body-2 font-weight-bold text-white">{{ feat.title }}</div>
              <div class="text-caption" style="color: #81C784;">{{ feat.desc }}</div>
            </div>
          </v-col>
        </v-row>

        <!-- Últimos resultados -->
        <div v-if="recentResults && recentResults.length" class="mb-8" style="max-width: 500px; width: 100%;">
          <div class="text-caption text-amber font-weight-bold mb-3 text-center">📋 ÚLTIMOS PARTIDOS</div>
          <v-card v-for="(r, i) in recentResults.slice(0, 3)" :key="i" class="mb-2 pa-3" color="rgba(255,255,255,0.06)" variant="flat" rounded="lg">
            <div class="d-flex align-center justify-space-between text-white">
              <span class="text-body-2 flex-grow-1 text-right">{{ r.home }}</span>
              <v-chip class="mx-3 font-weight-bold" color="amber" variant="flat" size="small">
                {{ r.homeScore }} - {{ r.awayScore }}
              </v-chip>
              <span class="text-body-2 flex-grow-1">{{ r.away }}</span>
            </div>
          </v-card>
        </div>

        <!-- Footer -->
        <div class="text-caption mt-4 pb-8" style="color: #66BB6A;">
          Hecho con ❤️ y ganas de meter gol · Grandes del Fútbol © {{ new Date().getFullYear() }}
        </div>
      </div>
    </v-main>
  </v-app>
</template>
