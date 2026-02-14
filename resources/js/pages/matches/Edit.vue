<script setup lang="ts">
import { Head, useForm, usePage, Link } from "@inertiajs/vue3";

const page = usePage<{
    match: any;
    tournaments: any[];
    teams: any[];
    venues: any[];
    groups?: any[];
    errors?: Record<string, string>;
}>();

const match = page.props.match;
const tournaments = page.props.tournaments || [];
const teams = page.props.teams || [];
const venues = page.props.venues || [];
const groups = page.props.groups || [];

const formatDateForInput = (dateStr: string | null) => {
    if (!dateStr) return "";
    const date = new Date(dateStr);
    return date.toISOString().slice(0, 16);
};

const form = useForm({
    tournament_id: match.tournament_id,
    home_team_id: match.home_team_id,
    away_team_id: match.away_team_id,
    venue_id: match.venue_id,
    match_date: formatDateForInput(match.match_date),
    matchday: match.round || 1,
    group_id: match.group_id || null,
});

const submit = () => form.put(`/matches/${match.id}`);
</script>

<template>
    <Head><title>Editar Partido</title></Head>
    <div>
        <h1 class="text-h4 font-weight-bold mb-6">Editar Partido</h1>
        <v-card>
            <v-card-text class="pa-6">
                <v-form @submit.prevent="submit">
                    <v-select
                        v-model="form.tournament_id"
                        :items="tournaments"
                        item-title="name"
                        item-value="id"
                        label="Torneo *"
                        :error-messages="form.errors.tournament_id"
                        class="mb-4"
                        required
                    />

                    <v-select
                        v-if="groups.length > 0"
                        v-model="form.group_id"
                        :items="groups"
                        item-title="name"
                        item-value="id"
                        label="Grupo (opcional)"
                        clearable
                        :error-messages="form.errors.group_id"
                        class="mb-4"
                    />

                    <v-row>
                        <v-col cols="12" md="6">
                            <v-select
                                v-model="form.home_team_id"
                                :items="teams"
                                item-title="name"
                                item-value="id"
                                label="Equipo Local *"
                                :error-messages="form.errors.home_team_id"
                                required
                            />
                        </v-col>
                        <v-col cols="12" md="6">
                            <v-select
                                v-model="form.away_team_id"
                                :items="teams"
                                item-title="name"
                                item-value="id"
                                label="Equipo Visitante *"
                                :error-messages="form.errors.away_team_id"
                                required
                            />
                        </v-col>
                    </v-row>

                    <v-row>
                        <v-col cols="12" md="6">
                            <v-text-field
                                v-model="form.match_date"
                                label="Fecha y Hora *"
                                type="datetime-local"
                                :error-messages="form.errors.match_date"
                                required
                            />
                        </v-col>
                        <v-col cols="12" md="6">
                            <v-text-field
                                v-model.number="form.matchday"
                                label="Jornada *"
                                type="number"
                                min="1"
                                :error-messages="form.errors.matchday"
                                required
                            />
                        </v-col>
                    </v-row>

                    <v-select
                        v-model="form.venue_id"
                        :items="venues"
                        item-title="name"
                        item-value="id"
                        label="Sede *"
                        :error-messages="form.errors.venue_id"
                        class="mb-4"
                        required
                    />

                    <div class="d-flex ga-4">
                        <v-btn
                            type="submit"
                            color="primary"
                            :loading="form.processing"
                            :disabled="form.processing"
                            >Guardar</v-btn
                        >
                        <Link
                            :href="`/matches/${match.id}`"
                            class="text-decoration-none"
                        >
                            <v-btn variant="outlined">Cancelar</v-btn>
                        </Link>
                    </div>
                </v-form>
            </v-card-text>
        </v-card>
    </div>
</template>
