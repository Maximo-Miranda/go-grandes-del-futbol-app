<script setup lang="ts">
import { ref, computed } from "vue";
import { Head, useForm, Link, router } from "@inertiajs/vue3";

const props = defineProps<{ team: any; errors?: Record<string, string> }>();

const form = useForm({
    name: props.team.name,
    color: props.team.color || "#4CAF50",
    contact_phone: props.team.contact_phone || "",
    logo: null as File | null,
    remove_logo: false,
});

const logoPreview = ref<string | null>(null);
const currentLogoUrl = computed(() => props.team.logo_url || null);

const handleLogoChange = (e: Event) => {
    const target = e.target as HTMLInputElement;
    const file = target.files?.[0];
    if (file) {
        form.logo = file;
        form.remove_logo = false;
        const reader = new FileReader();
        reader.onload = (e) => {
            logoPreview.value = e.target?.result as string;
        };
        reader.readAsDataURL(file);
    }
};

const removeLogo = () => {
    form.logo = null;
    form.remove_logo = true;
    logoPreview.value = null;
};

const displayLogo = computed(() => {
    if (form.remove_logo) return null;
    if (logoPreview.value) return logoPreview.value;
    return currentLogoUrl.value;
});

const submit = () => {
    router.post(
        `/teams/${props.team.id}`,
        {
            _method: "PUT",
            name: form.name,
            color: form.color,
            contact_phone: form.contact_phone,
            logo: form.logo,
            remove_logo: form.remove_logo ? "true" : "",
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
    <Head><title>Editar Equipo</title></Head>
    <div>
        <h1 class="text-h4 font-weight-bold mb-6">Editar: {{ team.name }}</h1>
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
                                    v-if="displayLogo"
                                    :src="displayLogo"
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
                                        displayLogo
                                            ? "Cambiar Logo"
                                            : "Subir Logo"
                                    }}
                                </v-btn>
                                <v-btn
                                    v-if="displayLogo"
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
                            >Guardar</v-btn
                        >
                        <Link
                            :href="`/teams/${team.id}`"
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
