APP_NAME=gobid
REGION=us-east-1
VPC_ID=vpc-0887b4a025cf8cf30

create-sg:
	@if ! aws ec2 describe-security-groups --filter "Name=group-name, Values=$(SG_NAME)" --region $(REGION) --query "SecurityGroups[*].GroupId" --output text | grep -qE 'sg-'; then \
		echo "Creating security group $(SG_NAME)..."; \
		aws ec2 create-security-group \
			--group-name $(SG_NAME) \
			--description "Allow Postgres for Gobid" \
			--vpc-id $(VPC_ID) \
			--region $(REGION); \
	else \
		echo "Security group $(SG_NAME) already exists."; \
	fi;
	SG_ID=$$(aws ec2 describe-security-groups --filter "Name=group-name, Values=$(SG_NAME)" --regiom $(REGION) --query "SecurityGroups[*].GroupId" --output text); \
	echo "Authorizing port 5432 on SG $$SG_ID..."; \
	aws ec2 authorize-security-group-ingress \
		--group-id $$SG_ID \
		--protocol tcp \
		--port 5432 \
		--cidr 0.0.0.0/0 \
		--region $(REGION); || echo "Ingress rule already exists or failed silently."