<script setup lang="ts">
import { Head, useForm, Link } from "@inertiajs/vue3";

const form = useForm({
  name: "",
  nickname: "",
  document_id: "",
  phone: "",
  position: "",
});

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
          <v-text-field
            v-model="form.name"
            label="Nombre Completo *"
            :error-messages="form.errors.name"
            class="mb-4"
            required
          />
          <v-text-field
            v-model="form.nickname"
            label="Apodo"
            :error-messages="form.errors.nickname"
            class="mb-4"
          />
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.document_id"
                label="Documento de Identidad"
                :error-messages="form.errors.document_id"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.phone"
                label="Teléfono"
                :error-messages="form.errors.phone"
              />
            </v-col>
          </v-row>
          <v-select
            v-model="form.position"
            :items="positions"
            label="Posición"
            :error-messages="form.errors.position"
            clearable
            class="mb-4"
          />
          <div class="d-flex ga-4">
            <v-btn type="submit" color="primary" :loading="form.processing">Crear Jugador</v-btn>
            <Link href="/players" class="text-decoration-none">
              <v-btn variant="outlined">Cancelar</v-btn>
            </Link>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>
