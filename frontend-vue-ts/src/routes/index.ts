import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "home",
    component: () =>
      import("../views/home/Index.vue"),
  },
  {
    path: "/sign-up",
    name: "signup",
    component: () =>
      import("../views/auth/SignUp.vue"),
  },
  {
    path: "/sign-in",
    name: "signin",
    component: () =>
      import("../views/auth/SignIn.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;