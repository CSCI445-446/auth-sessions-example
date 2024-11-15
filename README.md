# Echo Authentication, sessions, and cookies example

This is a full end-to-end example of using sessions, cookies, and a simple (insecure!) authentication system in Echo.

There are four routes:
- `/`
- `/dashboard`
- `/login`
- `/logout`

The root route (`/`) displays a simple form that POSTs to `/login`.

Once you've logged in, the `/login` route redirects to `/dashboard`.

The `/dashboard` displays a message based on the user's role and provides a log out button.

The `/logout` route destroys the session cookie by setting its MaxAge to -1 and then redirects back to `/`
