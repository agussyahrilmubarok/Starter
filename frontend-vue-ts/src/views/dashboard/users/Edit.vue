<script setup lang="ts">
import { ref, reactive, watch } from "vue";
import { useRouter, useRoute } from "vue-router";
import { type AxiosError } from "axios";
import { toast } from "vue-sonner";
import { useUserById } from "../../../composables/user/useUserById";
import { useUserUpdate } from "../../../composables/user/useUserUpdate";
import { type ValidationErrors } from "../../../types/common";
import { type UserResponse } from "../../../types/user";
import DashboardLayout from "../../../layouts/DashboardLayout.vue";

document.title = "Edit User";

const route = useRoute();
const router = useRouter();

const id = route.params.id as string;
const { data: user } = useUserById(id);

const { mutate, isPending } = useUserUpdate();

const name = ref<string>("");
const email = ref<string>("");
const password = ref<string>("");
const errors = reactive<ValidationErrors>({});

watch(
  user,
  (newUser) => {
    if (newUser) {
      name.value = newUser.name;
      email.value = newUser.email;
    }
  },
  { immediate: true },
);

const updateUser = () => {
  mutate(
    {
      id,
      payload: {
        name: name.value,
        email: email.value,
        password: password.value,
      },
    },
    {
      onSuccess: (data: UserResponse) => {
        toast.success(data?.message || "Edit user successfully");
        router.push("/dashboard/users");
      },
      onError: (error: Error) => {
        const axiosError = error as AxiosError<{
          errors: Record<string, string>;
        }>;
        const validationErrors = axiosError.response?.data?.errors;
        if (validationErrors) {
          Object.assign(errors, validationErrors);
        } else {
          toast.error("Something went wrong");
        }
      },
    },
  );
};
</script>

<template>
  <DashboardLayout>
    <div class="card border-0 rounded-4 shadow-sm">
      <div class="card-header">EDIT USER</div>
      <div class="card-body">
        <p v-if="!user">Loading...</p>

        <form v-else @submit.prevent="updateUser">
          <div class="form-group mb-3">
            <label class="mb-1 fw-bold">Full Name</label>
            <input
              v-model="name"
              type="text"
              class="form-control"
              placeholder="Full Name"
            />
            <div v-if="errors.name" class="alert alert-danger mt-2 rounded-4">
              {{ errors.name }}
            </div>
          </div>

          <div class="form-group mb-3">
            <label class="mb-1 fw-bold">Email address</label>
            <input
              v-model="email"
              type="email"
              class="form-control"
              placeholder="Email Address"
            />
            <div v-if="errors.email" class="alert alert-danger mt-2 rounded-4">
              {{ errors.email }}
            </div>
          </div>

          <div class="form-group mb-3">
            <label class="mb-1 fw-bold">Password</label>
            <input
              v-model="password"
              type="password"
              class="form-control"
              placeholder="Password"
            />
            <div
              v-if="errors.password"
              class="alert alert-danger mt-2 rounded-4"
            >
              {{ errors.password }}
            </div>
          </div>

          <button
            type="submit"
            class="btn btn-md btn-primary rounded-4 shadow-sm border-0"
            :disabled="isPending"
          >
            {{ isPending ? "Updating..." : "Update" }}
          </button>

          <RouterLink
            to="/dashboard/users"
            class="btn btn-md btn-secondary rounded-4 shadow-sm border-0 ms-2"
          >
            Cancel
          </RouterLink>
        </form>
      </div>
    </div>
  </DashboardLayout>
</template>
