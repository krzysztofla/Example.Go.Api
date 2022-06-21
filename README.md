# Example.Go.Api
Example go api. Learning sandbox for coding golang api and keeping all the good practices.

# Example Bank Sanbox
Just a simple api to play with go features for api's 
# Database Schema

### Schema definition
Generated by https://dbdiagram.io/d
Populated by https://tableplus.com

![image description](/pkoo.png)
---
```sql
// Creating tables
Table account as A {
  id bigserial [pk, not null] // auto-increment
  owner varchar
  balance bigint
  created_at timestampz [default: `now()`]
  country_code int
  
  indexes {
    owner
  }
}

Table entries {
  id bigserial [pk]
  name varchar
  account_id bigint [ref: > A.id, not null]
  amount bigint
  created_at timestampz [default: `now()`]
  
  indexes {
    account_id
  }
 }


Table transfers {
  id bigserial [pk]
  name varchar
  from_account_id bigint [ref: > A.id, not null]
  to_account_id bigint [ref: > A.id, not null]
  amount bigint
  created_at timestampz [default: `now()`]
  
  indexes {
    from_account_id
    to_account_id
    (from_account_id, to_account_id)
  }
  }

Enum Currency {
  USD 
  EUR
  PLN
}
```

- Run database inside docker container https://hub.docker.com/_/postgres
```docker
docker run --name postgresgobank -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres

//create db

docker exec -it postgresgobank createdb --username=postgres --owner=postgres simple_bank

```

