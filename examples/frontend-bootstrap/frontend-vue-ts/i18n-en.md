# i18n — English (en)

This document contains all strings that need to be internationalized from the `frontend-vue-ts` codebase, grouped by page/component. Keys are structured using flat dot-notation namespacing for compatibility with popular libraries such as `vue-i18n`.

---

## Key Convention

```
<namespace>.<section>.<identifier>
```

Example: `signIn.form.emailLabel` → namespace `signIn`, section `form`, identifier `emailLabel`.

---

## 1. Common / Shared

> Used in more than one place — place under the `common` namespace.

| Key | Value (EN) | Used In |
|-----|-----------|---------|
| `common.loading` | `Loading...` | SignIn, UserEdit |
| `common.saving` | `Saving...` | UserCreate |
| `common.updating` | `Updating...` | UserEdit |
| `common.deleting` | `Deleting...` | UsersIndex |
| `common.cancel` | `Cancel` | UserCreate, UserEdit |
| `common.somethingWentWrong` | `Something went wrong` | SignIn, SignUp, UserCreate, UserEdit |
| `common.fullName` | `Full Name` | SignUp, UserCreate, UserEdit, Profile |
| `common.emailAddress` | `Email Address` | SignIn, SignUp, UserCreate, UserEdit, Profile |
| `common.password` | `Password` | SignIn, SignUp, UserCreate, UserEdit |

---

## 2. Document Titles (useDocumentTitle)

> Strings passed to `useDocumentTitle()` — appear in the browser tab as `"<title> — <APP_NAME>"`.

| Key | Value (EN) | Component |
|-----|-----------|-----------|
| `pageTitle.home` | `Home` | `views/home/Index.vue` |
| `pageTitle.signIn` | `Sign In` | `views/auth/SignIn.vue` |
| `pageTitle.signUp` | `Sign Up` | `views/auth/SignUp.vue` |
| `pageTitle.dashboard` | `Dashboard` | `views/dashboard/Index.vue` |
| `pageTitle.profile` | `Profile` | `views/dashboard/profile/Index.vue` |
| `pageTitle.users` | `Users` | `views/dashboard/users/Index.vue` |
| `pageTitle.createUser` | `Create User` | `views/dashboard/users/Create.vue` |
| `pageTitle.editUser` | `Edit User` | `views/dashboard/users/Edit.vue` |
| `pageTitle.notFound` | `404` | `views/common/NotFound.vue` |

---

## 3. Home Page

> `views/home/Index.vue`

| Key | Value (EN) |
|-----|-----------|
| `home.badge` | `Vue + TypeScript` |
| `home.getStarted` | `Get Started` |
| `home.signIn` | `Sign In` |

---

## 4. Auth — Sign In

> `views/auth/SignIn.vue`

| Key | Value (EN) |
|-----|-----------|
| `signIn.heading` | `Sign In` |
| `signIn.subheading` | `Welcome back! Please enter your details` |
| `signIn.form.emailLabel` | `Email Address` |
| `signIn.form.emailPlaceholder` | `Enter your email` |
| `signIn.form.passwordLabel` | `Password` |
| `signIn.form.passwordPlaceholder` | `Enter your password` |
| `signIn.form.submitButton` | `Sign In` |
| `signIn.form.submitButtonLoading` | `Loading...` |
| `signIn.footer.noAccount` | `Don't have an account?` |
| `signIn.footer.signUpLink` | `Sign Up` |
| `signIn.toast.success` | `Signed in successfully` |

---

## 5. Auth — Sign Up

> `views/auth/SignUp.vue`

| Key | Value (EN) |
|-----|-----------|
| `signUp.heading` | `Create Account` |
| `signUp.subheading` | `Sign up to continue` |
| `signUp.form.nameLabel` | `Full Name` |
| `signUp.form.namePlaceholder` | `Enter your full name` |
| `signUp.form.emailLabel` | `Email Address` |
| `signUp.form.emailPlaceholder` | `Enter your email` |
| `signUp.form.passwordLabel` | `Password` |
| `signUp.form.passwordPlaceholder` | `Enter your password` |
| `signUp.form.submitButton` | `Sign Up` |
| `signUp.form.submitButtonLoading` | `Loading...` |
| `signUp.footer.hasAccount` | `Already have an account?` |
| `signUp.footer.signInLink` | `Sign In` |
| `signUp.toast.success` | `Signed up successfully` |

---

## 6. Dashboard

> `views/dashboard/Index.vue`

| Key | Value (EN) |
|-----|-----------|
| `dashboard.cardHeader` | `DASHBOARD` |
| `dashboard.welcome` | `Welcome,` |

---

## 7. Profile

> `views/dashboard/profile/Index.vue`

| Key | Value (EN) |
|-----|-----------|
| `profile.cardHeader` | `MY PROFILE` |
| `profile.noUserData` | `User data not found. Please log in again.` |
| `profile.fields.fullName` | `FULL NAME` |
| `profile.fields.email` | `EMAIL` |
| `profile.fields.userId` | `USER ID` |
| `profile.fields.memberSince` | `MEMBER SINCE` |

---

## 8. Users — Index

> `views/dashboard/users/Index.vue`

| Key | Value (EN) |
|-----|-----------|
| `users.cardHeader` | `USERS` |
| `users.addButton` | `ADD USER` |
| `users.table.colFullName` | `Full Name` |
| `users.table.colEmail` | `Email Address` |
| `users.table.colActions` | `Actions` |
| `users.table.editButton` | `EDIT` |
| `users.table.deleteButton` | `DELETE` |
| `users.table.deletingButton` | `DELETING...` |
| `users.alert.loading` | `Loading...` |
| `users.confirm.delete` | `Are you sure you want to delete this user?` |

---

## 9. Users — Create

> `views/dashboard/users/Create.vue`

| Key | Value (EN) |
|-----|-----------|
| `userCreate.cardHeader` | `ADD USER` |
| `userCreate.form.nameLabel` | `Full Name` |
| `userCreate.form.namePlaceholder` | `Full Name` |
| `userCreate.form.emailLabel` | `Email Address` |
| `userCreate.form.emailPlaceholder` | `Email Address` |
| `userCreate.form.passwordLabel` | `Password` |
| `userCreate.form.passwordPlaceholder` | `Password` |
| `userCreate.form.submitButton` | `Save` |
| `userCreate.form.submitButtonLoading` | `Saving...` |
| `userCreate.toast.success` | `User created successfully` |

---

## 10. Users — Edit

> `views/dashboard/users/Edit.vue`

| Key | Value (EN) |
|-----|-----------|
| `userEdit.cardHeader` | `EDIT USER` |
| `userEdit.form.nameLabel` | `Full Name` |
| `userEdit.form.namePlaceholder` | `Full Name` |
| `userEdit.form.emailLabel` | `Email Address` |
| `userEdit.form.emailPlaceholder` | `Email Address` |
| `userEdit.form.passwordLabel` | `Password` |
| `userEdit.form.passwordPlaceholder` | `Leave blank to keep current password` |
| `userEdit.form.submitButton` | `Update` |
| `userEdit.form.submitButtonLoading` | `Updating...` |
| `userEdit.loading` | `Loading...` |
| `userEdit.toast.success` | `User updated successfully` |

---

## 11. Not Found (404)

> `views/common/NotFound.vue`

| Key | Value (EN) |
|-----|-----------|
| `notFound.heading` | `Page Not Found` |
| `notFound.description` | `The page you are looking for does not exist.` |
| `notFound.backButton` | `Back to Home` |

---

## 12. NavBar

> `components/dashboard/NavBar.vue`

| Key | Value (EN) |
|-----|-----------|
| `navbar.toggleAriaLabel` | `Toggle navigation` |
| `navbar.myProfile` | `My Profile` |
| `navbar.signOut` | `Sign Out` |

---

## 13. SideBar

> `components/dashboard/SideBar.vue`

| Key | Value (EN) |
|-----|-----------|
| `sidebar.menuHeader` | `MAIN MENU` |
| `sidebar.nav.dashboard` | `Dashboard` |
| `sidebar.nav.users` | `Users` |

---

## Key Summary

| Namespace | Key Count |
|-----------|-----------|
| `common` | 9 |
| `pageTitle` | 9 |
| `home` | 3 |
| `signIn` | 11 |
| `signUp` | 12 |
| `dashboard` | 2 |
| `profile` | 6 |
| `users` | 10 |
| `userCreate` | 10 |
| `userEdit` | 11 |
| `notFound` | 3 |
| `navbar` | 3 |
| `sidebar` | 3 |
| **Total** | **92** |

---

## Implementation Notes

- Keys under `common.*` **should not be duplicated** in other namespaces — simply call `t('common.fullName')` from anywhere.
- Toast strings that fall back from the API (e.g. `data?.message || "Signed in successfully"`) — only the fallback portion needs to be i18n'd; strings from the server are typically already in the language chosen by the backend.
- The `confirm(...)` call in `users.confirm.delete` uses the native browser dialog — consider replacing it with a custom modal so the text can be properly translated.
- `profile.fields.*` values are uppercase because they are rendered as uppercase labels in the UI — you may store them in lowercase and apply `text-transform: uppercase` via CSS instead, depending on team preference.
- **Bug note:** `SignUp.vue` has a copy-paste bug — the toast fallback reads `"Signed in successfully"` instead of `"Signed up successfully"`. Fix this when implementing i18n.
