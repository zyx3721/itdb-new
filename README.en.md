<div align="center">

<h1>ITDB</h1>

<p><b>Full-lifecycle IT asset management</b> — hardware · software · invoices · vendors · files · contracts · locations · racks</p>

<p><a href="README.md">简体中文</a> · <b>English</b></p>

Rack sheets, warranty dates, software licences and contract renewals usually live in a handful of spreadsheets and a shared drive.  
ITDB pulls them into one **self-hosted** system: a Go backend, a React console and a single-file SQLite database.  
One machine, one `docker run`, or a single binary downloaded from Releases.  
Your inventory, attachments and backups stay on your own disk; the repository holds no tokens, no passwords and no real hostnames.

<p>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue?labelColor=1f2937" alt="MIT License"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go&logoColor=white&labelColor=1f2937" alt="Go 1.25+"></a>
  <a href="https://react.dev/"><img src="https://img.shields.io/badge/React-19-61DAFB?logo=react&logoColor=white&labelColor=1f2937" alt="React 19"></a>
  <a href="https://www.sqlite.org/"><img src="https://img.shields.io/badge/SQLite-single%20file-003B57?logo=sqlite&logoColor=white&labelColor=1f2937" alt="SQLite"></a>
  <img src="https://img.shields.io/badge/permissions-44-059669?labelColor=1f2937" alt="44 permissions">
</p>

<p>
  <b><a href="#preview">Preview</a></b> ·
  <a href="#what-it-does">What it does</a> ·
  <a href="#how-it-works">How it works</a> ·
  <a href="#tech-stack">Tech stack</a> ·
  <a href="#quick-start">Quick start</a> ·
  <a href="#deployment">Deployment</a> ·
  <a href="#permission-model">Permissions</a> ·
  <a href="#data-and-security">Security</a> ·
  <a href="#faq">FAQ</a> ·
  <a href="#repository-layout">Layout</a> ·
  <a href="#documentation">Docs</a>
</p>

</div>

---

ITDB is a full rewrite of [zyx3721/itdb](https://github.com/zyx3721/itdb/): the asset domain model and legacy-database import are kept, while the stack moves to a Go + React frontend/backend split with rebuilt permissions, auditing, backup and label printing. The backend uses a pure-Go SQLite driver, so there is **no CGO, no GCC and no external database**; the frontend is a React 19 console served with Nitro SSR after build.

## Preview

### Sign in

An internal system with no public registration. Local passwords, AD/LDAP and WeCom (WeChat Work) QR-code sign-in are supported, with password recovery over a captcha plus an email code.

![Sign-in page](.github/images/itdb-login.jpg)

### Dashboard

A read-only overview: asset totals, status distribution and recent activity. The sidebar shows or hides each entry according to the signed-in user's permissions, so actions you cannot perform never appear.

![Dashboard](.github/images/itdb-home.jpg)

## What it does

- **Full asset lifecycle** — eight resource types: hardware, software, invoices, vendors, files, contracts, locations and racks. Hardware carries serial numbers, network details, warranty and cost records and can link to software, invoices, contracts, files and other hardware; contracts support types/subtypes, renewals and event history; locations support floor-plan uploads with clickable areas; racks render U positions and front/back views.
- **Reference dictionaries** — six dictionaries: hardware types, contract types (with subtypes), status types (custom colours), file types, departments and tags, all with Excel template download, bulk import and export. Built-in rows are protected by both id and name.
- **Users and permissions** — local, AD/LDAP and WeCom QR-code sign-in (direct or unified auth-center mode), JWT sessions, WeCom account binding and an email password-recovery flow; three built-in roles (`admin` / `operator` / `viewer`), custom roles picking from 44 permissions, and user groups for bulk grants. The console hides entries you cannot use and the backend checks every request.
- **Audit log** — sign-in/out, asset changes, dictionary maintenance, configuration, backup, import and label printing are recorded with module, action, target, result and detail, all searchable, filterable and exportable. Saves with no actual change write nothing, and business-rule rejections add no failure noise.
- **Backup and migration** — manual backups can bundle the files the database actually references; scheduled backups run on a five-field cron expression and are pruned by retention days; import accepts `.db` and `.zip` and converts legacy databases automatically.
- **Label printing** — a QR label designer with several label-sheet presets, batch preview and printing.
- **Reporting and browsing** — a dashboard summary, built-in reports with XLSX/XLS/CSV/TXT export and an asset navigation tree by type, department, user or vendor.
- **System settings** — branding, security timings (password-recovery rate limits, login-failure lockout), scheduled-backup parameters, users/groups/roles, AD/LDAP and WeCom authentication (direct or unified auth-center mode), email notification, with connectivity and test-mail checks.

**It is not** a CMDB discovery tool and not a monitoring platform. ITDB manages *inventory + contracts + licences + locations*: it does not scan networks, collect metrics or log in to managed devices.

The same rack inventory, two ways of running it:

| Situation | Spreadsheets + shared drive | With ITDB |
| --- | --- | --- |
| Adding a device | Anyone can edit; versions live in filenames | Only permitted users edit; the change lands in the audit log |
| Who changed it | Unrecoverable | Module/action/target/result recorded in full |
| Warranty expiry | You have to remember | Warranty and contract end dates in one place |
| Software licences | Seat counts drift from reality | Licence usage derived from asset links; over-allocation is rejected |
| Rack position | Drawn once, stale forever | U-position and front/back views; conflicts are rejected |
| Access control | One folder permission for everyone | 44 permissions checked per endpoint |
| Moving hosts | Copy folders, then fix references | Copy `itdb.db` plus `data/files` |

## How it works

```text
        Browser
           │  http
           ▼
  ┌──────────────────────────────────────┐
  │  Nginx (inside the container or host) │
  │  /assets/ · /       → frontend SSR    │
  │  /api/ · /swagger/  → Go backend      │
  └──────────────────────────────────────┘
        │                        │
        ▼                        ▼
  React 19 console           Go 1.25 backend
  Nitro SSR :5173            chi :8080
                                 │
                        ┌────────┴─────────┐
                        ▼                  ▼
                data/itdb.db         data/files
                49 tables            uploaded files
```

- **Who does what** — the console only reads current database state to render assets and statistics; every mutation (asset CRUD, configuration, backup, import) runs permission checks in the backend and writes an audit record.
- **Where data lives** — the single-file SQLite database is the source of truth, uploaded files go to `data/files`, backups to `data/backups` and runtime logs to `data/logs`.
- **Sensitive values** — the JWT signing key persists in the `system_secrets` table, LDAP bind passwords are encrypted with a master key before storage, and API responses mask them.
- **External surface** — only health checks, sign-in, branding and the password-recovery flow are anonymous; every other endpoint requires `Authorization: Bearer <token>`.

## Tech stack

| Layer | Choice |
| --- | --- |
| Backend language | Go 1.25+ |
| HTTP router | [go-chi/chi](https://github.com/go-chi/chi) v5 |
| Database | SQLite ([modernc.org/sqlite](https://gitlab.com/cznic/sqlite), pure Go, no CGO) |
| Auth and crypto | JWT (golang-jwt/v5), AD/LDAP (go-ldap/ldap v3), WeCom OAuth QR-code sign-in (direct / unified auth-center SSO), bcrypt password hashing, AES for sensitive settings |
| Export and search | [excelize](https://github.com/qax-os/excelize) v2, mozillazg/go-pinyin |
| API docs | swag + http-swagger (Swagger UI) |
| Frontend framework | React 19 + TanStack Start / Router / Query |
| Language and build | TypeScript 5 + Vite 7 + Nitro |
| Styling and UI | Tailwind CSS v4, Radix UI, lucide-react, sonner |
| Charts and QR codes | Recharts, qrcode |
| Runtime packaging | Docker (Nginx + Supervisor) |

## Quick start

Local development needs **Go 1.25+** and **Node.js 20+**.

```bash
git clone https://github.com/zyx3721/itdb-new.git
cd itdb-new
```

**Backend**

```bash
cd backend
go mod download
cp .env.example .env      # adjust the listen address and secret as needed
go run cmd/server/main.go
```

The backend listens on `http://localhost:8080` and creates the database plus the default administrator `admin / admin123` on first start.

**Frontend** (in a second terminal)

```bash
cd frontend
npm install
npm run dev
```

The dev server runs on `http://localhost:5173` and proxies `/api` to `http://127.0.0.1:8080` (override with `VITE_API_BASE_URL` in `frontend/.env`).

Open `http://localhost:5173`, sign in with `admin / admin123`, and **change the password immediately**. Swagger lives at `http://localhost:8080/swagger/index.html`.

## Deployment

Two supported paths only: **Docker** (recommended) and **release binaries**. Source builds and bare systemd runs are documented as background detail in [chapter 4 of the full manual](docs/manual.md).

### Option 1: Docker

The image bundles the Go backend, Nginx and the Nitro SSR frontend under Supervisor, exposing only port 80. No clone needed — one command brings it up (replace `ITDB_JWT_SECRET` with a long random string):

```bash
docker run -d \
  --name itdb \
  --restart unless-stopped \
  -p 80:80 \
  -e ITDB_JWT_SECRET=replace-with-a-long-random-string \
  -e ITDB_HISTORY_LIMIT=1000 \
  -e ITDB_CORS_ORIGINS=* \
  -v "$(pwd)/data:/app/data" \
  registry.cn-shenzhen.aliyuncs.com/zyx3721/itdb-new:latest
```

> On Windows PowerShell use `-v "${PWD}/data:/app/data"` instead.

The command creates `data/` in the current directory and mounts it at `/app/data` inside the container:

```text
data/
├── itdb.db      # SQLite database (created on first start)
├── files/       # uploaded attachments
├── backups/     # manual and scheduled backups
└── logs/        # backend and frontend logs
```

Available environment variables:

| Variable | Default | Meaning |
| --- | --- | --- |
| `ITDB_JWT_SECRET` | empty | JWT signing key; when empty a random key is generated on first start and persisted in the database, so setting it explicitly is recommended |
| `ITDB_SESSION_TTL_HOURS` | `12` | Login session (JWT) lifetime in hours, positive integer |
| `ITDB_HISTORY_LIMIT` | `1000` | Audit history retention count; only limits rows kept in the history table, not rows returned by the audit log API |
| `ITDB_CORS_ORIGINS` | `*` | Allowed origins, comma-separated |
| `ITDB_SERVER_ADDR` | `127.0.0.1:8080` | Backend listen address; Nginx proxies to it inside the container, so leave it at the default |
| `ITDB_DB_PATH` | `data/itdb.db` | SQLite database path, relative to `/app` inside the container |
| `ITDB_UPLOAD_DIR` | `data/files` | Upload directory, relative to `/app` inside the container |

To run an image you built yourself, `deploy/Dockerfile` needs the repository sources:

```bash
git clone https://github.com/zyx3721/itdb-new.git && cd itdb-new/deploy
docker build -f Dockerfile -t itdb-new:latest --build-arg ALPINE_MIRROR=mirrors.aliyun.com ..
```

Then change the image name at the end of the command above to `itdb-new:latest` and recreate the container.

Day-to-day operations:

```bash
docker ps --filter name=itdb        # status
docker logs -f itdb                 # live logs
docker restart itdb                 # restart
docker stop itdb                    # stop
docker stop itdb && docker rm itdb  # stop and remove (data stays in ./data)

# Upgrade: pull the new image, then recreate the container with the command above
docker pull registry.cn-shenzhen.aliyuncs.com/zyx3721/itdb-new:latest
```

**Access**

- Console: `http://your-host/`, default account `admin / admin123`
- API docs: `http://your-host/swagger/index.html`
- Health check: `http://your-host/health`

To terminate TLS or host several sites on the host machine, publish a non-80 port (change `-p 80:80` to `-p 8080:80`) and put your own Nginx in front of the container. Full examples are in [section 3.7 of the manual](docs/manual.md).

### Option 2: Release binaries

Head to [GitHub Releases](https://github.com/zyx3721/itdb-new/releases), download the archive matching your operating system and CPU architecture, then follow the steps below to verify, unpack, configure and start.

**Which archive to download**

| Your machine | Download |
| --- | --- |
| Linux x86_64 | `itdb_<version>_linux_amd64.tar.gz` |
| Linux ARM64 (Kunpeng, Phytium, …) | `itdb_<version>_linux_arm64.tar.gz` |
| macOS on Intel | `itdb_<version>_darwin_amd64.tar.gz` |
| macOS on Apple silicon | `itdb_<version>_darwin_arm64.tar.gz` |
| Windows x86_64 | `itdb_<version>_windows_amd64.zip` |
| Windows ARM64 | `itdb_<version>_windows_arm64.zip` |
| Web console (needed on every platform) | `itdb-frontend_<version>.tar.gz` |
| Checksums | `SHA256SUMS` |

A backend archive contains the `itdb` executable (`itdb.exe` on Windows), `.env.example` and `README.txt`; the frontend archive contains the Nitro SSR `.output` bundle. The binaries have no external runtime dependencies and run as-is; the SSR frontend needs Node.js installed on the target machine.

**1. Verify the download**

```bash
VERSION=1.0.0
mkdir -p /data/itdb && cd /data/itdb
sha256sum -c SHA256SUMS
```

**2. Unpack**

```bash
mkdir -p backend frontend/.output
tar -xzf itdb_${VERSION}_linux_amd64.tar.gz -C backend --strip-components=1
tar -xzf itdb-frontend_${VERSION}.tar.gz -C frontend/.output
```

Resulting layout:

```text
/data/itdb/
├── backend/
│   ├── itdb            # backend binary
│   ├── .env.example
│   └── data/           # created on first start: itdb.db and files/
└── frontend/
    └── .output/
        ├── public/     # browser assets
        └── server/
            └── index.mjs   # Nitro SSR entry
```

**3. Configure and start the backend**

```bash
cd /data/itdb/backend
cp .env.example .env
vim .env               # set ITDB_JWT_SECRET at minimum
./itdb
```

The backend listens on `127.0.0.1:8080` and creates `data/itdb.db` and `data/files` relative to the working directory.

The backend also accepts command-line flags; a flag passed explicitly takes precedence over environment variables and the `.env` file. Run `./itdb -v` to print version information (version, commit, build date) and `./itdb -h` for the full list:

| Flag | Environment variable | Description |
| --- | --- | --- |
| `-addr` | `ITDB_SERVER_ADDR` | Backend listen address |
| `-port` | `PORT` | Backend listen port, used when `-addr` is not set |
| `-db` | `ITDB_DB_PATH` | SQLite database path |
| `-upload-dir` | `ITDB_UPLOAD_DIR` | Upload storage directory |
| `-jwt-secret` | `ITDB_JWT_SECRET` | JWT signing secret |
| `-history-limit` | `ITDB_HISTORY_LIMIT` | Operation history retention count; only limits rows kept in the history table, not rows returned by the audit log API |
| `-session-ttl` | `ITDB_SESSION_TTL_HOURS` | Session (JWT) validity in hours |
| `-cors-origins` | `ITDB_CORS_ORIGINS` | Allowed CORS origins, comma-separated |
| `-env` | — | Path to a custom `.env` file |
| `-v`, `-version` | — | Print version information and exit |

For a persistent service, use systemd:

```ini
# /etc/systemd/system/itdb-backend.service
[Unit]
Description=ITDB Backend
After=network.target

[Service]
Type=simple
WorkingDirectory=/data/itdb/backend
ExecStart=/data/itdb/backend/itdb
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```

```bash
systemctl daemon-reload && systemctl enable --now itdb-backend
```

**4. Start the SSR frontend**

```bash
cd /data/itdb/frontend
SSR_API_ORIGIN=http://127.0.0.1:8080 HOST=127.0.0.1 PORT=5173 node .output/server/index.mjs
```

When starting up, the SSR process fetches brand settings from the backend and renders the site name and icon into the first-frame HTML. The backend address is resolved in this order: the `SSR_API_ORIGIN` environment variable at runtime → a `.env` file in any directory from the `index.mjs` entry upward (the first existing one wins; it can share the same `.env` as the backend in the deploy root) → the default `http://127.0.0.1:8080` (in development mode it falls back to `VITE_API_BASE_URL` first, sharing the origin with the dev proxy). Make sure the address is reachable from the SSR process; otherwise the first frame falls back to the default brand before switching to the configured one, and the SSR process log prints an `[itdb-ssr] fetch brand failed` warning.

**5. Put Nginx in front**

```nginx
server {
    listen 80;
    server_name your-domain.com;
    client_max_body_size 500m;

    # /assets/ is also a page-route prefix; only take over static build files by extension
    location ~* ^/assets/.+\.(js|mjs|css|map|json|svg|png|jpe?g|gif|webp|ico|woff2?|ttf)$ {
        root /data/itdb/frontend/.output/public;
        try_files $uri =404;
        expires 1y;
        add_header Cache-Control "public, immutable";
    }

    location /api/ {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_read_timeout 600s;
    }

    location /swagger/ {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
    }

    location = /health {
        proxy_pass http://127.0.0.1:8080/api/health;
    }

    location / {
        proxy_pass http://127.0.0.1:5173;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

Pages must be served by `node .output/server/index.mjs`; **pointing a static root at `.output/public` alone breaks server-rendered pages**. Full examples including HTTPS and an 80 → 443 redirect are in [section 4.4 of the manual](docs/manual.md).

**6. Access**

Same as Docker: console at `http://your-domain.com` (`admin / admin123`), API docs at `/swagger/index.html`, health check at `/health`.

## Permission model

Every endpoint checks the caller's role permissions and returns 403 when they are missing; the `admin` user (`usertype=0`) holds all permissions. Permission keys look like `module.resource.action`, 44 in total.

| Identity | Default permissions |
| --- | --- |
| Default administrator `admin` | All 44; cannot be renamed, disabled or deleted |
| Built-in role `admin` | All 44 |
| Built-in role `operator` | View + manage on all eight assets and six dictionaries, the three label permissions, reports view + export, read on browse/audit and read on the four system settings; no settings management |
| Built-in role `viewer` | Read-only everywhere: eight assets, six dictionaries, label preview, reports, browse, audit and settings |
| Custom roles / user groups | Pick from the 44; ticking any `manage` automatically adds the matching `read` |

By module:

- **Assets** — `assets.{items|software|invoices|agents|files|contracts|locations|racks}.read` / `.manage`
- **Dictionaries** — `dictionaries.{itemtypes|contracttypes|statustypes|filetypes|dpttypes|tags}.read` / `.manage`
- **Labels** — `labels.preview` (preview), `labels.print` (print, implies preview), `labels.manage` (presets, implies preview)
- **Reports** — `reports.read` (view), `reports.manage` (export, implies view); **browse** — `browse.read`
- **Audit** — `audit.read` (view), `audit.manage` (export, implies view)
- **Settings** — `settings.{base|users|auth|notifications}.read` / `.manage`; database backup download and import require `settings.base.manage`
- **Aggregates** — `GET /api/bootstrap` and `GET /api/dashboard/summary` require any read permission

## Data and security

```text
Console accounts sign in, change configuration, manage assets
        +
Bearer tokens checked against 44 permissions per endpoint, 403 otherwise
        +
JWT key and LDAP bind password encrypted at rest, masked in responses
        +
Repository never carries tokens / passwords / real hostnames / customer names
```

- **Change the default password first** — update `admin`'s password right after the first deployment.
- **Set the JWT secret explicitly** — when unset, a random key is generated and persisted (dropping the database invalidates every session); multi-instance deployments must set `ITDB_JWT_SECRET`.
- **Enable HTTPS** — terminate TLS in Nginx in production; see [section 4.4.2 of the manual](docs/manual.md).
- **Tighten CORS** — configure `ITDB_CORS_ORIGINS` for your origins instead of leaving `*`.
- **Restrict access** — limit the reachable IP range through Nginx or a firewall.
- **Protect the data directory** — `data/` holds the database, attachments and backups; grant it only to the service account. `.env` is already excluded by `.gitignore`.
- **Keep an off-site copy** — add an off-site backup on top of the built-in scheduler.

## API docs

The backend ships Swagger/OpenAPI, so the live reference is available once the service is up:

- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- **OpenAPI JSON**: `http://localhost:8080/swagger/doc.json`
- **Health checks**: `GET /health`, `GET /api/health`

Anonymous endpoints are limited to `POST /api/auth/login`, `GET /api/auth/providers`, `GET /api/auth/wecom/authorize`, `POST /api/auth/wecom/callback`, `POST /api/auth/wecom/sso/callback`, `GET /api/auth/password-reset/captcha`, `POST /api/auth/password-reset/verify`, `POST /api/auth/password-reset/send`, `POST /api/auth/password-reset/confirm`, `GET /api/public/base`, `GET /health` and `GET /api/health`; everything else needs `Authorization: Bearer <token>`.

Example sign-in request:

```json
{
  "username": "admin",
  "password": "admin123",
  "mode": "local"
}
```

The complete endpoint list grouped by module (auth, vendors, hardware, software, invoices, contracts, files, locations, racks, dictionaries, labels, reports, browse, audit, backup, import, settings) is in [chapter 5 of the manual](docs/manual.md).

After changing endpoints, regenerate the Swagger artifacts from `backend/`:

```bash
swag init -g cmd/server/main.go -o docs
```

## Database

A single-file SQLite database, by default at `backend/data/itdb.db`, with 49 tables (36 created during initialisation, 13 created on demand by the settings module).

| Group | Tables |
| --- | --- |
| Core business | `items`, `software`, `contracts`, `invoices`, `files`, `agents`, `users`, `locations`, `racks`, `locareas`, `labelpapers`, `tags`, `actions`, `contractevents` |
| Relations | `item2inv`, `item2soft`, `item2file`, `itemlink`, `contract2item`, `contract2soft`, `contract2inv`, `contract2file`, `invoice2file`, `soft2inv`, `software2file`, `tag2item`, `tag2software` |
| Dictionaries | `itemtypes`, `contracttypes`, `contractsubtypes`, `dpttypes`, `statustypes`, `filetypes` |
| System and settings | `settings`, `system_secrets`, `history`, `settings_base`, `settings_email`, `settings_auth_providers`, `settings_roles`, `settings_role_status`, `settings_user_profiles`, `settings_user_groups`, `settings_user_group_members`, `settings_user_group_roles`, `settings_user_roles` |
| Password recovery | `password_reset_captchas`, `password_reset_requests`, `password_reset_send_log` |

Column-level detail is in [chapter 6 of the manual](docs/manual.md).

## FAQ

**I forgot the administrator password.**

Stop the backend and update the database directly (recommended, no data loss):

```bash
sqlite3 backend/data/itdb.db "UPDATE users SET pass = 'admin123' WHERE username = 'admin';"
```

The backend upgrades the plaintext password to bcrypt on the next start.

**How long do JWTs live?**

12 hours by default; adjust with the `ITDB_SESSION_TTL_HOURS` environment variable (in hours, positive integer) and restart the backend. Sessions are also recorded in the server-side `user_sessions` table: signing out deletes the corresponding session immediately and invalidates the token, and expired sessions are cleaned up automatically on the next login; resetting `ITDB_JWT_SECRET` invalidates all issued sessions immediately.

**Can I migrate by copying the database file?**

Yes. SQLite is a single file: stop the service, then copy `itdb.db` and `data/files`. The built-in import feature can also swap databases online.

**How do I migrate from the legacy PHP ITDB?**

Use *System settings → Data backup → Import database* and upload the old `.db` or a `.zip` including attachments. The old schema is detected and converted: only assets and dictionaries are migrated (hardware maintenance logs are not migrated), users are merged by username (existing accounts keep their current data; newly imported users get the builtin admin or viewer role based on their legacy user type), while system configuration, user role profiles and audit logs stay untouched. A pre-import backup is created first (uploaded files referenced by the current database are bundled into a zip) and bundled files are cleaned up after a successful import; the import is rejected with a prompt when the database file is locked by an external tool. See [section 7.4 of the manual](docs/manual.md) for details.

**How do I enable LDAP sign-in?**

First fill in the LDAP connection parameters under *System settings → Authentication* and enable it; then create a user whose name matches the LDAP `sAMAccountName`. If the local user table has no such user, a correct LDAP password alone still cannot sign in.

**How do I enable WeCom QR-code sign-in?**

Create a self-built app in the WeCom admin console and note the AgentID and Secret; set this system's domain as the trusted callback domain under *Web authorization & JS-SDK* and add this service's egress IP to the trusted IPs. Then fill in the Corp ID, AgentID and Secret under *System settings → Authentication → WeCom* and enable it (the callback prefix can be left empty to infer from the current access address). Users sign in with username and password first, bind their WeCom account via *Bind WeCom* in the top-right menu, and can then choose WeCom QR-code sign-in on the login page; the login page renders the QR code inline by default and automatically falls back to full-page redirection when the configuration misbehaves.

If your organization already runs the unified auth center (wecom-auth-center), switch *Authentication mode* to *Unified auth center* in the WeCom settings and fill in the auth-center URL, app ID and app secret; on the auth-center side just add an `apps` entry for this system with `domain` (this system's external address) and `callback_path` (`/wecom-qr-callback`, the shared relay route for both the embedded QR code and full-page flow, and the auth-center login page must allow being embedded in an iframe by this system). Multiple internal systems can then share one WeCom app configuration while binding and sign-in behave exactly the same.

**Where do scheduled backups go?**

`data/backups/`. When no uploaded file is referenced the output is `itdb-YYYYMMDD-HHMMSS.db`; when attachments are referenced it becomes a `.zip` of the same name (database at the archive root, attachments under `files/`). Expired files are pruned by retention days. Enable it and set the cron expression under *System settings → Base configuration → Data backup*.

**Is there an upload size limit?**

The backend imposes none by default; behind Nginx set `client_max_body_size` (500m in the examples).

**Why can't I serve the built frontend as static files only?**

Because the console is a TanStack Start + Nitro SSR application: pages are rendered by `node .output/server/index.mjs`, and a static root only serves assets such as `/assets/`.

More questions are covered in [chapter 7 of the manual](docs/manual.md).

## Repository layout

```text
itdb/
├── backend/                 Go backend service
│   ├── cmd/server/          service entry point
│   ├── config/              environment and runtime configuration
│   ├── docs/                generated Swagger/OpenAPI artifacts
│   ├── internal/            domain models, repositories, workflows, security, build info
│   ├── pkg/database/        SQLite connection and row-scanning infrastructure
│   ├── router/              route assembly, middleware, permission checks, Swagger annotations
│   │   ├── assets/          asset endpoints (hardware/software/invoices/vendors/files/contracts/locations/racks/dictionaries/tags/reports/browse)
│   │   ├── auth/            auth endpoints (sign-in, password recovery, password change, users)
│   │   ├── common/          responses, auth context, uploads, database initialisation
│   │   ├── settings/        settings endpoints (base config, users/roles/groups, auth, email)
│   │   └── system/          system endpoints (database backup, import, audit history)
│   ├── data/                SQLite database, uploads and backups (created at runtime)
│   └── .env.example         environment template
├── deploy/                  Docker image build and compose files
├── docs/                    full manual and English documentation
├── frontend/                React console
│   └── src/
│       ├── components/      layout, boot screen, dialogs, base UI
│       ├── features/        auth, assets, audit, dashboard, settings, tools
│       ├── lib/             session auth, branding, exports, shared utilities
│       ├── routes/          TanStack Router pages
│       ├── router.tsx       router instance
│       ├── server.ts        server entry
│       ├── start.ts         client entry
│       └── styles.css       global styles and theme variables
├── scripts/                 ops and online update scripts
├── .github/                 GitHub Actions workflows and preview images
├── .dockerignore            Docker build ignore rules
├── .gitignore               Git ignore rules
├── LICENSE
├── README.md                简体中文
├── README.en.md             English (this file)
└── docs/manual.md           Full manual (every endpoint, table and deployment detail)
```

## Documentation

| Start here | Then |
| --- | --- |
| [Quick start](#quick-start) | Run the backend and frontend locally; default account and ports |
| [Deployment](#deployment) | Docker and release binaries, environment variables, reverse proxy |
| [Permission model](#permission-model) | How the 44 permissions are grouped and what each built-in role holds |
| [Full manual](docs/manual.md) | Every endpoint, all 49 tables, Nginx and HTTPS examples, migration and troubleshooting (Chinese) |
| [简体中文 README](README.md) | The same content in Chinese |

## Releases

| Version | Date | Changelog |
| --- | --- | --- |
| v1.4.2 | 2026-10-22 | [verchanglog/v1.4.2.md](verchanglog/v1.4.2.md) |
| v1.4.1 | 2026-09-24 | [verchanglog/v1.4.1.md](verchanglog/v1.4.1.md) |
| v1.4.0 | 2026-09-24 | [verchanglog/v1.4.0.md](verchanglog/v1.4.0.md) |
| v1.3.7 | 2026-09-23 | [verchanglog/v1.3.7.md](verchanglog/v1.3.7.md) (Chinese) |
| v1.3.6 | 2026-09-23 | [verchanglog/v1.3.6.md](verchanglog/v1.3.6.md) (Chinese) |
| v1.3.5 | 2026-09-22 | [verchanglog/v1.3.5.md](verchanglog/v1.3.5.md) (Chinese) |
| v1.3.4 | 2026-09-22 | [verchanglog/v1.3.4.md](verchanglog/v1.3.4.md) (Chinese) |
| v1.3.3 | 2026-09-22 | [verchanglog/v1.3.3.md](verchanglog/v1.3.3.md) (Chinese) |
| v1.3.2 | 2026-09-19 | [verchanglog/v1.3.2.md](verchanglog/v1.3.2.md) (Chinese) |
| v1.3.1 | 2026-09-19 | [verchanglog/v1.3.1.md](verchanglog/v1.3.1.md) (Chinese) |
| v1.3.0 | 2026-09-19 | [verchanglog/v1.3.0.md](verchanglog/v1.3.0.md) (Chinese) |
| v1.2.1 | 2026-09-18 | [verchanglog/v1.2.1.md](verchanglog/v1.2.1.md) (Chinese) |
| v1.2.0 | 2026-09-18 | [verchanglog/v1.2.0.md](verchanglog/v1.2.0.md) (Chinese) |
| v1.1.2 | 2026-09-18 | [verchanglog/v1.1.2.md](verchanglog/v1.1.2.md) (Chinese) |
| v1.1.1 | 2026-09-18 | [verchanglog/v1.1.1.md](verchanglog/v1.1.1.md) (Chinese) |
| v1.1.0 | 2026-09-17 | [verchanglog/v1.1.0.md](verchanglog/v1.1.0.md) (Chinese) |
| v1.0.0 | 2026-09-17 | [verchanglog/v1.0.0.md](verchanglog/v1.0.0.md) (Chinese) |

Build assets and release notes for every version live on [GitHub Releases](https://github.com/zyx3721/itdb-new/releases).

## Acknowledgements

Thanks to these open-source projects and communities:

- [go-chi/chi](https://github.com/go-chi/chi) — lightweight, fast Go HTTP router
- [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) — pure-Go SQLite driver
- [swaggo/swag](https://github.com/swaggo/swag) — Swagger documentation generator
- [excelize](https://github.com/qax-os/excelize) — Excel read/write library
- [TanStack](https://tanstack.com/) — Router / Query / Start suite
- [Tailwind CSS](https://tailwindcss.com/) — utility-first CSS framework

And to the original [zyx3721/itdb](https://github.com/zyx3721/itdb/) for the domain model this rewrite builds on.

## License

Released under the [MIT License](LICENSE): use, copy, modify, merge, publish, distribute, sublicense and sell freely, as long as the copyright and licence notice are kept in all copies or substantial portions.

## Contact

- **Email**: 416685476@qq.com
- **GitHub Issues**: [zyx3721/itdb-new/issues](https://github.com/zyx3721/itdb-new/issues)
- **Project home**: [github.com/zyx3721/itdb-new](https://github.com/zyx3721/itdb-new)

---

**⭐ If this project helps you, a star is appreciated!**
