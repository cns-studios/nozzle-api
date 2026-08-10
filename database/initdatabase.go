package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const migrationSQL = `
do $$begin
  create schema if not exists public;
  
  create table if not exists users (
    id serial primary key not null,
    discord bigint null unique,
    github int null
  );
  
  if not exists (
    select 1 from information_schema.columns 
    where table_schema = 'public' 
    and table_name = 'users' 
    and column_name = 'discord'
  ) then
    alter table users add column discord bigint null unique;
  end if;
  
  if not exists (
    select 1 from information_schema.columns 
    where table_schema = 'public' 
    and table_name = 'users' 
    and column_name = 'github'
  ) then
    alter table users add column github int null;
  end if;
  
  create table if not exists audit (
    id bigserial primary key not null,
    userid int not null references users(id),
    action varchar not null,
    details jsonb not null,
    timestamp bigint not null
  );
  
  if not exists (
    select 1 from information_schema.columns 
    where table_schema = 'public' 
    and table_name = 'audit' 
    and column_name = 'userid'
  ) then
    alter table audit add column userid int not null references users(id);
  end if;
  
  if not exists (
    select 1 from information_schema.columns 
    where table_schema = 'public' 
    and table_name = 'audit' 
    and column_name = 'action'
  ) then
    alter table audit add column action varchar not null;
  end if;
  
  if not exists (
    select 1 from information_schema.columns 
    where table_schema = 'public' 
    and table_name = 'audit' 
    and column_name = 'details'
  ) then
    alter table audit add column details jsonb not null;
  end if;
  
  if not exists (
    select 1 from information_schema.columns 
    where table_schema = 'public' 
    and table_name = 'audit' 
    and column_name = 'timestamp'
  ) then
    alter table audit add column timestamp bigint not null;
  end if;
  
exception when others then
  raise notice 'Migration completed or already applied: %', sqlerrm;
end $$;
`

func RunMigration(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, migrationSQL)
	if err != nil {
		return fmt.Errorf("migration execution failed: %w", err)
	}
	return nil
}
