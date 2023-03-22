# Appchain-explorer setup
Appchain-explorer works over existing blockscout database and uses redis instance for addiotional indexing,  so we should install all 4 components one by one.

#### Postgress database
Filled by blockscout service,  used in appchain-explorer directly
It is important to use Postgress version 3.16,  in other case blockscout compatibility issues are possible
Command line to start is 
`postgres -c 'max_connections=1000'`
Docker image `postgres:13.6` can be used

#### Blockscout service
Detailed deployment instructions: <https://docs.blockscout.com/for-developers/manual-deployment>


1. Provide Postgress database and RPC node connections with environment variables:

```
export DATABASE_URL=<postgresql://user:password@localhost:5432/blockscout>
export ETHEREUM_JSONRPC_VARIANT = 'geth'
export BLOCK_TRANSFORMER = 'clique'
export COIN = Your_Coin_Symbols
export COIN_NAME = Your_Coin_Name
export ETHEREUM_JSONRPC_HTTP_URL = <http://host.docker.internal:23611/>
export ECTO_USE_SSL = 'false'
```

2. Get Blockscout sources    

```
git clone <https://github.com/blockscout/blockscout>
cd blockscout
```

3. Install Mix dependencies and compile them 

```
mix do deps.get, local.rebar --force, deps.compile
```

4. Generate a new secret_key_base for the DB by setting a corresponding ENV var:

export SECRET_KEY_BASE='New_Generated_key'
In order to generate a new secret_key_base run `mix phx.gen.secret`

5.Create database and start service:
```
mix do ecto.create, ecto.migrate, phx.server'
```
Docker image `blockscout/blockscout:4.1.5` can be used

#### Redis instance
Command to start:
```
redis-server --requirepass Your_Password_Here
```
Docker image `redis` can be used

#### Appchain-explorer
Use Docker image `ankrnetwork/sidechain-explorer:v0.0.3`
Connect it with installed dependencies with environment vars:
 DATABASE_POSTGRES_URL: "<postgresql://user:password@localhost:5432/blockscout?sslmode=disable">
 DATABASE_REDIS_URL: "redis_host:redis_port"
 DATABASE_REDIS_PASSWORD: "redis_passord"
 GATEWAY_HTTP_ADDRESS: ":port"
 STAKING_ETH1_URL: "address_to_staking_service"
 REACT_APP_API_ENDPOINT: "/"
 REACT_APP_DEFAULT_NETWORK: "default"



