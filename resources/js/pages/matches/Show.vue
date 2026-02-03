<script setup lang="ts">
import { ref } from "vue";
import { Head, usePage, Link, useForm } from "@inertiajs/vue3";

const page = usePage<{ match: any; events: any[] }>();
const match = page.props.match;
const events = page.props.events || [];

const showResultDialog = ref(false);
const resultForm = useForm({
  home_score: match.home_score || 0,
  away_score: match.away_score || 0,
});

const formatDate = (dateStr: string | null) => {
  if (!dateStr) return "Fecha por definir";
  return new Date(dateStr).toLocaleDateString("es-CO", {
    weekday: "long",
    day: "numeric",
    month: "long",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
};

const submitResult = () => {
  resultForm.post(`/matches/${match.id}/result`, {
    preserveScroll: true,
    onSuccess: () => {
      showResultDialog.value = false;
    },
  });
};

const statusColors: Record<string, string> = {
  scheduled: "info",
  live: "warning",
  completed: "success",
  cancelled: "error",
};

const statusLabels: Record<string, string> = {
  scheduled: "Programado",
  live: "En vivo",
  completed: "Finalizado",
  cancelled: "Cancelado",
};
</script>

<template>
  <Head><title>Partido - {{ match.home_team?.name }} vs {{ match.away_team?.name }}</title></Head>
  <div>
    <Link href="/matches" class="text-decoration-none">
      <v-btn variant="text" prepend-icon="mdi-arrow-left" class="mb-4">Volver a Partidos</v-btn>
    </Link>

    <v-card class="mb-6">
      <v-card-text class="text-center pa-8">
        <div class="text-caption text-medium-emphasis mb-2">
          <Link :href="`/tournaments/${match.tournament?.id}`" class="text-primary">
            {{ match.tournament?.name }}
          </Link>
        </div>
        <div class="text-caption text-medium-emphasis mb-4">{{ formatDate(match.match_date) }}</div>

        <v-row align="center" justify="center">
          <v-col cols="4" class="text-center">
            <Link :href="`/teams/${match.home_team?.id}`" class="text-decoration-none">
              <v-avatar size="64" color="primary" class="mb-2">
                <v-icon size="32">mdi-shield</v-icon>
              </v-avatar>
              <div class="text-h6 font-weight-bold">{{ match.home_team?.name || "TBD" }}</div>
            </Link>
          </v-col>

          <v-col cols="4" class="text-center">
            <div v-if="match.status === 'completed'" class="text-h3 font-weight-bold">
              {{ match.home_score }} - {{ match.away_score }}
            </div>
            <div v-else class="text-h5 text-medium-emphasis">vs</div>
            <v-chip size="small" :color="statusColors[match.status] || 'grey'" class="mt-2">
              {{ statusLabels[match.status] || match.status }}
            </v-chip>
          </v-col>

          <v-col cols="4" class="text-center">
            <Link :href="`/teams/${match.away_team?.id}`" class="text-decoration-none">
              <v-avatar size="64" color="secondary" class="mb-2">
                <v-icon size="32">mdi-shield</v-icon>
              </v-avatar>
              <div class="text-h6 font-weight-bold">{{ match.away_team?.name || "TBD" }}</div>
            </Link>
          </v-col>
        </v-row>

        <div v-if="match.venue" class="mt-4 text-medium-emphasis">
          <v-icon size="small">mdi-map-marker</v-icon>
          {{ match.venue.name }}
        </div>
      </v-card-text>

      <v-card-actions class="justify-center pb-4">
        <Link :href="`/matches/${match.id}/edit`" class="text-decoration-none">
          <v-btn variant="outlined" prepend-icon="mdi-pencil">Editar</v-btn>
        </Link>
        <v-btn
          v-if="match.status !== 'completed'"
          color="success"
          prepend-icon="mdi-scoreboard"
          @click="showResultDialog = true"
        >
          Registrar Resultado
        </v-btn>
      </v-card-actions>
    </v-card>

    <v-card v-if="events.length > 0">
      <v-card-title>Eventos del Partido</v-card-title>
      <v-list>
        <v-list-item v-for="e in events" :key="e.id">
          <template #prepend>
            <v-avatar size="32" color="grey-lighten-3">
              <span class="text-caption font-weight-bold">{{ e.minute }}'</span>
            </v-avatar>
          </template>
          <v-list-item-title>{{ e.event_type }}</v-list-item-title>
          <v-list-item-subtitle>{{ e.player?.name }}</v-list-item-subtitle>
        </v-list-item>
      </v-list>
    </v-card>

    <!-- Dialog para registrar resultado -->
    <v-dialog v-model="showResultDialog" max-width="400">
      <v-card>
        <v-card-title>Registrar Resultado</v-card-title>
        <v-card-text>
          <v-row>
            <v-col cols="6" class="text-center">
              <p class="text-caption font-weight-medium mb-2">{{ match.home_team?.name }}</p>
              <v-text-field
                v-model.number="resultForm.home_score"
                type="number"
                min="0"
                variant="outlined"
                hide-details
                class="text-center"
              />
            </v-col>
            <v-col cols="6" class="text-center">
              <p class="text-caption font-weight-medium mb-2">{{ match.away_team?.name }}</p>
              <v-text-field
                v-model.number="resultForm.away_score"
                type="number"
                min="0"
                variant="outlined"
                hide-details
                class="text-center"
              />
            </v-col>
          </v-row>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="showResultDialog = false">Cancelar</v-btn>
          <v-btn color="success" :loading="resultForm.processing" @click="submitResult">
            Guardar Resultado
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
