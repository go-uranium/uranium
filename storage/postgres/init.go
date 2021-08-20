package postgres

func (pg *Postgres) Init() error {
	err := pg.initUser()
	if err != nil {
		return err
	}

}

var SQLInitUser = `CREATE TABLE IF NOT EXISTS "user"(
    	uid SERIAL PRIMARY KEY NOT NULL,
    	username VARCHAR(20) NOT NULL,
    	lowercase VARCHAR(20) NOT NULL UNIQUE,
    	electrons int4 NOT NULL,
    	admin int2 NOT NULL,
    	created timestamptz NOT NULL,
    	deleted BOOLEAN NOT NULL
	);
	CREATE TABLE IF NOT EXISTS user_auth(
		uid SERIAL REFERENCES "user" NOT NULL,
		email VARCHAR(320) NOT NULL UNIQUE,
		password bytea(32) NOT NULL,
		security_email VARCHAR(320),
		two_factor BOOLEAN NOT NULL,
		locked BOOLEAN NOT NULL,
		locked_till timestamptz,
		disabled BOOLEAN NOT NULL
	);
	CREATE TABLE IF NOT EXISTS user_profile(
		uid SERIAL REFERENCES "user" NOT NULL,
		name VARCHAR(30) NOT NULL,
		bio VARCHAR(255),
		location VARCHAR(30),
		birthday int4,
		email VARCHAR(320),
		social text NOT NULL
	);
`

func (pg *Postgres) initUser() error {
	_, err := pg.db.Exec(SQLInitUser)
	return err
}

var SQLInitPost = `CREATE TABLE IF NOT EXISTS "user"(
    	uid SERIAL PRIMARY KEY NOT NULL,
    	username VARCHAR(20) NOT NULL,
    	lowercase VARCHAR(20) NOT NULL UNIQUE,
    	electrons int4 NOT NULL,
    	admin int2 NOT NULL,
    	created timestamptz NOT NULL,
    	deleted BOOLEAN NOT NULL
	);
	CREATE TABLE IF NOT EXISTS user_auth(
		uid SERIAL REFERENCES "user" NOT NULL,
		email VARCHAR(320) NOT NULL UNIQUE,
		password bytea(32) NOT NULL,
		security_email VARCHAR(320),
		two_factor BOOLEAN NOT NULL,
		locked BOOLEAN NOT NULL,
		locked_till timestamptz,
		disabled BOOLEAN NOT NULL
	);
	CREATE TABLE IF NOT EXISTS user_profile(
		uid SERIAL REFERENCES "user" NOT NULL,
		name VARCHAR(30) NOT NULL,
		bio VARCHAR(255),
		location VARCHAR(30),
		birthday int4,
		email VARCHAR(320),
		social text NOT NULL
	);
`

func (pg *Postgres) initPost() error {
	_, err := pg.db.Exec(SQLInitPost)
	return err
}
