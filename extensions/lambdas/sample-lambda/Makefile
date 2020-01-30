NAMESPACE?=dev
NAME?=sample-lambda

create-function:
	kubeless function deploy $(NAME) --namespace $(NAMESPACE) --from-file $(NAME).js --dependencies ./package.json --label app=$(NAME) --runtime nodejs8 --handler handler.main --dryrun -o yaml > ./deployment/$(NAME).yaml

expose-function:
	../../../scripts/create-api.sh "./deployment/$(NAME)-api.yaml" $(NAME) $(NAMESPACE)
