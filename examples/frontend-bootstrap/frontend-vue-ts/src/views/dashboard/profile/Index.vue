<script setup lang="ts">
import { computed } from "vue";
import { useAuthStore } from "../../../stores/auth";
import DashboardLayout from "../../../layouts/DashboardLayout.vue";
import useDocumentTitle from "../../../composables/common/useDocumentTitle";

useDocumentTitle("Profile");

const authStore = useAuthStore();

const user = computed(() => authStore.user);
</script>

<template>
  <DashboardLayout>
    <div v-if="!user" class="card border-0 rounded-4 shadow-sm">
      <div class="card-body text-center text-muted py-5">
        User data not found. Please log in again.
      </div>
    </div>

    <div v-else class="card border-0 rounded-4 shadow-sm">
      <div class="card-header fw-semibold">MY PROFILE</div>
      <div class="card-body">
        <div class="text-center mb-4">
          <div
            class="rounded-circle d-inline-flex align-items-center justify-content-center bg-dark text-white fw-bold"
            :style="{ width: '80px', height: '80px', fontSize: '32px' }"
          >
            {{ user.name?.charAt(0).toUpperCase() ?? "?" }}
          </div>
          <h5 class="mt-3 mb-0">{{ user.name }}</h5>
          <small class="text-muted">{{ user.email }}</small>
        </div>

        <hr />

        <div class="row g-3">
          <div class="col-12 col-md-6">
            <label class="form-label text-muted small fw-semibold"
              >FULL NAME</label
            >
            <p class="form-control-plaintext fw-medium">{{ user.name }}</p>
          </div>
          <div class="col-12 col-md-6">
            <label class="form-label text-muted small fw-semibold">EMAIL</label>
            <p class="form-control-plaintext fw-medium">{{ user.email }}</p>
          </div>
          <div class="col-12 col-md-6">
            <label class="form-label text-muted small fw-semibold"
              >USER ID</label
            >
            <p class="form-control-plaintext text-muted small">{{ user.id }}</p>
          </div>
          <div class="col-12 col-md-6">
            <label class="form-label text-muted small fw-semibold"
              >MEMBER SINCE</label
            >
            <p class="form-control-plaintext fw-medium">
              {{
                new Date(user.created_at).toLocaleDateString("en-US", {
                  day: "numeric",
                  month: "long",
                  year: "numeric",
                })
              }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>
