<script setup lang="ts">
defineProps<{
  title: string
  subtitle?: string
}>()
</script>

<template>
  <div class="d-flex flex-column flex-sm-row justify-space-between align-start align-sm-center mb-6 ga-4">
    <div>
      <h1 class="text-h4 font-weight-medium">{{ title }}</h1>
      <p v-if="subtitle" class="text-body-2 text-medium-emphasis mt-1">
        {{ subtitle }}
      </p>
    </div>
    <div class="d-flex ga-2">
      <slot name="actions" />
    </div>
  </div>
</template>
