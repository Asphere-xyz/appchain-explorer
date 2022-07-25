#!/bin/bash

PG_URL="postgres://postgres:@rpc.dev-01.bas.ankr.com:7432/blockscout?sslmode=disable"

#mkdir -p shared/entity
#xo schema "$PG_URL" -o ./shared/entity

#xo query "$PG_URL" --append --trim --strip -o shared/entity -T Block <<ENDSQL
#SELECT * FROM "blocks"
#ORDER BY "timestamp" DESC
#LIMIT %%limit uint64%%
#ENDSQL
#
#xo query "$PG_URL" --append --trim --strip -o shared/entity -T Block <<ENDSQL
#SELECT * FROM "blocks"
#WHERE %%number uint64%% >= number
#ORDER BY "timestamp" DESC
#LIMIT %%limit uint64%%
#ENDSQL
#
#xo query "$PG_URL" --append --trim --strip -o shared/entity -T Transaction <<ENDSQL
#SELECT * FROM "transactions"
#ORDER BY "updated_at" DESC
#LIMIT %%limit uint64%%
#ENDSQL
#
#xo query "$PG_URL" --append --trim --strip -o shared/entity -T Transaction <<ENDSQL
#SELECT * FROM "transactions"
#WHERE %%ts string%% <= updated_at
#ORDER BY "updated_at" DESC
#LIMIT %%limit uint64%%
#ENDSQL
#
#xo query "$PG_URL" --append --trim --strip -o shared/entity -T AddressTokenBalance <<ENDSQL
#SELECT * FROM "address_token_balances"
#WHERE address_hash = %%hash []byte%%
#ENDSQL
#
#xo query "$PG_URL" --append --trim --strip -o shared/entity -T TokenTransfer <<ENDSQL
#SELECT * FROM "token_transfers"
#WHERE transaction_hash = %%hash []byte%%
#ENDSQL
#
#xo query "$PG_URL" --append --trim --strip -o shared/entity -T ContractMethod <<ENDSQL
#SELECT * FROM "contract_methods"
#WHERE identifier = %%identifier int32%%
#ENDSQL

xo query "$PG_URL" --append --trim --strip -o shared/entity --one -T TokenTransfer <<ENDSQL
SELECT COUNT(transaction_hash) FROM "token_transfers"
WHERE transaction_hash = %%hash []byte%%
ENDSQL
