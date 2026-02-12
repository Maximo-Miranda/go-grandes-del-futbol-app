<script setup lang="ts">
import { ref } from "vue";
import { Head, usePage, router, Link } from "@inertiajs/vue3";
import ConfirmDialog from "@/components/ConfirmDialog.vue";

const page = usePage<{ venues: any[] }>();
const venues = page.props.venues || [];

const showDeleteDialog = ref(false);
const venueToDelete = ref<string | null>(null);
const deleting = ref(false);

const confirmDelete = (uuid: string) => {
    venueToDelete.value = uuid;
    showDeleteDialog.value = true;
};

const deleteVenue = () => {
    if (!venueToDelete.value) return;
    deleting.value = true;
    router.delete(`/venues/${venueToDelete.value}`, {
        onSuccess: () => {
            router.visit(window.location.href, { preserveScroll: true });
        },
        onFinish: () => {
            deleting.value = false;
            showDeleteDialog.value = false;
            venueToDelete.value = null;
        },
    });
};

const cancelDelete = () => {
    showDeleteDialog.value = false;
    venueToDelete.value = null;
};
</script>

<template>
    <Head><title>Sedes</title></Head>
    <div>
        <div class="d-flex justify-space-between align-center mb-6">
            <h1 class="text-h4 font-weight-bold">Sedes</h1>
            <Link href="/venues/create" class="text-decoration-none">
                <v-btn color="primary" prepend-icon="mdi-plus"
                    >Nueva Sede</v-btn
                >
            </Link>
        </div>

        <v-card>
            <v-table v-if="venues.length > 0">
                <thead>
                    <tr>
                        <th style="width: 60px"></th>
                        <th>Nombre</th>
                        <th>Dirección</th>
                        <th>Acciones</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="v in venues" :key="v.uuid">
                        <td>
                            <v-avatar
                                size="40"
                                color="grey-lighten-3"
                                rounded="lg"
                            >
                                <v-img
                                    v-if="v.photo_url"
                                    :src="v.photo_url"
                                    cover
                                />
                                <v-icon v-else size="24" color="grey-lighten-1"
                                    >mdi-stadium</v-icon
                                >
                            </v-avatar>
                        </td>
                        <td class="font-weight-medium">{{ v.name }}</td>
                        <td>{{ v.address || "—" }}</td>
                        <td>
                            <Link
                                :href="`/venues/${v.uuid}/edit`"
                                class="text-decoration-none"
                            >
                                <v-btn
                                    icon="mdi-pencil"
                                    size="small"
                                    variant="text"
                                />
                            </Link>
                            <v-btn
                                icon="mdi-delete"
                                size="small"
                                variant="text"
                                color="error"
                                @click="confirmDelete(v.uuid)"
                            />
                        </td>
                    </tr>
                </tbody>
            </v-table>
            <v-card-text v-else class="text-center pa-8">
                <v-icon size="64" color="grey-lighten-1" class="mb-4"
                    >mdi-map-marker</v-icon
                >
                <p class="text-medium-emphasis">No hay sedes registradas</p>
                <Link href="/venues/create" class="text-decoration-none">
                    <v-btn color="primary" variant="tonal" class="mt-4"
                        >Crear Sede</v-btn
                    >
                </Link>
            </v-card-text>
        </v-card>

        <ConfirmDialog
            v-model="showDeleteDialog"
            title="¿Eliminar sede?"
            message="Esta acción no se puede deshacer."
            confirm-text="Eliminar"
            :loading="deleting"
            @confirm="deleteVenue"
            @cancel="cancelDelete"
        />
    </div>
</template>
