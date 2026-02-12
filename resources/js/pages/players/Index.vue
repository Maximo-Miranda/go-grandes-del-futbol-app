<script setup lang="ts">
import { Head, Link } from "@inertiajs/vue3";
import { computed } from "vue";

// Use defineProps for reactive SPA behavior
const props = defineProps<{ players: any[] }>();

// Computed to ensure reactivity
const players = computed(() => props.players || []);

const getPhotoUrl = (player: any): string | undefined => {
    return player.photo_url || undefined;
};

const getPositionLabel = (position: string) => {
    const labels: Record<string, string> = {
        goalkeeper: "Portero",
        defender: "Defensa",
        midfielder: "Mediocampista",
        forward: "Delantero",
    };
    return labels[position] || position || "—";
};

const getPositionColor = (position: string) => {
    const colors: Record<string, string> = {
        goalkeeper: "amber",
        defender: "blue",
        midfielder: "green",
        forward: "red",
    };
    return colors[position] || "grey";
};
</script>

<template>
    <Head><title>Jugadores</title></Head>
    <div>
        <div class="d-flex justify-space-between align-center mb-6">
            <h1 class="text-h4 font-weight-bold">Jugadores</h1>
        </div>
        <v-card>
            <v-table v-if="players.length > 0">
                <thead>
                    <tr>
                        <th style="width: 60px"></th>
                        <th>Nombre</th>
                        <th>Apodo</th>
                        <th>Posición</th>
                        <th>Teléfono</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="p in players" :key="p.id">
                        <td>
                            <v-avatar
                                size="40"
                                :color="getPositionColor(p.position)"
                            >
                                <v-img
                                    v-if="getPhotoUrl(p)"
                                    :src="getPhotoUrl(p)"
                                    cover
                                />
                                <span
                                    v-else
                                    class="text-white text-body-2 font-weight-bold"
                                >
                                    {{
                                        (p.nickname || p.user?.name || "?")
                                            .charAt(0)
                                            .toUpperCase()
                                    }}
                                </span>
                            </v-avatar>
                        </td>
                        <td>
                            <Link
                                :href="`/players/${p.id}`"
                                class="text-primary font-weight-medium"
                            >
                                {{ p.user?.name || "—" }}
                            </Link>
                        </td>
                        <td>{{ p.nickname || "—" }}</td>
                        <td>
                            <v-chip
                                v-if="p.position"
                                :color="getPositionColor(p.position)"
                                size="small"
                                variant="tonal"
                            >
                                {{ getPositionLabel(p.position) }}
                            </v-chip>
                            <span v-else>—</span>
                        </td>
                        <td>{{ p.phone || "—" }}</td>
                    </tr>
                </tbody>
            </v-table>
            <v-card-text v-else class="text-center pa-8">
                <v-icon size="64" color="grey-lighten-1" class="mb-4"
                    >mdi-account-group</v-icon
                >
                <p class="text-medium-emphasis">
                    No hay jugadores registrados aún.
                </p>
            </v-card-text>
        </v-card>
    </div>
</template>
