# ITDB — IT 完整手册

> **文档说明**：本文档由原 `README.md` 重命名而来，保留了全部细节（完整接口清单、49 张表结构、Nginx 与 HTTPS 示例、源码编译部署、旧库迁移与排障），作为项目的完整手册使用。
>
> 项目概览、功能说明与部署入口请回到 [README.md](../README.md)。部署方式以 README 中给出的 **Docker 部署** 与 **Release 二进制部署** 两种为准；本文第三、四章内的源码编译与直接运行步骤属于过程化细节，仅作参考。

基于 [zyx3721/itdb](https://github.com/zyx3721/itdb/) 重新进行全面优化的 IT 资产管理系统，使用 Go + React 前后端分离架构重新实现。支持硬件设备、软件许可、合同、单据、文件、机架、地点等资产的全生命周期管理。

## 目录

- [一、项目介绍](#一项目介绍)
- [二、本地开发快速启动](#二本地开发快速启动)
- [三、Docker Compose 快速部署（推荐）](#三docker-compose-快速部署推荐)
- [四、生产环境部署](#四生产环境部署)
- [五、API 文档](#五api-文档)
- [六、数据库说明](#六数据库说明)
- [七、常见问题](#七常见问题)
- [八、安全建议](#八安全建议)
- [九、许可证](#九许可证)
- [十、致谢](#十致谢)
- [十一、联系方式](#十一联系方式)

# 一、项目介绍

## 1.1 项目简介

ITDB 是面向企业与机房场景的 IT 资产全生命周期管理平台，由 Go 后端与 React 控制台组成，覆盖硬件、软件、单据、代理、文件、合同、地点与机架八类资产，以及配套的资料字典、标签打印、统计报表与资产导航。

平台提供细粒度权限控制（内置三角色与自定义角色共 44 项权限）、本地与 AD/LDAP 双模式登录、找回密码邮件流程、六大模块的中文审计日志，以及手动/定时数据备份与旧项目数据库兼容导入。后端使用 SQLite（纯 Go 驱动，无 CGO 依赖）保存业务数据与系统配置，JWT 签名密钥与 LDAP 绑定密码等敏感信息以主密钥加密后落库。

平台面向需要统一 IT 台账的运维团队。控制台读取数据库当前状态展示资产与统计；资产增删改、系统配置、备份与导入等变更由后端执行权限校验并写入审计记录，业务数据与上传文件均保存在本机数据目录，可整体备份与迁移。

## 1.2 项目预览

|               项目登录页                |
| :-------------------------------------: |
| ![login](../.github/images/itdb-login.jpg) |

|                 项目首页                 |
| :--------------------------------------: |
| ![home](../.github/images/itdb-home.jpg) |

## 1.3 核心功能

- **资产全生命周期管理**：硬件、软件、单据、代理、文件、合同、地点、机架八类资源；硬件支持序列号、网络信息、维保与成本记录，并可关联软件、单据、合同、文件与内部硬件；合同支持类型/子类型、续签与事件历史；地点支持平面图上传与区域热区标注；机架支持 U 位与正反面可视化。
- **资料管理**：硬件类型、合同类型（含子类型）、状态类型（自定义颜色）、文件类型、所属部门、标记六类字典，支持 Excel 模板下载、批量导入与导出；内置数据受编号与名称双重保护；标记可关联硬件与软件并统计关联数。
- **用户认证与访问控制**：本地密码与 AD/LDAP 双模式登录、JWT 会话、找回密码邮件流程；内置 admin/operator/viewer 三角色，自定义角色按资产管理、资料管理、打印标签、统计报表、资产导航、审计日志、系统设置共 44 项权限划分，前端隐藏无权操作、后端逐接口校验。
- **审计日志**：登录注销、资产增删改、字典维护、系统配置、备份导入、标签打印等操作按模块、操作、目标、结果与详情全量记录，支持搜索、筛选与导出；无实际修改的保存不写入日志，重名创建等业务规则拒绝不产生失败噪音。
- **数据备份与迁移**：手动备份可勾选按数据库实际引用的文件一并打包 zip；定时备份按五段 Cron 计划执行并按保留天数自动清理；导入支持 .db 与 .zip，兼容旧项目数据库自动转换（仅迁移资产、资料与用户数据，系统配置保持当前默认）。
- **标签打印**：QR 码标签设计器与多种标签纸预设，支持批量预览与打印。
- **统计与导航**：仪表盘资产概览、内置统计报表（支持导出 XLSX/XLS/CSV/TXT）、资产导航树。
- **系统配置**：品牌标识、找回密码限流与登录失败锁定等安全时效、企业微信扫码有效期、定时备份参数、用户/用户组/角色、AD/LDAP 与企业微信认证、邮件通知配置，并提供认证连通性与邮件发送测试。
- **API 文档**：集成 swag 与 Swagger UI，可通过 `/swagger/index.html` 查看接口、参数、鉴权和响应定义。

## 1.4 数据与安全边界

SQLite 数据库（`data/itdb.db`）是平台业务数据的事实来源，上传文件保存于 `data/files`，手动与定时备份输出到 `data/backups`。JWT 签名密钥持久化于 `system_secrets` 表，LDAP 绑定密码以主密钥加密后存储，日志输出不包含中文以外的敏感内容。

公开访问仅限健康检查、登录、品牌信息与找回密码流程接口；其余接口均需 Bearer Token，并由后端权限中间件按 44 项权限逐接口校验。登录（含失败）、资产与配置变更、备份导入等关键操作均写入审计日志，删除被引用数据、创建重名等业务规则拒绝由前端明确提示、不产生审计噪音。

## 1.5 技术栈

### 1.5.1 后端

- **语言**：Go 1.25+
- **HTTP**：go-chi/chi v5
- **数据库**：SQLite（modernc.org/sqlite 纯 Go 驱动，无 CGO 依赖）
- **认证与加密**：JWT（golang-jwt/v5）、AD/LDAP（go-ldap/ldap v3）、bcrypt 密码散列与 AES 敏感配置加密（golang.org/x/crypto）
- **数据导出与检索**：xuri/excelize v2、mozillazg/go-pinyin
- **API 文档**：swag + http-swagger

### 1.5.2 前端

- **框架**：React 19 + TanStack Start / Router / Query
- **语言与构建**：TypeScript 5 + Vite 7
- **样式与组件**：Tailwind CSS v4、Radix UI、lucide-react、sonner
- **图表与二维码**：Recharts、qrcode

## 1.6 项目结构

```bash
itdb/
├── backend/                 Go 后端服务
│   ├── cmd/server/          服务入口
│   ├── config/              环境变量与运行配置加载
│   ├── docs/                Swagger/OpenAPI 生成文件
│   ├── internal/            领域模型、数据仓储、本地化、安全能力与构建信息
│   ├── pkg/database/        SQLite 连接与行转换基础设施
│   ├── router/              HTTP 路由装配、全局中间件、权限校验与 Swagger 注解
│   │   ├── assets/          资产域接口（硬件/软件/单据/代理/文件/合同/地点/机架/字典/标记/报表）
│   │   ├── auth/            认证域接口（登录、找回密码、用户管理）
│   │   ├── common/          响应输出、上下文认证、文件上传与数据库初始化等公共能力
│   │   ├── settings/        系统配置域接口（基础配置、用户/角色/群组、认证与邮件）
│   │   └── system/          系统域接口（数据库备份、导入与审计历史）
│   ├── data/                SQLite 数据库、上传文件与备份目录（运行时生成）
│   └── .env.example         环境变量模板
├── deploy/                  Docker Compose、Dockerfile 与 Nginx 部署文件
├── frontend/                React 控制台
│   └── src/
│       ├── components/      布局、启动页、弹窗与基础 UI 组件
│       ├── features/        认证、资产、系统配置等业务域组件
│       ├── lib/             会话认证、审计上报、数据导出与通用工具
│       ├── routes/          TanStack Router 页面路由
│       ├── router.tsx       路由实例
│       ├── server.ts        服务端入口
│       ├── start.ts         客户端入口
│       └── styles.css       全局样式与主题变量
├── scripts/                 运维与在线更新脚本
├── .github/                 GitHub Actions 工作流与项目预览图
├── .dockerignore            Docker 构建忽略规则
├── .gitignore               Git 忽略规则
├── LICENSE
└── README.md                项目说明文档
```

# 二、本地开发快速启动

## 2.1 环境要求

- Go 1.25+（后端）
- Node.js 20+

> 后端使用纯 Go SQLite 驱动（`modernc.org/sqlite`），无需安装 GCC 或 CGO 环境。

## 2.2 克隆项目

```bash
git clone https://github.com/zyx3721/itdb-new.git /data/itdb
cd /data/itdb
```

## 2.3 后端配置与启动

1. 进入后端目录下载相关依赖：

```bash
cd backend
go mod download
```

2. 配置环境变量：

```bash
# 步骤1：复制模板文件
cp .env.example .env

# 步骤2：编辑 .env，按实际环境修改监听地址、密钥等信息
vim .env
# 后端监听地址
ITDB_SERVER_ADDR=127.0.0.1:8080

# 数据库与上传目录
ITDB_DB_PATH=./data/itdb.db
ITDB_UPLOAD_DIR=./data/files

# 鉴权与接口行为
# 留空则启动时自动生成随机密钥并持久化到数据库（删除数据库后所有会话自动失效）
ITDB_JWT_SECRET=itdb-change-me
ITDB_SESSION_TTL_HOURS=12
ITDB_HISTORY_LIMIT=1000
ITDB_CORS_ORIGINS=*
```

环境变量说明：

|               变量               |      默认值      |                         说明                         |
| :------------------------------: | :--------------: | :--------------------------------------------------: |
|        `ITDB_SERVER_ADDR`        | `127.0.0.1:8080` |                       监听地址                       |
|          `ITDB_DB_PATH`          |  `data/itdb.db`  |                  SQLite 数据库路径                   |
|        `ITDB_UPLOAD_DIR`         |   `data/files`   |                   上传文件存储目录                   |
|        `ITDB_JWT_SECRET`         | `itdb-change-me` |            JWT 签名密钥，生产环境务必设置            |
|     `ITDB_SESSION_TTL_HOURS`     |       `12`       |            登录会话（JWT）有效期，单位小时            |
|       `ITDB_HISTORY_LIMIT`       |      `1000`      | 操作历史保留条数；仅控制历史表保留数量，不限制审计日志接口返回条数 |
|        `ITDB_CORS_ORIGINS`       |       `*`        |            允许的跨域来源，多个用逗号分隔            |

后端同时支持命令行参数，显式传入的参数优先于环境变量与 `.env` 文件；`./itdb -v` 可查看版本信息（版本、commit、构建时间），`./itdb -h` 查看全部参数：

|             参数              |        等价环境变量        |                       说明                        |
| :---------------------------: | :------------------------: | :-----------------------------------------------: |
|            `-addr`            |     `ITDB_SERVER_ADDR`     |                      监听地址                      |
|            `-port`            |           `PORT`           |          监听端口，未设置 `-addr` 时生效           |
|             `-db`             |       `ITDB_DB_PATH`       |                  SQLite 数据库路径                 |
|         `-upload-dir`         |      `ITDB_UPLOAD_DIR`     |                   上传文件存储目录                 |
|         `-jwt-secret`         |      `ITDB_JWT_SECRET`     |                    JWT 签名密钥                    |
|        `-history-limit`       |     `ITDB_HISTORY_LIMIT`   | 操作历史保留条数；仅控制历史表保留数量，不限制审计日志接口返回条数 |
|         `-session-ttl`        |   `ITDB_SESSION_TTL_HOURS` |              登录会话有效期（小时）               |
|         `-cors-origins`       |      `ITDB_CORS_ORIGINS`   |            允许的跨域来源，多个用逗号分隔         |
|            `-env`             |             无             |             指定 `.env` 配置文件路径              |
|        `-v`、`-version`       |             无             |                  显示版本信息并退出               |

3. 运行后端服务：

```bash
# 方式1：前台运行（终端关闭则服务停止）
go run cmd/server/main.go

# 方式2：后台运行（日志输出到 app.log）
nohup go run cmd/server/main.go > app.log 2>&1 &
```

后端服务默认运行在 `http://localhost:8080` ，如需指定地址和端口，可修改环境变量文件内的 `ITDB_SERVER_ADDR` 参数，或直接使用 `-addr` 命令行参数（优先级高于环境变量）。首次启动会自动创建数据库和默认管理员账户 `admin / admin123` 。

## 2.4 前端配置与启动

1. 进入前端目录下载相关依赖：

```bash
cd frontend
npm install
```

2. 配置 API 地址（可选）：

```bash
# 配置说明：
# - 后端端口 = 8080：无需创建 .env 文件（默认值为 http://127.0.0.1:8080）
# - 后端端口 ≠ 8080：需要创建 .env 文件（指定正确端口，例如后端端口改为 8090）
#   创建 .env 文件，例如：
echo "VITE_API_BASE_URL=http://localhost:8080" > .env
```

3. 启动前端服务：

```bash
# 方式1：前台运行（终端关闭则服务停止）
npm run dev
# 如果要指定外部访问和监听端口，可执行例如：
npm run dev -- --host --port 5173

# 方式2：后台运行（日志输出到 itdb-frontend.log）
nohup npm run dev > itdb-frontend.log 2>&1 &
```

前端服务默认运行在 `http://localhost:5173/` 。

## 2.5 访问系统

- **首页**：`http://localhost:5173`
  - **默认用户名**：`admin`
  - **默认密码**：`admin123`
- **API 文档**：`http://localhost:8080/swagger/index.html`

# 三、Docker Compose 快速部署（推荐）

## 3.1 部署目录结构

Docker Compose 部署相关文件统一放在 `deploy/` 目录下。`certflow` 单镜像内包含 Go 后端、Nginx 和前端 Nitro SSR 服务，并通过 Supervisor 管理多进程。

仓库内置文件结构：

```bash
deploy/
├── Dockerfile            # 多阶段镜像构建：前端构建、后端构建、运行时镜像
├── docker-compose.yml    # 服务编排配置
├── entrypoint.sh         # 容器启动入口，交给 Supervisor 拉起各进程
├── nginx.conf            # 容器内 Nginx 配置，负责静态资源、API、SSE 和页面反代
├── supervisord.conf      # 容器内多进程管理配置
├── .env.example          # 环境变量模板
```

首次部署时需要从 `.env.example` 复制生成 `.env`，运行后会在 `deploy/` 下生成持久化目录：

```bash
deploy/
├── .env                  # 实际环境变量文件
├── data/                 # 应用数据挂载目录
│   ├── itdb.db           # SQLite 数据库（首次启动自动创建）
│   ├── files/            # 上传的附件文件
│   ├── backups/          # 自动备份文件
│   └── logs/             # 运行日志
```

镜像构建时会分别生成前端 `.output` 产物和后端二进制；运行时由 Supervisor 同时管理 Go 后端、Nginx 和前端 Nitro SSR 服务。

运行时只复制前端 `.output` 产物，并在 `/app/frontend` 执行 `node .output/server/index.mjs`。Nginx 直接托管 `.output/public/assets` 等静态资源，并将 `/api/`、`/api/events` 和 `/swagger/` 反向代理到后端。

## 3.2 准备配置文件

进入 `deploy` 目录，创建 `.env` 环境变量文件：

```bash
cd deploy
vim .env
```

`.env` 文件内容参考：

```bash
# 鉴权与接口行为
ITDB_JWT_SECRET=itdb-change-me
ITDB_SESSION_TTL_HOURS=12
ITDB_HISTORY_LIMIT=1000
ITDB_CORS_ORIGINS=*
```

**配置参数说明详情见 [2.3](#23-后端配置与启动)。**

## 3.3 构建镜像（可选）

如果不想使用阿里云镜像仓库的镜像，可直接在本地手动构建（默认使用阿里云镜像仓库地址）：

```bash
# 在 deploy/ 目录下构建（构建上下文为项目根目录）
cd deploy
docker build \
  -f Dockerfile \
  -t itdb-new:latest \
  --build-arg ALPINE_MIRROR=mirrors.aliyun.com \
  ..
```

然后修改 `deploy/docker-compose.yml` 中 `itdb` 服务的 `image` 字段为 `itdb-new:latest` 。

## 3.4 启动服务

```bash
cd deploy
docker compose up -d
```

## 3.5 服务管理

```bash
# 查看服务状态
docker compose ps

# 查看实时日志
docker compose logs -f itdb

# 重启 itdb 服务
docker compose restart itdb

# 停止所有服务
docker compose down

# 停止并删除数据卷（谨慎！数据会丢失）
docker compose down -v
```

## 3.6 访问系统

服务启动后，访问以下地址：

- **首页**：`http://your-domain.com`
  - **默认用户名**：`admin`
  - **默认密码**：`admin123`
- **API 文档**：`http://your-domain.com/swagger/index.html`
- **健康检查**：`https://your-domain.com/health`

## 3.7 宿主机 Nginx 反代（可选）

如需通过宿主机 Nginx 统一配置公网域名、HTTPS 证书或多站点入口，可将 `deploy/docker-compose.yml` 中的端口映射改为非 80 端口（如 `8080:80`），再由宿主机 Nginx 反向代理到容器内 Nginx。

此时请求链路为：

```text
浏览器 -> 宿主机 Nginx -> itdb 容器内 Nginx -> Go 后端 / 前端 SSR 服务
```

### 3.7.1 HTTP 示例

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # 限制上传文件大小（可选）
    client_max_body_size 500m;

    # 日志配置
    access_log /usr/local/nginx/logs/itdb-access.log;
    error_log /usr/local/nginx/logs/itdb-error.log warn;

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # 超时配置
        proxy_connect_timeout 600s;
        proxy_send_timeout 600s;
        proxy_read_timeout 600s;
    }
}
```

### 3.7.2 HTTPS 示例

> HTTPS 示例（含 80→443 跳转，请替换证书路径）：

```nginx
# HTTP 80端口配置，自动重定向到HTTPS
server {
    listen 80;
    server_name your-domain.com;   # 修改为你的域名/主机名，例如：itdb.cn
    return 301 https://$host$request_uri;
}

# itdb 站点 HTTPS 配置
server {
    # listen 443 ssl http2;  # Nginx 1.25 以下版本写法
    listen 443 ssl;
    http2 on;
    server_name your-domain.com;   # 修改为你的域名/主机名，例如：itdb.cn

    # 证书路径（替换为实际证书文件）
    ssl_certificate     /usr/local/nginx/ssl/your-domain.com.pem;  # 例如：/usr/local/nginx/ssl/itdb.cn.pem
    ssl_certificate_key /usr/local/nginx/ssl/your-domain.com.key;  # 例如：/usr/local/nginx/ssl/itdb.cn.key

    # SSL安全优化
    ssl_protocols              TLSv1.2 TLSv1.3;
    ssl_prefer_server_ciphers  on;
    ssl_ciphers                ECDHE-RSA-AES256-GCM-SHA512:DHE-RSA-AES256-GCM-SHA512:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES256-GCM-SHA384;
    ssl_session_timeout        10m;
    ssl_session_cache          shared:SSL:10m;

    # 限制上传文件大小（可选）
    client_max_body_size 500m;

    # 日志配置
    access_log /usr/local/nginx/logs/itdb-access.log;
    error_log /usr/local/nginx/logs/itdb-error.log warn;

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # 超时配置
        proxy_connect_timeout 600s;
        proxy_send_timeout 600s;
        proxy_read_timeout 600s;
    }
}
```

# 四、生产环境部署

## 4.1 克隆项目

```bash
git clone https://github.com/zyx3721/itdb-new.git /data/itdb
cd /data/itdb
```

## 4.2 后端构建与配置

1. 进入后端目录下载相关依赖：

```bash
cd backend
go mod download
```

2. 配置环境变量：

```bash
# 步骤1：复制模板文件
cp .env.example .env

# 步骤2：编辑 .env，按实际环境修改监听地址、密钥等信息
vim .env
# 后端监听地址
ITDB_SERVER_ADDR=127.0.0.1:8080

# 数据库与上传目录
ITDB_DB_PATH=./data/itdb.db
ITDB_UPLOAD_DIR=./data/files

# 鉴权与接口行为
ITDB_JWT_SECRET=itdb-change-me
ITDB_HISTORY_LIMIT=1000
ITDB_CORS_ORIGINS=*
```

**配置参数说明详情见 [2.3](#23-后端配置与启动)。**

3. 构建后端可执行文件：

```bash
go build -o itdb-backend cmd/server/main.go
```

4. 运行后端服务： 

```bash
# 方式1：前台运行（终端关闭则服务停止）
./itdb-backend

# 方式2：后台运行（日志输出到 app.log）
nohup ./itdb-backend > app.log 2>&1 &
```

`./itdb-backend -v` 可查看版本信息（版本、commit、构建时间），`-h` 查看全部命令行参数，各参数优先级高于同名环境变量，详见 [2.3](#23-后端配置与启动) 的命令行参数表。

```bash
# 方法3：加入 systemd 管理启动运行
# 服务配置参考如下，请自行修改相应目录路径
cat > /etc/systemd/system/itdb-backend.service <<EOF
[Unit]
Description=ITDB Backend Golang Service
After=network.target network-online.target
Wants=network-online.target

[Service]
Type=simple
WorkingDirectory=/data/itdb/backend
ExecStart=/data/itdb/backend/itdb-backend
Restart=on-failure
RestartSec=5
LimitNOFILE=65535
StandardOutput=journal
StandardError=journal
SyslogIdentifier=itdb-backend

[Install]
WantedBy=multi-user.target
EOF

# 重载服务配置并启动
systemctl daemon-reload
systemctl start itdb-backend

# 设置开机自启
systemctl enable --now itdb-backend
```

## 4.3 前端构建与配置

1. 进入前端目录下载相关依赖：

```bash
cd frontend
npm install
```

2. 构建前端项目：

```bash
npm run build
```

构建产物在 `.output` 目录。当前前端使用 TanStack React Start + Nitro，构建后会生成可直接运行的 Node 服务端入口和静态资源目录：

- `.output/server/index.mjs`：生产环境 Node SSR 入口；
- `.output/public/`：浏览器静态资源，包含 JS、CSS、favicon 等文件；
- 生产环境浏览器侧的 `/api/` 请求统一通过 Nginx 将 `/api/` 反向代理到后端，前端无需单独配置 API 地址；但 SSR 进程启动时会主动访问后端拉取品牌配置，把站点名与图标直出进首帧 HTML。

因此生产部署时需要先启动 `.output/server/index.mjs`，再由 Nginx 将页面请求反向代理到该前端服务；不要只把 `.output/public` 配置为 Nginx 静态根目录，否则服务端渲染页面无法正常返回。

3. 启动前端 SSR 服务：

```bash
# 方式1：前台运行（终端关闭则服务停止）
HOST=127.0.0.1 PORT=5173 npm run start

# 方式2：后台运行（日志输出到 itdb-frontend.log）
nohup env HOST=127.0.0.1 PORT=5173 npm run start > itdb-frontend.log 2>&1 &
```

SSR 拉取品牌配置的后端地址按以下顺序确定：运行时环境变量 `SSR_API_ORIGIN` → 入口文件 `index.mjs` 所在目录向上任意一层的 `.env` 文件（首个存在的生效，可与后端共用部署根目录的同一份 `.env`）→ 默认 `http://127.0.0.1:8080`（开发模式下会先回退到 `VITE_API_BASE_URL`，与 dev 代理同源）。请确保该地址对 SSR 进程可达，否则首屏会先显示默认品牌、加载后再切换为配置值，且 SSR 进程日志会输出 `[itdb-ssr] fetch brand failed` 警告。
```

## 4.4 配置Nginx反向代理

在服务器上准备前端目录（例如 `/data/itdb/frontend/.output`），**将本地 `.output` 目录中的所有文件和子目录整体上传到该目录**，保持 `public/` 与 `server/` 结构不变，例如：

```bash
/data/itdb/frontend/.output/
├── public/
│   ├── assets/             # 前端浏览器端 JS/CSS 静态资源
│   └── favicon.svg         # 站点图标
└── server/
    └── index.mjs           # Nitro 生产 SSR 入口
```

上传完成后，在 `.output` 所属的前端项目目录执行 `HOST=127.0.0.1 PORT=5173 npm run start` 启动前端服务。Nginx 的 `/` 请求应反向代理到该服务，例如下方示例中的 `127.0.0.1:5173`；`/api/` 和 `/swagger/` 仍反向代理到 Go 后端 `127.0.0.1:8080`。

`/assets/` 下带扩展名的构建产物（JS/CSS/字体/图片等）可由 Nginx 直接读取 `.output/public` 返回，避免静态资源经过前端 SSR 服务，并可为带 hash 的构建资源启用长期缓存；`/favicon.svg` 同理。注意资产模块的页面路由同样位于 `/assets/` 前缀下（如 `/assets/software`），因此静态资源 location 必须按扩展名正则匹配，不能用 `^~ /assets/` 前缀匹配加 `=404`，否则这些页面的整页请求（新标签打开、刷新、直链）会被静态块拦截返回 404；也不要改用 `try_files $uri @ssr` 回退，该块的一年期强缓存响应头会同样作用于回退返回的 SSR 页面。`/crl/`、`/ocsp` 与 `/ocsp/` 必须直接反向代理到 Go 后端，保留原始路径、请求方法和 `Content-Type`，以支持 CRL 分发、RFC 6960 二进制请求及 JSON 状态查询。示例中的 `root /data/certflow/admin/.output/public;` 请按实际上传目录替换。

### 4.4.1 HTTP 示例

> 配置 Nginx （按需替换域名/路径/证书），`HTTP 示例` ：

```nginx
server {
    listen 80;
    server_name your-domain.com;   # 修改为你的域名/主机名，例如：itdb.cn
    
    # 限制上传文件大小（可选）
    client_max_body_size 500m;
    
    # 日志配置
    access_log /usr/local/nginx/logs/itdb-access.log;
    error_log /usr/local/nginx/logs/itdb-error.log warn;
    
    # 前端静态构建资源：按扩展名匹配，/assets/ 下的页面路由不受影响
    location ~* ^/assets/.+\.(js|mjs|css|map|json|svg|png|jpe?g|gif|webp|ico|woff2?|ttf)$ {
        root /data/itdb/frontend/.output/public;
        try_files $uri =404;
        access_log off;
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
    
    # 站点图标
    location = /favicon.svg {
        root /data/itdb/frontend/.output/public;
        try_files $uri =404;
        access_log off;
        expires 7d;
        add_header Cache-Control "public";
    }
    
    # 后端 API 反向代理
    location /api/ {
        proxy_pass http://127.0.0.1:8080;  # 与后端 API 相同地址
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_connect_timeout 60s;
        proxy_send_timeout 300s;
        proxy_read_timeout 300s;
    }
    
    # 后端 API 文档
    location /swagger/ {
        proxy_pass http://127.0.0.1:8080;  # 与后端 API 相同地址
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
    
    # 前端 Nitro SSR 服务
    location / {
        proxy_pass http://127.0.0.1:5173;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # 健康检查
    location = /health {
        proxy_pass http://127.0.0.1:8080/api/health;
    }
}
```

### 4.4.2 HTTPS 示例

> HTTPS 示例（含 80→443 跳转，请替换证书路径）：

```nginx
# HTTP 80端口配置，自动重定向到HTTPS
server {
    listen 80;
    server_name your-domain.com;   # 修改为你的域名/主机名，例如：itdb.cn
    return 301 https://$host$request_uri;
}

# itdb 站点 HTTPS 配置
server {
    # listen 443 ssl http2;  # Nginx 1.25 以下版本写法
    listen 443 ssl;
    http2 on;
    server_name your-domain.com;   # 修改为你的域名/主机名，例如：itdb.cn

    # 证书路径（替换为实际证书文件）
    ssl_certificate     /usr/local/nginx/ssl/your-domain.com.pem;  # 例如：/usr/local/nginx/ssl/itdb.cn.pem
    ssl_certificate_key /usr/local/nginx/ssl/your-domain.com.key;  # 例如：/usr/local/nginx/ssl/itdb.cn.key
    
    # SSL安全优化
    ssl_protocols              TLSv1.2 TLSv1.3;
    ssl_prefer_server_ciphers  on;
    ssl_ciphers                ECDHE-RSA-AES256-GCM-SHA512:DHE-RSA-AES256-GCM-SHA512:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES256-GCM-SHA384;
    ssl_session_timeout        10m;
    ssl_session_cache          shared:SSL:10m;
    
    # 限制上传文件大小（可选）
    client_max_body_size 500m;

    # 日志配置
    access_log /usr/local/nginx/logs/itdb-access.log;
    error_log /usr/local/nginx/logs/itdb-error.log warn;
    
    # 前端静态构建资源：按扩展名匹配，/assets/ 下的页面路由不受影响
    location ~* ^/assets/.+\.(js|mjs|css|map|json|svg|png|jpe?g|gif|webp|ico|woff2?|ttf)$ {
        root /data/certflow/admin/.output/public;
        try_files $uri =404;
        access_log off;
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
    
    # 站点图标
    location = /favicon.svg {
        root /data/certflow/admin/.output/public;
        try_files $uri =404;
        access_log off;
        expires 7d;
        add_header Cache-Control "public";
    }
    
    # 后端 API 反向代理
    location /api/ {
        proxy_pass http://127.0.0.1:8080;  # 与后端 API 相同地址
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_connect_timeout 60s;
        proxy_send_timeout 300s;
        proxy_read_timeout 300s;
    }
    
    # 后端 API 文档
    location /swagger/ {
        proxy_pass http://127.0.0.1:8080;  # 与后端 API 相同地址
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
    
    # 前端 Nitro SSR 服务
    location / {
        proxy_pass http://127.0.0.1:5173;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # 健康检查
    location = /health {
        proxy_pass http://127.0.0.1:8080/api/health;
    }
}
```

## 4.5 访问系统

服务启动后，访问以下地址：

- **首页**：`http://your-domain.com`
  - **默认用户名**：`admin`
  - **默认密码**：`admin123`
- **API 文档**：`http://your-domain.com/swagger/index.html`
- **健康检查**：`http://your-domain.com/health`

# 五、API 文档

后端已集成 Swagger/OpenAPI 文档，启动后可通过以下地址查看在线接口文档：

- **Swagger UI**：`http://localhost:8080/swagger/index.html`
- **OpenAPI JSON**：`http://localhost:8080/swagger/doc.json`
- **健康检查**：`GET /health`、`GET /api/health`

除 `POST /api/auth/login`、`GET /api/auth/providers`、`GET /api/auth/password-reset/captcha`、`POST /api/auth/password-reset/verify`、`POST /api/auth/password-reset/send`、`POST /api/auth/password-reset/confirm`、`GET /api/public/base`、`GET /health` 和 `GET /api/health` 外，其他接口均需要在请求头中携带 `Authorization: Bearer <token>`。

## 5.1 权限模型

接口按角色权限逐个校验，无权限返回 403；`admin` 用户（usertype=0）拥有全部权限。权限 Key 形如 `模块.资源.操作`：

- **资产管理**：`assets.{items|software|invoices|agents|files|contracts|locations|racks}.read` / `.manage`，资产域所有 GET 接口要求对应资源 `read`，POST/PUT/DELETE 要求对应资源 `manage`
- **资料管理**：`dictionaries.{itemtypes|contracttypes|statustypes|filetypes|dpttypes|tags}.read` / `.manage`，`GET /api/dictionaries` 仅返回有查看权限的字典分组，字典写入接口按 URL 中的 `{name}` 校验对应 `manage`（`contractsubtypes` 归属 `contracttypes`）
- **打印标签**：`labels.preview`（标签数据/预设读取/预览生成）、`labels.print`（隐含 preview）、`labels.manage`（预设保存与删除，隐含 preview）
- **统计报表**：`reports.read`（查看）、`reports.manage`（导出，隐含 read）；**资产导航**：`browse.read`；**审计日志**：`audit.read`（查看）、`audit.manage`（导出，隐含 read）
- **系统设置**：`settings.{base|users|auth|notifications}.read` / `.manage`；数据库备份下载与导入要求 `settings.base.manage`
- **聚合接口**：`GET /api/bootstrap`、`GET /api/dashboard/summary` 要求持有任一只读类权限
- **隐含规则**：勾选任一 `manage` 或操作类权限时，服务端保存角色与登录下发权限时会自动补全其隐含的查看权限（如 `labels.print` 隐含 `labels.preview` 与 `assets.items.read`）

## 5.2 接口清单

以下接口清单与 Swagger 文档（`/swagger/index.html`）完全一致，按模块分组列出；除标注“无需认证”的接口外，均需在请求头携带 `Authorization: Bearer <token>`。

登录请求示例：

```json
{
  "username": "admin",
  "password": "admin123",
  "mode": "local"
}
```

### 5.2.1 代理

- `GET /api/agents` - 获取厂商/代理商列表：获取代理（硬件厂商、软件厂商、供应商、采购方、承包方）列表，支持关键字搜索
- `POST /api/agents` - 创建厂商/代理商：创建代理；同一名称（忽略大小写与首尾空格）不允许重复
- `GET /api/agents/{id}` - 获取厂商/代理商详情：按编号获取代理详情
- `PUT /api/agents/{id}` - 更新厂商/代理商：更新代理；取消仍被引用的类型会被整体拒绝
- `DELETE /api/agents/{id}` - 删除厂商/代理商：删除代理并置空硬件、软件、单据中的引用；仍被引用时返回 409

### 5.2.2 认证

- `POST /api/auth/change-password` - 修改当前用户密码：修改当前登录用户的密码，成功后需重新登录
- `POST /api/auth/login` - 用户登录：用户登录，支持本地密码与 AD/LDAP 两种方式，成功返回令牌与用户信息（含企微绑定状态 `wecomBound`）；同一账号连续密码失败达到「安全时效」配置的锁定次数后返回 429 并提示剩余等待分钟数，锁定时长内即使密码正确也会被拒绝，登录成功、找回密码重置成功或锁定窗口过期后重新计数，内置管理员 admin 不受限
- `POST /api/auth/logout` - 登出当前会话：退出当前会话，并写入用户注销审计
- `GET /api/auth/me` - 获取当前用户：获取当前登录用户的资料、直接角色、有效角色与权限清单
- `GET /api/auth/password-reset/captcha` - 获取找回密码图形验证码：获取找回密码图形验证码，无需认证
- `POST /api/auth/password-reset/confirm` - 确认找回密码：凭邮箱验证码完成找回密码，重置账号密码并写入审计，若该账号存在登录失败锁定记录则一并解除，并退出该账号全部登录会话
- `POST /api/auth/password-reset/send` - 发送找回密码验证码：向校验通过的邮箱发送找回密码验证码，受发送冷却与限流窗口约束
- `POST /api/auth/password-reset/verify` - 校验找回密码身份：校验用户名与图形验证码，换取找回密码流程令牌
- `GET /api/auth/providers` - 获取公开认证方式：获取登录页可用的认证方式与找回密码开关，无需认证
- `GET /api/auth/wecom/authorize` - 获取企业微信扫码登录地址：生成企业微信 Web 扫码登录页地址（含防伪 state）与内嵌二维码渲染参数（`iframe_url`、回跳路径 `/wecom-qr-callback`），登录页默认内嵌渲染二维码，配置异常时自动回退整页跳转；需已启用企业微信认证，无需认证
- `POST /api/auth/wecom/callback` - 企业微信扫码登录回调：校验 state 后用授权码换取企业微信成员身份，按绑定关系登录并返回令牌与用户信息（含 `wecomBound`）；未绑定时返回 401，无需认证
- `POST /api/auth/wecom/sso/callback` - 统一认证中心扫码登录回调：统一认证（SSO）模式下校验认证中心回跳的一次性 ticket 后按绑定关系登录并返回令牌与用户信息（含 `wecomBound`）；需已启用统一认证模式，无需认证
- `GET /api/auth/wecom/bind-url` - 获取企业微信绑定扫码地址：为当前登录用户生成企业微信绑定扫码地址，state 绑定当前用户
- `POST /api/auth/wecom/bind` - 绑定企业微信账号：校验绑定 state 后用授权码换取企业微信成员 userid，与当前用户建立绑定；该企微已绑定其他用户时返回 409
- `POST /api/auth/wecom/sso/bind` - 统一认证中心扫码绑定：统一认证（SSO）模式下校验认证中心一次性 ticket 后与当前登录用户建立绑定；该企微已绑定其他用户时返回 409
- `DELETE /api/auth/wecom/bind` - 解绑企业微信账号：解除当前登录用户的企业微信绑定
- `GET /api/public/base` - 获取公开品牌标识：获取登录页与启动屏使用的品牌标识，无需认证

### 5.2.3 备份

- `GET /api/backups/database` - 下载数据库备份：下载数据库备份（itdb-日期.db）；携带 files=1 时按数据库实际引用的文件打包为 zip（文件位于包内 files 目录），引用文件全部不存在或未引用时仅导出 db 文件，缺失文件写入备份审计

### 5.2.4 启动数据

- `GET /api/bootstrap` - 获取前端启动字典数据：获取前端启动所需的字典、引用选项与权限数据

### 5.2.5 浏览树

- `GET /api/browse/tree` - 获取浏览树：按类型、部门、用户、代理等维度获取资产导航树

### 5.2.6 合同

- `GET /api/contracts` - 获取合同列表：获取合同列表，支持关键字搜索
- `POST /api/contracts` - 创建合同：创建合同；合同标题不允许重复
- `GET /api/contracts/{id}` - 获取合同详情：按编号获取合同详情及全部关联
- `PUT /api/contracts/{id}` - 更新合同：更新合同；备件、事件历史随表单一并保存，重名标题返回 409
- `DELETE /api/contracts/{id}` - 删除合同：删除合同；被下级合同引用时返回 409

### 5.2.7 合同事件

- `GET /api/contracts/next-event-id` - 获取下一个合同事件编号：获取下一个可用的合同事件编号
- `GET /api/contracts/{id}/events` - 获取合同事件：获取指定合同的事件历史列表
- `POST /api/contracts/{id}/events` - 创建合同事件：为指定合同新增一条事件历史
- `PUT /api/contracts/{id}/events/{eventId}` - 更新合同事件：更新指定的事件历史
- `DELETE /api/contracts/{id}/events/{eventId}` - 删除合同事件：删除指定的事件历史

### 5.2.8 仪表盘

- `GET /api/dashboard/summary` - 获取仪表盘统计：获取仪表盘各资源总数统计，需要任意查看权限

### 5.2.9 字典

- `GET /api/dictionaries` - 获取所有字典数据：一次性获取硬件类型、部门、状态、标记等全部字典数据（按权限过滤可见分类）
- `POST /api/dictionaries/{name}` - 创建字典行：创建字典行；重名返回 409；携带 audit=0 时跳过逐行审计（批量导入专用）
- `PUT /api/dictionaries/{name}/{id}` - 更新字典行：更新字典行；内容无任何变化时不写审计
- `DELETE /api/dictionaries/{name}/{id}` - 删除字典行：删除字典行；仍被引用时返回 409 且不写审计

### 5.2.10 文件

- `GET /api/files` - 获取文件列表：获取文件列表，支持关键字搜索
- `POST /api/files` - 上传文件：上传文件（multipart），支持关联到硬件、软件、单据、合同
- `GET /api/files/{id}` - 获取文件详情：按编号获取文件详情及全部关联
- `PUT /api/files/{id}` - 更新文件：更新文件信息，可选择替换文件内容；无变化时不写审计
- `DELETE /api/files/{id}` - 删除文件：删除文件；仍被关联引用时返回 409
- `GET /api/files/{id}/download` - 下载文件：下载文件内容

### 5.2.11 健康检查

- `GET /api/health` - API 健康检查：探测服务进程是否存活，无需认证
- `GET /health` - 健康检查：探测服务进程是否存活，无需认证

### 5.2.12 审计日志

- `GET /api/history` - 获取审计日志：导入类事件的目标仅返回前两个名称，完整清单通过 targetTitle 返回供悬浮提示；默认返回全部保留记录，`limit` 可选且不设上限（`-1` 等同于不限制）
- `POST /api/history/clear` - 清空审计日志：按前端筛选结果传入待清空的编号清单批量删除，清空成功后自动补记一条「清空审计日志」审计事件（归属「审计日志」模块，目标为清除条数）；需要审计管理权限
- `POST /api/history/events` - 上报前端审计事件：前端执行数据导出（资产、字典、统计报表）、硬件维护日志导出、资料字典导入或标签打印后上报审计事件；事件类型白名单校验并按所属资源校验权限

### 5.2.13 数据库导入

- `POST /api/import/database` - 导入 SQLite 数据库：上传 .db 或 zip 替换当前数据库，压缩包内的上传文件恢复到上传目录，导入前自动预备份（当前库引用的上传文件一并打包为 zip，未引用文件时仅备份数据库），导入成功后清理已打包的上传文件；旧版平台数据库自动转换，仅迁移资产管理与资料管理数据，用户按用户名与当前库合并（同名用户保留现状），系统配置、用户角色档案与审计历史从当前库恢复；与当前项目结构一致的数据库则完整替换；当前数据库文件被外部工具占用时返回 409 并提示先关闭再导入

### 5.2.14 单据

- `GET /api/invoices` - 获取单据列表：获取单据（发票）列表，支持关键字搜索
- `POST /api/invoices` - 创建单据：创建单据
- `GET /api/invoices/{id}` - 获取单据详情：按编号获取单据详情及全部关联
- `PUT /api/invoices/{id}` - 更新单据：更新单据；无任何修改时不写审计
- `DELETE /api/invoices/{id}` - 删除单据：删除单据及其关联

### 5.2.15 硬件资产

- `GET /api/items` - 获取硬件资产列表：获取硬件资产列表，支持关键字搜索与全量分页参数
- `POST /api/items` - 创建硬件资产：创建硬件资产；机架位置冲突、软件授权超限等会返回 409
- `GET /api/items/{id}` - 获取硬件资产详情：按编号获取硬件资产详情及全部关联
- `PUT /api/items/{id}` - 更新硬件资产：更新硬件资产；无任何修改时不写审计，维护日志记录变更项
- `DELETE /api/items/{id}` - 删除硬件资产：删除硬件资产并清理关联、维护日志与标记关联；被交换机引用时返回 409
- `POST /api/items/{id}/tags` - 关联或移除硬件标签：关联或解除硬件标记（action=add/remove）

### 5.2.16 硬件操作记录

- `GET /api/items/{id}/actions` - 获取硬件操作记录：获取硬件的维护操作记录

### 5.2.17 标签打印

- `GET /api/labels/items` - 获取标签打印资产：获取可打印标签的硬件清单，支持关键字搜索与排序
- `GET /api/labels/presets` - 获取标签纸预设：获取标签纸预设列表
- `POST /api/labels/presets` - 保存标签纸预设（同名时更新）：保存标签预设；同名覆盖更新，无任何修改时不写审计
- `DELETE /api/labels/presets/{id}` - 删除标签纸预设：删除标签纸预设
- `POST /api/labels/preview` - 预览标签打印：按选中硬件生成标签预览数据，并写入打印标签审计

### 5.2.18 地点

- `GET /api/locations` - 获取地点列表：获取地点列表，支持关键字搜索
- `POST /api/locations` - 创建地点：创建地点；审计详情包含保存时提交的区域清单
- `GET /api/locations/{id}` - 获取地点详情：按编号获取地点详情
- `PUT /api/locations/{id}` - 更新地点：更新地点；区域页签的最终清单随表单提交，无任何修改时不写审计
- `DELETE /api/locations/{id}` - 删除地点：删除地点；仍被机架或硬件引用时返回 409
- `GET /api/locations/{id}/floorplan` - 查看地点平面图：查看地点平面图图片

### 5.2.19 地点区域

- `GET /api/locations/next-area-id` - 获取下一个区域编号：获取下一个可用的区域编号
- `GET /api/locations/{id}/areas` - 获取地点区域：获取地点下的全部区域
- `POST /api/locations/{id}/areas` - 创建地点区域：为地点新增区域
- `PUT /api/locations/{id}/areas/{areaId}` - 更新地点区域：更新地点区域
- `DELETE /api/locations/{id}/areas/{areaId}` - 删除地点区域：删除地点区域；仍被引用时返回 409

### 5.2.20 机架

- `GET /api/racks` - 获取机架列表：获取机架列表，支持关键字搜索
- `POST /api/racks` - 创建机架：创建机架
- `GET /api/racks/{id}` - 获取机架详情：按编号获取机架详情
- `PUT /api/racks/{id}` - 更新机架：更新机架；无任何修改时不写审计
- `DELETE /api/racks/{id}` - 删除机架：删除机架；仍被硬件引用时返回 409

### 5.2.21 报表

- `GET /api/reports` - 获取报表列表：获取可用的统计报表清单
- `GET /api/reports/{name}` - 执行报表：执行指定报表并返回数据与图表配置

### 5.2.22 系统设置

- `GET /api/settings/auth-provider` - 获取认证配置：获取 AD/LDAP 认证配置（敏感字段脱敏）
- `PUT /api/settings/auth-provider` - 保存 AD/LDAP 认证配置
- `GET /api/settings/auth/wecom` - 获取企业微信认证配置：含认证方式（direct/sso）与对应配置项；Secret 与应用密钥不回显，仅返回是否已配置
- `PUT /api/settings/auth/wecom` - 保存企业微信认证配置：认证方式 direct 启用时要求企业 ID、AgentID 与 Secret 完整，sso 启用时要求认证中心地址、应用标识与应用密钥完整，密钥加密存储
- `GET /api/settings/base` - 获取系统基础配置：获取系统基础配置
- `PUT /api/settings/base` - 更新系统基础配置：更新系统基础配置；携带 section（brand/security/backup）时按区块局部更新，审计目标细分为品牌标识、安全时效、数据备份
- `GET /api/settings/email` - 获取邮件配置：获取邮件配置（敏感字段脱敏）
- `PUT /api/settings/email` - 保存邮件配置
- `POST /api/settings/email` - 发送测试邮件
- `GET /api/settings/roles` - 获取设置角色列表：获取角色列表及权限集合
- `POST /api/settings/roles` - 创建角色
- `PUT /api/settings/roles/{id}` - 更新角色
- `DELETE /api/settings/roles/{id}` - 删除角色
- `POST /api/settings/roles/{id}/disabled` - 启用或禁用角色
- `GET /api/settings/user-groups` - 获取设置用户组列表：获取用户群组列表及成员与角色
- `POST /api/settings/user-groups` - 创建用户群组
- `PUT /api/settings/user-groups/{id}` - 更新用户群组
- `DELETE /api/settings/user-groups/{id}` - 删除用户群组
- `GET /api/settings/users` - 获取设置用户列表：获取用户列表及有效角色与权限
- `POST /api/settings/users` - 创建用户
- `PUT /api/settings/users/{id}` - 更新用户
- `DELETE /api/settings/users/{id}` - 删除用户
- `POST /api/settings/users/{id}/disabled` - 启用或禁用用户

### 5.2.23 软件许可

- `GET /api/software` - 获取软件许可列表：获取软件许可列表，支持关键字搜索
- `POST /api/software` - 创建软件许可：创建软件许可；标题加版本不允许重复，授权口径校验失败返回 409
- `GET /api/software/{id}` - 获取软件许可详情：按编号获取软件许可详情及全部关联
- `PUT /api/software/{id}` - 更新软件许可：更新软件许可；无任何修改时不写审计
- `DELETE /api/software/{id}` - 删除软件许可：删除软件许可及其关联
- `POST /api/software/{id}/tags` - 关联或移除软件标签：关联或解除软件标记（action=add/remove）

### 5.2.24 标记

- `GET /api/tags/next-id` - 获取下一个标记编号：获取下一个可用的标记编号
- `GET /api/tags/{id}/items` - 获取标记关联硬件：获取标记关联的硬件清单（仅现存硬件）
- `GET /api/tags/{id}/software` - 获取标记关联软件：获取标记关联的软件清单（仅现存软件）

# 六、数据库说明

使用 SQLite 单文件数据库，默认路径 `backend/data/itdb.db`，共 49 张表（36 张随初始化创建、13 张由系统配置模块按需创建）。

## 6.1 核心业务表

| 表名 | 说明 |
|:----:|:----:|
| `items` | 硬件资产（核心表，含 SN、IP、机架位置、CPU/RAM/HD 等字段） |
| `software` | 软件许可证 |
| `contracts` | 合同 |
| `invoices` | 单据 |
| `files` | 文件附件 |
| `agents` | 代理（硬件/软件厂商、供应商、采购方、承包方） |
| `users` | 系统用户 |
| `locations` | 地点（机房/楼层） |
| `racks` | 机架 |
| `locareas` | 地点区域（平面图热区） |
| `labelpapers` | 标签纸预设 |
| `tags` | 标记 |
| `actions` | 硬件维护日志 |
| `contractevents` | 合同事件历史 |

## 6.2 关联表

| 表名 | 说明 |
|:----:|:----:|
| `item2inv` | 硬件 ↔ 单据 |
| `item2soft` | 硬件 ↔ 软件（含安装日期） |
| `item2file` | 硬件 ↔ 文件 |
| `itemlink` | 硬件 ↔ 硬件互联 |
| `contract2item` | 合同 ↔ 硬件 |
| `contract2soft` | 合同 ↔ 软件 |
| `contract2inv` | 合同 ↔ 单据 |
| `contract2file` | 合同 ↔ 文件 |
| `invoice2file` | 单据 ↔ 文件 |
| `soft2inv` | 软件 ↔ 单据 |
| `software2file` | 软件 ↔ 文件 |
| `tag2item` | 标记 ↔ 硬件 |
| `tag2software` | 标记 ↔ 软件 |

## 6.3 字典表

| 表名 | 说明 |
|:----:|:----:|
| `itemtypes` | 硬件类型 |
| `contracttypes` | 合同类型 |
| `contractsubtypes` | 合同子类型 |
| `dpttypes` | 部门 |
| `statustypes` | 资产状态（含颜色） |
| `filetypes` | 文件类型 |

## 6.4 系统表

| 表名 | 说明 |
|:----:|:----:|
| `settings` | AD/LDAP 连接参数（单行表，绑定密码加密存储） |
| `system_secrets` | JWT 签名密钥（单行表） |
| `history` | 审计日志（记录模块、操作、目标、结果、详情与原始 SQL） |
| `settings_base` | 基础配置（品牌标识、安全时效、备份参数，单行 JSON） |
| `settings_email` | 邮件通知配置 |
| `settings_auth_providers` | AD/LDAP 与企业微信认证配置 |
| `settings_user_wecom` | 用户与企业微信账号绑定关系（user_id 主键，wecom_userid 唯一） |
| `settings_roles` | 用户角色与权限集合 |
| `settings_role_status` | 角色启用/禁用状态 |
| `settings_user_profiles` | 用户资料（邮箱、启用状态、来源、最后登录时间） |
| `settings_user_groups` | 用户群组 |
| `settings_user_group_members` | 群组成员关系 |
| `settings_user_group_roles` | 群组授予的角色 |
| `settings_user_roles` | 用户直接授予的角色 |
| `password_reset_captchas` | 找回密码图形验证码 |
| `password_reset_requests` | 找回密码流程请求（流程令牌与邮箱验证码） |
| `password_reset_send_log` | 找回密码验证码发送记录（按邮箱频率限制） |
| `login_failure_log` | 登录密码失败记录（登录失败锁定计数，成功登录或找回密码重置后清零） |

# 七、常见问题

## 7.1 忘记管理员密码怎么办？

可通过 SQLite 命令行工具直接重置密码（推荐，不会丢失数据）：

```bash
# 停止后端服务后执行
sqlite3 backend/data/itdb.db "UPDATE users SET pass = 'admin123' WHERE username = 'admin';"
```

重启后端服务后，使用 `admin / admin123` 登录，系统会自动将明文密码升级为加密存储。

如果无法使用 sqlite3 工具，也可以删除数据库文件 `backend/data/itdb.db` 并重启服务，但这会清空所有数据，仅建议在全新部署时使用。

## 7.2 如何修改 JWT 有效期？

默认 12 小时，通过环境变量 `ITDB_SESSION_TTL_HOURS`（单位小时，正整数）调整：在 `.env` 或运行环境中设置后重启后端生效。会话同时记录于服务端 `user_sessions` 表：主动注销会立即删除对应会话并使令牌失效，已过期的会话在下次登录时自动清理；重设 `ITDB_JWT_SECRET` 可使全部已签发会话立即失效。

## 7.3 数据库文件可以直接复制迁移吗？

可以。SQLite 是单文件数据库，停止后端服务后直接复制 `itdb.db` 文件即可完成迁移。也可以通过系统内置的数据库导入功能在线替换。

## 7.4 如何从旧版 PHP ITDB 迁移？

旧项目数据库（含旧版 PHP ITDB 导出的 .db）可通过「系统配置 → 数据备份 → 数据库导入」上传 .db 或包含上传文件的 .zip 自动完成迁移，适合重复导入以同步旧项目最新数据：系统会识别旧库结构并自动转换，仅迁移资产管理与资料管理数据（旧「用户描述」对应现「显示名称」，新导入用户的邮箱留空后可手动补填），用户按用户名与当前系统合并——当前系统已有的同名用户保留现状（邮箱、密码、角色等不被覆盖），仅导入旧项目中新增的用户，并保持其原编号以便旧数据关联，新导入用户按旧库用户类型自动关联内置角色（管理员类型关联 admin，普通类型关联 viewer 只读角色，导入后可在角色管理中调整），避免导入后无任何权限；

系统配置（基础配置、认证配置、通知配置、角色权限、用户组、用户档案）与审计日志全部保持当前系统现状不被清空；

状态类型自动从 1 重新编号并同步硬件状态引用，硬件类型重建为当前默认：编号 1-5 固定为内置类型（仅「服务器」默认支持软件），旧库其余类型按原编号顺序从 6 连续追加，与内置或已追加类型重名的自动忽略，硬件记录的类型编号同步映射到新编号；

硬件维护日志（actions 操作记录）不迁移，导入后保持为空；英文内置名（文件类型/合同类型）自动转为中文。

以上合并与保留规则仅针对旧版数据库，与当前项目结构一致的数据库导入时仍然完整替换。导入前会自动预备份当前数据库，当前库引用的上传文件存在时会连同文件一并打包为 zip（文件放在包内 files 目录），未引用文件时仅备份数据库文件，导入成功后自动清理已打包的上传文件；当前数据库文件被外部工具占用时会提示先关闭再导入。

## 7.5 如何启用 LDAP 登录？

LDAP 登录需要两步配置：

1. 在「系统配置」的「认证配置」中配置 LDAP 服务器地址、Base DN、Bind DN 等连接参数，并启用 LDAP 认证
2. 在「系统配置 → 用户管理」中创建与 LDAP 账号同名的用户（用户名必须与 LDAP 中的 `sAMAccountName` 一致）

登录时用户选择「LDAP」模式，系统会先在本地用户表中查找该用户名，再通过 LDAP 服务器验证密码。如果本地用户表中不存在对应用户，即使 LDAP 密码正确也无法登录。

## 7.6 自动备份存储在哪里？

定时备份存储在 `backend/data/backups/` 目录：未引用上传文件（或引用文件全部不存在）时为 `itdb-YYYYMMDD-HHMMSS.db`；引用到上传文件时打包为 `itdb-YYYYMMDD-HHMMSS.zip`（数据库在包根目录，文件在包内 files 目录）。过期文件按基础配置中的保留天数自动清理。

在 系统配置 → 基础配置 → 数据备份 中启用「启用定时备份」并配置五段 Cron 计划（分 时 日 月 周），例如 `0 0 * * *` 表示每天 0 点自动备份。

## 7.7 上传文件大小有限制吗？

后端默认无大小限制，但如果使用 Nginx 反向代理，需要配置 `client_max_body_size`（参考上方 Nginx 配置示例）。

## 7.8 如何启用企业微信扫码登录？

企业微信扫码登录需要三步配置：

1. 在企业微信管理后台创建自建应用，记录「企业 ID（corpid）」「应用 AgentID」「应用 Secret」，并将系统站点的访问域名配置为该应用的「可信域名」（Web 登录授权回调所需）
2. 在「系统配置 → 认证配置 → 企业微信」中填写企业 ID、AgentID 与 Secret，并按需填写「回调地址前缀」（企业外部可访问的站点地址，如 `https://itdb.example.com`，留空则按用户当前访问地址自动推断），保存并启用
3. 启用后登录页出现「企业微信」登录方式，选择后跳转企业微信扫码页；扫码确认后按绑定关系自动登录

用户与企微账号的对应关系通过「绑定」维护：先用账号密码登录，在右上角用户菜单点击「绑定企微」扫码完成绑定（绑定后可解绑）。未绑定的企微账号扫码登录会被拒绝并提示先绑定；绑定关系在数据库 `settings_user_wecom` 表中维护，旧库迁移导入时会自动保留。扫码授权的有效期默认 5 分钟，可在「系统配置 → 基础配置 → 安全时效」中调整（1-60 分钟），超时需重新扫码。

企业微信认证支持两种方式，可在认证配置的「认证方式」中切换，两套凭据分别保存、互不影响：

- **直连企业微信**（默认）：本系统直接持有企微应用凭据，按上述步骤配置即可
- **统一认证中心**：若企业已部署统一认证中心（wecom-auth-center），将认证方式切换为「统一认证中心」，填写认证中心地址、应用标识与应用密钥（应用密钥与认证中心 `apps` 配置的 `app_secret` 一致）。认证中心 `config.yaml` 中为本系统新增 `apps` 条目：`domain` 填本系统外部访问地址、`callback_path` 填 `/wecom-qr-callback`（本系统扫码中转路由，内嵌二维码与整页跳转共用，顶层整页回跳会自动转发登录页，登录或绑定交互与 `/login` 时代一致；同时认证中心登录页需允许被本系统页面以 iframe 嵌入，未配置时登录页自动回退整页跳转），修改后重启认证中心生效。此模式下企微扫码由认证中心代理完成，认证中心携带一次性 ticket 回跳本系统中转路由完成登录或绑定，绑定关系与会话机制与直连模式一致，多个内部系统可共用同一套企微应用配置

# 八、安全建议

1. **修改默认密码**：首次部署后立即修改 `admin` 账户的默认密码
2. **设置 JWT 密钥**：未配置 `ITDB_JWT_SECRET` 时，系统会在首次启动自动生成随机密钥并持久化到数据库（删除或重建数据库后所有已签发会话自动失效）；多实例部署或需固定密钥时，请在 `.env` 中显式设置 `ITDB_JWT_SECRET`
3. **启用 HTTPS**：生产环境建议通过 Nginx 配置 SSL 证书，启用 HTTPS 访问
4. **限制访问来源**：通过 Nginx 或防火墙限制系统的访问 IP 范围
5. **定期备份**：虽然系统已有每日自动备份，建议额外配置异地备份策略
6. **文件目录权限**：确保 `backend/data/` 目录权限合理，避免非授权访问数据库和上传文件
7. **环境变量安全**：`.env` 文件包含敏感信息，确保不被提交到版本控制（已在 `.gitignore` 中排除）
8. **CORS 配置**：生产环境按需配置 `ITDB_CORS_ORIGINS`，避免设置为 `*`

# 九、许可证

本项目采用 [MIT License](../LICENSE) 开源协议。

MIT License 是一个宽松的开源许可证，允许您自由地使用、复制、修改、合并、发布、分发、再许可和/或销售本软件的副本。唯一的要求是在所有副本或重要部分中保留版权声明和许可声明。

# 十、致谢

感谢以下开源项目和技术社区的支持：

- [go-chi/chi](https://github.com/go-chi/chi) - 轻量高效的 Go HTTP 路由框架
- [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) - 纯 Go 实现的 SQLite 驱动
- [swaggo/swag](https://github.com/swaggo/swag) - Swagger 文档生成工具
- [excelize](https://github.com/qax-os/excelize) - Excel 文件读写库
- [TanStack](https://tanstack.com/) - Router / Query / Start 前端框架套件
- [Tailwind CSS](https://tailwindcss.com/) - 原子化 CSS 框架

特别感谢所有为本项目贡献代码、提出建议和报告问题的开发者。

# 十一、联系方式

如果您在使用过程中遇到问题，或有任何建议和反馈，欢迎通过以下方式联系：

- **Email**: 416685476@qq.com
- **GitHub Issues**: [https://github.com/zyx3721/itdb-new/issues](https://github.com/zyx3721/itdb-new/issues)
- **项目主页**: [https://github.com/zyx3721/itdb-new](https://github.com/zyx3721/itdb-new)

---

**⭐ 如果这个项目对您有帮助，欢迎 Star 支持！**

