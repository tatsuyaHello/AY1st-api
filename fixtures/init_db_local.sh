#!/bin/sh

echo "データベースのリセット実行開始"
time (
docker-compose exec mysql mysql -uroot --default-character-set=utf8 -e "DROP DATABASE IF EXISTS \`${DATABASE_NAME}\`;"
docker-compose exec mysql mysql -uroot --default-character-set=utf8 -e "CREATE DATABASE \`${DATABASE_NAME}\` DEFAULT CHARACTER SET utf8;"

echo run sql file fixtures/db_schema.sql
docker-compose exec mysql bash -c "mysql ${DATABASE_NAME} -uroot --default-character-set=utf8 < /var/fixtures/db_schema.sql"

)
echo "データベースのリセット実行完了"
