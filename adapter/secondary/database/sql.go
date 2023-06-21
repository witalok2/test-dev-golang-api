package database

const (
	SqlListClient = `
		SELECT 
			id,
			name,
			last_name,
			contact,
			address,
			birthday,
			cpf,
			created_at,
			updated_at
		FROM client
		WHERE deleted_at IS NULL
			ORDER BY created_at
			LIMIT $1 OFFSET $2
	`

	SqlGetClientById = `
		SELECT 
			id,
			name,
			last_name,
			contact,
			address,
			birthday,
			cpf,
			created_at,
			updated_at
		FROM client
		WHERE id = $1 AND deleted_at IS NULL
	`

	SqlDeleteClient = `
		UPDATE client
		SET
			name = 'Deleted_' || NOW()::timestamp,
			last_name = '',
			contact = '',
			cpf = null,
			updated_at = NOW(),
			deleted_at = NOW()
		WHERE
			id = $1
			AND deleted_at IS NULL
	`

	SqlCountClient = `SELECT COUNT(1) FROM client WHERE deleted_at IS NULL`
)
