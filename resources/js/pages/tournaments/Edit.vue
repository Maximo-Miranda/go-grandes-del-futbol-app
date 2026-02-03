<script setup lang="ts">
import { Head, useForm, usePage } from "@inertiajs/vue3";

const page = usePage<{ tournament: any; venues: any[] }>();
const tournament = page.props.tournament;
const venues = page.props.venues || [];

const form = useForm({
  name: tournament.name, description: tournament.description || "",
  format: tournament.format, game_type: tournament.game_type,
  start_date: tournament.start_date?.split("T")[0] || "",
  end_date: tournament.end_date?.split("T")[0] || "",
  venue_id: tournament.venue_id,
});

const formats = [
  { title: "Todos contra todos", value: "round_robin" },
  { title: "Eliminación directa", value: "knockout" },
  { title: "Grupos + Eliminación", value: "group_knockout" },
];
const gameTypes = [
  { title: "5 vs 5", value: "5v5" },
  { title: "7 vs 7", value: "7v7" },
  { title: "11 vs 11", value: "11v11" },
];

const submit = () => form.put(`/tournaments/${tournament.id}`);
</script>

<template>
  <Head><title>Editar Torneo</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Editar Torneo</h1>
    <v-card>
      <v-card-text class="pa-6">
        <v-form @submit.prevent="submit">
          <v-text-field v-model="form.name" label="Nombre del Torneo" :error-messages="form.errors.name" class="mb-4" />
          <v-textarea v-model="form.description" label="Descripción" rows="3" class="mb-4" />
          <v-row>
            <v-col cols="12" md="6"><v-select v-model="form.format" :items="formats" label="Formato" class="mb-4" /></v-col>
            <v-col cols="12" md="6"><v-select v-model="form.game_type" :items="gameTypes" label="Tipo de Juego" class="mb-4" /></v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6"><v-text-field v-model="form.start_date" label="Fecha Inicio" type="date" class="mb-4" /></v-col>
            <v-col cols="12" md="6"><v-text-field v-model="form.end_date" label="Fecha Fin" type="date" class="mb-4" /></v-col>
          </v-row>
          <v-select v-model="form.venue_id" :items="venues" item-title="name" item-value="id" label="Sede" clearable class="mb-4" />
          <div class="d-flex ga-4">
            <v-btn type="submit" color="primary" :loading="form.processing">Guardar Cambios</v-btn>
            <v-btn variant="outlined" :href="`/tournaments/${tournament.id}`">Cancelar</v-btn>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>
