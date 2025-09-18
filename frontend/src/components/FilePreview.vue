<script lang="ts" setup>
import { ref } from "vue"
import DataTable from "primevue/datatable"
import Column from "primevue/column"
import Tag from "primevue/tag"
import { MediaMetadata, MetadataFormatter } from "../utils/templateProcessor"

interface PreviewItem {
  id: number
  originalFilename: string
  newFilename: string
  metadata: MediaMetadata
  status: "ready" | "warning" | "error"
  message?: string
}

interface Props {
  previewFiles: PreviewItem[]
}

interface Emits {
  (e: "row-reorder", files: PreviewItem[]): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const template = ref("[title] ([year]) - S[##]E[##]")

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

const onRowReorder = (event: any) => {
  emit("row-reorder", event.value)
}
</script>

<template>
  <div class="file-preview">
    <div class="file-preview-header">
      <h3>Rename Preview</h3>
    </div>

    <DataTable
      :value="props.previewFiles"
      :scrollable="true"
      scroll-height="100%"
      reorderableRows
      @row-reorder="onRowReorder"
      class="preview-table"
      :pt="{
        table: { style: 'min-width: 100%' },
        bodyRow: { style: 'height: 4rem' }
      }"
    >
      <Column :rowReorder="true" header="" style="width: 3rem" />

      <!-- <Column field="originalFilename" header="Original" style="width: 40%">
        <template #body="{ data }">
          <div class="original-filename-cell">
            <span class="original-filename">{{ data.originalFilename }}</span>
          </div>
        </template>
      </Column> -->

      <Column field="newFilename" header="New Filename" style="min-width: 200px">
        <template #body="{ data }">
          <div class="new-filename-cell">
            <span class="new-filename">{{ data.newFilename }}</span>
            <small class="metadata-display">{{ MetadataFormatter.formatMetadataDisplay(data.metadata) }}</small>
          </div>
        </template>
      </Column>

      <Column field="status" header="Status" style="width: 120px">
        <template #body="{ data }">
          <div class="status-cell">
            <Tag
              :value="data.status"
              :severity="getStatusSeverity(data.status)"
              class="status-tag"
            />
          </div>
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
  font-size: 0.9em;
}

.file-preview-header {
  padding: 0.75rem;
  border-bottom: 1px solid var(--p-surface-border);
  --bg: light-dark(var(--p-surface-100), var(--p-surface-800));
  background-color: color-mix(in srgb, var(--bg), transparent 50%);
}

.file-preview-header h3 {
  margin: 0;
  font-size: 0.75rem;
  font-weight: 500;
  text-transform: uppercase;
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

.original-filename-cell {
  display: flex;
  align-items: center;
  height: 100%;
  padding: 0.5rem 0;
}

.original-filename {
  color: var(--p-text-muted-color);
  font-size: 0.9rem;
  line-height: 1.2;
}

.new-filename-cell {
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 100%;
  padding: 0.5rem 0;
}

.new-filename {
  font-weight: 500;
  margin-bottom: 0.25rem;
  line-height: 1.2;
  text-wrap: nowrap;
}

.metadata-display {
  color: var(--p-text-muted-color);
  font-size: 0.875rem;
  line-height: 1.2;
}

.status-cell {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}

.status-tag {
  font-size: 0.75rem;
  padding: 0.25rem 0.5rem;
}
</style>
