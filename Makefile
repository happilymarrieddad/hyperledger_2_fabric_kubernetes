CURRENT_HYPERLEDGER_VERSION=2.2.1
CURRENT_FABRIC_CA_VERSION=1.4.9
CURRENT_DB=order-manager

install:
	curl -sSL https://bit.ly/2ysbOFE | bash -s -- $(CURRENT_HYPERLEDGER_VERSION) $(CURRENT_FABRIC_CA_VERSION)

install_deps:
	go get -u github.com/pressly/goose/cmd/goose

install_deps.linux: install_deps
	# Kubectl
	curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
	chmod +x ./kubectl
	sudo mv ./kubectl /usr/local/bin/kubectl

	# Minikube
	curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
	sudo dpkg -i minikube_latest_amd64.deb
	rm minikube_latest_amd64.deb

install_deps.osx: install_deps
	brew install kubectl
	brew install minikube

# USAGE - make name=AddSomeDBTable db.migrate.create
db.migrate.create:
	goose -dir api/db/migrations create $(name) sql

db.migrate.up:
	goose -dir api/db/migrations postgres "postgres://postgres:postgres@localhost:5432/${CURRENT_DB}?connect_timeout=180&sslmode=disable" up

db.migrate.down:
	goose -dir api/db/migrations postgres "postgres://postgres:postgres@localhost:5432/${CURRENT_DB}?connect_timeout=180&sslmode=disable" reset
