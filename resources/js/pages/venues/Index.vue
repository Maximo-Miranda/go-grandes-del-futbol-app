<script setup lang="ts">
import { Head, usePage, router } from "@inertiajs/vue3";

const page = usePage<{ venues: any[] }>();
const venues = page.props.venues || [];

const deleteVenue = (id: number) => {
  if (confirm("¿Eliminar esta sede?")) {
    router.delete(`/venues/${id}`);
  }
};
</script>

<template>
  <Head><title>Sedes</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">Sedes</h1>
      <v-btn color="primary" href="/venues/create" prepend-icon="mdi-plus">Nueva Sede</v-btn>
    </div>

    <v-card>
      <v-table v-if="venues.length > 0">
        <thead>
          <tr>
            <th>Nombre</th>
            <th>Dirección</th>
            <th>Acciones</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="v in venues" :key="v.id">
            <td class="font-weight-medium">{{ v.name }}</td>
            <td>{{ v.address || "—" }}</td>
            <td>
              <v-btn icon="mdi-pencil" size="small" variant="text" :href="`/venues/${v.id}/edit`" />
              <v-btn icon="mdi-delete" size="small" variant="text" color="error" @click="deleteVenue(v.id)" />
            </td>
          </tr>
        </tbody>
      </v-table>
      <v-card-text v-else class="text-center text-medium-emphasis">
        No hay sedes registradas
      </v-card-text>
    </v-card>
  </div>
</template>
