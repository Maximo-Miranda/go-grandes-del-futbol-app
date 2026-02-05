<script setup lang="ts">
import { ref, computed } from "vue";
import { Head, useForm, Link, router } from "@inertiajs/vue3";

const props = defineProps<{ player: any; errors?: Record<string, string> }>();

const form = useForm({
  name: props.player.name,
  nickname: props.player.nickname || "",
  document_id: props.player.document_id || "",
  phone: props.player.phone || "",
  position: props.player.position || "",
  photo: null as File | null,
  remove_photo: false,
});

const photoPreview = ref<string | null>(null);
const currentPhotoUrl = computed(() => {
  if (props.player.photo) {
    return `/players/photo/${props.player.photo.split('/').pop()}`;
  }
  return null;
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
  router.post(`/players/${props.player.id}`, {
    _method: 'PUT',
    name: form.name,
    nickname: form.nickname,
    document_id: form.document_id,
    phone: form.phone,
    position: form.position,
    photo: form.photo,
    remove_photo: form.remove_photo ? 'true' : '',
  }, {
    forceFormData: true,
  });
};
</script>

<template>
  <Head><title>Editar Jugador</title></Head>
  <div>
    <h1 class="text-h4 font-weight-bold mb-6">Editar: {{ player.name }}</h1>
    <v-card>
      <v-card-text class="pa-6">
        <v-form @submit.prevent="submit">
          <!-- Photo Upload -->
          <div class="mb-6">
            <label class="text-subtitle-2 d-block mb-2">Foto del Jugador</label>
            <div class="d-flex align-center ga-4">
              <v-avatar size="100" color="grey-lighten-3">
                <v-img v-if="displayPhoto" :src="displayPhoto" cover />
                <v-icon v-else size="48" color="grey-lighten-1">mdi-account</v-icon>
              </v-avatar>
              <div>
                <v-btn 
                  variant="outlined" 
                  prepend-icon="mdi-camera"
                  @click="($refs.photoInput as HTMLInputElement).click()"
                >
                  {{ displayPhoto ? 'Cambiar Foto' : 'Subir Foto' }}
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
                <p class="text-caption text-medium-emphasis mt-2">
                  Formatos: JPG, PNG, WebP. Max 5MB.
                </p>
              </div>
            </div>
          </div>

          <v-text-field
            v-model="form.name"
            label="Nombre Completo *"
            :error-messages="form.errors.name"
            class="mb-4"
            required
          />
          <v-text-field
            v-model="form.nickname"
            label="Apodo"
            :error-messages="form.errors.nickname"
            class="mb-4"
          />
          <v-row>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.document_id"
                label="Documento de Identidad"
                :error-messages="form.errors.document_id"
              />
            </v-col>
            <v-col cols="12" md="6">
              <v-text-field
                v-model="form.phone"
                label="Teléfono"
                :error-messages="form.errors.phone"
              />
            </v-col>
          </v-row>
          <v-select
            v-model="form.position"
            :items="positions"
            label="Posición"
            :error-messages="form.errors.position"
            clearable
            class="mb-4"
          />
          <div class="d-flex ga-4">
            <v-btn type="submit" color="primary" :loading="form.processing">Guardar</v-btn>
            <Link :href="`/players/${player.id}`" class="text-decoration-none">
              <v-btn variant="outlined">Cancelar</v-btn>
            </Link>
          </div>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>
