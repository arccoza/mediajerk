<script lang="ts" setup>
import { computed } from "vue";
import { SelectFiles } from "../../wailsjs/go/main/App"
import type { main } from "../../wailsjs/go/models"
export type FileInfo = main.FileInfo

const { accept = { "All Files": "*.*" }, multiple = false } = defineProps<{
  title?: string
  accept?: Record<string, string>
  multiple?: boolean
}>()

const emit = defineEmits<{
  select: [files: main.FileInfo[]]
}>()

const filters = computed(() => {
  return Object.entries(accept).map(([k, v]) => ({ displayName: k, pattern: v })) as main.FileFilter[]
})

const selectFiles = async () => {
  try {
    const files = await SelectFiles({
      title: "Select Media Files",
      filters: filters.value,
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
