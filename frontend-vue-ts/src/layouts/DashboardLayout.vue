<script setup lang="ts">
import { useRouter } from "vue-router";
import { useAuthStore } from "../stores/auth";
import NavBar from "../components/dashboard/NavBar.vue";
import SideBar from "../components/dashboard/SideBar.vue";

const router = useRouter();
const authStore = useAuthStore();

const handleSignOut = () => {
  authStore.signOut();
  router.push({ name: "signin" });
};
</script>

<template>
  <div>
    <NavBar
      :userEmail="authStore.user?.email || 'User'"
      :onSignOut="handleSignOut"
    />
    <div class="container mt-4 mb-5">
      <div class="row g-4">
        <div class="col-12 col-md-3">
          <SideBar />
        </div>
        <div class="col-12 col-md-9">
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>
