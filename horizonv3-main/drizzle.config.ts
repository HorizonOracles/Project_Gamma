import { defineConfig } from "drizzle-kit";
import { join } from 'path';

const dbPath = process.env.DATABASE_PATH || join(process.cwd(), 'local.db');

export default defineConfig({
  out: "./migrations",
  schema: "./shared/schema.ts",
  dialect: "sqlite",
  dbCredentials: {
    url: dbPath,
  },
});
