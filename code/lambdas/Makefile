BRANCH := $(shell git branch --show-current)

create-function:
	kubeless function deploy $(NAME) --namespace $(BRANCH) --from-file $(NAME)/$(NAME).js --dependencies $(NAME)/package.json --label app=$(NAME) --runtime nodejs8 --handler handler.main --dryrun -o yaml > ../../deployment/lambdas/$(NAME).yaml

expose-function:
	../../scripts/create-api.sh "../../deployment/lambdas/$(NAME)-api.yaml" $(NAME) $(BRANCH)
