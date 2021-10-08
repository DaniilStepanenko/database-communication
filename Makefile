test-db: data/dvdrental
	#docker-compose up -d db

data/dvdrental: data/dvdrental.tar
	mkdir -p ./data/dvdrental
	tar xvf --directory=./data/dvdrental/ ./data/dvdrental.tar

data/dvdrental.tar:
	unzip -d data/ data/dvdrental.zip
