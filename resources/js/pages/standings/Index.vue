<script setup lang="ts">
import { computed } from "vue";
import { Head, router, Link } from "@inertiajs/vue3";

const props = defineProps<{
    standings: any[];
    tournaments: any[];
    tournamentId: string;
}>();

const goToTournament = (v: number | null) => {
    if (v) {
        router.get("/standings", { tournament_id: v }, { preserveState: true });
    } else {
        router.get("/standings", {}, { preserveState: true });
    }
};

const hasGroups = computed(() =>
    props.standings.some((s) => s.group_id && s.group),
);

const groupedStandings = computed(() => {
    if (!hasGroups.value) return null;

    const groups: Record<string, { name: string; standings: any[] }> = {};

    for (const s of props.standings) {
        if (s.group_id && s.group) {
            const key = String(s.group_id);
            if (!groups[key]) {
                groups[key] = { name: s.group.name, standings: [] };
            }
            groups[key].standings.push(s);
        }
    }

    for (const g of Object.values(groups)) {
        g.standings.sort(
            (a: any, b: any) =>
                b.points - a.points || b.goal_difference - a.goal_difference,
        );
    }

    return Object.values(groups);
});
</script>

<template>
    <Head><title>Tabla de Posiciones</title></Head>
    <div>
        <h1 class="text-h4 font-weight-bold mb-2">Tabla de Posiciones</h1>
        <p class="text-medium-emphasis mb-6">Clasificación por torneo</p>

        <v-card class="mb-4">
            <v-card-text>
                <v-autocomplete
                    label="Buscar torneo"
                    :items="props.tournaments"
                    item-title="name"
                    item-value="id"
                    :model-value="
                        props.tournamentId ? Number(props.tournamentId) : null
                    "
                    clearable
                    placeholder="Escribe para buscar..."
                    no-data-text="No se encontraron torneos"
                    @update:model-value="goToTournament"
                />
            </v-card-text>
        </v-card>

        <v-card v-if="props.tournamentId">
            <!-- Grouped standings -->
            <template v-if="groupedStandings && groupedStandings.length > 0">
                <div
                    v-for="(g, idx) in groupedStandings"
                    :key="idx"
                    class="mb-4"
                >
                    <v-card-title class="text-h6">{{ g.name }}</v-card-title>
                    <v-table>
                        <thead>
                            <tr>
                                <th>#</th>
                                <th>Equipo</th>
                                <th>PJ</th>
                                <th>G</th>
                                <th>E</th>
                                <th>P</th>
                                <th>GF</th>
                                <th>GC</th>
                                <th>DG</th>
                                <th class="font-weight-bold">Pts</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(s, i) in g.standings" :key="s.id">
                                <td>{{ i + 1 }}</td>
                                <td class="font-weight-medium">
                                    <Link
                                        :href="`/teams/${s.team?.id}`"
                                        class="text-primary"
                                    >
                                        {{ s.team?.name }}
                                    </Link>
                                </td>
                                <td>{{ s.played }}</td>
                                <td>{{ s.won }}</td>
                                <td>{{ s.drawn }}</td>
                                <td>{{ s.lost }}</td>
                                <td>{{ s.goals_for }}</td>
                                <td>{{ s.goals_against }}</td>
                                <td>
                                    {{ s.goal_difference > 0 ? "+" : ""
                                    }}{{ s.goal_difference }}
                                </td>
                                <td class="font-weight-bold text-primary">
                                    {{ s.points }}
                                </td>
                            </tr>
                        </tbody>
                    </v-table>
                </div>
            </template>

            <!-- Flat standings -->
            <v-table v-else-if="props.standings.length > 0">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Equipo</th>
                        <th>PJ</th>
                        <th>G</th>
                        <th>E</th>
                        <th>P</th>
                        <th>GF</th>
                        <th>GC</th>
                        <th>DG</th>
                        <th class="font-weight-bold">Pts</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(s, i) in props.standings" :key="s.id">
                        <td>{{ i + 1 }}</td>
                        <td class="font-weight-medium">
                            <Link
                                :href="`/teams/${s.team?.id}`"
                                class="text-primary"
                            >
                                {{ s.team?.name }}
                            </Link>
                        </td>
                        <td>{{ s.played }}</td>
                        <td>{{ s.won }}</td>
                        <td>{{ s.drawn }}</td>
                        <td>{{ s.lost }}</td>
                        <td>{{ s.goals_for }}</td>
                        <td>{{ s.goals_against }}</td>
                        <td>
                            {{ s.goal_difference > 0 ? "+" : ""
                            }}{{ s.goal_difference }}
                        </td>
                        <td class="font-weight-bold text-primary">
                            {{ s.points }}
                        </td>
                    </tr>
                </tbody>
            </v-table>

            <!-- Empty state -->
            <v-card-text v-else class="text-center pa-8">
                <v-icon size="64" color="grey-lighten-1" class="mb-4"
                    >mdi-podium</v-icon
                >
                <p class="text-medium-emphasis">
                    No hay clasificaciones para este torneo.
                </p>
            </v-card-text>
        </v-card>
    </div>
</template>
