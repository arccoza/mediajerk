export interface MediaMetadata {
  title?: string
  year?: number
  season?: number
  episode?: number
  episodeTitle?: string
  seriesName?: string
  type?: "movie" | "tv" | "unknown"
}

export interface FileTemplate {
  pattern: string
  name: string
  description: string
}

export const defaultTemplates: FileTemplate[] = [
  {
    name: "TV Show Default",
    pattern: "[title] ([year]) - S[##]E[##]",
    description: "Standard TV show format with title, year, season and episode",
  },
  {
    name: "Movie Default",
    pattern: "[title] ([year])",
    description: "Standard movie format with title and year",
  },
  {
    name: "TV Show Extended",
    pattern: "[title] ([year]) - S[##]E[##] - [episode_title]",
    description: "Extended TV show format including episode title",
  },
]

export class MetadataFormatter {
  static formatMetadataDisplay(metadata: MediaMetadata): string {
    if (metadata.type === "tv") {
      // TV format: "01|02, Show name, Episode title"
      const season = metadata.season?.toString().padStart(2, "0") || "??"
      const episode = metadata.episode?.toString().padStart(2, "0") || "??"
      const showName = metadata.title || "Unknown Show"
      const episodeTitle = metadata.episodeTitle || "Unknown Episode"

      return `${season}|${episode}, ${showName}, ${episodeTitle}`
    } else if (metadata.type === "movie") {
      // Movie format: "Series name, Movie name" or just "Movie name"
      const movieName = metadata.title || "Unknown Movie"

      if (metadata.seriesName) {
        return `${metadata.seriesName}, ${movieName}`
      }

      return movieName
    }

    // Fallback for unknown type
    return metadata.title || "Unknown"
  }
}

export class TemplateProcessor {
  static applyTemplate(
    template: string,
    metadata: MediaMetadata,
    originalFilename: string,
  ): { newFilename: string; status: "ready" | "warning" | "error"; message?: string } {
    try {
      let result = template

      // Extract file extension
      const extension = originalFilename.split(".").pop() || ""

      // Replace template variables
      result = result.replace(/\[title\]/g, metadata.title || "Unknown")
      result = result.replace(/\[year\]/g, metadata.year?.toString() || "Unknown")

      // Handle season/episode formatting
      if (metadata.season !== undefined) {
        result = result.replace(/\[##\]/g, metadata.season.toString().padStart(2, "0"))
        result = result.replace(/\[#\]/g, metadata.season.toString())
      }

      if (metadata.episode !== undefined) {
        const episodeMatch = result.match(/S\d+E\[##?\]/)
        if (episodeMatch) {
          result = result.replace(/E\[##\]/g, `E${metadata.episode.toString().padStart(2, "0")}`)
          result = result.replace(/E\[#\]/g, `E${metadata.episode.toString()}`)
        }
      }

      // Clean up any remaining template variables
      result = result.replace(/\[episode_title\]/g, "")
      result = result.replace(/\s+/g, " ").trim()

      // Add extension
      if (extension) {
        result += `.${extension}`
      }

      // Determine status
      let status: "ready" | "warning" | "error" = "ready"
      let message: string | undefined

      if (!metadata.title || !metadata.year) {
        status = "warning"
        message = "Missing metadata - title or year not found"
      }

      if (
        template.includes("[##]") &&
        (metadata.season === undefined || metadata.episode === undefined)
      ) {
        status = "warning"
        message = "Missing season/episode information"
      }

      return { newFilename: result, status, message }
    } catch (error) {
      return {
        newFilename: originalFilename,
        status: "error",
        message: `Template processing failed: ${error}`,
      }
    }
  }

  static parseTemplate(template: string): string[] {
    const variables = template.match(/\[[^\]]+\]/g) || []
    return variables.map((v) => v.slice(1, -1))
  }

  static validateTemplate(template: string): { valid: boolean; errors: string[] } {
    const errors: string[] = []

    // Check for unmatched brackets
    const openBrackets = (template.match(/\[/g) || []).length
    const closeBrackets = (template.match(/\]/g) || []).length

    if (openBrackets !== closeBrackets) {
      errors.push("Unmatched brackets in template")
    }

    // Check for valid variable names
    const validVariables = ["title", "year", "season", "episode", "episode_title", "#", "##"]
    const variables = TemplateProcessor.parseTemplate(template)

    for (const variable of variables) {
      if (!validVariables.includes(variable)) {
        errors.push(`Unknown template variable: [${variable}]`)
      }
    }

    return {
      valid: errors.length === 0,
      errors,
    }
  }
}
