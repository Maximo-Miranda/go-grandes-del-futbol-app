<script setup lang="ts">
import { Head, useForm } from "@inertiajs/vue3";
import PageHeader from "@/components/PageHeader.vue";
import { fieldErrors } from "@/utils/errors";

const props = defineProps<{ player: any; errors?: Record<string, string | Record<string, string>> }>();

const form = useForm({
    name: props.player.name,
    nickname: props.player.nickname || "",
    document_id: props.player.document_id || "",
    phone: props.player.phone || "",
    position: props.player.position || "",
});

const positions = ["Portero", "Defensa", "Mediocampista", "Delantero"];
const submit = () => form.put(`/players/${props.player.id}`);
</script>

<template>
    <Head title="Editar Jugador" />
    <PageHeader :title="`Editar: ${player.name}`" />
    <v-card class="pa-6">
        <v-form @submit.prevent="submit">
            <v-text-field v-model="form.name" label="Nombre completo" :error-messages="fieldErrors(errors?.name)" class="mb-2" />
            <v-text-field v-model="form.nickname" label="Apodo" class="mb-2" />
            <v-row>
                <v-col cols="12" md="6"><v-text-field v-model="form.document_id" label="Documento de identidad" /></v-col>
                <v-col cols="12" md="6"><v-text-field v-model="form.phone" label="Teléfono" /></v-col>
            </v-row>
            <v-select v-model="form.position" label="Posición" :items="positions" clearable class="mb-4" />
            <div class="d-flex ga-2">
                <v-btn type="submit" color="primary" :loading="form.processing" class="text-none">Guardar</v-btn>
                <v-btn :href="`/players/${player.id}`" variant="text" class="text-none">Cancelar</v-btn>
            </div>
        </v-form>
    </v-card>
</template>
