<script lang="ts" setup>
import { Button, Dialog, Divider, Toolbar } from "primevue"
import { useCssModule, ref } from "vue"
import { useFiles } from "../composables/useFiles"
import FilePicker from "./FilePicker.vue"
import WindowControls from "./WindowControls.vue"


// Use files composable
const { addFiles, clearSelection, hasSelection, selectedCount, removeSelectedFiles } = useFiles()
// const styles = useCssModule()

const metaSearchVisible = ref(false)
const templateEditorVisible = ref(false)

const showMetaSearch = () => {
  metaSearchVisible.value = true
}

const showTemplateEditor = () => {
  templateEditorVisible.value = true
}

const refresh = () => {
  // TODO: Implement refresh logic
  console.log("Refresh clicked")
}

const showSettings = () => {
  // TODO: Implement settings logic
  console.log("Settings clicked")
}

const showHelp = () => {
  // TODO: Implement help logic
  console.log("Help clicked")
}
</script>

<template>
  <Toolbar class="app-header" style="--wails-draggable:drag">
    <template #start>
      <div class="toolbar-start">
        <svg viewBox="0 0 512 512" fill="none" xmlns="http://www.w3.org/2000/svg" style="width: 1.5em; height: 1.5em;">
          <path d="M142 122H356C356 349 199 458 142 389C85.7484 320.906 214 188 443 231" stroke-width="40"
            stroke-linecap="round" style="stroke: light-dark(black, white);" />
        </svg>&nbsp;
        <FilePicker :accept="{ 'Video Files': '*.mp4;*.mkv;*.avi;*.mov;*.wmv;*.flv;*.webm' }" @select="addFiles">
          <Button icon="pi pi-plus" label="Add Files" severity="secondary" size="small" class="mr-2" />
        </FilePicker>
        <div class="button-group">
          <Button icon="pi pi-ban" severity="secondary" size="small" class="mr-2 contrast-button"
            :disabled="!hasSelection" @click="clearSelection" />
          <Button :label="`${selectedCount}`" severity="secondary" size="small" style="min-width: 3em;" />
          <Button icon="pi pi-trash" severity="secondary" size="small" class="mr-2 danger-button"
            :disabled="!hasSelection" @click="removeSelectedFiles" />
        </div>
        <Divider layout="vertical" />
        <Button icon="pi pi-search" label="Fetch Metadata" @click="showMetaSearch" severity="secondary" size="small"
          class="mr-2" />
        <Button icon="pi pi-file-edit" label="Edit Template" @click="showTemplateEditor" severity="secondary"
          size="small" class="mr-2" />
        <Divider layout="vertical" />
        <Button icon="pi pi-play" label="Rename" @click="refresh" size="small" />
      </div>
    </template>

    <template #end>
      <div class="toolbar-end">
        <Button icon="pi pi-cog" label="Settings" @click="showSettings" severity="secondary" text size="small"
          class="mr-2" />
        <Button icon="pi pi-question-circle" label="Help" @click="showHelp" severity="secondary" text size="small" />
        <WindowControls />
      </div>
    </template>
  </Toolbar>

  <!-- MetaSearch Modal -->
  <Dialog v-model:visible="metaSearchVisible" modal header="Fetch Metadata" :style="{ width: '50rem' }"
    :breakpoints="{ '1199px': '75vw', '575px': '90vw' }">
    <div class="modal-placeholder">
      <p>MetaSearch component will be implemented here.</p>
      <p>This will allow searching and fetching metadata from TMDB for your media files.</p>
    </div>
  </Dialog>

  <!-- Template Editor Modal -->
  <Dialog v-model:visible="templateEditorVisible" modal header="Edit Rename Template" :style="{ width: '50rem' }"
    :breakpoints="{ '1199px': '75vw', '575px': '90vw' }">
    <div class="modal-placeholder">
      <p>FsTemplateEditor component will be implemented here.</p>
      <p>This will allow editing filename templates like:</p>
      <code>[title] ([year]) - S[##]E[##]</code>
    </div>
  </Dialog>
</template>

<style scoped>
.app-header {
  border: none;
  border-bottom: 1px solid var(--p-toolbar-border-color);
  border-bottom: 1px solid color-mix(in srgb, var(--p-toolbar-border-color), transparent 25%);
  border-radius: 0;
  background-color: color-mix(in srgb, light-dark(var(--p-surface-50), var(--p-surface-800)), transparent 60%);
}

.app-header :deep(.p-toolbar) {
  border-radius: 0;
  border: none;
  border-bottom: 1px solid var(--p-toolbar-border-color);
}

.toolbar-start {
  display: flex;
  gap: 0.5rem;
  align-items: center;
}

/* .remove-button {
  background: color-mix(in srgb, var(--p-button-danger-background), transparent 30%);
  border-color: transparent;
} */

.select-counter {
  font-family: monospace;
  font-size: inherit;
  font-weight: 500;
  width: 1.25em;
  height: 1.25em;
  padding: 0;
  margin: 0;
}

.button-group {
  display: flex;
  flex-flow: row nowrap;
}

.button-group > :first-child {
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
}

.button-group> :not(:first-child):not(:last-child) {
  border-radius: 0;
}

.button-group > :last-child {
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
}

.danger-button {
  /* background: color-mix(in srgb, var(--p-button-danger-background), transparent 70%); */
  color: var(--p-button-danger-background);
}

.danger-button:not(:disabled):hover {
  /* background: color-mix(in srgb, var(--p-button-danger-background), transparent 50%); */
  color: var(--p-button-danger-background);
  /* border-color: transparent; */
}

.contrast-button {
  /* color: var(--p-button-danger-color); */
  color: var(--p-button-contrast-background);
}

.toolbar-end {
  display: flex;
  gap: 0.5rem;
  align-items: center;
}

.modal-placeholder {
  padding: 1rem 0;
}

.modal-placeholder code {
  background-color: var(--p-surface-100);
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  font-family: monospace;
}
</style>
