# UOB Library Management System

<<<<<<< HEAD
A university library project with an animated presentation of the library scenario and a student-facing catalog demo. The interface is built with HTML, CSS, and JavaScript, with a small Go server for local hosting.

## Features

- Browse a sample library catalog and search for books
- View sample availability, reservations, and account information
- Switch between light and dark themes
- Explore an animated presentation explaining the library workflow
- View separate student, librarian, and administrator roles

> **Demo note:** This is a front-end sample. Sign-in, loans, reservations, notifications, fines, and reports are demonstrations only; data is not saved to a database.

## Run locally with Go

You need [Go 1.22 or later](https://go.dev/dl/).

1. Download or clone this repository.
2. Open a terminal in the project folder.
3. Run:

   ```bash
   go run .
   ```

4. Open the app at [http://localhost:8081/app](http://localhost:8081/app).
5. Open the presentation at [http://localhost:8081/slides](http://localhost:8081/slides).

On Windows, you can also double-click `start.bat`.

## Preview with Node.js (optional)

If Node.js is installed, run `node preview.mjs` from the project folder. Then open:

- App: [http://localhost:8082/app](http://localhost:8082/app)
- Slides: [http://localhost:8082/slides](http://localhost:8082/slides)

## Publish the static site with GitHub Pages

The `docs` folder contains the static version prepared for GitHub Pages. Push the project files to the `main` branch, then in the repository settings:

1. Open **Settings → Pages**.
2. Under **Build and deployment**, choose **Deploy from a branch**.
3. Select branch **main** and folder **/docs**, then click **Save**.
4. Wait for the Pages deployment to finish. GitHub will show the published URL in the Pages settings.

For this repository, the expected site URL is <https://shaideralawi-ui.github.io/uob-library/>. The catalog is the home page, and the animated presentation is at `/slides/`. Leave **Custom domain** blank unless you own and have configured a domain.

GitHub Pages serves the static demo only. It does not run the Go server or provide a database, so the demo features do not save user or circulation data.

## Project structure

```text
app/                    Student-facing catalog demo
  assets/               App styles and scripts
presentation/           Animated library scenario slides
  assets/               Presentation styles and scripts
docs/                   Static files for GitHub Pages
main.go                 Go web server
preview.mjs             Optional Node.js preview server
start.bat               Windows Go launch script
```

## Technology

- HTML, CSS, and JavaScript
- Go standard library (`net/http`) for local serving
- GitHub Pages for the static hosted demo
=======
A university library scenario presentation and student-facing app sample.

>>>>>>> 983b47c62e199c2c94f62014c5162ffc2b2336f9
