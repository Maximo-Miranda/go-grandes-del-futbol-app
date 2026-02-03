<script setup lang="ts">
import { Head, useForm } from "@inertiajs/vue3";

const form = useForm({ name: "", nickname: "", document_id: "", phone: "", position: "" });

const positions = [
  { title: "Portero", value: "goalkeeper" },
  { title: "Defensa", value: "defender" },
  { title: "Mediocampista", value: "midfielder" },
  { title: "Delantero", value: "forward" },
];

const submit = () => form.post("/players");
</script>

<template>
  <Head><title>Crear Jugador</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Crear Jugador</h1>
    <v-card>
      <v-card-text class="pa-6">
        <v-form @submit.prevent="submit">
          <v-text-field v-model="form.name" label="Nombre Completo" :error-messages="form.errors.name" class="mb-4" />
          <v-text-field v-model="form.nickname" label="Apodo" class="mb-4" />
          <v-text-field v-model="form.document_id" label="Documento de Identidad" class="mb-4" />
          <v-text-field v-model="form.phone" label="Teléfono" class="mb-4" />
          <v-select v-model="form.position" :items="positions" label="Posición" clearable class="mb-4" />
          <div class="d-flex ga-4">
            <v-btn type="submit" color="primary" :loading="form.processing">Crear Jugador</v-btn>
            <v-btn variant="outlined" href="/players">Cancelar</v-btn>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>
