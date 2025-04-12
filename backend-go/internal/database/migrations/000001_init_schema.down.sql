-- Drop indexes
DROP INDEX IF EXISTS idx_progress_category;
DROP INDEX IF EXISTS idx_progress_user;
DROP INDEX IF EXISTS idx_results_quiz;
DROP INDEX IF EXISTS idx_results_user;
DROP INDEX IF EXISTS idx_answers_question;
DROP INDEX IF EXISTS idx_questions_quiz;
DROP INDEX IF EXISTS idx_quizzes_difficulty;
DROP INDEX IF EXISTS idx_quizzes_category;
DROP INDEX IF EXISTS idx_users_provider;
DROP INDEX IF EXISTS idx_users_email;

-- Drop tables in correct order due to foreign key constraints
DROP TABLE IF EXISTS user_progress;
DROP TABLE IF EXISTS quiz_results;
DROP TABLE IF EXISTS answers;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS quizzes;
DROP TABLE IF EXISTS users;

-- Drop UUID extension
DROP EXTENSION IF EXISTS "uuid-ossp"; 