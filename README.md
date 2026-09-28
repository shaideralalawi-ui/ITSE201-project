# UOB Library

A university library scenario presentation and student-facing app sample. The project contains a Go standard-library server and a static-site version for GitHub Pages.

## Local preview

Run `preview.mjs` with Node.js, then open `http://localhost:8082/app` or `http://localhost:8082/slides`.

## Publish with GitHub Pages

The `docs` folder is already arranged as a static site. After pushing the project to GitHub, open **Settings → Pages**, select **Deploy from a branch**, choose `main` and `/docs`, then save. GitHub Pages will show the public site URL there. The app sample will be the home page, and the presentation will be at `/slides/`.

The GitHub Pages version does not run Go. That is suitable for this sample because the catalog search, theme toggle, and demo buttons run in the browser. Borrowing, reservations, account authentication, and member records are not connected to a server or database.

## Host a Go backend later

For real accounts, saved loans, and persistent records, deploy the Go server to an app host and connect it to a database. GitHub can keep the source and trigger deployments, but it does not keep a Go server running itself. Render's free web services sleep after inactivity, so use a paid always-on service if the backend must stay awake continuously.
