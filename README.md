# Finix

这是一个面相小型团队的资金流动管理平台

## 部署方式

1. 确认安装 Golang, Python 与 node.js 环境

```bash
# 安装 Golang
sudo apt-get install golang-go

# 安装 node.js
sudo apt-get install nodejs

# 安装 Python 3.x
sudo apt-get install python3
```

2. 克隆项目到本地

```bash
git clone https://github.com/Finix-Protocol/Finix.git
```

> 注意, 目前项目正处于开发阶段, 调试请克隆 dev 分支
>
> ```bash
> git clone -b dev https://github.com/Finix-Protocol/Finix.git
> ```
>
> 3. 安装依赖

```bash
cd frontend # 进入前端目录
npm install # 安装前端依赖
npm run dev # 启动前端服务器
# 或者执行 yarn dev
```

> 注意: 此时可能需要再打开一个终端

```bash
# 请确保您正在项目根目录

python make.py run # 运行自动化脚本
```

> 注意, 根据操作系统不同, 命令可能不同, 请自行修改, 有疑问请提 issue, 谢谢!
