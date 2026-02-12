<script setup lang="ts">
import { ref } from "vue";
import { Head, usePage, router, Link } from "@inertiajs/vue3";
import ConfirmDialog from "@/components/ConfirmDialog.vue";

const page = usePage<{ tournaments: any[] }>();
const tournaments = page.props.tournaments || [];

const statusColors: Record<string, string> = {
    draft: "grey",
    active: "success",
    completed: "primary",
    cancelled: "error",
};
const statusLabels: Record<string, string> = {
    draft: "Borrador",
    active: "Activo",
    completed: "Finalizado",
    cancelled: "Cancelado",
};
const formatLabels: Record<string, string> = {
    round_robin: "Todos contra todos",
    knockout: "Eliminación directa",
    group_knockout: "Grupos + Eliminación",
};
const gameTypeLabels: Record<string, string> = {
    "5v5": "5 vs 5",
    "7v7": "7 vs 7",
    "11v11": "11 vs 11",
};

const formatDate = (dateStr: string | null) => {
    if (!dateStr) return "—";
    return new Date(dateStr).toLocaleDateString("es-CO", {
        day: "numeric",
        month: "short",
        year: "numeric",
    });
};

const showDeleteDialog = ref(false);
const tournamentToDelete = ref<number | null>(null);
const deleting = ref(false);

const confirmDelete = (id: number) => {
    tournamentToDelete.value = id;
    showDeleteDialog.value = true;
};

const deleteTournament = () => {
    if (!tournamentToDelete.value) return;
    deleting.value = true;
    router.delete(`/tournaments/${tournamentToDelete.value}`, {
        onSuccess: () => {
            router.visit(window.location.href, { preserveScroll: true });
        },
        onFinish: () => {
            deleting.value = false;
            showDeleteDialog.value = false;
            tournamentToDelete.value = null;
        },
    });
};

const cancelDelete = () => {
    showDeleteDialog.value = false;
    tournamentToDelete.value = null;
};
</script>

<template>
    <Head><title>Torneos</title></Head>
    <div>
        <div class="d-flex justify-space-between align-center mb-6">
            <h1 class="text-h4 font-weight-bold">Torneos</h1>
            <Link href="/tournaments/create" class="text-decoration-none">
                <v-btn color="primary" prepend-icon="mdi-plus"
                    >Nuevo Torneo</v-btn
                >
            </Link>
        </div>
        <v-card>
            <v-table v-if="tournaments.length > 0">
                <thead>
                    <tr>
                        <th>Nombre</th>
                        <th>Formato</th>
                        <th>Tipo</th>
                        <th>Fechas</th>
                        <th>Estado</th>
                        <th>Sede</th>
                        <th>Acciones</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="t in tournaments" :key="t.id">
                        <td>
                            <Link
                                :href="`/tournaments/${t.id}`"
                                class="text-primary font-weight-medium"
                            >
                                {{ t.name }}
                            </Link>
                        </td>
                        <td>{{ formatLabels[t.format] || t.format }}</td>
                        <td>
                            {{ gameTypeLabels[t.game_type] || t.game_type }}
                        </td>
                        <td class="text-medium-emphasis">
                            {{ formatDate(t.start_date) }}
                            <span v-if="t.end_date">
                                — {{ formatDate(t.end_date) }}</span
                            >
                        </td>
                        <td>
                            <v-chip
                                :color="statusColors[t.status]"
                                size="small"
                            >
                                {{ statusLabels[t.status] || t.status }}
                            </v-chip>
                        </td>
                        <td>{{ t.venue?.name || "—" }}</td>
                        <td>
                            <Link
                                :href="`/tournaments/${t.id}/edit`"
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
                                @click="confirmDelete(t.id)"
                            />
                        </td>
                    </tr>
                </tbody>
            </v-table>
            <v-card-text v-else class="text-center text-medium-emphasis pa-8">
                <v-icon size="64" color="grey-lighten-1" class="mb-4"
                    >mdi-trophy</v-icon
                >
                <p>No hay torneos creados aún. ¡Crea el primero!</p>
                <Link href="/tournaments/create" class="text-decoration-none">
                    <v-btn color="primary" variant="tonal" class="mt-4"
                        >Crear Torneo</v-btn
                    >
                </Link>
            </v-card-text>
        </v-card>

        <ConfirmDialog
            v-model="showDeleteDialog"
            title="¿Eliminar torneo?"
            message="Esta acción no se puede deshacer. Se eliminará el torneo y toda su información asociada."
            confirm-text="Eliminar"
            :loading="deleting"
            @confirm="deleteTournament"
            @cancel="cancelDelete"
        />
    </div>
</template>
