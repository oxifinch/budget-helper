# Budget Helper (working title)

This web app is my final exam project as a web developer student at Glimåkra Folkhögskola, written in Go with a strong focus on back-end development. 

## Overview
### Core functionality
Budget Helper aims to provide a easier and more visually appealing alternative to using a spreadsheet to keep track of personal finances, while still keeping the process relatively simple, private and manual. The user will be able to allocate a monthly budget, log transactions and organize them by categories, and keep track of recurring expenses and income using a simple and easy to use interface in the browser. User authentication will use only a username and password, and all information will be kept in a SQLite database, which the user should be able to export as a CSV file for backups. 

### Features
#### Must have
- [ ] User authentication with username and password
- [ ] Session management so that users don't have to sign in again every time
- [ ] Server-side rendering with HTMX
- [ ] Clear MVC structure, separation of database API and request handling
- [ ] User gets data ONLY belonging to the authenticated user
- [ ] User can create a monthly budget, and allocate money for each category if they wish 
- [ ] User can add new transactions and assign a category to them
- [ ] User can see some visual representation of their monthly spending 
- [ ] Clean and consistent UI styling

#### Nice to have
- [ ] User can see a chart showing their remaining funds over the specified period
- [ ] User can export their budget sheet as a CSV file
- [ ] User can change their username and password
---

### Tech stack
The app will be written mainly in Go, handling everything including authentication, sessions, HTTP requests and database connection. Server-side rendering will be done with HTMX, also served via Go. The database used will be SQLite. The goal of this project, besides learning, is to have as few dependencies as possible, using only external Go libraries when necessary, but doing most of the work using only standard library functions. 

### Learning goals
I want to learn more in detail about how sessions and HTTP requests work, and so, to the extent that the project time allows me, I want to build as much as I can without relying on external libraries. I also want to explore other options beyond the usual HTML, CSS and JavaScript + framework tech stack, and see how much of a full-stack application I can make with almost exclusively back-end technologies. 
