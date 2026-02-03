<script setup lang="ts">
import { Head, useForm } from "@inertiajs/vue3";
import PageHeader from "@/components/PageHeader.vue";
import { fieldErrors } from "@/utils/errors";

const props = defineProps<{ team: any; errors?: Record<string, string | Record<string, string>> }>();

const form = useForm({ name: props.team.name, color: props.team.color || "#2E7D32", contact_phone: props.team.contact_phone || "" });
const submit = () => form.put(`/teams/${props.team.id}`);
</script>

<template>
    <Head title="Editar Equipo" />
    <PageHeader :title="`Editar: ${team.name}`" />
    <v-card class="pa-6">
        <v-form @submit.prevent="submit">
            <v-text-field v-model="form.name" label="Nombre del equipo" :error-messages="fieldErrors(errors?.name)" class="mb-2" />
            <v-text-field v-model="form.color" label="Color" type="color" class="mb-2" />
            <v-text-field v-model="form.contact_phone" label="Teléfono de contacto" class="mb-4" />
            <div class="d-flex ga-2">
                <v-btn type="submit" color="primary" :loading="form.processing" class="text-none">Guardar</v-btn>
                <v-btn :href="`/teams/${team.id}`" variant="text" class="text-none">Cancelar</v-btn>
            </div>
        </v-form>
    </v-card>
</template>
