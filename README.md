# Save My RPG Server

## Description

Save My RPG is a desktop application that allows users to share their game saves with a group of friends. Originally built for Baulders Gate 3, this application could easily be adapted to a generic file sharing application.

<b>UPDATE</b>

Due to cost, the functionality of this software will cease. 

## Technical Details

* Go
* PostgreSQL
* Docker
* Bunny CDN

### Code Details

* `src`
    * `dal` Data Access Layer for PostGreSQL
    * `smrpg` Server Code
        * `cdn.go` Interface code for Bunny CDN
    * `config.json` Config File for server
* `Dockerfile.production` Used to build  savemyrpg server
* `db_initialize.sql` Used to Initialize the PostgreSQL database  
* `<TOKEN>` is used in code to represent a code secret/token


## Motivation

I wanted a way for my friends and I to be able to mess with our characters while the host of the game was not available, but in truth I wanted the practice of building a app from start to finish.

## Installation

https://savemyrpg.cloud

## Misc

If your interested in the client code, please follow the link below
<br>
https://github.com/bshafer93/SaveMyRPGClient
