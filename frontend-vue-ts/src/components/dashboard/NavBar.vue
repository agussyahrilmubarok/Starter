<script setup lang="ts">
import { APP_NAME } from "../../constants/app";

const props = withDefaults(
  defineProps<{
    userEmail?: string;
    onSignOut?: () => void;
  }>(),
  {
    userEmail: "User",
  },
);
</script>

<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark px-3 shadow-sm">
    <RouterLink class="navbar-brand fw-bold fs-5" to="/dashboard">
      ⚡ {{ APP_NAME }}
    </RouterLink>

    <button
      class="navbar-toggler"
      type="button"
      data-bs-toggle="collapse"
      data-bs-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent"
      aria-expanded="false"
      aria-label="Toggle navigation"
    >
      <span class="navbar-toggler-icon" />
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav ms-auto">
        <li class="nav-item dropdown">
          <a
            class="nav-link dropdown-toggle text-light"
            href="#"
            role="button"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            👤 {{ props.userEmail }}
          </a>
          <ul class="dropdown-menu dropdown-menu-end">
            <li>
              <RouterLink class="dropdown-item" to="/dashboard/profile">
                My Profile
              </RouterLink>
            </li>
            <li><hr class="dropdown-divider" /></li>
            <li>
              <!-- [CONVERT] onClick={onSignOut} → @click="props.onSignOut?.()" -->
              <button
                class="dropdown-item text-danger"
                @click="props.onSignOut?.()"
              >
                Sign Out
              </button>
            </li>
          </ul>
        </li>
      </ul>
    </div>
  </nav>
</template>
