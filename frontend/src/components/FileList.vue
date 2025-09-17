<script lang="ts" setup>
import DataTable from "primevue/datatable"
import Column from "primevue/column"
import Button from "primevue/button"

interface FileItem {
  id: number
  filename: string
  path: string
  size?: string
}

interface Props {
  files: FileItem[]
}

interface Emits {
  (e: "row-reorder", files: FileItem[]): void
  (e: "remove-file", id: number): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const removeFile = (id: number) => {
  emit("remove-file", id)
}

const onRowReorder = (event: any) => {
  emit("row-reorder", event.value)
}
</script>

<template>
  <div class="file-list">
    <div class="file-list-header">
      <h3>Original Files</h3>
    </div>

    <DataTable
      :value="props.files"
      :scrollable="true"
      scroll-height="100%"
      reorderableRows
      @row-reorder="onRowReorder"
      class="file-table"
      :pt="{
        table: { style: 'min-width: 100%' },
        bodyRow: { style: 'height: 4rem' }
      }"
    >
      <Column :rowReorder="true" header="" style="width: 3rem" />

      <Column field="filename" header="Filename" :sortable="true" style="min-width: 200px">
        <template #body="{ data }">
          <div class="filename-cell">
            <span class="filename">{{ data.filename }}</span>
            <small class="file-path">{{ data.path }}</small>
          </div>
        </template>
      </Column>

      <Column field="size" header="Size" :sortable="true" style="width: 100px">
        <template #body="{ data }">
          <div class="size-cell">
            <span class="file-size">{{ data.size }}</span>
          </div>
        </template>
      </Column>

      <Column header="Actions" style="width: 80px">
        <template #body="{ data }">
          <div class="actions-cell">
            <Button
              icon="pi pi-trash"
              severity="danger"
              text
              rounded
              @click="removeFile(data.id)"
              class="remove-btn"
            />
          </div>
        </template>
      </Column>
    </DataTable>
  </div>
</template>

<style scoped>
.file-list {
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 1em;
}

.file-list-header {
  padding: 0.75rem;
  border-bottom: 1px solid var(--p-surface-border);
  --bg: light-dark(var(--p-surface-100), var(--p-surface-800));
  background-color: color-mix(in srgb, var(--bg), transparent 50%);
}

.file-list-header h3 {
  margin: 0;
  font-size: 0.75rem;
  font-weight: 500;
  text-transform: uppercase;
}

.file-table {
  flex: 1;
  border: none;
}

.filename-cell {
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 100%;
  padding: 0.5rem 0;
}

.filename {
  font-weight: 500;
  margin-bottom: 0.25rem;
  line-height: 1.2;
  text-wrap: nowrap;
}

.file-path {
  color: var(--p-text-muted-color);
  font-size: 0.875rem;
  line-height: 1.2;
}

.size-cell {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}

.file-size {
  font-family: monospace;
  color: var(--p-text-muted-color);
}

.actions-cell {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}

.remove-btn {
  width: 2rem;
  height: 2rem;
}
</style>
