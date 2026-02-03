<script setup lang="ts">
import { ref } from "vue";
import { useForm, Head, Link } from "@inertiajs/vue3";

const form = useForm({
  name: "",
  email: "",
  password: "",
  password_confirmation: "",
});
const showPassword = ref(false);

const submit = () => {
  form.post("/auth/register");
};
</script>

<template>
  <Head><title>Registrarse</title></Head>
  <v-card elevation="4">
    <v-card-title class="text-center pt-6">
      <Link href="/" class="text-decoration-none">
        <div class="text-h4 mb-1">⚽</div>
      </Link>
      <div class="text-h5 font-weight-bold">Crear Cuenta</div>
    </v-card-title>
    <v-card-text class="pa-6">
      <v-form @submit.prevent="submit">
        <v-text-field
          v-model="form.name"
          label="Nombre *"
          prepend-inner-icon="mdi-account"
          :error-messages="form.errors.name"
          class="mb-4"
          required
        />
        <v-text-field
          v-model="form.email"
          label="Correo Electrónico *"
          type="email"
          prepend-inner-icon="mdi-email"
          :error-messages="form.errors.email"
          class="mb-4"
          required
        />
        <v-text-field
          v-model="form.password"
          label="Contraseña *"
          :type="showPassword ? 'text' : 'password'"
          prepend-inner-icon="mdi-lock"
          :append-inner-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
          @click:append-inner="showPassword = !showPassword"
          :error-messages="form.errors.password"
          class="mb-4"
          required
        />
        <v-text-field
          v-model="form.password_confirmation"
          label="Confirmar Contraseña *"
          :type="showPassword ? 'text' : 'password'"
          prepend-inner-icon="mdi-lock-check"
          :error-messages="form.errors.password_confirmation"
          class="mb-4"
          required
        />
        <v-btn type="submit" color="primary" size="large" block :loading="form.processing">
          Registrarse
        </v-btn>
      </v-form>
    </v-card-text>
    <v-card-actions class="justify-center pb-6">
      <span class="text-body-2">¿Ya tienes cuenta?</span>
      <Link href="/auth/login" class="text-decoration-none">
        <v-btn variant="text" color="primary" size="small">Iniciar Sesión</v-btn>
      </Link>
    </v-card-actions>
  </v-card>
</template>
