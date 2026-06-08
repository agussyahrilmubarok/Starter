<script setup lang="ts">
import { useAuthStore } from "../../stores/auth";
import DashboardLayout from "../../layouts/DashboardLayout.vue";
import useDocumentTitle from "../../composables/common/useDocumentTitle";

useDocumentTitle("Dashboard");

const authStore = useAuthStore();
</script>

<template>
  <DashboardLayout>
    <div class="card border-0 rounded-4 shadow-sm">
      <div class="card-header">DASHBOARD</div>
      <div class="card-body">
        Welcome, <strong>{{ authStore.user?.name || "" }}</strong>
      </div>
    </div>
  </DashboardLayout>
</template>
