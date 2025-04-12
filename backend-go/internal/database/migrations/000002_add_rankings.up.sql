-- Create global_rankings table to track overall user performance
CREATE TABLE global_rankings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    total_score DECIMAL(10,2) DEFAULT 0,
    quizzes_completed INTEGER DEFAULT 0,
    average_score DECIMAL(5,2) DEFAULT 0,
    total_time_spent INTEGER DEFAULT 0, -- in seconds
    rank INTEGER,
    percentile DECIMAL(5,2),
    ranking_period VARCHAR(20) NOT NULL, -- 'weekly', 'monthly', 'all_time'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, ranking_period)
);

-- Create category_rankings table to track performance by category
CREATE TABLE category_rankings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    category VARCHAR(50) NOT NULL,
    total_score DECIMAL(10,2) DEFAULT 0,
    quizzes_completed INTEGER DEFAULT 0,
    average_score DECIMAL(5,2) DEFAULT 0,
    rank INTEGER,
    percentile DECIMAL(5,2),
    ranking_period VARCHAR(20) NOT NULL, -- 'weekly', 'monthly', 'all_time'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, category, ranking_period)
);

-- Create achievements table
CREATE TABLE achievements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50),
    criteria JSONB NOT NULL, -- Store achievement criteria
    icon_url VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create user_achievements table to track earned achievements
CREATE TABLE user_achievements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    achievement_id UUID REFERENCES achievements(id),
    earned_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    progress JSONB, -- Store progress towards achievement
    UNIQUE(user_id, achievement_id)
);

-- Create benchmarks table for category/difficulty combinations
CREATE TABLE benchmarks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category VARCHAR(50) NOT NULL,
    difficulty VARCHAR(20) NOT NULL,
    average_score DECIMAL(5,2) DEFAULT 0,
    median_score DECIMAL(5,2) DEFAULT 0,
    percentile_75 DECIMAL(5,2) DEFAULT 0,
    percentile_90 DECIMAL(5,2) DEFAULT 0,
    total_attempts INTEGER DEFAULT 0,
    average_completion_time INTEGER DEFAULT 0, -- in seconds
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(category, difficulty)
);

-- Add indexes for better query performance
CREATE INDEX idx_global_rankings_score ON global_rankings(total_score DESC);
CREATE INDEX idx_global_rankings_period ON global_rankings(ranking_period, total_score DESC);
CREATE INDEX idx_category_rankings_score ON category_rankings(category, total_score DESC);
CREATE INDEX idx_category_rankings_period ON category_rankings(ranking_period, category, total_score DESC);
CREATE INDEX idx_user_achievements_user ON user_achievements(user_id);
CREATE INDEX idx_benchmarks_category_difficulty ON benchmarks(category, difficulty);

-- Add function to update rankings
CREATE OR REPLACE FUNCTION update_rankings()
RETURNS TRIGGER AS $$
BEGIN
    -- Update global rankings
    WITH ranked_users AS (
        SELECT 
            user_id,
            total_score,
            RANK() OVER (ORDER BY total_score DESC) as new_rank,
            PERCENT_RANK() OVER (ORDER BY total_score DESC) * 100 as new_percentile
        FROM global_rankings
        WHERE ranking_period = NEW.ranking_period
    )
    UPDATE global_rankings gr
    SET 
        rank = ru.new_rank,
        percentile = ru.new_percentile,
        updated_at = CURRENT_TIMESTAMP
    FROM ranked_users ru
    WHERE gr.user_id = ru.user_id
    AND gr.ranking_period = NEW.ranking_period;

    -- Update category rankings
    WITH ranked_categories AS (
        SELECT 
            user_id,
            category,
            total_score,
            RANK() OVER (PARTITION BY category ORDER BY total_score DESC) as new_rank,
            PERCENT_RANK() OVER (PARTITION BY category ORDER BY total_score DESC) * 100 as new_percentile
        FROM category_rankings
        WHERE ranking_period = NEW.ranking_period
    )
    UPDATE category_rankings cr
    SET 
        rank = rc.new_rank,
        percentile = rc.new_percentile,
        updated_at = CURRENT_TIMESTAMP
    FROM ranked_categories rc
    WHERE cr.user_id = rc.user_id
    AND cr.category = rc.category
    AND cr.ranking_period = NEW.ranking_period;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create triggers to automatically update rankings
CREATE TRIGGER update_global_rankings
AFTER INSERT OR UPDATE ON global_rankings
FOR EACH ROW
EXECUTE FUNCTION update_rankings();

CREATE TRIGGER update_category_rankings
AFTER INSERT OR UPDATE ON category_rankings
FOR EACH ROW
EXECUTE FUNCTION update_rankings(); 