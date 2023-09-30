package link

const (
	saveRequest = `
		INSERT INTO links(short_suffix, url, expiration_date) 
			VALUES ($1, $2, $3)
			RETURNING id::text;
	`

	selectBySuffixRequest = `
		SELECT id::text, short_suffix, url, clicks, expiration_date FROM links 
			WHERE short_suffix=$1 AND expiration_date>=NOW();
	`

	selectByLinkRequest = `
		SELECT id::text, short_suffix, url, clicks, expiration_date FROM links
			WHERE link=$1 AND expiration_date>=NOW();
	`

	selectByIDRequest = `
		SELECT id::text, short_suffix, url, clicks, expiration_date FROM links
			WHERE id::text=$1;
	`

	deleteByIDRequest = `
		DELETE FROM links 
			WHERE id::text=$1;
	`

	incrementClicksBySuffixRequest = `
		UPDATE links
			SET clicks = clicks+1
			WHERE short_suffix=$1 AND expiration_date>=NOW();
	`
)
