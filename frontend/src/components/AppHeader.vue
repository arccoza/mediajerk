<script lang="ts" setup>
import { ref } from "vue"
import Button from "primevue/button"
import Toolbar from "primevue/toolbar"
import Dialog from "primevue/dialog"

const metaSearchVisible = ref(false)
const templateEditorVisible = ref(false)

const showMetaSearch = () => {
  metaSearchVisible.value = true
}

const showTemplateEditor = () => {
  templateEditorVisible.value = true
}

const addFiles = () => {
  // TODO: Implement file selection logic
  console.log("Add files clicked")
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
        <Button
          icon="pi pi-plus"
          label="Add Files"
          @click="addFiles"
          severity="secondary"
          size="small"
          class="mr-2"
        />
        <Button
          icon="pi pi-search"
          label="Fetch Metadata"
          @click="showMetaSearch"
          severity="secondary"
          size="small"
          class="mr-2"
        />
        <Button
          icon="pi pi-file-edit"
          label="Edit Template"
          @click="showTemplateEditor"
          severity="secondary"
          size="small"
          class="mr-2"
        />
        <Button
          icon="pi pi-play"
          label="Rename"
          @click="refresh"
          size="small"
        />
      </div>
    </template>

    <template #end>
      <div class="toolbar-end">
        <Button
          icon="pi pi-cog"
          label="Settings"
          @click="showSettings"
          severity="secondary"
          text
          size="small"
          class="mr-2"
        />
        <Button
          icon="pi pi-question-circle"
          label="Help"
          @click="showHelp"
          severity="secondary"
          text
          size="small"
        />
      </div>
    </template>
  </Toolbar>

  <!-- MetaSearch Modal -->
  <Dialog
    v-model:visible="metaSearchVisible"
    modal
    header="Fetch Metadata"
    :style="{ width: '50rem' }"
    :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
  >
    <div class="modal-placeholder">
      <p>MetaSearch component will be implemented here.</p>
      <p>This will allow searching and fetching metadata from TMDB for your media files.</p>
    </div>
  </Dialog>

  <!-- Template Editor Modal -->
  <Dialog
    v-model:visible="templateEditorVisible"
    modal
    header="Edit Rename Template"
    :style="{ width: '50rem' }"
    :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
  >
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
  border-radius: 0;
}

.app-header :deep(.p-toolbar) {
  border-radius: 0;
  border: none;
  border-bottom: 1px solid var(--p-toolbar-border-color);
}

.toolbar-start {
  display: flex;
  gap: 0.5rem;
}

.toolbar-end {
  display: flex;
  gap: 0.5rem;
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
