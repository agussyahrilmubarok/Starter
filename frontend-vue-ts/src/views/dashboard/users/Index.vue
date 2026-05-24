<script setup lang="ts">
import { useQueryClient } from "@tanstack/vue-query";
import { useUsers } from "../../../composables/user/useUsers";
import { useUserDelete } from "../../../composables/user/useUserDelete";
import { type User } from "../../../types/user";
import DashboardLayout from "../../../layouts/DashboardLayout.vue";
import useDocumentTitle from "../../../composables/common/useDocumentTitle";

useDocumentTitle("Users");

const { data: users, isLoading, isError, error } = useUsers();
const queryClient = useQueryClient();
const { mutate, isPending } = useUserDelete();

const handleDelete = (id: string) => {
  if (confirm("Are you sure you want to delete this user?")) {
    mutate(id, {
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["users"] });
      },
    });
  }
};
</script>

<template>
  <DashboardLayout>
    <div class="card border-0 rounded-4 shadow-sm">
      <div
        class="card-header d-flex justify-content-between align-items-center"
      >
        <span>USERS</span>
        <RouterLink
          to="/dashboard/users/create"
          class="btn btn-sm btn-success rounded-4 shadow-sm border-0"
        >
          ADD USER
        </RouterLink>
      </div>
      <div class="card-body">
        <div v-if="isLoading" class="alert alert-info text-center">
          Loading...
        </div>

        <div v-if="isError" class="alert alert-danger text-center">
          Error: {{ error?.message }}
        </div>

        <div class="table-responsive">
          <table class="table table-bordered">
            <thead class="bg-dark text-white">
              <tr>
                <th scope="col">Full Name</th>
                <th scope="col">Email Address</th>
                <th scope="col" style="width: 20%">Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="user in users as User[]" :key="user.id">
                <td>{{ user.name }}</td>
                <td>{{ user.email }}</td>
                <td class="text-center">
                  <RouterLink
                    :to="`/dashboard/users/edit/${user.id}`"
                    class="btn btn-sm btn-primary rounded-4 shadow-sm border-0 me-2"
                  >
                    EDIT
                  </RouterLink>
                  <button
                    @click="handleDelete(user.id)"
                    :disabled="isPending"
                    class="btn btn-sm btn-danger rounded-4 shadow-sm border-0"
                  >
                    {{ isPending ? "DELETING..." : "DELETE" }}
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>
