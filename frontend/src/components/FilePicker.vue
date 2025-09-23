<script lang="ts" setup>
import { computed } from "vue";
import { SelectFiles } from "../../wailsjs/go/main/App"
import type { main, frontend } from "../../wailsjs/go/models"


const { accept = { "All Files": "*.*" }, multiple = false } = defineProps<{
  title?: string
  accept?: Record<string, string>
  multiple?: boolean
}>()

const emit = defineEmits<{
  select: [files: string[]]
}>()

const filters = computed(() => {
  return Object.entries(accept).map(([k, v]) => ({ DisplayName: k, Pattern: v })) as frontend.FileFilter[]
})

const selectFiles = async () => {
  try {
    const files = await SelectFiles({
      Title: "Select Media Files",
      Filters: filters.value,
    } as main.FileDialogOptions)

    emit("select", files)
  } catch (error) {
    console.error('Error selecting files:', error)
  }
}
</script>
<template>
  <div class="file-picker" @click="selectFiles">
    <slot></slot>
  </div>
</template>
<style lang="css" scoped>
.file-picker {
  display: contents;
}
</style>
