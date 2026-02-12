<script setup lang="ts">
import { ref } from "vue";
import { Head, usePage, router, Link } from "@inertiajs/vue3";

const page = usePage<{ team: any; players: any[] }>();
const team = page.props.team;
const players = page.props.players || [];

const showAddPlayerDialog = ref(false);
const selectedPlayerId = ref<number | null>(null);
const availablePlayers = ref<any[]>([]);
const loadingPlayers = ref(false);

const openAddPlayerDialog = async () => {
    loadingPlayers.value = true;
    try {
        const response = await fetch(`/teams/${team.id}/available-players`);
        const data = await response.json();
        availablePlayers.value = data.players || [];
    } catch (e) {
        console.error(e);
    }
    loadingPlayers.value = false;
    showAddPlayerDialog.value = true;
};

const addPlayer = () => {
    if (selectedPlayerId.value) {
        router.post(
            `/teams/${team.id}/players`,
            { player_id: selectedPlayerId.value },
            {
                preserveScroll: true,
                onSuccess: () => {
                    showAddPlayerDialog.value = false;
                    selectedPlayerId.value = null;
                    // Refresh the page data to update the list
                    router.visit(window.location.href, {
                        preserveScroll: true,
                    });
                },
            },
        );
    }
};

const getPlayerDisplayName = (player: any) => {
    const name = player.user?.name || "Sin nombre";
    return player.nickname ? `${name} (${player.nickname})` : name;
};

const removePlayer = (playerId: number) => {
    if (confirm("¿Quitar este jugador del equipo?")) {
        router.delete(`/teams/${team.id}/players/${playerId}`, {
            preserveScroll: true,
            onSuccess: () => {
                router.visit(window.location.href, { preserveScroll: true });
            },
        });
    }
};
</script>

<template>
    <Head
        ><title>{{ team.name }}</title></Head
    >
    <div>
        <div class="d-flex justify-space-between align-center mb-6">
            <div class="d-flex align-center ga-4">
                <v-avatar size="64" :color="team.color || 'primary'">
                    <v-icon size="32">mdi-shield</v-icon>
                </v-avatar>
                <div>
                    <h1 class="text-h4 font-weight-bold">{{ team.name }}</h1>
                    <p v-if="team.contact_phone" class="text-medium-emphasis">
                        📞 {{ team.contact_phone }}
                    </p>
                </div>
            </div>
            <Link :href="`/teams/${team.id}/edit`" class="text-decoration-none">
                <v-btn color="primary" prepend-icon="mdi-pencil">Editar</v-btn>
            </Link>
        </div>

        <v-row>
            <v-col cols="12" md="8">
                <v-card>
                    <v-card-title
                        class="d-flex justify-space-between align-center"
                    >
                        <span>Jugadores ({{ players.length }})</span>
                        <v-btn
                            size="small"
                            color="primary"
                            variant="tonal"
                            prepend-icon="mdi-plus"
                            @click="openAddPlayerDialog"
                        >
                            Agregar Jugador
                        </v-btn>
                    </v-card-title>
                    <v-table v-if="players.length > 0">
                        <thead>
                            <tr>
                                <th>Jugador</th>
                                <th>Apodo</th>
                                <th>Posición</th>
                                <th>Dorsal</th>
                                <th></th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr
                                v-for="tp in players"
                                :key="tp.player?.id || tp.id"
                            >
                                <td>
                                    <Link
                                        :href="`/players/${tp.player?.id || tp.player_id}`"
                                        class="text-primary font-weight-medium"
                                    >
                                        {{ tp.player?.user?.name || "—" }}
                                    </Link>
                                </td>
                                <td>{{ tp.player?.nickname || "—" }}</td>
                                <td>{{ tp.player?.position || "—" }}</td>
                                <td>{{ tp.jersey_number || "—" }}</td>
                                <td>
                                    <v-btn
                                        icon="mdi-close"
                                        size="x-small"
                                        variant="text"
                                        color="error"
                                        @click="
                                            removePlayer(
                                                tp.player?.id || tp.player_id,
                                            )
                                        "
                                    />
                                </td>
                            </tr>
                        </tbody>
                    </v-table>
                    <v-card-text
                        v-else
                        class="text-center text-medium-emphasis"
                    >
                        Este equipo aún no tiene jugadores
                    </v-card-text>
                </v-card>
            </v-col>

            <v-col cols="12" md="4">
                <v-card>
                    <v-card-title>Estadísticas</v-card-title>
                    <v-card-text>
                        <div class="d-flex justify-space-between mb-2">
                            <span class="text-medium-emphasis">Jugadores</span>
                            <span class="font-weight-bold">{{
                                players.length
                            }}</span>
                        </div>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>

        <!-- Dialog para agregar jugador -->
        <v-dialog v-model="showAddPlayerDialog" max-width="400">
            <v-card>
                <v-card-title>Agregar Jugador al Equipo</v-card-title>
                <v-card-text>
                    <v-select
                        v-model="selectedPlayerId"
                        :items="
                            availablePlayers.map((p) => ({
                                title: getPlayerDisplayName(p),
                                value: p.id,
                            }))
                        "
                        label="Seleccionar Jugador"
                        :loading="loadingPlayers"
                        :disabled="availablePlayers.length === 0"
                    />
                    <p
                        v-if="availablePlayers.length === 0 && !loadingPlayers"
                        class="text-medium-emphasis"
                    >
                        No hay jugadores disponibles. Primero crea jugadores.
                    </p>
                </v-card-text>
                <v-card-actions>
                    <v-spacer />
                    <v-btn variant="text" @click="showAddPlayerDialog = false"
                        >Cancelar</v-btn
                    >
                    <v-btn
                        color="primary"
                        :disabled="!selectedPlayerId"
                        @click="addPlayer"
                        >Agregar</v-btn
                    >
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>
