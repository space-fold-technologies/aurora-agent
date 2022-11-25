#!/usr/bin/env bash

protoc -I protos/ protos/agents.proto --go_out=app/domain/agents