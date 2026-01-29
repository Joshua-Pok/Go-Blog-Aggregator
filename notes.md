<!--markdownlint-disable-->


# Postgres

Postgres itself is a server. It listens for requests on a port (default is 5432) and responds to those requests


sudo -iu postgres initdb -D /var/lib/postgres/data


We need to do this when installing postgres to initialize a user called postgres and store data files in /var/lib/postgres/data


-iu simply means run command as postgres system user


we need to do sudo -iu postgres to and then psql to access postgres shell


# Goose

We store versions of our db in schema

<db_version><change>.sql


we can just use goose up and down to move between versions

# RSS


RSS stands for Really simple syndication and is a way to get the latest content from a website in a structured format. Most content websites have a RSS feed


RSS is a specified structure of XML


We need to unmarshal the document into a struct
