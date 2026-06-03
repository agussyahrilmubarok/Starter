TODO: Internationalization (i18n) Implementation
================================================
Scope: HTML templates only (frontend-facing text)
      Backend Go code, flash messages, and error logs stay in English.

[ ] 2. Create translation files
      - Create locales/ directory at project root
      - Add locales/en.toml (default/fallback)
      - Add locales/id.toml (Bahasa Indonesia)
      - Cover all strings listed in item 4 below

[ ] 3. Set up i18n in templates
      - Add template FuncMap with "T" translation function
        e.g. {{ T "welcome_message" }} → "Welcome," / "Selamat Datang,"
      - Register FuncMap in controller.go before LoadTemplate()
      - Pass active locale into every render() call via gin.H

[ ] 4. Replace hardcoded strings in templates with translation keys

      [ ] 4a. layouts/dashboard_layout.html
            - "MAIN MENU"
            - "Dashboard"
            - "Users"
            - "My Profile"
            - "Sign Out"

      [ ] 4b. layouts/default_layout.html
            - (no user-visible hardcoded text currently)

      [ ] 4c. auth/sign_in_index.html
            - "Sign In"
            - "Welcome back! Please enter your details"
            - "Email Address"
            - "Password"
            - "Enter your email"
            - "Enter your password"
            - "Sign In" (button)
            - "Don't have an account?"
            - "Sign Up" (link)

      [ ] 4d. auth/sign_up_index.html
            - "Create Account"
            - "Sign up to continue"
            - "Full Name"
            - "Email Address"
            - "Password"
            - "Enter your full name"
            - "Enter your email"
            - "Enter your password"
            - "Sign Up" (button)
            - "Already have an account?"
            - "Sign In" (link)

      [ ] 4e. dashboard/dashboard_index.html
            - "DASHBOARD"
            - "Welcome," (was "Selamat Datang,")

      [ ] 4f. dashboard/profile/profile_index.html
            - "MY PROFILE"
            - "FULL NAME"
            - "EMAIL"
            - "USER ID"
            - "MEMBER SINCE"
            - "User data not found. Please log in again."

      [ ] 4g. dashboard/users/users_index.html
            - "USERS"
            - "ADD USER"
            - "Full Name"
            - "Email Address"
            - "Actions"
            - "EDIT"
            - "DELETE"
            - "No users found"
            - "Are you sure you want to delete this user?" (JS confirm)

      [ ] 4h. dashboard/users/users_add.html
            - "ADD USER"
            - "Full Name"
            - "Email Address"
            - "Password"
            - "Save" (button)
            - "Cancel" (button)

      [ ] 4i. dashboard/users/users_edit.html
            - "EDIT USER"
            - "Full Name"
            - "Email Address"
            - "Password"
            - "Leave blank if you don't want to change the password"
            - "Update" (button)
            - "Cancel" (button)

[ ] 5. Implement locale detection middleware
      - Create internal/delivery/web/middleware/locale.go
      - Detect locale from: cookie (user preference) → Accept-Language header → fallback to "en"
      - Store active locale in gin.Context (c.Set("locale", "id"))

[ ] 6. Add language switcher UI
      - Add switcher to both layout files (default_layout.html, dashboard_layout.html)
      - On switch: set cookie and redirect back to current page
      - Add a POST /language route to handle the switch

[ ] 7. Documentation
      - Update README.md: list supported languages
      - Document how to add a new locale file