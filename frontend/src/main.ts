import { createApp } from "vue"
import App from "./App.vue"
import "./style.css"

// @ts-expect-error
import Aura from "@primeuix/themes/aura"
import PrimeVue from "primevue/config"
import "primeicons/primeicons.css"

const app = createApp(App)

app.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {
      darkModeSelector: "system",
      cssLayer: false,
    },
  },
})

app.mount("#app")
