<script setup lang="ts">
import { Head, usePage, router, Link } from "@inertiajs/vue3";

const page = usePage<{ venues: any[] }>();
const venues = page.props.venues || [];

const deleteVenue = (id: number) => {
  if (confirm("¿Eliminar esta sede?")) {
    router.delete(`/venues/${id}`, {
      onSuccess: () => {
        router.reload({ only: ["venues"] });
      },
    });
  }
};
</script>

<template>
  <Head><title>Sedes</title></Head>
  <div>
    <div class="d-flex justify-space-between align-center mb-6">
      <h1 class="text-h4 font-weight-bold">Sedes</h1>
      <Link href="/venues/create" class="text-decoration-none">
        <v-btn color="primary" prepend-icon="mdi-plus">Nueva Sede</v-btn>
      </Link>
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
              <Link :href="`/venues/${v.id}/edit`" class="text-decoration-none">
                <v-btn icon="mdi-pencil" size="small" variant="text" />
              </Link>
              <v-btn icon="mdi-delete" size="small" variant="text" color="error" @click="deleteVenue(v.id)" />
            </td>
          </tr>
        </tbody>
      </v-table>
      <v-card-text v-else class="text-center pa-8">
        <v-icon size="64" color="grey-lighten-1" class="mb-4">mdi-map-marker</v-icon>
        <p class="text-medium-emphasis">No hay sedes registradas</p>
        <Link href="/venues/create" class="text-decoration-none">
          <v-btn color="primary" variant="tonal" class="mt-4">Crear Sede</v-btn>
        </Link>
      </v-card-text>
    </v-card>
  </div>
</template>
