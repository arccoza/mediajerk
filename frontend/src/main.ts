import { createApp } from "vue"
import App from "./App.vue"
import "./style.css"

import PrimeVue from "primevue/config"
// @ts-ignore
import Aura from "@primeuix/themes/aura"
import "primeicons/primeicons.css"

const app = createApp(App)

app.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {
      darkModeSelector: "system",
      cssLayer: false
    }
  },
})

app.mount("#app")
