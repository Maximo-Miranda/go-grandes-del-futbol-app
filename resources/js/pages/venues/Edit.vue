<script setup lang="ts">
import { Head, useForm, usePage, Link } from "@inertiajs/vue3";

const page = usePage<{ venue: any; errors?: Record<string, string> }>();
const venue = page.props.venue;

const form = useForm({
  name: venue.name || "",
  address: venue.address || "",
});

const submit = () => {
  form.put(`/venues/${venue.id}`, {
    preserveScroll: true,
  });
};
</script>

<template>
  <Head><title>Editar Sede</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Editar Sede</h1>
    <v-card>
      <v-card-text class="pa-6">
        <v-form @submit.prevent="submit">
          <v-text-field
            v-model="form.name"
            label="Nombre de la Sede *"
            :error-messages="form.errors.name"
            class="mb-4"
            required
          />
          <v-text-field
            v-model="form.address"
            label="Dirección"
            :error-messages="form.errors.address"
            class="mb-4"
          />
          <div class="d-flex ga-4">
            <v-btn type="submit" color="primary" :loading="form.processing">Guardar</v-btn>
            <Link href="/venues" class="text-decoration-none">
              <v-btn variant="outlined">Cancelar</v-btn>
            </Link>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>
