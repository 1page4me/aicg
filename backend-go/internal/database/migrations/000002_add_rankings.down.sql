-- Drop triggers first
DROP TRIGGER IF EXISTS update_global_rankings ON global_rankings;
DROP TRIGGER IF EXISTS update_category_rankings ON category_rankings;

-- Drop function
DROP FUNCTION IF EXISTS update_rankings();

-- Drop indexes
DROP INDEX IF EXISTS idx_benchmarks_category_difficulty;
DROP INDEX IF EXISTS idx_user_achievements_user;
DROP INDEX IF EXISTS idx_category_rankings_period;
DROP INDEX IF EXISTS idx_category_rankings_score;
DROP INDEX IF EXISTS idx_global_rankings_period;
DROP INDEX IF EXISTS idx_global_rankings_score;

-- Drop tables in correct order due to foreign key constraints
DROP TABLE IF EXISTS benchmarks;
DROP TABLE IF EXISTS user_achievements;
DROP TABLE IF EXISTS achievements;
DROP TABLE IF EXISTS category_rankings;
DROP TABLE IF EXISTS global_rankings; 