HOSTNAME=registry.terraform.io
NAMESPACE=statsig-io
NAME=statsig
BINARY=terraform-provider-${NAME}
OS_ARCH=darwin_arm64
VERSION=`cat statsig/version`

default: install

version:
	echo ${VERSION}

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

setup-test-projects:
	cd tests/setup && go install && go run main.go && cd ../..

testacc: install
	TF_ACC=1 TIER=$(TIER) go test ./tests -v ./... -timeout 120m

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	TF_ACC=1 go test $(TEST) -v ./... -sweep=all

# Overwrites internal/provider_code_spec.json, which carries six hand-added
# stringplanmodifier.UseStateForUnknown() plan_modifiers blocks: on `id` for
# dynamic_config, experiment, gate, metric and segment, and on `key` for keys.
# Without them the resource id is unknown in the plan whenever the configuration
# does not set it, so the provider sends an empty id and updates and deletes
# address the collection endpoint instead of the resource. That is the reported
# bug this provider fix exists to close. Neither internal/openapi_spec.json nor
# internal/generator_config.yml can express a plan modifier, so
# tfplugingen-openapi cannot re-emit them, and JSON accepts no comment, so the
# spec file cannot carry this warning itself. Re-apply the six blocks by hand
# after running this. TestAccUpdatePathAddressesTheResource and
# TestAccMetricUpdateRefusesAnEmptyId catch the loss; without this note the
# failure reads as unrelated codegen churn.
generate-provider:
	tfplugingen-openapi generate \
    --config internal/generator_config.yml \
    --output internal/provider_code_spec.json \
    internal/openapi_spec.json

# Rewrites internal/*/*_resource_gen.go, which carry eleven hand-added IsNull
# and IsUnknown guards in ToObjectValue. Without them a nested collection the
# configuration never mentions becomes a known empty list, the request then
# carries it, and the Console API reads a present field as an instruction to
# clear the value. Each of the eleven sites carries its own comment in place,
# unlike internal/provider_code_spec.json above.
# TestAccUnrelatedEditOmitsUnspecifiedNestedAttributes and
# TestAccExplicitlyEmptyNestedAttributesAreStillSent catch the loss; without
# this note the failure reads as unrelated codegen churn.
generate-resources:
	tfplugingen-framework generate resources \
		--input internal/provider_code_spec.json \
		--output internal

generate-scaffold:
	tfplugingen-framework scaffold resource \
    --name $(resource) \
    --force \
    --output-dir internal

generate-docs:
	cd tools && go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs && cd .. && tfplugindocs generate