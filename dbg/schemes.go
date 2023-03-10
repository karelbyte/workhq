package dbg

const TenantScheme = `CREATE TABLE IF NOT EXISTS tenants (
	id CHAR(36) NOT NULL,
    domain VARCHAR(255) NOT NULL,
	db VARCHAR(255),
	status TINYINT,
	created_at DATETIME NOT NULL DEFAULT current_timestamp(),
	update_at DATETIME NOT NULL DEFAULT current_timestamp(),
	deleted_at DATETIME,
	PRIMARY KEY (id),
	CONSTRAINT uc_tenants UNIQUE (id, domain)
 )
`


// const UsersScheme = `CREATE TABLE IF NOT EXISTS tenants (
// 	id CHAR(36) NOT NULL,
//     fullname VARCHAR(255) NOT NULL,
//     address VARCHAR(255),
//     email VARCHAR(255) NOT NULL,
//     phone VARCHAR(15),
// 	password VARCHAR(255),
// 	token VARCHAR(255),
// 	status TINYINT,
// 	created_at DATETIME NOT NULL DEFAULT current_timestamp(),
// 	update_at DATETIME NOT NULL DEFAULT current_timestamp(),
// 	deleted_at DATETIME,
// 	PRIMARY KEY (id),
// 	CONSTRAINT uc_users UNIQUE (id, email)
//  )
// `

// const ProductsScheme = `CREATE TABLE IF NOT EXISTS products (
// 	id CHAR(36) NOT NULL,
//     code VARCHAR(255) NOT NULL,
// 	name VARCHAR(255) NOT NULL,
//     description VARCHAR(255),
//     url VARCHAR(255),
// 	image VARCHAR(255),
// 	created_at DATETIME NOT NULL DEFAULT current_timestamp(),
// 	update_at DATETIME NOT NULL DEFAULT current_timestamp(),
// 	deleted_at DATETIME,
// 	PRIMARY KEY (id),
// 	CONSTRAINT pc_products UNIQUE (id, code)
//  )
// `

var SchemeEntities = []string{
	TenantScheme,
	// ProductsScheme,
}
