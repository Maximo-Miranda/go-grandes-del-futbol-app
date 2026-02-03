<script setup lang="ts">
import { Head, useForm, usePage, Link } from "@inertiajs/vue3";

const page = usePage<{ venues: any[]; errors?: Record<string, string> }>();
const venues = page.props.venues || [];

const form = useForm({
  name: "",
  description: "",
  format: "round_robin",
  game_type: "5v5",
  start_date: "",
  end_date: "",
  venue_id: null as number | null,
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

const submit = () => form.post("/tournaments");
</script>

<template>
  <Head><title>Crear Torneo</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Crear Torneo</h1>
    <v-card>
      <v-card-text class="pa-6">
        <v-form @submit.prevent="submit">
          <v-text-field
            v-model="form.name"
            label="Nombre del Torneo *"
            :error-messages="form.errors.name"
            class="mb-4"
            required
          />
          <v-textarea
            v-model="form.description"
            label="Descripción"
            rows="3"
            class="mb-4"
          />
          <v-row>
            <v-col cols="12" md="6">
              <v-select
                v-model="form.format"
                :items="formats"
                label="Formato *"
                :error-messages="form.errors.format"
                class="mb-4"
                required
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-select
                v-model="form.game_type"
                :items="gameTypes"
                label="Tipo de Juego *"
                :error-messages="form.errors.game_type"
                class="mb-4"
                required
              />
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.start_date"
                label="Fecha Inicio"
                type="date"
                :error-messages="form.errors.start_date"
                class="mb-4"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.end_date"
                label="Fecha Fin"
                type="date"
                :error-messages="form.errors.end_date"
                class="mb-4"
              />
            </v-col>
          </v-row>
          <v-select
            v-model="form.venue_id"
            :items="venues"
            item-title="name"
            item-value="id"
            label="Sede"
            clearable
            :error-messages="form.errors.venue_id"
            class="mb-4"
          />
          <div class="d-flex ga-4">
            <v-btn type="submit" color="primary" :loading="form.processing">Crear Torneo</v-btn>
            <Link href="/tournaments" class="text-decoration-none">
              <v-btn variant="outlined">Cancelar</v-btn>
            </Link>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>
