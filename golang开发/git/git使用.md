# git教程
## ssh key
```
<!-- 查看配置信息 -->
git config --global --list
<!-- 设置用户名 -->
git config --global user.name "username"
<!-- 设置邮箱 -->
git config --global user.email "email"
<!-- 生成密钥 -->
ssh-keygen -t rsa -C "email"
```
## 常用操作
```
<!-- 初始化仓库 -->
git init
<!-- 添加文件到暂存区  -->
git add . 
<!-- 将暂存区内容添加到仓库 -->
git commit -m "first commit"
<!-- 创建分支 -->
git branch -M dev
<!-- 添加远程仓库 origin 是远程仓库名-->
git remote add Ning-Qing/BetterGolanger https://github.com/Ning-Qing/BetterGolanger.git
<!-- 将本地仓库的的内容与远程仓库内容合并 -->
git push -u Ning-Qing/BetterGolanger dev
```

## 常见问题
