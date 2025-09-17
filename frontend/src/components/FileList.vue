<script lang="ts" setup>
import { ref } from "vue"
import DataTable from "primevue/datatable"
import Column from "primevue/column"
import Button from "primevue/button"

interface FileItem {
  id: number
  filename: string
  path: string
  size?: string
}

const files = ref<FileItem[]>([
  {
    id: 1,
    filename: "Game.of.Thrones.S01E01.Winter.is.Coming.mkv",
    path: "/media/tv-shows/",
    size: "1.2 GB",
  },
  {
    id: 2,
    filename: "Breaking.Bad.S01E01.Pilot.mp4",
    path: "/media/tv-shows/",
    size: "850 MB",
  },
])

const removeFile = (id: number) => {
  files.value = files.value.filter((file) => file.id !== id)
}
</script>

<template>
  <div class="file-list">
    <div class="file-list-header">
      <h3>Original Files</h3>
    </div>

    <DataTable
      :value="files"
      :scrollable="true"
      scroll-height="100%"
      class="file-table"
      :pt="{ table: { style: 'min-width: 100%' } }"
    >
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
          <span class="file-size">{{ data.size }}</span>
        </template>
      </Column>

      <Column header="Actions" style="width: 80px">
        <template #body="{ data }">
          <Button
            icon="pi pi-trash"
            severity="danger"
            text
            rounded
            @click="removeFile(data.id)"
            class="remove-btn"
          />
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
  padding: 0.5rem;
  border-bottom: 1px solid var(--p-surface-border);
  background-color: light-dark(var(--p-surface-100), var(--p-surface-800));
}

.file-list-header h3 {
  margin: 0;
  font-size: 1rem;
  font-weight: bold;
}

.file-table {
  flex: 1;
  border: none;
}

.filename-cell {
  display: flex;
  flex-direction: column;
}

.filename {
  font-weight: 500;
  margin-bottom: 0.25rem;
}

.file-path {
  color: var(--p-text-muted-color);
  font-size: 0.875rem;
}

.file-size {
  font-family: monospace;
  color: var(--p-text-muted-color);
}

.remove-btn {
  width: 2rem;
  height: 2rem;
}
</style>
