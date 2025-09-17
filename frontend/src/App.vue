<script lang="ts" setup>
import { ref } from "vue"
import AppHeader from "./components/AppHeader.vue"
import FileList from "./components/FileList.vue"
import FilePreview from "./components/FilePreview.vue"
import Splitter from "primevue/splitter"
import SplitterPanel from "primevue/splitterpanel"
import { MediaMetadata } from "./utils/templateProcessor"

interface FileItem {
  id: number
  filename: string
  path: string
  size?: string
}

interface PreviewItem {
  id: number
  originalFilename: string
  newFilename: string
  metadata: MediaMetadata
  status: "ready" | "warning" | "error"
  message?: string
}

// Shared data state
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
  {
    id: 3,
    filename: "The.Lord.of.the.Rings.The.Two.Towers.2002.mkv",
    path: "/media/movies/",
    size: "2.1 GB",
  },
  {
    id: 4,
    filename: "The.Matrix.1999.mp4",
    path: "/media/movies/",
    size: "1.8 GB",
  },
])

const previewFiles = ref<PreviewItem[]>([
  {
    id: 1,
    originalFilename: "Game.of.Thrones.S01E01.Winter.is.Coming.mkv",
    newFilename: "Game of Thrones (2011) - S01E01.mkv",
    metadata: {
      title: "Game of Thrones",
      year: 2011,
      season: 1,
      episode: 1,
      episodeTitle: "Winter is Coming",
      type: "tv"
    },
    status: "ready",
  },
  {
    id: 2,
    originalFilename: "Breaking.Bad.S01E01.Pilot.mp4",
    newFilename: "Breaking Bad (2008) - S01E01.mp4",
    metadata: {
      title: "Breaking Bad",
      year: 2008,
      season: 1,
      episode: 1,
      episodeTitle: "Pilot",
      type: "tv"
    },
    status: "ready",
  },
  {
    id: 3,
    originalFilename: "The.Lord.of.the.Rings.The.Two.Towers.2002.mkv",
    newFilename: "The Lord of the Rings - The Two Towers (2002).mkv",
    metadata: {
      title: "The Two Towers",
      year: 2002,
      seriesName: "The Lord of the Rings",
      type: "movie"
    },
    status: "ready",
  },
  {
    id: 4,
    originalFilename: "The.Matrix.1999.mp4",
    newFilename: "The Matrix (1999).mp4",
    metadata: {
      title: "The Matrix",
      year: 1999,
      type: "movie"
    },
    status: "ready",
  },
])

// Synchronized reordering functions
const onFilesReorder = (reorderedFiles: FileItem[]) => {
  files.value = reorderedFiles
  // Reorder preview files to match the same order by ID
  const fileOrder = reorderedFiles.map(f => f.id)
  previewFiles.value.sort((a, b) => fileOrder.indexOf(a.id) - fileOrder.indexOf(b.id))
}

const onPreviewReorder = (reorderedPreviews: PreviewItem[]) => {
  previewFiles.value = reorderedPreviews
  // Reorder files to match the same order by ID
  const previewOrder = reorderedPreviews.map(p => p.id)
  files.value.sort((a, b) => previewOrder.indexOf(a.id) - previewOrder.indexOf(b.id))
}

const removeFile = (id: number) => {
  files.value = files.value.filter(file => file.id !== id)
  previewFiles.value = previewFiles.value.filter(preview => preview.id !== id)
}
</script>

<template>
  <div class="app-container">
    <AppHeader />

    <div class="app-main">
      <Splitter style="height: 100%; border: none">
        <SplitterPanel :size="50" :min-size="30">
          <FileList
            :files="files"
            @row-reorder="onFilesReorder"
            @remove-file="removeFile"
          />
        </SplitterPanel>
        <SplitterPanel :size="50" :min-size="30">
          <FilePreview
            :previewFiles="previewFiles"
            @row-reorder="onPreviewReorder"
          />
        </SplitterPanel>
      </Splitter>
    </div>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body {
  height: 100%;
  font-family: var(--p-font-family);
  background-color: light-dark(var(--p-surface-0), var(--p-surface-900));
  color: var(--p-text-color);
}

#app {
  height: 100vh;
}

.app-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 0em;
}

.app-main {
  flex: 1;
  overflow: hidden;
  /* margin: var(--p-toolbar-padding); */
}
</style>
