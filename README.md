# EVM-Docker
api and EVM

测试分支在master分支中

要将本地修改的代码（修改部分）同步到远程仓库，请遵循以下步骤：

打开终端，使用 cd 命令导航到您的本地仓库：
bash
Copy code
cd /path/to/your/local/repo
将 /path/to/your/local/repo 替换为实际的本地仓库路径。

使用 git status 命令查看已修改的文件。这将显示您尚未暂存的更改：
bash
Copy code
git status
使用 git add 命令暂存您要提交的更改。您可以使用文件名暂存单个文件，或使用以下命令暂存所有更改：
bash
Copy code
git add .
使用 git commit 命令创建一个新的提交，并附上描述更改的提交信息：
bash
Copy code
git commit -m "Your commit message here"
将 Your commit message here 替换为实际的提交信息，以描述您所做的更改。

确保您已经将本地仓库关联到远程仓库。使用以下命令检查：
bash
Copy code
git remote -v
如果您看到 origin 关键字后面跟着您 GitHub 仓库的 URL，则表示已经关联。如果没有，使用以下命令将本地仓库关联到远程仓库：

bash
Copy code
git remote add origin https://github.com/your_username/your_repo.git
将 your_username 和 your_repo 替换为实际的 GitHub 用户名和仓库名。

将您的提交推送到远程仓库：
bash
Copy code
git push origin main
将 main 替换为您要推送的实际分支名称。这可能是 main 或 master，具体取决于您的仓库设置。

现在，您已经将本地修改的代码同步到远程仓库。在 GitHub 上查看仓库，您应该能看到已推送的提交。
