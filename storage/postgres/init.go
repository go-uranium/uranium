package postgres

func (pg *Postgres) Init() error {

}

var SQLInitUser = `CREATE TABLE IF NOT EXISTS "user"(
    	uid SERIAL PRIMARY KEY NOT NULL,
    	name VARCHAR(30) NOT NULL,     
    	username VARCHAR(20) NOT NULL,
    	electrons int4 NOT NULL,
    	permission int2 NOT NULL,
    	created_at TIMESTAMP NOT NULL
	);
	CREATE TABLE IF NOT EXISTS user_auth(
		uid SERIAL REFERENCES "user" NOT NULL,
		email VARCHAR(320) NOT NULL,
		password CHAR(64) NOT NULL,
		security_email VARCHAR(320) NOT NULL,
		locked TIMESTAMP NOT NULL,
		disabled BOOLEAN NOT NULL
	);
	CREATE TABLE IF NOT EXISTS user_profile(
		uid SERIAL REFERENCES "user" NOT NULL,
		bio VARCHAR(255) NOT NULL,
		
	);
`

func (pg *Postgres) initUser() error {
	_, err := pg.db.Exec(SQLInitUser)
	if err != nil {
		return err
	}
}
