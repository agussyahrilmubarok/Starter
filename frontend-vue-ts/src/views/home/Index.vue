<script setup lang="ts">
</script>

<template>
  <div class="min-vh-100 d-flex align-items-center justify-content-center">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-lg-8 col-sm-6">
          <div class="card border-0 rounded-5 shadow-lg">
            <div class="card-body p-5 text-center">
              <span class="badge bg-primary px-3 py-2 mb-3 rounded-pill">
                Vue + TypeScript
              </span>
              <h1 class="display-3 fw-bold mb-3">frontend-vue-ts</h1>
              <div
                class="d-flex flex-column flex-sm-row gap-3 justify-content-center"
              >
                <RouterLink
                  to="/sign-up"
                  class="btn btn-primary btn-lg px-5 rounded-pill"
                >
                  Get Started
                </RouterLink>
                <RouterLink
                  to="/sign-in"
                  class="btn btn-dark btn-lg px-5 rounded-pill"
                >
                  Sign In
                </RouterLink>
              </div>
              <div class="mt-4 text-dark opacity-50 small">
                starter.agussyahrilmubarok.github.io
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
