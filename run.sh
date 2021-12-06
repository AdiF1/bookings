#!/bin/bash

go build -o bookings cmd/web/*.go
./bookings -db_name=bookings2 -db_user=postgres -cache=false -production=false