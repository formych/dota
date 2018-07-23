#!/bin/bash

binName=(basename `pwd`)
go build -o release/bin/$binName
