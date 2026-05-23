import {
  createRouter,
  createWebHistory,
  type RouteRecordRaw,
} from "vue-router";
import { useAuthStore } from "../stores/auth";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "home",
    component: () => import("../views/home/Index.vue"),
  },
  {
    path: "/sign-up",
    name: "signup",
    meta: { guest: true },
    component: () => import("../views/auth/SignUp.vue"),
  },
  {
    path: "/sign-in",
    name: "signin",
    meta: { guest: true },
    component: () => import("../views/auth/SignIn.vue"),
  },
  {
    path: "/dashboard",
    name: "dashboard",
    meta: { requiresAuth: true },
    component: () => import("../views/dashboard/Index.vue"),
  },
  {
    path: "/dashboard/profile",
    name: "profile",
    meta: { requiresAuth: true },
    component: () => import("../views/dashboard/profile/Index.vue"),
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: () => import("../views/common/NotFound.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, _from, next) => {
  const authStore = useAuthStore();
  const isAuthenticated = !!authStore.token;

  if (to.meta.requiresAuth && !isAuthenticated) {
    return next({ name: "signin" });
  }

  if (to.meta.guest && isAuthenticated) {
    return next({ name: "dashboard" });
  }

  next();
});

export default router;
