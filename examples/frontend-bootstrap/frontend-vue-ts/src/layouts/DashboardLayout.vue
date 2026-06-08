<script setup lang="ts">
import NavBar from "../components/dashboard/NavBar.vue";
import SideBar from "../components/dashboard/SideBar.vue";
</script>

<template>
  <div>
    <NavBar />
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
