## GitHub同步协作
1. fork到自己的GitHub
2. clone 到自己本地 `git clone fork后的项目地址`
3. 查看远程仓库信息 `git remote -v`
4. 增加与原项目同步 `git remote add upstream 原项目地址`
5. 查看远程仓库信息 `git remote -v`

### 使fork项目与原项目同步
1. 在本地切换正确分支 `git checkout master`
2. 拉取远程更新并与本地仓库合并 `git pull --rebase upstream master`
3. 将更新推送给fork项目 `git push origin master`

## 分支

### 创建分支
`git checkout -b 分支名`

```
git branch 分支名

git checkout 分支名
```

### 切换分支

`git checkout 分支名`

### 查看分支

`git branch`

### 合并指定分支到当前分支

`git merge 分支名` 

### 删除分支

`git branch -d dev` 
强制删除：`git branch -D dev` 