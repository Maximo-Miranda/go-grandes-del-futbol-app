<script setup lang="ts">
import { computed } from "vue";
import { usePage } from "@inertiajs/vue3";

interface Flash { success?: string; error?: string; warning?: string }
const page = usePage();
const flash = computed(() => page.props.flash as Flash | undefined);
</script>

<template>
  <v-app>
    <v-app-bar elevation="2" color="primary" density="comfortable">
      <v-container class="d-flex justify-center align-center py-0">
        <a href="/" class="text-decoration-none">
          <div class="d-flex align-center">
            <span class="text-h5 mr-2">⚽</span>
            <span class="text-h6 font-weight-bold text-white">Grandes del Fútbol</span>
          </div>
        </a>
      </v-container>
    </v-app-bar>

    <v-main class="bg-grey-lighten-4">
      <v-container class="fill-height py-6" fluid>
        <v-row align="center" justify="center" class="fill-height">
          <v-col cols="12" sm="10" md="6" lg="4" xl="3" class="px-4">
            <v-slide-y-transition>
              <v-alert v-if="flash?.success" type="success" variant="tonal" class="mb-4" density="compact" closable>
                {{ flash.success }}
              </v-alert>
            </v-slide-y-transition>
            <v-slide-y-transition>
              <v-alert v-if="flash?.error" type="error" variant="tonal" class="mb-4" density="compact" closable>
                {{ flash.error }}
              </v-alert>
            </v-slide-y-transition>
            <slot />
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <v-footer color="transparent" class="text-center text-caption text-medium-emphasis pa-4">
      <v-row justify="center" no-gutters>
        <v-col cols="12">&copy; {{ new Date().getFullYear() }} Grandes del Fútbol</v-col>
      </v-row>
    </v-footer>
  </v-app>
</template>
