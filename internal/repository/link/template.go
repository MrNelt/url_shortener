package link

const (
	saveRequest = `
		INSERT INTO links(short_suffix, link, expiration_date) 
			VALUES ($1, $2, $3)
			RETURNING id;
	`

	selectBySuffixRequest = `
		SELECT id, short_suffix, url, clicks, expiration_date FROM links 
			WHERE short_suffix=$1;
	`

	selectByLinkRequest = `
		SELECT id, short_suffix, url, clicks, expiration_date FROM links  FROM links 
			WHERE link=$1;
	`

	selectByIDRequest = `
		SELECT id, short_suffix, url, clicks, expiration_date FROM links  FROM links 
			WHERE id=$1;
	`

	deleteByIDRequest = `
		DELETE FROM links 
			WHERE id=$1;
	`

	incrementClicksBySuffixRequest = `
		UPDATE links
			SET clicks = clicks+1
			WHERE short_suffix=$1;
	`
)
