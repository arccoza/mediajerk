import { computed, readonly, ref, shallowRef } from "vue"
import type { main } from "../../wailsjs/go/models"


export type FileInfo = main.FileInfo
export type MoreFileInfo = FileInfo & {
  newName?: string
  metadata?: any
}

// Shared reactive state
const filesMap = ref(new Map<string, MoreFileInfo>())
const selectedSet = ref(new Set<string>())

export function useFiles() {
  // Computed properties for selection
  const files = computed(() => [...filesMap.value.values()])
  const selectedFiles = computed(() => files.value.filter((v) => selectedSet.value.has(v.path)))
  const hasSelection = computed(() => selectedSet.value.size > 0)
  const selectedCount = computed(() => selectedSet.value.size)

  // Add files to the collection
  const addFiles = (newFiles: FileInfo[]) => {
    // // Filter out duplicates based on path
    for (const f of newFiles) {
      filesMap.value.set(f.path, f)
    }
  }

  // Selection management functions
  const selectFiles = (newSelection: FileInfo[]) => {
    selectedSet.value = new Set(newSelection.map((v) => v.path))
  }

  const clearSelection = () => {
    selectedSet.value.clear()
  }

  // Remove file by path
  const removeFile = (path: string) => {
    filesMap.value.delete(path)
    // Also remove from selection if it was selected
    selectedSet.value.delete(path)
  }

  // Remove currently selected files
  const removeSelectedFiles = () => {
    for (const path of selectedSet.value) {
      filesMap.value.delete(path)
    }

    clearSelection()
  }

  // Remove file by index
  const removeFileByIndex = (index: number) => {
    const path = [...filesMap.value.values()][index]?.path
    filesMap.value.delete(path)
  }

  // Clear all files
  const clearFiles = () => {
    filesMap.value.clear()
  }

  // Reorder files
  const reorderFiles = (reorderedFiles: FileInfo[]) => {
    clearFiles()
    addFiles(reorderedFiles)
  }

  // Get file count
  const fileCount = computed(() => filesMap.value.size)

  // Check if files list is empty
  const isEmpty = computed(() => filesMap.value.size === 0)

  return {
    // Read-only access to files
    filesMap,
    files,
    selectedFiles,

    // Computed properties
    fileCount,
    isEmpty,
    hasSelection,
    selectedCount,

    // Methods
    addFiles,
    removeFile,
    removeFileByIndex,
    removeSelectedFiles,
    clearFiles,
    reorderFiles,
    selectFiles,
    clearSelection,
  }
}
