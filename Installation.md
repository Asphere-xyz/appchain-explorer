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

5. Create database and start service:
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
```
DATABASE_POSTGRES_URL: "postgresql://user:password@localhost:5432/blockscout?sslmode=disable"
DATABASE_REDIS_URL: "redis_host:redis_port"
DATABASE_REDIS_PASSWORD: "redis_passord"
GATEWAY_HTTP_ADDRESS: ":port"
STAKING_ETH1_URL: "address_to_staking_service"
REACT_APP_API_ENDPOINT: "/"
REACT_APP_DEFAULT_NETWORK: "default"
```

### docker-compose file example:
```
version: '3.8'

services:
  db-pg:
    image: postgres:13.6
    restart: always
    container_name: 'inst_postgres'
    command: postgres -c 'max_connections=1000'
    environment:
        POSTGRES_PASSWORD: ''
        POSTGRES_USER: 'postgres'
        POSTGRES_HOST_AUTH_METHOD: 'trust'
    volumes:
      - ./postgres-data:/var/lib/postgresql/data

  blockscout-inst:
    depends_on:
      - db-pg
    image: blockscout/blockscout:4.1.5
    restart: always
    container_name: 'inst_blockscout'
    links:
      - db-pg:database
    command: 'mix do ecto.create, ecto.migrate, phx.server'
    extra_hosts:
      - 'host.docker.internal:host-gateway'
    env_file:
      -  ./envs/common-blockscout.env
    logging:
      driver: "json-file"
      options:
        max-size: "2048m"
    environment:
        ETHEREUM_JSONRPC_VARIANT: 'geth'
        BLOCK_TRANSFORMER: 'clique'
        COIN: 'BOMB'
        COIN_NAME: 'BOMB'
        ETHEREUM_JSONRPC_HTTP_URL: "Your RPC endpoint URL"
        DATABASE_URL: postgresql://postgres:@db-pg:5432/blockscout?ssl=false
        ECTO_USE_SSL: 'false'
        SECRET_KEY_BASE: 'Your generated key here'
    ports:
      - 127.0.0.1:4010:4000


  blockscout-redis-inst:
    image: redis
    command: redis-server --requirepass Your_redis_password
    container_name: 'inst_redis'
    extra_hosts:
      - 'host.docker.internal:host-gateway'
    volumes:
      - ./datadir/redis:/data
    restart: always

  blockscout-gateway-inst:
    depends_on:
      - blockscout-inst
    image: ankrnetwork/sidechain-explorer:v0.0.3
    container_name: 'inst_gateway'
    extra_hosts:
      - 'host.docker.internal:host-gateway'
    environment:
      DATABASE_POSTGRES_URL: "postgres://postgres:@db-inst:5432/blockscout?sslmode=disable"
      DATABASE_REDIS_URL: "blockscout-redis-inst:6379"
      DATABASE_REDIS_PASSWORD: "Your_redis_password"
      GATEWAY_HTTP_ADDRESS: ":9000"
      STAKING_ETH1_URL: "Staking_url_here"
      REACT_APP_API_ENDPOINT: "/"
      REACT_APP_DEFAULT_NETWORK: "default"
    ports:
      - 127.0.0.1:9010:9000
    restart: always
```

