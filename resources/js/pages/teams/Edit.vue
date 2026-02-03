<script setup lang="ts">
import { Head, useForm, Link } from "@inertiajs/vue3";

const props = defineProps<{ team: any; errors?: Record<string, string> }>();

const form = useForm({
  name: props.team.name,
  color: props.team.color || "#4CAF50",
  contact_phone: props.team.contact_phone || "",
});

const submit = () => form.put(`/teams/${props.team.id}`);
</script>

<template>
  <Head><title>Editar Equipo</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Editar: {{ team.name }}</h1>
    <v-card>
      <v-card-text class="pa-6">
        <v-form @submit.prevent="submit">
          <v-text-field
            v-model="form.name"
            label="Nombre del Equipo *"
            :error-messages="form.errors.name"
            class="mb-4"
            required
          />
          <v-text-field
            v-model="form.color"
            label="Color del Equipo"
            type="color"
            :error-messages="form.errors.color"
            class="mb-4"
          />
          <v-text-field
            v-model="form.contact_phone"
            label="Teléfono de Contacto"
            :error-messages="form.errors.contact_phone"
            class="mb-4"
          />
          <div class="d-flex ga-4">
            <v-btn type="submit" color="primary" :loading="form.processing">Guardar</v-btn>
            <Link :href="`/teams/${team.id}`" class="text-decoration-none">
              <v-btn variant="outlined">Cancelar</v-btn>
            </Link>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>
