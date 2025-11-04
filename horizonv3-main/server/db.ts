// DegenArena Database Connection - SQLite
import { drizzle } from 'drizzle-orm/better-sqlite3';
import Database from 'better-sqlite3';
import * as schema from "@shared/schema";
import { join } from 'path';

// Use local.db in development, or specified DATABASE_PATH
const dbPath = process.env.DATABASE_PATH || join(process.cwd(), 'local.db');

console.log(`Using SQLite database at: ${dbPath}`);

const sqlite = new Database(dbPath);

// Enable WAL mode for better performance
sqlite.pragma('journal_mode = WAL');

export const db = drizzle(sqlite, { schema });
