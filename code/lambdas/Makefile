create-function:
	kubeless function deploy $(NAME) --from-file $(NAME)/$(NAME).js --dependencies $(NAME)/package.json --label app=$(NAME) --runtime nodejs8 --handler handler.main --dryrun -o yaml > ../../resources/lambdas/$(NAME).yaml

expose-function:
	../../scripts/create-api.sh "../../resources/lambdas/$(NAME)-api.yaml" $(NAME)
