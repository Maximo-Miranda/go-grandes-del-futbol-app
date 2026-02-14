<script setup lang="ts">
import { ref } from "vue";
import { Head, useForm, Link, router } from "@inertiajs/vue3";

defineProps<{ errors?: Record<string, string> }>();

const form = useForm({
    name: "",
    color: "#4CAF50",
    contact_phone: "",
    logo: null as File | null,
});

const logoPreview = ref<string | null>(null);

const handleLogoChange = (e: Event) => {
    const target = e.target as HTMLInputElement;
    const file = target.files?.[0];
    if (file) {
        form.logo = file;
        const reader = new FileReader();
        reader.onload = (e) => {
            logoPreview.value = e.target?.result as string;
        };
        reader.readAsDataURL(file);
    }
};

const removeLogo = () => {
    form.logo = null;
    logoPreview.value = null;
};

const submit = () => {
    form.post("/teams", {
        forceFormData: true,
        preserveScroll: true,
        onSuccess: () => {
            router.visit("/teams");
        },
    });
};
</script>

<template>
    <Head><title>Crear Equipo</title></Head>
    <div>
        <h1 class="text-h4 font-weight-bold mb-6">Crear Equipo</h1>
        <v-card>
            <v-card-text class="pa-6">
                <v-form @submit.prevent="submit">
                    <!-- Logo Upload -->
                    <div class="mb-6">
                        <label class="text-subtitle-2 d-block mb-2"
                            >Logo del Equipo</label
                        >
                        <div class="d-flex align-center ga-4">
                            <v-avatar
                                size="100"
                                :color="form.color || 'primary'"
                            >
                                <v-img
                                    v-if="logoPreview"
                                    :src="logoPreview"
                                    cover
                                />
                                <v-icon v-else size="48" color="white"
                                    >mdi-shield-half-full</v-icon
                                >
                            </v-avatar>
                            <div>
                                <v-btn
                                    variant="outlined"
                                    prepend-icon="mdi-camera"
                                    @click="
                                        (
                                            $refs.logoInput as HTMLInputElement
                                        ).click()
                                    "
                                >
                                    {{
                                        logoPreview
                                            ? "Cambiar Logo"
                                            : "Subir Logo"
                                    }}
                                </v-btn>
                                <v-btn
                                    v-if="logoPreview"
                                    variant="text"
                                    color="error"
                                    icon="mdi-delete"
                                    @click="removeLogo"
                                    class="ml-2"
                                />
                                <input
                                    ref="logoInput"
                                    type="file"
                                    accept="image/jpeg,image/png,image/webp"
                                    style="display: none"
                                    @change="handleLogoChange"
                                />
                                <p
                                    class="text-caption text-medium-emphasis mt-2"
                                >
                                    Formatos: JPG, PNG, WebP. Max 1MB.
                                </p>
                                <p
                                    v-if="form.errors.logo"
                                    class="text-caption text-error mt-1"
                                >
                                    {{ form.errors.logo }}
                                </p>
                            </div>
                        </div>
                    </div>

                    <v-text-field
                        v-model="form.name"
                        label="Nombre del Equipo *"
                        :error-messages="form.errors.name"
                        class="mb-4"
                        required
                    />
                    <v-row>
                        <v-col cols="12" md="6">
                            <v-text-field
                                v-model="form.color"
                                label="Color del Equipo *"
                                type="color"
                                :error-messages="form.errors.color"
                                required
                            />
                        </v-col>
                        <v-col cols="12" md="6">
                            <v-text-field
                                v-model="form.contact_phone"
                                label="Teléfono de Contacto *"
                                :error-messages="form.errors.contact_phone"
                                required
                            />
                        </v-col>
                    </v-row>

                    <div class="d-flex ga-4">
                        <v-btn
                            type="submit"
                            color="primary"
                            :loading="form.processing"
                            :disabled="form.processing"
                            >Crear Equipo</v-btn
                        >
                        <Link href="/teams" class="text-decoration-none">
                            <v-btn variant="outlined">Cancelar</v-btn>
                        </Link>
                    </div>
                </v-form>
            </v-card-text>
        </v-card>
    </div>
</template>
