TODO: Internationalization (i18n) Implementation
================================================

[ ] 1. Choose i18n library
      - Evaluate: go-i18n (nicksnyder/go-i18n) or gotext
      - Recommended: go-i18n (widely used, supports pluralization, TOML/JSON/YAML)

[ ] 2. Set up i18n package structure
      - Create pkg/i18n/ package
      - Create locales/ directory at project root
      - Add translation files: locales/en.toml, locales/id.toml (etc.)

[ ] 3. Create translation keys for all hardcoded strings

      [ ] 3a. Controllers (Go backend messages)
            - "Failed to load users"
            - "Something went wrong"
            - "Email already exists"
            - "User not found"
            - "User added successfully!"
            - "User updated successfully!"
            - "User deleted successfully!"
            - "Failed to delete user"
            - "Sign up successfully!"
            - "Validation error"
            - "Internal server error"
            - "Bad request"
            - "Invalid email"
            - "Invalid password"

      [ ] 3b. HTML Templates
            - dashboard_index.html : "Selamat Datang," → "Welcome,"  (CURRENTLY IN INDONESIAN — fix first)
            - sign_in_index.html   : "Welcome back! Please enter your details"
            - sign_in_index.html   : "Don't have an account?"
            - sign_up_index.html   : "Create Account", "Sign up to continue"
            - sign_up_index.html   : "Already have an account?"
            - users_index.html     : "No users found", "Are you sure you want to delete this user?"
            - users_add.html       : "Add User", "Full Name", "Email Address", "Password"
            - users_edit.html      : "Edit User", "Leave blank if you don't want to change the password"
            - profile_index.html   : "MY PROFILE", "FULL NAME", "EMAIL", "USER ID", "MEMBER SINCE"
            - layouts              : "MAIN MENU", "Dashboard", "Users", "My Profile", "Sign Out"

      [ ] 3c. Validator error messages (pkg/validator/validator.go)
            - "is required"
            - "Invalid email format"
            - "already exists"
            - "must be at least %s characters"
            - "must be at most %s characters"
            - "must be a number"
            - "Invalid value"

[ ] 4. Implement locale detection middleware
      - Detect from: Accept-Language header (browser default)
      - Optionally: URL prefix (/en/, /id/) or query param (?lang=id)
      - Store active locale in gin.Context
      - Create internal/delivery/web/middleware/locale.go

[ ] 5. Pass translator to templates
      - Add a template function map: "T" func(key string) string
      - Register via gin's FuncMap before loading templates
      - Usage in templates: {{ T "welcome_message" }}

[ ] 6. Pass translator to controllers
      - Extract locale from gin.Context in each handler
      - Use locale to translate flash messages and error strings

[ ] 7. Add language switcher UI
      - Add dropdown or links in both layouts (default_layout.html, dashboard_layout.html)
      - Persist selected language in cookie or session

[ ] 8. Write tests
      - Unit test for locale middleware
      - Unit test for translated validator error messages
      - Smoke test each page renders correctly in each supported locale

[ ] 9. Documentation
      - Update README.md with supported languages
      - Document how to add a new language/translation file