<script setup lang="ts">
import { useAuthStore } from "../../stores/auth";
import DashboardLayout from "../../layouts/DashboardLayout.vue";

document.title = "Dashboard";

const authStore = useAuthStore();

</script>

<template>
  <dashboard-layout>
    <div class="card border-0 rounded-4 shadow-sm">
      <div class="card-header">DASHBOARD</div>
      <div class="card-body">
        Selamat Datang, <strong>{{ authStore.user?.name || "" }}</strong>
      </div>
    </div>
  </dashboard-layout>
</template>
