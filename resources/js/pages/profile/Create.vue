<script setup lang="ts">
import { ref } from "vue";
import { Head, useForm, usePage, router } from "@inertiajs/vue3";

defineProps<{ errors?: Record<string, string> }>();

const page = usePage();
const userName = (page.props.auth as any)?.user?.name || "";

const form = useForm({
    nickname: "",
    document_id: "",
    phone: "",
    position: "",
    photo: null as File | null,
});

const photoPreview = ref<string | null>(null);

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
        const reader = new FileReader();
        reader.onload = (e) => {
            photoPreview.value = e.target?.result as string;
        };
        reader.readAsDataURL(file);
    }
};

const removePhoto = () => {
    form.photo = null;
    photoPreview.value = null;
};

const submit = () => {
    form.post("/profile", {
        forceFormData: true,
        preserveScroll: true,
        onSuccess: () => {
            router.visit("/dashboard");
        },
    });
};
</script>

<template>
    <Head><title>Crear Perfil Deportivo</title></Head>
    <div>
        <div class="text-center mb-6">
            <v-icon size="64" color="primary" class="mb-4"
                >mdi-account-circle</v-icon
            >
            <h1 class="text-h4 font-weight-bold">
                Completa tu Perfil Deportivo
            </h1>
            <p class="text-medium-emphasis mt-2">
                Hola <strong>{{ userName }}</strong
                >, antes de continuar necesitas completar tu perfil deportivo.
            </p>
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
                                    v-if="photoPreview"
                                    :src="photoPreview"
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
                                        photoPreview
                                            ? "Cambiar Foto"
                                            : "Subir Foto"
                                    }}
                                </v-btn>
                                <v-btn
                                    v-if="photoPreview"
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

                    <v-btn
                        type="submit"
                        color="primary"
                        size="large"
                        block
                        :loading="form.processing"
                        :disabled="form.processing"
                    >
                        Crear Perfil Deportivo
                    </v-btn>
                </v-form>
            </v-card-text>
        </v-card>
    </div>
</template>
