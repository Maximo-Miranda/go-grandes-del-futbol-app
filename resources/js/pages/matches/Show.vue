<script setup lang="ts">
import { ref, computed } from "vue";
import { Head, usePage, Link, useForm, router } from "@inertiajs/vue3";

interface Player {
  id: number;
  name: string;
  nickname?: string;
  position?: string;
  photo?: string;
}

interface MatchEvent {
  id: number;
  match_id: number;
  player_id: number;
  team_id: number;
  event_type: string;
  minute: number;
  player?: Player;
  team?: { id: number; name: string };
}

interface MatchLineup {
  id: number;
  match_id: number;
  player_id: number;
  team_id: number;
  is_starter: boolean;
  minutes_played: number;
  player?: Player;
  team?: { id: number; name: string };
}

const page = usePage<{
  match: any;
  events: MatchEvent[];
  lineups: MatchLineup[];
  homeTeamPlayers: Player[];
  awayTeamPlayers: Player[];
}>();

const match = page.props.match;
const events = page.props.events || [];
const lineups = page.props.lineups || [];
const homeTeamPlayers = page.props.homeTeamPlayers || [];
const awayTeamPlayers = page.props.awayTeamPlayers || [];

// Dialogs
const showResultDialog = ref(false);
const showEventDialog = ref(false);
const showLineupDialog = ref(false);

// Forms
const resultForm = useForm({
  home_score: match.home_score || 0,
  away_score: match.away_score || 0,
});

const eventForm = useForm({
  player_id: null as number | null,
  team_id: null as number | null,
  event_type: "",
  minute: 0,
});

const lineupForm = useForm({
  player_id: null as number | null,
  team_id: null as number | null,
  is_starter: true,
  minutes_played: 90,
});

const selectedTeamForEvent = ref<"home" | "away" | null>(null);
const selectedTeamForLineup = ref<"home" | "away" | null>(null);

const eventTypes = [
  { title: "⚽ Gol", value: "goal" },
  { title: "🅰️ Asistencia", value: "assist" },
  { title: "🟨 Tarjeta Amarilla", value: "yellow_card" },
  { title: "🟥 Tarjeta Roja", value: "red_card" },
  { title: "🔄 Sustitución (Entra)", value: "sub_in" },
  { title: "🔄 Sustitución (Sale)", value: "sub_out" },
];

const playersForEvent = computed(() => {
  if (selectedTeamForEvent.value === "home") {
    eventForm.team_id = match.home_team_id;
    return homeTeamPlayers;
  } else if (selectedTeamForEvent.value === "away") {
    eventForm.team_id = match.away_team_id;
    return awayTeamPlayers;
  }
  return [];
});

const playersForLineup = computed(() => {
  if (selectedTeamForLineup.value === "home") {
    lineupForm.team_id = match.home_team_id;
    return homeTeamPlayers;
  } else if (selectedTeamForLineup.value === "away") {
    lineupForm.team_id = match.away_team_id;
    return awayTeamPlayers;
  }
  return [];
});

const homeLineups = computed(() => lineups.filter((l) => l.team_id === match.home_team_id));
const awayLineups = computed(() => lineups.filter((l) => l.team_id === match.away_team_id));

const homeEvents = computed(() => events.filter((e) => e.team_id === match.home_team_id));
const awayEvents = computed(() => events.filter((e) => e.team_id === match.away_team_id));

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
      router.visit(window.location.href, { preserveScroll: true });
    },
  });
};

const closeMatch = () => {
  if (confirm("¿Cerrar el partido y finalizar? Esta acción actualizará las clasificaciones.")) {
    router.post(`/matches/${match.id}/close`, {}, {
      onSuccess: () => {
        router.visit(window.location.href, { preserveScroll: true });
      },
    });
  }
};

const openEventDialog = (team: "home" | "away") => {
  selectedTeamForEvent.value = team;
  eventForm.reset();
  eventForm.team_id = team === "home" ? match.home_team_id : match.away_team_id;
  showEventDialog.value = true;
};

const submitEvent = () => {
  eventForm.post(`/matches/${match.id}/events`, {
    preserveScroll: true,
    onSuccess: () => {
      showEventDialog.value = false;
      eventForm.reset();
      router.visit(window.location.href, { preserveScroll: true });
    },
  });
};

const deleteEvent = (eventId: number) => {
  if (confirm("¿Eliminar este evento?")) {
    router.delete(`/matches/${match.id}/events/${eventId}`, {
      preserveScroll: true,
      onSuccess: () => {
        router.visit(window.location.href, { preserveScroll: true });
      },
    });
  }
};

const openLineupDialog = (team: "home" | "away") => {
  selectedTeamForLineup.value = team;
  lineupForm.reset();
  lineupForm.team_id = team === "home" ? match.home_team_id : match.away_team_id;
  lineupForm.is_starter = true;
  lineupForm.minutes_played = 90;
  showLineupDialog.value = true;
};

const submitLineup = () => {
  lineupForm.transform((data) => ({
    ...data,
    is_starter: data.is_starter ? "true" : "false",
  })).post(`/matches/${match.id}/lineups`, {
    preserveScroll: true,
    onSuccess: () => {
      showLineupDialog.value = false;
      lineupForm.reset();
      router.visit(window.location.href, { preserveScroll: true });
    },
  });
};

const removeLineup = (playerId: number) => {
  if (confirm("¿Eliminar jugador de la alineación?")) {
    router.delete(`/matches/${match.id}/lineups/${playerId}`, {
      preserveScroll: true,
      onSuccess: () => {
        router.visit(window.location.href, { preserveScroll: true });
      },
    });
  }
};

const getEventIcon = (type: string) => {
  const icons: Record<string, string> = {
    goal: "⚽",
    assist: "🅰️",
    yellow_card: "🟨",
    red_card: "🟥",
    sub_in: "🔼",
    sub_out: "🔽",
  };
  return icons[type] || "📋";
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

const getPhotoUrl = (player: Player): string | undefined => {
  if (player.photo) {
    return `/players/photo/${player.photo.split("/").pop()}`;
  }
  return undefined;
};
</script>

<template>
  <Head><title>Partido - {{ match.home_team?.name }} vs {{ match.away_team?.name }}</title></Head>
  <div>
    <Link href="/matches" class="text-decoration-none">
      <v-btn variant="text" prepend-icon="mdi-arrow-left" class="mb-4">Volver a Partidos</v-btn>
    </Link>

    <!-- Match Header -->
    <v-card class="mb-6">
      <v-card-text class="text-center pa-8">
        <div class="text-caption text-medium-emphasis mb-2">
          <Link :href="`/tournaments/${match.tournament?.id}`" class="text-primary">
            {{ match.tournament?.name }}
          </Link>
          <span v-if="match.round"> — Jornada {{ match.round }}</span>
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
            <div v-else class="text-h3 font-weight-bold text-medium-emphasis">
              {{ match.home_score }} - {{ match.away_score }}
            </div>
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

      <v-card-actions class="justify-center pb-4 flex-wrap ga-2">
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
        <v-btn
          v-if="match.status !== 'completed'"
          color="warning"
          prepend-icon="mdi-flag-checkered"
          @click="closeMatch"
        >
          Cerrar Partido
        </v-btn>
      </v-card-actions>
    </v-card>

    <!-- Lineups Section -->
    <v-row class="mb-6">
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title class="d-flex align-center justify-space-between">
            <span>{{ match.home_team?.name }} - Alineación</span>
            <v-btn
              v-if="match.status !== 'completed'"
              color="primary"
              size="small"
              prepend-icon="mdi-plus"
              @click="openLineupDialog('home')"
            >
              Agregar
            </v-btn>
          </v-card-title>
          <v-divider />
          <v-list v-if="homeLineups.length > 0" density="compact">
            <v-list-item v-for="l in homeLineups" :key="l.id">
              <template #prepend>
                <v-avatar size="32" color="primary">
                  <v-img v-if="l.player && getPhotoUrl(l.player)" :src="getPhotoUrl(l.player)" cover />
                  <span v-else class="text-white text-caption">{{ l.player?.name?.charAt(0) }}</span>
                </v-avatar>
              </template>
              <v-list-item-title>
                <Link :href="`/players/${l.player_id}`" class="text-primary">
                  {{ l.player?.nickname || l.player?.name }}
                </Link>
                <v-chip v-if="l.is_starter" size="x-small" color="green" class="ml-2">Titular</v-chip>
              </v-list-item-title>
              <v-list-item-subtitle>{{ l.minutes_played }}' minutos</v-list-item-subtitle>
              <template #append>
                <v-btn
                  v-if="match.status !== 'completed'"
                  icon="mdi-delete"
                  size="x-small"
                  variant="text"
                  color="error"
                  @click="removeLineup(l.player_id)"
                />
              </template>
            </v-list-item>
          </v-list>
          <v-card-text v-else class="text-center text-medium-emphasis">
            Sin alineación registrada
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" md="6">
        <v-card>
          <v-card-title class="d-flex align-center justify-space-between">
            <span>{{ match.away_team?.name }} - Alineación</span>
            <v-btn
              v-if="match.status !== 'completed'"
              color="primary"
              size="small"
              prepend-icon="mdi-plus"
              @click="openLineupDialog('away')"
            >
              Agregar
            </v-btn>
          </v-card-title>
          <v-divider />
          <v-list v-if="awayLineups.length > 0" density="compact">
            <v-list-item v-for="l in awayLineups" :key="l.id">
              <template #prepend>
                <v-avatar size="32" color="secondary">
                  <v-img v-if="l.player && getPhotoUrl(l.player)" :src="getPhotoUrl(l.player)" cover />
                  <span v-else class="text-white text-caption">{{ l.player?.name?.charAt(0) }}</span>
                </v-avatar>
              </template>
              <v-list-item-title>
                <Link :href="`/players/${l.player_id}`" class="text-primary">
                  {{ l.player?.nickname || l.player?.name }}
                </Link>
                <v-chip v-if="l.is_starter" size="x-small" color="green" class="ml-2">Titular</v-chip>
              </v-list-item-title>
              <v-list-item-subtitle>{{ l.minutes_played }}' minutos</v-list-item-subtitle>
              <template #append>
                <v-btn
                  v-if="match.status !== 'completed'"
                  icon="mdi-delete"
                  size="x-small"
                  variant="text"
                  color="error"
                  @click="removeLineup(l.player_id)"
                />
              </template>
            </v-list-item>
          </v-list>
          <v-card-text v-else class="text-center text-medium-emphasis">
            Sin alineación registrada
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Events Section -->
    <v-row>
      <v-col cols="12" md="6">
        <v-card>
          <v-card-title class="d-flex align-center justify-space-between">
            <span>{{ match.home_team?.name }} - Eventos</span>
            <v-btn
              v-if="match.status !== 'completed'"
              color="primary"
              size="small"
              prepend-icon="mdi-plus"
              @click="openEventDialog('home')"
            >
              Agregar
            </v-btn>
          </v-card-title>
          <v-divider />
          <v-list v-if="homeEvents.length > 0" density="compact">
            <v-list-item v-for="e in homeEvents" :key="e.id">
              <template #prepend>
                <v-avatar size="32" color="grey-lighten-3">
                  <span class="text-caption font-weight-bold">{{ e.minute }}'</span>
                </v-avatar>
              </template>
              <v-list-item-title>
                {{ getEventIcon(e.event_type) }}
                <Link :href="`/players/${e.player_id}`" class="text-primary ml-1">
                  {{ e.player?.nickname || e.player?.name }}
                </Link>
              </v-list-item-title>
              <template #append>
                <v-btn
                  v-if="match.status !== 'completed'"
                  icon="mdi-delete"
                  size="x-small"
                  variant="text"
                  color="error"
                  @click="deleteEvent(e.id)"
                />
              </template>
            </v-list-item>
          </v-list>
          <v-card-text v-else class="text-center text-medium-emphasis">
            Sin eventos registrados
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" md="6">
        <v-card>
          <v-card-title class="d-flex align-center justify-space-between">
            <span>{{ match.away_team?.name }} - Eventos</span>
            <v-btn
              v-if="match.status !== 'completed'"
              color="primary"
              size="small"
              prepend-icon="mdi-plus"
              @click="openEventDialog('away')"
            >
              Agregar
            </v-btn>
          </v-card-title>
          <v-divider />
          <v-list v-if="awayEvents.length > 0" density="compact">
            <v-list-item v-for="e in awayEvents" :key="e.id">
              <template #prepend>
                <v-avatar size="32" color="grey-lighten-3">
                  <span class="text-caption font-weight-bold">{{ e.minute }}'</span>
                </v-avatar>
              </template>
              <v-list-item-title>
                {{ getEventIcon(e.event_type) }}
                <Link :href="`/players/${e.player_id}`" class="text-primary ml-1">
                  {{ e.player?.nickname || e.player?.name }}
                </Link>
              </v-list-item-title>
              <template #append>
                <v-btn
                  v-if="match.status !== 'completed'"
                  icon="mdi-delete"
                  size="x-small"
                  variant="text"
                  color="error"
                  @click="deleteEvent(e.id)"
                />
              </template>
            </v-list-item>
          </v-list>
          <v-card-text v-else class="text-center text-medium-emphasis">
            Sin eventos registrados
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>

    <!-- Dialog: Record Result -->
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
          <p class="text-caption text-medium-emphasis mt-4">
            Esto marcará el partido como finalizado y actualizará las clasificaciones.
          </p>
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

    <!-- Dialog: Add Event -->
    <v-dialog v-model="showEventDialog" max-width="400">
      <v-card>
        <v-card-title>
          Agregar Evento - {{ selectedTeamForEvent === "home" ? match.home_team?.name : match.away_team?.name }}
        </v-card-title>
        <v-card-text>
          <v-select
            v-model="eventForm.player_id"
            :items="playersForEvent"
            item-title="name"
            item-value="id"
            label="Jugador *"
            variant="outlined"
            class="mb-4"
          />
          <v-select
            v-model="eventForm.event_type"
            :items="eventTypes"
            label="Tipo de Evento *"
            variant="outlined"
            class="mb-4"
          />
          <v-text-field
            v-model.number="eventForm.minute"
            type="number"
            min="0"
            max="120"
            label="Minuto"
            variant="outlined"
          />
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="showEventDialog = false">Cancelar</v-btn>
          <v-btn
            color="primary"
            :loading="eventForm.processing"
            :disabled="!eventForm.player_id || !eventForm.event_type"
            @click="submitEvent"
          >
            Agregar Evento
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- Dialog: Add Lineup -->
    <v-dialog v-model="showLineupDialog" max-width="400">
      <v-card>
        <v-card-title>
          Agregar a Alineación - {{ selectedTeamForLineup === "home" ? match.home_team?.name : match.away_team?.name }}
        </v-card-title>
        <v-card-text>
          <v-select
            v-model="lineupForm.player_id"
            :items="playersForLineup"
            item-title="name"
            item-value="id"
            label="Jugador *"
            variant="outlined"
            class="mb-4"
          />
          <v-checkbox v-model="lineupForm.is_starter" label="Titular" color="primary" class="mb-2" />
          <v-text-field
            v-model.number="lineupForm.minutes_played"
            type="number"
            min="0"
            max="120"
            label="Minutos Jugados"
            variant="outlined"
          />
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn variant="text" @click="showLineupDialog = false">Cancelar</v-btn>
          <v-btn
            color="primary"
            :loading="lineupForm.processing"
            :disabled="!lineupForm.player_id"
            @click="submitLineup"
          >
            Agregar Jugador
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>
