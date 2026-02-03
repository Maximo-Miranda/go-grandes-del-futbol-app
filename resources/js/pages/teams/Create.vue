<script setup lang="ts">
import { Head, useForm } from "@inertiajs/vue3";

const form = useForm({ name: "", color: "#4CAF50", contact_phone: "" });
const submit = () => form.post("/teams");
</script>

<template>
  <Head><title>Crear Equipo</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Crear Equipo</h1>
    <v-card>
      <v-card-text class="pa-6">
        <v-form @submit.prevent="submit">
          <v-text-field v-model="form.name" label="Nombre del Equipo" :error-messages="form.errors.name" class="mb-4" />
          <v-text-field v-model="form.color" label="Color del Equipo" type="color" class="mb-4" />
          <v-text-field v-model="form.contact_phone" label="Teléfono de Contacto" class="mb-4" />
          <div class="d-flex ga-4">
            <v-btn type="submit" color="primary" :loading="form.processing">Crear Equipo</v-btn>
            <v-btn variant="outlined" href="/teams">Cancelar</v-btn>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>
