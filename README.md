# shelfer API
書籍管理をするWebアプリのサーバ側です。  
ログイン認証やレンタル情報、書籍情報などをAPIとして提供します。

## Required
```
- jwt-go
- cors
- crypto
- gin
- gorm
- gotoenv
```

## Install
1. $ git clone
2. custome your environment  
    自身の情報に書き換える  
    2.1 $ cp .env.sample .env  

    .envに書き換えたdb情報などを書き換える  
    2.2 $ cp docker-compose.yml.sample docker-compose.yml    
3. $ docker-compose build
4. $ docker-compose up -d

## その他
- docker-compose down (立ち上げているコンテナを落とします)