<script setup lang="ts">
import { Head, useForm, usePage, Link, router } from "@inertiajs/vue3";

const page = usePage<{
    tournaments: any[];
    teams: any[];
    venues: any[];
    tournamentId?: string;
    errors?: Record<string, string>;
}>();

const tournaments = page.props.tournaments || [];
const teams = page.props.teams || [];
const venues = page.props.venues || [];

const form = useForm({
    tournament_id: page.props.tournamentId
        ? parseInt(page.props.tournamentId)
        : (null as number | null),
    home_team_id: null as number | null,
    away_team_id: null as number | null,
    venue_id: null as number | null,
    match_date: "",
    matchday: 1,
});

const submit = () => {
    form.post("/matches", {
        preserveScroll: true,
        onSuccess: () => {
            router.visit("/matches");
        },
    });
};
</script>

<template>
    <Head><title>Crear Partido</title></Head>
    <div>
        <h1 class="text-h4 font-weight-bold mb-6">Crear Partido</h1>
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
                            >Crear Partido</v-btn
                        >
                        <Link href="/matches" class="text-decoration-none">
                            <v-btn variant="outlined">Cancelar</v-btn>
                        </Link>
                    </div>
                </v-form>
            </v-card-text>
        </v-card>
    </div>
</template>
