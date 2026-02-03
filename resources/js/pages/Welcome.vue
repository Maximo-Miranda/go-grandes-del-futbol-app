<script setup lang="ts">
import { Head } from "@inertiajs/vue3";
</script>

<template>
  <Head><title>Grandes del Fútbol</title></Head>
  <v-app>
    <v-main>
      <div class="welcome-hero d-flex flex-column align-center justify-center" style="min-height: 100vh; background: linear-gradient(135deg, #1B5E20 0%, #4CAF50 50%, #81C784 100%);">
        <div class="text-center text-white">
          <div class="text-h1 mb-4">⚽</div>
          <h1 class="text-h2 font-weight-bold mb-2">Grandes del Fútbol</h1>
          <p class="text-h6 font-weight-light mb-8 text-white-darken-1">
            Gestiona torneos, equipos y jugadores en un solo lugar
          </p>
          <div class="d-flex ga-4 justify-center flex-wrap">
            <v-btn size="x-large" color="white" class="text-primary font-weight-bold" href="/login" prepend-icon="mdi-login">
              Iniciar Sesión
            </v-btn>
            <v-btn size="x-large" variant="outlined" color="white" href="/register" prepend-icon="mdi-account-plus">
              Registrarse
            </v-btn>
          </div>
        </div>
        <div class="mt-16 text-center text-white-darken-2">
          <v-row class="justify-center ga-8">
            <v-col cols="auto" class="text-center">
              <v-icon size="48" class="mb-2">mdi-trophy</v-icon>
              <div class="text-body-1">Torneos</div>
            </v-col>
            <v-col cols="auto" class="text-center">
              <v-icon size="48" class="mb-2">mdi-shield-half-full</v-icon>
              <div class="text-body-1">Equipos</div>
            </v-col>
            <v-col cols="auto" class="text-center">
              <v-icon size="48" class="mb-2">mdi-account-group</v-icon>
              <div class="text-body-1">Jugadores</div>
            </v-col>
            <v-col cols="auto" class="text-center">
              <v-icon size="48" class="mb-2">mdi-format-list-numbered</v-icon>
              <div class="text-body-1">Clasificación</div>
            </v-col>
          </v-row>
        </div>
      </div>
    </v-main>
  </v-app>
</template>
