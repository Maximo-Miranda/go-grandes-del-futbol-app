<script setup lang="ts">
import { ref, computed } from "vue";
import { Head, usePage, router, Link } from "@inertiajs/vue3";
import ConfirmDialog from "@/components/ConfirmDialog.vue";

const page = usePage<{
    tournament: any;
    teams: any[];
    matches: any[];
    standings: any[];
    availableTeams: any[];
}>();
const tournament = page.props.tournament;
const teams = page.props.teams || [];
const matches = page.props.matches || [];
const standings = page.props.standings || [];
const availableTeams = ref(page.props.availableTeams || []);

const isGroupKnockout = tournament.format === "group_knockout";
const groups = tournament.groups || [];

// Add team dialog
const showAddTeamDialog = ref(false);
const selectedTeamId = ref<number | null>(null);
const loadingTeams = ref(false);

// Group dialogs
const showCreateGroupDialog = ref(false);
const newGroupName = ref("");
const showEditGroupDialog = ref(false);
const editingGroup = ref<{ id: number; name: string } | null>(null);
const editGroupName = ref("");
const showDeleteGroupDialog = ref(false);
const groupToDelete = ref<number | null>(null);
const deletingGroup = ref(false);

const formatDate = (dateStr: string | null) => {
    if (!dateStr) return "Sin fecha";
    return new Date(dateStr).toLocaleDateString("es-CO", {
        weekday: "short",
        day: "numeric",
        month: "short",
        hour: "2-digit",
        minute: "2-digit",
    });
};

// Team management
const openAddTeamDialog = async () => {
    loadingTeams.value = true;
    try {
        const response = await fetch(
            `/tournaments/${tournament.id}/available-teams`,
        );
        const data = await response.json();
        availableTeams.value = data.teams || [];
    } catch (e) {
        console.error(e);
    }
    loadingTeams.value = false;
    showAddTeamDialog.value = true;
};

const addTeam = () => {
    if (selectedTeamId.value) {
        router.post(
            `/tournaments/${tournament.id}/teams`,
            { team_id: selectedTeamId.value },
            {
                preserveScroll: true,
                onSuccess: () => {
                    showAddTeamDialog.value = false;
                    selectedTeamId.value = null;
                    router.visit(window.location.href, {
                        preserveScroll: true,
                    });
                },
            },
        );
    }
};

const removeTeam = (teamId: number) => {
    if (confirm("¿Quitar este equipo del torneo?")) {
        router.delete(`/tournaments/${tournament.id}/teams/${teamId}`, {
            preserveScroll: true,
            onSuccess: () => {
                router.visit(window.location.href, { preserveScroll: true });
            },
        });
    }
};

// Group management
const createGroup = () => {
    if (newGroupName.value.trim()) {
        router.post(
            `/tournaments/${tournament.id}/groups`,
            { name: newGroupName.value.trim() },
            {
                preserveScroll: true,
                onSuccess: () => {
                    showCreateGroupDialog.value = false;
                    newGroupName.value = "";
                    router.visit(window.location.href, {
                        preserveScroll: true,
                    });
                },
            },
        );
    }
};

const openEditGroupDialog = (group: { id: number; name: string }) => {
    editingGroup.value = group;
    editGroupName.value = group.name;
    showEditGroupDialog.value = true;
};

const updateGroup = () => {
    if (editingGroup.value && editGroupName.value.trim()) {
        router.put(
            `/tournaments/${tournament.id}/groups/${editingGroup.value.id}`,
            { name: editGroupName.value.trim() },
            {
                preserveScroll: true,
                onSuccess: () => {
                    showEditGroupDialog.value = false;
                    editingGroup.value = null;
                    router.visit(window.location.href, {
                        preserveScroll: true,
                    });
                },
            },
        );
    }
};

const confirmDeleteGroup = (groupId: number) => {
    groupToDelete.value = groupId;
    showDeleteGroupDialog.value = true;
};

const deleteGroup = () => {
    if (groupToDelete.value) {
        deletingGroup.value = true;
        router.delete(
            `/tournaments/${tournament.id}/groups/${groupToDelete.value}`,
            {
                preserveScroll: true,
                onSuccess: () => {
                    router.visit(window.location.href, {
                        preserveScroll: true,
                    });
                },
                onFinish: () => {
                    deletingGroup.value = false;
                    showDeleteGroupDialog.value = false;
                    groupToDelete.value = null;
                },
            },
        );
    }
};

const assignTeamToGroup = (teamId: number, groupId: number | null) => {
    router.put(
        `/tournaments/${tournament.id}/teams/${teamId}/group`,
        { group_id: groupId || "" },
        {
            preserveScroll: true,
            onSuccess: () => {
                router.visit(window.location.href, { preserveScroll: true });
            },
        },
    );
};

// Group select items for team assignment
const groupSelectItems = computed(() => {
    return [
        { title: "Sin grupo", value: null },
        ...groups.map((g: any) => ({ title: g.name, value: g.id })),
    ];
});

// Grouped standings
const groupedStandings = computed(() => {
    if (!isGroupKnockout || groups.length === 0) return null;

    const result: Record<string, { name: string; standings: any[] }> = {};

    for (const s of standings) {
        if (s.group_id && s.group) {
            const key = String(s.group_id);
            if (!result[key]) {
                result[key] = { name: s.group.name, standings: [] };
            }
            result[key].standings.push(s);
        }
    }

    return Object.values(result);
});
</script>

<template>
    <Head
        ><title>{{ tournament.name }}</title></Head
    >
    <div>
        <div class="d-flex justify-space-between align-center mb-6">
            <h1 class="text-h4 font-weight-bold">{{ tournament.name }}</h1>
            <div class="d-flex ga-2">
                <Link
                    :href="`/matches/create?tournament_id=${tournament.id}`"
                    class="text-decoration-none"
                >
                    <v-btn
                        color="success"
                        variant="tonal"
                        prepend-icon="mdi-soccer"
                        >Nuevo Partido</v-btn
                    >
                </Link>
                <Link
                    :href="`/tournaments/${tournament.id}/edit`"
                    class="text-decoration-none"
                >
                    <v-btn color="primary" prepend-icon="mdi-pencil"
                        >Editar</v-btn
                    >
                </Link>
            </div>
        </div>

        <v-row>
            <v-col cols="12" md="8">
                <v-card class="mb-4">
                    <v-card-title>Información</v-card-title>
                    <v-card-text>
                        <p v-if="tournament.description" class="mb-4">
                            {{ tournament.description }}
                        </p>
                        <div class="d-flex flex-wrap ga-2">
                            <v-chip>{{ tournament.format }}</v-chip>
                            <v-chip>{{ tournament.game_type }}</v-chip>
                            <v-chip
                                :color="
                                    tournament.status === 'active'
                                        ? 'success'
                                        : 'grey'
                                "
                                >{{ tournament.status }}</v-chip
                            >
                            <v-chip
                                v-if="tournament.venue"
                                color="info"
                                prepend-icon="mdi-map-marker"
                                >{{ tournament.venue.name }}</v-chip
                            >
                            <v-chip
                                v-if="tournament.start_date"
                                prepend-icon="mdi-calendar"
                            >
                                {{
                                    new Date(
                                        tournament.start_date,
                                    ).toLocaleDateString("es-CO")
                                }}
                                <span v-if="tournament.end_date">
                                    —
                                    {{
                                        new Date(
                                            tournament.end_date,
                                        ).toLocaleDateString("es-CO")
                                    }}</span
                                >
                            </v-chip>
                        </div>
                    </v-card-text>
                </v-card>

                <v-card v-if="matches.length > 0" class="mb-4">
                    <v-card-title
                        class="d-flex justify-space-between align-center"
                    >
                        <span>Partidos ({{ matches.length }})</span>
                        <Link
                            :href="`/matches?tournament_id=${tournament.id}`"
                            class="text-decoration-none"
                        >
                            <v-btn size="small" variant="text">Ver todos</v-btn>
                        </Link>
                    </v-card-title>
                    <v-list>
                        <Link
                            v-for="m in matches"
                            :key="m.id"
                            :href="`/matches/${m.id}`"
                            class="text-decoration-none"
                        >
                            <v-list-item>
                                <v-list-item-title
                                    class="d-flex justify-space-between"
                                >
                                    <span
                                        >{{ m.home_team?.name || "TBD" }} vs
                                        {{ m.away_team?.name || "TBD" }}</span
                                    >
                                    <v-chip
                                        v-if="m.status === 'completed'"
                                        size="x-small"
                                        color="success"
                                    >
                                        {{ m.home_score }} - {{ m.away_score }}
                                    </v-chip>
                                </v-list-item-title>
                                <v-list-item-subtitle
                                    >{{ formatDate(m.match_date) }} —
                                    {{ m.status }}</v-list-item-subtitle
                                >
                            </v-list-item>
                        </Link>
                    </v-list>
                </v-card>
                <v-card v-else class="mb-4">
                    <v-card-text class="text-center pa-8 text-medium-emphasis">
                        <v-icon size="48" color="grey-lighten-1" class="mb-2"
                            >mdi-soccer-field</v-icon
                        >
                        <p>No hay partidos programados</p>
                        <Link
                            :href="`/matches/create?tournament_id=${tournament.id}`"
                            class="text-decoration-none"
                        >
                            <v-btn
                                color="primary"
                                variant="tonal"
                                size="small"
                                class="mt-2"
                                >Crear Partido</v-btn
                            >
                        </Link>
                    </v-card-text>
                </v-card>
            </v-col>

            <v-col cols="12" md="4">
                <!-- Groups card (only for group_knockout) -->
                <v-card v-if="isGroupKnockout" class="mb-4">
                    <v-card-title
                        class="d-flex justify-space-between align-center"
                    >
                        <span>Grupos ({{ groups.length }})</span>
                        <v-btn
                            size="small"
                            color="primary"
                            variant="tonal"
                            prepend-icon="mdi-plus"
                            @click="showCreateGroupDialog = true"
                        >
                            Crear
                        </v-btn>
                    </v-card-title>
                    <v-list density="compact">
                        <v-list-item v-for="g in groups" :key="g.id">
                            <v-list-item-title>{{ g.name }}</v-list-item-title>
                            <template #append>
                                <v-btn
                                    icon="mdi-pencil"
                                    size="x-small"
                                    variant="text"
                                    @click="openEditGroupDialog(g)"
                                />
                                <v-btn
                                    icon="mdi-close"
                                    size="x-small"
                                    variant="text"
                                    color="error"
                                    @click="confirmDeleteGroup(g.id)"
                                />
                            </template>
                        </v-list-item>
                        <v-list-item v-if="groups.length === 0">
                            <v-list-item-title class="text-medium-emphasis"
                                >Sin grupos. Crea grupos para organizar los
                                equipos.</v-list-item-title
                            >
                        </v-list-item>
                    </v-list>
                </v-card>

                <!-- Teams card -->
                <v-card class="mb-4">
                    <v-card-title
                        class="d-flex justify-space-between align-center"
                    >
                        <span>Equipos ({{ teams.length }})</span>
                        <v-btn
                            size="small"
                            color="primary"
                            variant="tonal"
                            prepend-icon="mdi-plus"
                            @click="openAddTeamDialog"
                        >
                            Agregar
                        </v-btn>
                    </v-card-title>
                    <v-list density="compact">
                        <v-list-item
                            v-for="t in teams"
                            :key="t.team?.id || t.id"
                        >
                            <template #prepend>
                                <v-avatar size="32" color="primary">
                                    <v-icon>mdi-shield</v-icon>
                                </v-avatar>
                            </template>
                            <Link
                                :href="`/teams/${t.team?.id || t.team_id}`"
                                class="text-decoration-none"
                            >
                                <v-list-item-title class="text-primary">{{
                                    t.team?.name || t.name
                                }}</v-list-item-title>
                            </Link>
                            <v-list-item-subtitle
                                v-if="isGroupKnockout && groups.length > 0"
                            >
                                <v-select
                                    :model-value="t.group_id || null"
                                    :items="groupSelectItems"
                                    density="compact"
                                    variant="plain"
                                    hide-details
                                    style="max-width: 140px"
                                    @update:model-value="
                                        assignTeamToGroup(
                                            t.team?.id || t.team_id,
                                            $event,
                                        )
                                    "
                                />
                            </v-list-item-subtitle>
                            <template #append>
                                <v-btn
                                    icon="mdi-close"
                                    size="x-small"
                                    variant="text"
                                    color="error"
                                    @click="removeTeam(t.team?.id || t.team_id)"
                                />
                            </template>
                        </v-list-item>
                        <v-list-item v-if="teams.length === 0">
                            <v-list-item-title class="text-medium-emphasis"
                                >Sin equipos aún</v-list-item-title
                            >
                        </v-list-item>
                    </v-list>
                </v-card>

                <!-- Standings card -->
                <v-card v-if="standings.length > 0">
                    <v-card-title>Clasificación</v-card-title>

                    <!-- Grouped standings -->
                    <template
                        v-if="groupedStandings && groupedStandings.length > 0"
                    >
                        <div
                            v-for="(g, idx) in groupedStandings"
                            :key="idx"
                            class="mb-2"
                        >
                            <v-card-subtitle
                                class="font-weight-bold text-uppercase"
                                >{{ g.name }}</v-card-subtitle
                            >
                            <v-table density="compact">
                                <thead>
                                    <tr>
                                        <th>#</th>
                                        <th>Equipo</th>
                                        <th>Pts</th>
                                        <th>PJ</th>
                                        <th>DG</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr
                                        v-for="(s, i) in g.standings"
                                        :key="s.id"
                                    >
                                        <td>{{ i + 1 }}</td>
                                        <td>
                                            <Link
                                                :href="`/teams/${s.team?.id}`"
                                                class="text-primary"
                                                >{{ s.team?.name }}</Link
                                            >
                                        </td>
                                        <td class="font-weight-bold">
                                            {{ s.points }}
                                        </td>
                                        <td>{{ s.played }}</td>
                                        <td>
                                            {{ s.goal_difference > 0 ? "+" : ""
                                            }}{{ s.goal_difference }}
                                        </td>
                                    </tr>
                                </tbody>
                            </v-table>
                        </div>
                    </template>

                    <!-- Flat standings -->
                    <v-table v-else density="compact">
                        <thead>
                            <tr>
                                <th>#</th>
                                <th>Equipo</th>
                                <th>Pts</th>
                                <th>PJ</th>
                                <th>DG</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(s, i) in standings" :key="s.id">
                                <td>{{ i + 1 }}</td>
                                <td>
                                    <Link
                                        :href="`/teams/${s.team?.id}`"
                                        class="text-primary"
                                        >{{ s.team?.name }}</Link
                                    >
                                </td>
                                <td class="font-weight-bold">{{ s.points }}</td>
                                <td>{{ s.played }}</td>
                                <td>
                                    {{ s.goal_difference > 0 ? "+" : ""
                                    }}{{ s.goal_difference }}
                                </td>
                            </tr>
                        </tbody>
                    </v-table>
                </v-card>
            </v-col>
        </v-row>

        <!-- Add Team Dialog -->
        <v-dialog v-model="showAddTeamDialog" max-width="400">
            <v-card>
                <v-card-title>Agregar Equipo al Torneo</v-card-title>
                <v-card-text>
                    <v-select
                        v-model="selectedTeamId"
                        :items="availableTeams"
                        item-title="name"
                        item-value="id"
                        label="Seleccionar Equipo"
                        :loading="loadingTeams"
                        :disabled="availableTeams.length === 0"
                    />
                    <p
                        v-if="availableTeams.length === 0 && !loadingTeams"
                        class="text-medium-emphasis"
                    >
                        No hay equipos disponibles.
                        <Link href="/teams/create" class="text-primary"
                            >Crea uno primero</Link
                        >.
                    </p>
                </v-card-text>
                <v-card-actions>
                    <v-spacer />
                    <v-btn variant="text" @click="showAddTeamDialog = false"
                        >Cancelar</v-btn
                    >
                    <v-btn
                        color="primary"
                        :disabled="!selectedTeamId"
                        @click="addTeam"
                        >Agregar</v-btn
                    >
                </v-card-actions>
            </v-card>
        </v-dialog>

        <!-- Create Group Dialog -->
        <v-dialog v-model="showCreateGroupDialog" max-width="400">
            <v-card>
                <v-card-title>Crear Grupo</v-card-title>
                <v-card-text>
                    <v-text-field
                        v-model="newGroupName"
                        label="Nombre del Grupo *"
                        placeholder="Ej: Grupo A"
                        autofocus
                    />
                </v-card-text>
                <v-card-actions>
                    <v-spacer />
                    <v-btn variant="text" @click="showCreateGroupDialog = false"
                        >Cancelar</v-btn
                    >
                    <v-btn
                        color="primary"
                        :disabled="!newGroupName.trim()"
                        @click="createGroup"
                        >Crear</v-btn
                    >
                </v-card-actions>
            </v-card>
        </v-dialog>

        <!-- Edit Group Dialog -->
        <v-dialog v-model="showEditGroupDialog" max-width="400">
            <v-card>
                <v-card-title>Editar Grupo</v-card-title>
                <v-card-text>
                    <v-text-field
                        v-model="editGroupName"
                        label="Nombre del Grupo *"
                        autofocus
                    />
                </v-card-text>
                <v-card-actions>
                    <v-spacer />
                    <v-btn variant="text" @click="showEditGroupDialog = false"
                        >Cancelar</v-btn
                    >
                    <v-btn
                        color="primary"
                        :disabled="!editGroupName.trim()"
                        @click="updateGroup"
                        >Guardar</v-btn
                    >
                </v-card-actions>
            </v-card>
        </v-dialog>

        <!-- Delete Group Confirm Dialog -->
        <ConfirmDialog
            v-model="showDeleteGroupDialog"
            title="¿Eliminar grupo?"
            message="Los equipos y partidos asignados a este grupo quedarán sin grupo. Esta acción no se puede deshacer."
            confirm-text="Eliminar"
            :loading="deletingGroup"
            @confirm="deleteGroup"
            @cancel="
                showDeleteGroupDialog = false;
                groupToDelete = null;
            "
        />
    </div>
</template>
