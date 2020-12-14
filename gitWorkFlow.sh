#!/bin/bash

# 获取当前分支
function getCurrentBranch() {
	currentBranch=$(git branch | grep "*")
	currentBranch=${currentBranch/* /}
	echo ${currentBranch}
	return 0
}

# 舍弃本地修改，将本地对应的分支，合到指定版本
function mergeToAppointBranch() {
	echo $1
	desBranch=$1
	if [[ -z ${desBranch} ]]; then
	        echo "分支输入错误${desBranch}"
	        return 1
	fi
	echo "要合并到这个${desBranch}分支"

	forbid_branch=("master")
	if [[ "${forbid_branch[@]}" =~ ${desBranch} ]]; then
	        echo "此分支禁止合并";
	        return 1
	fi
	currentBranch=$(getCurrentBranch)
	echo "当前分支为${currentBranch}"

	# 删除本地修改
	echo "准备删除本地代码的修改，请确认修改代码已上传到对应分支"
	sleep 3s

	# 不包括删除、新增的文件
	git stash && git stash drop
	echo "删除本地修改成功"
	git branch -D ${desBranch} && git pull
	echo "拉取成功"
	git checkout -b ${desBranch} origin/${desBranch} && git pull origin ${currentBranch} && git push origin ${desBranch} && git checkout ${currentBranch}
	return 0
}

# merge冲突后，继续提交
function mergeContinue() {
	branch=$(getCurrentBranch)
	git merge --continue
	git push origin ${branch}	
	return 0
}

# rebase 指定分支
function rebaseToAppointBranch() {
	echo $1
        desBranch=$1
        if [[ -z ${desBranch} ]]; then
                echo "分支输入错误${desBranch}"
                return 1
        fi
        echo "当前分支rebase这个${desBranch}分支"

	curBranch=$(getCurrentBranch)
 	echo "当前分支为${curBranch}"

 	# 不包括删除、新增的文件
        git stash && git stash drop
        echo "删除本地修改成功"

	git checkout ${desBranch} && git pull && git checkout ${curBranch} && git rebase ${desBranch} && git push -f origin ${curBranch}
	return 0
}

# rebase冲突后，继续提交
function rebaseContinue() {
	curBranch=$(getCurrentBranch)
	echo ${curBranch}
	git rebase --continue
	git push -f origin ${curBranch}
	return 0
}

function main() {
	case $1 in
		"mc")
			# 继续合并
			mergeContinue
		;;
		"mt")
			# 合并当前分支到指定分支
			mergeToAppointBranch $2
		;;
		"ct")
			echo "branch is $(getCurrentBranch)"
		;;
		"rt")
			rebaseToAppointBranch $2
		;;
		"rc")
			# rebase continue
			rebaseContinue
		;;
		*)
			echo "输入指令，请重新输入"
		;;
	esac
}

main $1 $2
