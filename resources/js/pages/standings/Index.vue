<script setup lang="ts">
import { Head } from "@inertiajs/vue3";
import PageHeader from "@/components/PageHeader.vue";

const props = defineProps<{ standings: any[]; tournaments: any[]; tournamentId: string }>();

const goToTournament = (v: number | null) => {
    window.location.href = v ? `/standings?tournament_id=${v}` : '/standings';
};
</script>

<template>
    <Head title="Posiciones" />
    <PageHeader title="Tabla de Posiciones" subtitle="Clasificación por torneo" />

    <v-card class="mb-4 pa-4">
        <v-select
            label="Filtrar por torneo"
            :items="props.tournaments"
            item-title="name"
            item-value="id"
            :model-value="props.tournamentId ? Number(props.tournamentId) : null"
            clearable
            @update:model-value="goToTournament"
        />
    </v-card>

    <v-card>
        <v-table v-if="props.standings.length > 0">
            <thead>
                <tr>
                    <th>#</th><th>Equipo</th><th>Grupo</th><th>PJ</th><th>G</th><th>E</th><th>P</th><th>GF</th><th>GC</th><th>DG</th><th class="font-weight-bold">Pts</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(s, i) in props.standings" :key="s.id">
                    <td>{{ i + 1 }}</td>
                    <td class="font-weight-medium">
                        <a :href="`/teams/${s.team?.id}`" class="text-primary">{{ s.team?.name }}</a>
                    </td>
                    <td>{{ s.group?.name || '—' }}</td>
                    <td>{{ s.played }}</td>
                    <td>{{ s.won }}</td>
                    <td>{{ s.drawn }}</td>
                    <td>{{ s.lost }}</td>
                    <td>{{ s.goals_for }}</td>
                    <td>{{ s.goals_against }}</td>
                    <td>{{ s.goal_difference }}</td>
                    <td class="font-weight-bold text-primary">{{ s.points }}</td>
                </tr>
            </tbody>
        </v-table>
        <v-card-text v-else class="text-center text-medium-emphasis pa-6">
            Selecciona un torneo para ver las posiciones.
        </v-card-text>
    </v-card>
</template>
