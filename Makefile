migrate:
	goose -dir migrations mysql "root:@/realcast-auth" up