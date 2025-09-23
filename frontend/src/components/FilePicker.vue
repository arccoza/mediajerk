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

// let files: File[]
// let input = document.createElement("input")

// input.type = "file"
// input.onchange = (ev: Event) => {
//   // @ts-ignore
//   files = Array.from(input.files)
//   console.log(files)
// }

// // Make input attributes reactive to prop changes
// watchEffect(() => {
//   input.accept = accept
//   input.multiple = multiple
//   console.log(input)
// })
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
