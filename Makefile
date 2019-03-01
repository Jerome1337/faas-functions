deploy:
	faas-cli build -f functions.yml
	faas-cli push -f functions.yml
	faas-cli deploy -f functions.yml
