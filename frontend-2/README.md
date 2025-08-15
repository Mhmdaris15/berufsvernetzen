# Berufsvernetzen Frontend

This is the frontend for the Berufsvernetzen application, built with Next.js, React, TypeScript, and Tailwind CSS.

## Prerequisites

Before you begin, ensure you have the following installed:
- [Node.js](https://nodejs.org/) (LTS version recommended)
- [npm](https://www.npmjs.com/) (comes with Node.js)

## Environment Variables

This project uses environment variables for configuration. Create a `.env.local` file in the `frontend-2` directory and add the following variables.

```
# Clerk credentials (https://clerk.com/)
CLERK_SECRET_KEY=
NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=
NEXT_PUBLIC_CLERK_SIGN_IN_URL=/sign-in

# Optional
DATABASE_URL=
LOGTAIL_SOURCE_TOKEN=
NEXT_PUBLIC_APP_URL=http://localhost:3000
```

## Getting Started

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd berufsvernetzen/frontend-2
    ```

2.  **Install dependencies:**
    ```bash
    npm install
    ```

3.  **Run the development server:**
    ```bash
    npm run dev
    ```
    Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## Available Scripts

In the `frontend-2` directory, you can run several commands:

- `npm run dev`: Runs the app in development mode.
- `npm run build`: Builds the app for production.
- `npm start`: Starts a production server.
- `npm test`: Runs unit tests with Vitest.
- `npm run test:e2e`: Runs end-to-end tests with Playwright.
- `npm run lint`: Lints the code using ESLint.
- `npm run format`: Formats the code with Prettier.

## Docker Usage

The application can also be run using Docker Compose.

1.  **Build and start the container:**
    From the `frontend-2` directory, run:
    ```bash
    docker-compose up --build
    ```
    This will build the Next.js application image and start the container. The application will be available at [http://localhost:3000](http://localhost:3000).
