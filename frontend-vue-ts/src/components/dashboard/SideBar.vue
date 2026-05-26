<script setup lang="ts"></script>

<template>
  <div class="card border-0 rounded-4 shadow-sm">
    <div class="card-header">MAIN MENU</div>
    <div class="card-body">
      <div class="list-group">
        <RouterLink
          to="/dashboard"
          class="list-group-item list-group-item-action"
          exact-active-class="active"
        >
          Dashboard
        </RouterLink>

        <RouterLink
          to="/dashboard/users"
          class="list-group-item list-group-item-action"
          active-class="active"
        >
          Users
        </RouterLink>
      </div>
    </div>
  </div>
</template>
