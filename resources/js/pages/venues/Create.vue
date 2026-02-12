<script setup lang="ts">
import { ref } from "vue";
import { Head, useForm, Link, router } from "@inertiajs/vue3";

const form = useForm({
    name: "",
    address: "",
    photo: null as File | null,
});

const photoPreview = ref<string | null>(null);

const MAX_FILE_SIZE = 1048576; // 1 MB
const ALLOWED_TYPES = ["image/jpeg", "image/png", "image/webp"];

const handlePhotoChange = (e: Event) => {
    const target = e.target as HTMLInputElement;
    const file = target.files?.[0];
    if (file) {
        if (!ALLOWED_TYPES.includes(file.type)) {
            form.setError(
                "photo" as any,
                "La foto debe ser una imagen válida (jpg, jpeg, png, webp)",
            );
            target.value = "";
            return;
        }
        if (file.size > MAX_FILE_SIZE) {
            form.setError("photo" as any, "La foto no puede exceder 1 MB");
            target.value = "";
            return;
        }
        form.clearErrors("photo" as any);
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
    form.post("/venues", {
        forceFormData: true,
        preserveScroll: true,
        onSuccess: () => {
            router.visit("/venues");
        },
    });
};
</script>

<template>
    <Head><title>Nueva Sede</title></Head>
    <div>
        <h1 class="text-h4 font-weight-bold mb-6">Nueva Sede</h1>
        <v-card>
            <v-card-text class="pa-6">
                <v-form @submit.prevent="submit">
                    <!-- Photo Upload -->
                    <div class="mb-6">
                        <label class="text-subtitle-2 d-block mb-2"
                            >Foto de la Sede</label
                        >
                        <div class="d-flex align-center ga-4">
                            <v-avatar
                                size="100"
                                color="grey-lighten-3"
                                rounded="lg"
                            >
                                <v-img
                                    v-if="photoPreview"
                                    :src="photoPreview"
                                    cover
                                />
                                <v-icon v-else size="48" color="grey-lighten-1"
                                    >mdi-stadium</v-icon
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
                                    Formatos: JPG, PNG, WebP. Max 1MB.
                                </p>
                                <p
                                    v-if="form.errors.photo"
                                    class="text-caption text-error mt-1"
                                >
                                    {{ (form.errors as any).photo }}
                                </p>
                            </div>
                        </div>
                    </div>

                    <v-text-field
                        v-model="form.name"
                        label="Nombre de la Sede *"
                        :error-messages="form.errors.name"
                        class="mb-4"
                        required
                    />
                    <v-text-field
                        v-model="form.address"
                        label="Dirección *"
                        :error-messages="form.errors.address"
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
                            Crear Sede
                        </v-btn>
                        <Link href="/venues" class="text-decoration-none">
                            <v-btn
                                variant="outlined"
                                :disabled="form.processing"
                                >Cancelar</v-btn
                            >
                        </Link>
                    </div>
                </v-form>
            </v-card-text>
        </v-card>
    </div>
</template>
