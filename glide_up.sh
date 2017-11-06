#!/bin/bash

sudo chown $USER:$USER -R ./pg_data
glide up
sudo chown 990:root -R ./pg_data
