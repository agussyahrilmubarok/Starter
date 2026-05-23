<script setup lang="ts">
import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
import { type AxiosError } from "axios";
import { toast } from "vue-sonner";
import { useSignIn } from "../../composables/auth/useSignIn";
import { useAuthStore } from "../../stores/auth";
import { type ValidationErrors } from "../../types/common";
import { type AuthUser, type SignInResponse } from "../../types/auth";

document.title = "Sign In";

const router = useRouter();
const { mutate, isPending } = useSignIn();
const authStore = useAuthStore();

const email = ref<string>("");
const password = ref<string>("");
const errors = reactive<ValidationErrors>({});

const handleSignIn = () => {
  mutate(
    {
      email: email.value,
      password: password.value,
    },
    {
      onSuccess: (data: SignInResponse) => {
        toast.success(data?.message || "Sign in successfully");

        const d = data.data;
        const token = d.token;
        const user: AuthUser = {
          id: d.user.id,
          name: d.user.name,
          email: d.user.email,
          created_at: d.user.created_at,
          updated_at: d.user.updated_at,
        };
        authStore.setAuth(token, user);

        router.push("/dashboard");
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
  <div class="min-vh-100 d-flex align-items-center justify-content-center">
    <div class="col-11 col-sm-10 col-md-8 col-lg-5">
      <div class="card border-0 rounded-4 shadow-lg">
        <div class="card-body p-4 p-md-5">
          <h3 class="fw-bold text-center mb-2">Sign In</h3>
          <p class="text-secondary text-center mb-4">
            Welcome back! Please enter your details
          </p>

          <form @submit.prevent="handleSignIn">
            <div class="row">
              <div class="col-md-12 mb-3">
                <div class="form-group">
                  <label class="form-label fw-semibold">Email Address</label>
                  <input
                    v-model="email"
                    type="email"
                    :class="`form-control rounded-3 ${errors.email ? 'is-invalid' : ''}`"
                    placeholder="Enter your email"
                  />
                  <div v-if="errors.email" class="invalid-feedback">
                    {{ errors.email }}
                  </div>
                </div>
              </div>
            </div>

            <div class="row">
              <div class="form-group">
                <label class="form-label fw-semibold">Password</label>
                <input
                  v-model="password"
                  type="password"
                  :class="`form-control rounded-3 ${errors.password ? 'is-invalid' : ''}`"
                  placeholder="Enter your password"
                />
                <div v-if="errors.password" class="invalid-feedback">
                  {{ errors.password }}
                </div>
              </div>
            </div>

            <button
              type="submit"
              class="btn btn-primary w-100 rounded-4 mt-4"
              :disabled="isPending"
            >
              {{ isPending ? "Loading..." : "Sign In" }}
            </button>

            <div class="text-center mt-3">
              <span class="text-secondary">Don't have an account?</span>
              <RouterLink
                to="/sign-up"
                class="text-decoration-none fw-semibold ms-1"
              >
                Sign Up
              </RouterLink>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>
