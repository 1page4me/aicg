-- Create database if it doesn't exist
SELECT 'CREATE DATABASE aicg'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'aicg');

-- Connect to the database
\c aicg;

-- Create user if it doesn't exist
DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'postgres') THEN
      CREATE USER postgres WITH PASSWORD 'your_password';
   END IF;
END
$do$;

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE aicg TO postgres;
ALTER USER postgres WITH SUPERUSER; 