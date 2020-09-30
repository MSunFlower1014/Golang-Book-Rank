# Golang-Book-Rank
golang语言的起点排行


linux中发布流程：  
1. 切换到项目main方法目录 src/QidianRank.go  
2. 设置生成linux可执行文件，shell中执行 -> set GOARCH=amd64 ; set GOOS=linux;  
3. 编译 -> go build QidianRank.go  
4. 将文件 QidianRank 复制到linux主机  
5. 增加执行权限 -> chmod 777 QidianRank  
6. 执行 -> nohup ./QidianRank go > book.log &  