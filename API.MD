# [WIP] [DRAFT] USER API DEFINITION

Var:
|Name|Value|Explanation|
|---|---|---|
|uid|uid|id == 0 stands for current user(self)|
|username|username|username only|
|user|username or email|for login|

## POST /user/
### POST /user/login
Permission: Public
### POST /user/signup/start
Permission: Public
### POST /user/signup/finish
Permission: Public
### POST /user/forget
Permission: Public
### POST /user/admin/new
Permission: AdminOnly

## /user/:uid/
### GET /user/:uid/info
Permission: User

### GET /user/:uid/basic
Permission: User

### GET /user/:uid/profile
Permission: User

### GET /user/:uid/sudo/auth
Permission: SudoAndAdmin

### GET /user/:uid/sudo/totp
Permission: SudoAndAdmin

## GET /user/username/:username/
### GET /user/username/:username/info
Permission: User

### GET /user/username/:username/basic
Permission: User

### GET /user/username/:username/profile
Permission: User

## /user/:uid/
### POST /user/:uid/avatar
Permission: SelfAndAdmin

### POST /user/:uid/profile
Permission: SelfAndAdmin

### POST /user/:uid/sudo/email
Permission: SudoAndAdmin

### POST /user/:uid/sudo/password
Permission: SudoAndAdmin

### POST /user/:uid/sudo/security_email
Permission: SudoAndAdmin

### POST /user/:uid/mod/locked
Permission: ModAndAdmin

### POST /user/:uid/admin/disabled
Permission: AdminOnly

### POST /user/:uid/admin/electrons
Permission: AdminOnly

### POST /user/:uid/admin/mod
Permission: AdminOnly

## DELETE /user/:uid/
### DELETE /user/:uid/sudo/user
Permission: SudoAdnAdmin

