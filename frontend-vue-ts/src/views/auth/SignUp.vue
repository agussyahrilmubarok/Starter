<script setup lang="ts">
import { ref, reactive } from "vue";
import { useRouter } from "vue-router";
import { type AxiosError } from "axios";
import { toast } from "vue-sonner";
import { useSignUp } from "../../composables/auth/useSignUp";
import { type ValidationErrors } from "../../types/common";
import useDocumentTitle from "../../composables/common/useDocumentTitle";
import { SignUpResponse } from "../../types/auth";

useDocumentTitle("Sign Up");

const router = useRouter();
const { mutate, isPending } = useSignUp();

const name = ref<string>("");
const email = ref<string>("");
const password = ref<string>("");
const errors = reactive<ValidationErrors>({});

const handleSignUp = () => {
  mutate(
    {
      name: name.value,
      email: email.value,
      password: password.value,
    },
    {
      onSuccess: (data: SignUpResponse) => {
        toast.success(data?.message || "Signed in successfully");
        router.push({ name: "signin" });
      },
      onError: (error: Error) => {
        const axiosError = error as AxiosError<{
          errors: Record<string, string>;
        }>;
        const validationErrors = axiosError.response?.data?.errors;
        if (validationErrors) {
          Object.assign(errors, validationErrors)
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
          <h3 class="fw-bold text-center mb-2">Create Account</h3>
          <p class="text-secondary text-center mb-4">Sign up to continue</p>

          <form @submit.prevent="handleSignUp">
          
            <div class="row">
              <div class="col-md-12 mb-3">
                <div class="form-group">
                  <label class="form-label fw-semibold">Full Name</label>
                  <input
                    v-model="name"
                    type="text"
                    :class="`form-control rounded-3 ${errors.name ? 'is-invalid' : ''}`"
                    placeholder="Enter your full name"
                  />
                  <div v-if="errors.name" class="invalid-feedback">
                    {{ errors.name }}
                  </div>
                </div>
              </div>
            </div>

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
              {{ isPending ? "Loading..." : "Sign Up" }}
            </button>

            <div class="text-center mt-3">
              <span class="text-secondary">Already have an account?</span>
              <RouterLink
                to="/sign-in"
                class="text-decoration-none fw-semibold ms-1"
              >
                Sign In
              </RouterLink>
            </div>
            
          </form>
          
        </div>
      </div>
    </div>
  </div>
</template>
