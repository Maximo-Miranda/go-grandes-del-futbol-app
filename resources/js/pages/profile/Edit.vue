<script setup lang="ts">
import { ref, computed, watch } from "vue";
import { Head, useForm, usePage, router } from "@inertiajs/vue3";

const props = defineProps<{ player: any; errors?: Record<string, string> }>();

const page = usePage();
const userName = computed(() => (page.props.auth as any)?.user?.name || "");

const form = useForm({
    nickname: props.player.nickname || "",
    document_id: props.player.document_id || "",
    phone: props.player.phone || "",
    position: props.player.position || "",
    photo: null as File | null,
    remove_photo: false,
});

watch(
    () => props.player,
    (newPlayer) => {
        form.nickname = newPlayer.nickname || "";
        form.document_id = newPlayer.document_id || "";
        form.phone = newPlayer.phone || "";
        form.position = newPlayer.position || "";
        form.photo = null;
        form.remove_photo = false;
        photoPreview.value = null;
    },
    { deep: true },
);

const photoPreview = ref<string | null>(null);
const currentPhotoUrl = computed(() => {
    return props.player.photo_url || null;
});

const positions = [
    { title: "Portero", value: "goalkeeper" },
    { title: "Defensa", value: "defender" },
    { title: "Mediocampista", value: "midfielder" },
    { title: "Delantero", value: "forward" },
];

const handlePhotoChange = (e: Event) => {
    const target = e.target as HTMLInputElement;
    const file = target.files?.[0];
    if (file) {
        form.photo = file;
        form.remove_photo = false;
        const reader = new FileReader();
        reader.onload = (e) => {
            photoPreview.value = e.target?.result as string;
        };
        reader.readAsDataURL(file);
    }
};

const removePhoto = () => {
    form.photo = null;
    form.remove_photo = true;
    photoPreview.value = null;
};

const displayPhoto = computed(() => {
    if (form.remove_photo) return null;
    if (photoPreview.value) return photoPreview.value;
    return currentPhotoUrl.value;
});

const submit = () => {
    router.post(
        "/profile/update",
        {
            _method: "PUT",
            nickname: form.nickname,
            document_id: form.document_id,
            phone: form.phone,
            position: form.position,
            photo: form.photo,
            remove_photo: form.remove_photo ? "true" : "",
        },
        {
            forceFormData: true,
            preserveScroll: true,
            onError: (errors) => {
                Object.keys(errors).forEach((key) => {
                    form.setError(key as any, errors[key]);
                });
            },
            onStart: () => {
                form.processing = true;
            },
            onFinish: () => {
                form.processing = false;
            },
        },
    );
};
</script>

<template>
    <Head><title>Editar Perfil Deportivo</title></Head>
    <div>
        <div class="text-center mb-6">
            <h1 class="text-h4 font-weight-bold">Editar Perfil Deportivo</h1>
            <p class="text-medium-emphasis mt-2">{{ userName }}</p>
        </div>

        <v-card max-width="600" class="mx-auto">
            <v-card-text class="pa-6">
                <v-form @submit.prevent="submit">
                    <!-- Photo Upload -->
                    <div class="mb-6">
                        <label class="text-subtitle-2 d-block mb-2"
                            >Foto de Jugador</label
                        >
                        <div class="d-flex align-center ga-4">
                            <v-avatar size="100" color="grey-lighten-3">
                                <v-img
                                    v-if="displayPhoto"
                                    :src="displayPhoto"
                                    cover
                                />
                                <v-icon v-else size="48" color="grey-lighten-1"
                                    >mdi-account</v-icon
                                >
                            </v-avatar>
                            <div>
                                <v-btn
                                    variant="outlined"
                                    prepend-icon="mdi-camera"
                                    @click="
                                        (
                                            $refs.photoInput as HTMLInputElement
                                        ).click()
                                    "
                                >
                                    {{
                                        displayPhoto
                                            ? "Cambiar Foto"
                                            : "Subir Foto"
                                    }}
                                </v-btn>
                                <v-btn
                                    v-if="displayPhoto"
                                    variant="text"
                                    color="error"
                                    icon="mdi-delete"
                                    @click="removePhoto"
                                    class="ml-2"
                                />
                                <input
                                    ref="photoInput"
                                    type="file"
                                    accept="image/jpeg,image/png,image/webp"
                                    style="display: none"
                                    @change="handlePhotoChange"
                                />
                                <p
                                    class="text-caption text-medium-emphasis mt-2"
                                >
                                    Formatos: JPG, PNG, WebP. Max 5MB.
                                </p>
                            </div>
                        </div>
                    </div>

                    <v-text-field
                        v-model="form.nickname"
                        label="Apodo *"
                        :error-messages="form.errors.nickname"
                        class="mb-4"
                        required
                    />

                    <v-row>
                        <v-col cols="12" md="6">
                            <v-text-field
                                v-model="form.document_id"
                                label="Documento de Identidad *"
                                :error-messages="form.errors.document_id"
                                required
                            />
                        </v-col>
                        <v-col cols="12" md="6">
                            <v-text-field
                                v-model="form.phone"
                                label="Teléfono *"
                                :error-messages="form.errors.phone"
                                required
                            />
                        </v-col>
                    </v-row>

                    <v-select
                        v-model="form.position"
                        :items="positions"
                        label="Posición *"
                        :error-messages="form.errors.position"
                        class="mb-4"
                        required
                    />

                    <div class="d-flex ga-4">
                        <v-btn
                            type="submit"
                            color="primary"
                            :loading="form.processing"
                            :disabled="form.processing"
                        >
                            Guardar Cambios
                        </v-btn>
                        <v-btn
                            variant="outlined"
                            :disabled="form.processing"
                            @click="router.visit('/dashboard')"
                        >
                            Cancelar
                        </v-btn>
                    </div>
                </v-form>
            </v-card-text>
        </v-card>
    </div>
</template>
