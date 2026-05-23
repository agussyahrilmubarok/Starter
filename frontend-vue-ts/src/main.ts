import { createApp } from "vue";
import { VueQueryPlugin } from "@tanstack/vue-query";
import { createPinia } from "pinia";
import App from "./App.vue";
import routes from "./routes";

const app = createApp(App);
app.use(VueQueryPlugin);
app.use(createPinia());
app.use(routes);
app.mount("#app");
