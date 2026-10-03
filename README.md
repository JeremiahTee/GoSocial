# GoSocial

A REST API for a small social network, written in Go on PostgreSQL.

## What it does

- Users: registration with bcrypt-hashed passwords, profiles, follow and unfollow
- Posts: create, read, update, delete, with tags and comments
- Feed: posts from the people a user follows, with pagination, sorting and filtering
- Optimistic concurrency on post updates, using a version column
- GIN indexes for search over titles, tags and comments
- Query timeouts on every database call
- Request validation, consistent JSON error responses, structured logging with zap
- Swagger docs generated from handler annotations

## Stack

Go 1.23 · chi · PostgreSQL 16 · golang-migrate · swaggo · zap · air

## Layout

    cmd/api          HTTP handlers, routing, middleware
    cmd/migrate      SQL migrations and the seed command
    internal/store   repository layer (users, posts, comments, followers)
    internal/db      connection pool and seed data
    docs             generated Swagger spec

## Status

A learning project. Registration is in place; token authentication and
email invitations are in progress.
