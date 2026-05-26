import { createApp } from "vue";
import { createPinia } from "pinia";
import { VueQueryPlugin } from "@tanstack/vue-query";
import App from "./App.vue";
import routes from "./routes";

const app = createApp(App);

app.use(createPinia());
app.use(VueQueryPlugin);
app.use(routes);

app.mount("#app");
