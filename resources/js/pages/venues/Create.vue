<script setup lang="ts">
import { Head, useForm, Link } from "@inertiajs/vue3";

const form = useForm({
  name: "",
  address: "",
});

const submit = () => form.post("/venues");
</script>

<template>
  <Head><title>Nueva Sede</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Nueva Sede</h1>
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
            <v-btn type="submit" color="primary" :loading="form.processing">Crear Sede</v-btn>
            <Link href="/venues" class="text-decoration-none">
              <v-btn variant="outlined">Cancelar</v-btn>
            </Link>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>
