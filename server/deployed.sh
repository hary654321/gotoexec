#!/bin/bash
###
 # @Description: 
 # @Version: 2.0
 # @Autor: ABing
 # @Date: 2024-06-17 16:49:08
 # @LastEditors: lhl
 # @LastEditTime: 2024-06-20 11:17:36
### 
go build -o mmkz main.go
scp /mnt/hgfs/go-pro/muma/server/mmkz root@172.16.130.160:/var/www/muma
scp /mnt/hgfs/go-pro/muma/server/conf.toml root@172.16.130.160:/var/www/muma
