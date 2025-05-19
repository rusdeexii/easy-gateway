DO
$$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'pguser') THEN
        CREATE USER pguser WITH PASSWORD 'pgpass';
    END IF;
END
$$;
