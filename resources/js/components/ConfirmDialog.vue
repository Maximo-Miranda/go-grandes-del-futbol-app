<script setup lang="ts">
defineProps<{
  modelValue: boolean
  title?: string
  message?: string
  confirmText?: string
  cancelText?: string
  loading?: boolean
}>()

defineEmits<{
  'update:modelValue': [value: boolean]
  confirm: []
  cancel: []
}>()
</script>

<template>
  <v-dialog :model-value="modelValue" max-width="400" @update:model-value="$emit('update:modelValue', $event)">
    <v-card>
      <v-card-title class="text-h6">{{ title || '¿Estás seguro?' }}</v-card-title>
      <v-card-text>{{ message || 'Esta acción no se puede deshacer.' }}</v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn variant="text" @click="$emit('cancel')">{{ cancelText || 'Cancelar' }}</v-btn>
        <v-btn color="error" :loading="loading" @click="$emit('confirm')">{{ confirmText || 'Eliminar' }}</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
