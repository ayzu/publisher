gen:
	oapi-codegen --include-tags articles --package main swagger.json > gen.go

medium-test-list:
	http "https://api.medium.com/v1/users/${MEDIUMUSERID}/publications" Authorization:"Bearer ${MEDIUMAPIKEY}"