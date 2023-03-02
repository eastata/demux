


# Demux API
API for Demux [github.com/eastata/demux/pkg/demux](https://github.com/eastata/demux/tree/main/pkg/demux)
  

## Informations

### Version

0.0.1

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## Access control

### Security Requirements
  * api_key: []

## All endpoints

###  operations

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /job_submit | [job submit](#job-submit) | Submit the job for summing the list of int64 |
  


## Paths

### <span id="job-submit"></span> Submit the job for summing the list of int64 (*JobSubmit*)

```
POST /job_submit
```

This will submit the job to demux

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| data | `body` | []int64 (formatted integer) | `[]int64` | |  | | Send a json body in a request with a key "data" that must be a list of int64 |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#job-submit-200) | OK | JobUUID |  | [schema](#job-submit-200-schema) |

#### Responses


##### <span id="job-submit-200"></span> 200 - JobUUID
Status: OK

###### <span id="job-submit-200-schema"></span> Schema
   
  

[JobID](#job-id)

## Models

### <span id="job-id"></span> JobID


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Id | uuid (formatted string)| `strfmt.UUID` |  | | in: string | `{\"id\": \"75a9e835-5cd6-4499-bd2a-a066e335b963\"}` |


