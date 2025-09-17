<script lang="ts" setup>
import { ref, computed } from "vue"
import DataTable from "primevue/datatable"
import Column from "primevue/column"
import Tag from "primevue/tag"

interface PreviewItem {
  id: number
  originalFilename: string
  newFilename: string
  status: "ready" | "warning" | "error"
  message?: string
}

const template = ref("[title] ([year]) - S[##]E[##]")

const previewFiles = ref<PreviewItem[]>([
  {
    id: 1,
    originalFilename: "Game.of.Thrones.S01E01.Winter.is.Coming.mkv",
    newFilename: "Game of Thrones (2011) - S01E01.mkv",
    status: "ready",
  },
  {
    id: 2,
    originalFilename: "Breaking.Bad.S01E01.Pilot.mp4",
    newFilename: "Breaking Bad (2008) - S01E01.mp4",
    status: "warning",
    message: "Metadata not found",
  },
])

const getStatusSeverity = (status: string) => {
  switch (status) {
    case "ready":
      return "success"
    case "warning":
      return "warning"
    case "error":
      return "danger"
    default:
      return "info"
  }
}
</script>

<template>
  <div class="file-preview">
    <div class="file-preview-header">
      <h3>Rename Preview</h3>
    </div>

    <DataTable
      :value="previewFiles"
      :scrollable="true"
      scroll-height="100%"
      class="preview-table"
      :pt="{ table: { style: 'min-width: 100%' } }"
    >
      <Column field="originalFilename" header="Original" style="width: 40%">
        <template #body="{ data }">
          <span class="original-filename">{{ data.originalFilename }}</span>
        </template>
      </Column>

      <Column field="newFilename" header="New Filename" style="width: 40%">
        <template #body="{ data }">
          <div class="new-filename-cell">
            <span class="new-filename">{{ data.newFilename }}</span>
            <small v-if="data.message" class="rename-message">{{ data.message }}</small>
          </div>
        </template>
      </Column>

      <Column field="status" header="Status" style="width: 20%">
        <template #body="{ data }">
          <Tag
            :value="data.status"
            :severity="getStatusSeverity(data.status)"
            class="status-tag"
          />
        </template>
      </Column>
    </DataTable>
  </div>
</template>

<style scoped>
.file-preview {
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 1em;
}

.file-preview-header {
  padding: 0.5rem;
  border-bottom: 1px solid var(--p-surface-border);
  background-color: light-dark(var(--p-surface-100), var(--p-surface-800));
}

.file-preview-header h3 {
  margin: 0;
  font-size: 1rem;
  font-weight: bold;
}

.template-info code {
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  font-family: monospace;
  font-size: 0.8rem;
}

.preview-table {
  flex: 1;
  border: none;
}

.original-filename {
  color: var(--p-text-muted-color);
  font-size: 0.9rem;
}

.new-filename-cell {
  display: flex;
  flex-direction: column;
}

.new-filename {
  font-weight: 500;
  margin-bottom: 0.25rem;
}

.rename-message {
  color: var(--p-orange-500);
  font-size: 0.8rem;
}

.status-tag {
  font-size: 0.75rem;
  padding: 0.25rem 0.5rem;
}
</style>
