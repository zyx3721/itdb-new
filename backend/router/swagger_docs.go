package router

// swaggerHealthRoot documents GET /health.
// @Summary 健康检查
// @Tags 健康检查
// @Produce json
// @Success 200 {object} swaggerHealthResponse
// @Description 探测服务进程是否存活，无需认证
// @Router /health [get]
func swaggerHealthRoot() {}

// swaggerHealthAPI documents GET /api/health.
// @Summary API 健康检查
// @Tags 健康检查
// @Produce json
// @Success 200 {object} swaggerHealthResponse
// @Description 探测服务进程是否存活，无需认证
// @Router /api/health [get]
func swaggerHealthAPI() {}

// swaggerPublicAuthProviders documents GET /api/auth/providers.
// @Summary 获取公开认证方式
// @Tags 认证
// @Produce json
// @Success 200 {object} swaggerPublicAuthProvidersResponse
// @Description 获取登录页可用的认证方式与找回密码开关，无需认证
// @Router /api/auth/providers [get]
func swaggerPublicAuthProviders() {}

// swaggerLogin documents POST /api/auth/login.
// @Summary 用户登录
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body authLoginRequest true "登录信息"
// @Success 200 {object} authLoginResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 401 {object} swaggerErrorResponse
// @Failure 429 {object} swaggerErrorResponse
// @Description 用户登录，支持本地密码与 AD/LDAP 两种方式，成功返回令牌与用户信息；连续密码失败达到安全时效配置的锁定次数后返回 429，admin 不受限
// @Router /api/auth/login [post]
func swaggerLogin() {}

// swaggerWecomAuthorize documents GET /api/auth/wecom/authorize.
// @Summary 获取企业微信扫码登录地址
// @Tags 认证
// @Produce json
// @Success 200 {object} swaggerWecomAuthorizeResponse
// @Failure 400 {object} swaggerErrorResponse
// @Description 生成企业微信 Web 扫码登录页地址（含防伪 state）与内嵌二维码渲染参数（iframe_url、回跳路径 /wecom-qr-callback），前端可选择内嵌渲染或整页跳转；需已启用企业微信认证
// @Router /api/auth/wecom/authorize [get]
func swaggerWecomAuthorize() {}

// swaggerWecomCallback documents POST /api/auth/wecom/callback.
// @Summary 企业微信扫码登录回调
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body swaggerWecomCallbackRequest true "扫码回调载荷"
// @Success 200 {object} authLoginResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 401 {object} swaggerErrorResponse
// @Description 校验 state 后用授权码换取企业微信成员身份，按绑定关系登录并返回令牌与用户信息；未绑定时返回 401
// @Router /api/auth/wecom/callback [post]
func swaggerWecomCallback() {}

// swaggerWecomBindURL documents GET /api/auth/wecom/bind-url.
// @Summary 获取企业微信绑定扫码地址
// @Tags 认证
// @Produce json
// @Success 200 {object} swaggerWecomAuthorizeResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 401 {object} swaggerErrorResponse
// @Description 为当前登录用户生成企业微信绑定扫码地址，state 绑定当前用户；需已启用企业微信认证
// @Security BearerAuth
// @Router /api/auth/wecom/bind-url [get]
func swaggerWecomBindURL() {}

// swaggerWecomBind documents POST /api/auth/wecom/bind.
// @Summary 绑定企业微信账号
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body swaggerWecomCallbackRequest true "扫码回调载荷"
// @Success 200 {object} swaggerWecomBindResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 401 {object} swaggerErrorResponse
// @Failure 403 {object} swaggerErrorResponse
// @Failure 409 {object} swaggerErrorResponse
// @Description 校验绑定 state 后用授权码换取企业微信成员 userid，与当前用户建立绑定；该企微已绑定其他用户时返回 409
// @Security BearerAuth
// @Router /api/auth/wecom/bind [post]
func swaggerWecomBind() {}

// swaggerWecomUnbind documents DELETE /api/auth/wecom/bind.
// @Summary 解绑企业微信账号
// @Tags 认证
// @Produce json
// @Success 200 {object} swaggerOKResponse
// @Failure 401 {object} swaggerErrorResponse
// @Description 解除当前登录用户的企业微信绑定
// @Security BearerAuth
// @Router /api/auth/wecom/bind [delete]
func swaggerWecomUnbind() {}

// swaggerWecomSSOCallback documents POST /api/auth/wecom/sso/callback.
// @Summary 统一认证中心扫码登录回调
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body swaggerWecomSSORequest true "认证中心回跳凭证"
// @Success 200 {object} authLoginResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 401 {object} swaggerErrorResponse
// @Description 企业微信统一认证（SSO）模式：校验认证中心回跳的一次性 ticket 后按绑定关系登录并返回令牌与用户信息；需已启用统一认证模式
// @Router /api/auth/wecom/sso/callback [post]
func swaggerWecomSSOCallback() {}

// swaggerWecomSSOBind documents POST /api/auth/wecom/sso/bind.
// @Summary 统一认证中心扫码绑定
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body swaggerWecomSSORequest true "认证中心回跳凭证"
// @Success 200 {object} swaggerWecomBindResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 401 {object} swaggerErrorResponse
// @Failure 409 {object} swaggerErrorResponse
// @Description 企业微信统一认证（SSO）模式：校验认证中心一次性 ticket 后与当前登录用户建立绑定；该企微已绑定其他用户时返回 409
// @Security BearerAuth
// @Router /api/auth/wecom/sso/bind [post]
func swaggerWecomSSOBind() {}

// swaggerMe documents GET /api/auth/me.
// @Summary 获取当前用户
// @Tags 认证
// @Produce json
// @Success 200 {object} swaggerGenericObject
// @Failure 401 {object} swaggerErrorResponse
// @Description 获取当前登录用户的资料、直接角色、有效角色与权限清单
// @Security BearerAuth
// @Router /api/auth/me [get]
func swaggerMe() {}

// swaggerLogout documents POST /api/auth/logout.
// @Summary 登出当前会话
// @Tags 认证
// @Produce json
// @Success 200 {object} swaggerOKResponse
// @Description 删除服务端会话记录使令牌立即失效，并写入用户注销审计
// @Security BearerAuth
// @Router /api/auth/logout [post]
func swaggerLogout() {}

// swaggerBootstrap documents GET /api/bootstrap.
// @Summary 获取前端启动字典数据
// @Tags 启动数据
// @Produce json
// @Success 200 {object} swaggerGenericObject
// @Description 获取前端启动所需的字典、引用选项与权限数据
// @Security BearerAuth
// @Router /api/bootstrap [get]
func swaggerBootstrap() {}

// swaggerDashboardSummary documents GET /api/dashboard/summary.
// @Summary 获取仪表盘统计
// @Tags 仪表盘
// @Produce json
// @Success 200 {object} swaggerDashboardSummaryResponse
// @Description 获取仪表盘各资源总数统计，需要任意查看权限
// @Security BearerAuth
// @Router /api/dashboard/summary [get]
func swaggerDashboardSummary() {}

// swaggerHistory documents GET /api/history.
// @Summary 获取审计日志
// @Tags 审计日志
// @Produce json
// @Param limit query int false "返回条数（默认返回全部保留记录；传 -1 表示不限制，不设上限）"
// @Success 200 {object} swaggerAuditHistoryResponse
// @Description 获取审计日志；导入类事件的目标仅返回前两个名称，完整清单通过 targetTitle 返回供悬浮提示
// @Security BearerAuth
// @Router /api/history [get]
func swaggerHistory() {}

// swaggerHistoryClear documents POST /api/history/clear.
// @Summary 清空审计日志
// @Tags 审计日志
// @Accept json
// @Produce json
// @Param body body swaggerHistoryClearRequest true "待清空的审计日志编号清单"
// @Success 200 {object} swaggerHistoryClearResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 403 {object} swaggerErrorResponse
// @Description 按前端筛选结果批量清空审计日志；清空成功后自动补记一条「清空审计日志」审计事件（归属「审计日志」模块，目标为清除条数）
// @Security BearerAuth
// @Router /api/history/clear [post]
func swaggerHistoryClear() {}

// swaggerBackupDatabase documents GET /api/backups/database.
// @Summary 下载数据库备份
// @Tags 备份
// @Produce application/octet-stream
// @Success 200 {file} file
// @Description 下载数据库备份（itdb-日期.db）；携带 files=1 时按数据库实际引用的文件打包为 zip（文件位于包内 files 目录），引用文件全部不存在或未引用时仅导出 db 文件，缺失文件写入备份审计
// @Security BearerAuth
// @Router /api/backups/database [get]
func swaggerBackupDatabase() {}

func swaggerBackupFull() {}

// swaggerSettingsBaseGet documents GET /api/settings/base.
// @Summary 获取系统基础配置
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerSystemBaseConfig
// @Failure 403 {object} swaggerErrorResponse
// @Description 获取系统基础配置
// @Security BearerAuth
// @Router /api/settings/base [get]
func swaggerSettingsBaseGet() {}

// swaggerSettingsBaseUpdate documents PUT /api/settings/base.
// @Summary 更新系统基础配置
// @Tags 系统设置
// @Accept json
// @Produce json
// @Param body body swaggerSystemBaseConfig true "系统基础配置"
// @Success 200 {object} swaggerSystemBaseConfig
// @Failure 400 {object} swaggerErrorResponse
// @Failure 403 {object} swaggerErrorResponse
// @Description 更新系统基础配置；携带 section（brand/security/backup）时按区块局部更新，审计目标细分为品牌标识、安全时效、数据备份
// @Security BearerAuth
// @Router /api/settings/base [put]
func swaggerSettingsBaseUpdate() {}

// swaggerPublicBaseGet documents GET /api/public/base.
// @Summary 获取公开品牌标识
// @Tags 认证
// @Produce json
// @Success 200 {object} swaggerPublicBrandingResponse
// @Description 获取登录页与启动屏使用的品牌标识，无需认证
// @Router /api/public/base [get]
func swaggerPublicBaseGet() {}

// swaggerItemsList documents GET /api/items.
// @Summary 获取硬件资产列表
// @Tags 硬件资产
// @Produce json
// @Param search query string false "搜索关键词"
// @Param limit query int false "返回数量，-1 表示全部"
// @Param offset query int false "偏移量"
// @Success 200 {array} swaggerGenericObject
// @Description 获取硬件资产列表，支持关键字搜索与全量分页参数
// @Security BearerAuth
// @Router /api/items [get]
func swaggerItemsList() {}

// swaggerItemsGet documents GET /api/items/{id}.
// @Summary 获取硬件资产详情
// @Tags 硬件资产
// @Produce json
// @Param id path int true "硬件资产 ID"
// @Success 200 {object} swaggerGenericObject
// @Failure 404 {object} swaggerErrorResponse
// @Description 按编号获取硬件资产详情及全部关联
// @Security BearerAuth
// @Router /api/items/{id} [get]
func swaggerItemsGet() {}

// swaggerItemsCreate documents POST /api/items.
// @Summary 创建硬件资产
// @Tags 硬件资产
// @Accept json
// @Produce json
// @Param body body itemPayload true "硬件资产"
// @Success 201 {object} swaggerIDResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 403 {object} swaggerErrorResponse
// @Description 创建硬件资产；机架位置冲突、软件授权超限等会返回 409
// @Security BearerAuth
// @Router /api/items [post]
func swaggerItemsCreate() {}

// swaggerItemsUpdate documents PUT /api/items/{id}.
// @Summary 更新硬件资产
// @Tags 硬件资产
// @Accept json
// @Produce json
// @Param id path int true "硬件资产 ID"
// @Param body body itemPayload true "硬件资产"
// @Success 200 {object} swaggerIDResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 403 {object} swaggerErrorResponse
// @Description 更新硬件资产；无任何修改时不写审计，维护日志记录变更项
// @Security BearerAuth
// @Router /api/items/{id} [put]
func swaggerItemsUpdate() {}

// swaggerItemsDelete documents DELETE /api/items/{id}.
// @Summary 删除硬件资产
// @Tags 硬件资产
// @Produce json
// @Param id path int true "硬件资产 ID"
// @Success 200 {object} swaggerOKResponse
// @Failure 403 {object} swaggerErrorResponse
// @Description 删除硬件资产并清理关联、维护日志与标记关联；被交换机引用时返回 409
// @Security BearerAuth
// @Router /api/items/{id} [delete]
func swaggerItemsDelete() {}

// swaggerItemTagsMutate documents POST /api/items/{id}/tags.
// @Summary 关联或移除硬件标签
// @Tags 硬件资产
// @Accept json
// @Produce json
// @Param id path int true "硬件资产 ID"
// @Param body body tagMutationPayload true "标签操作"
// @Success 200 {object} swaggerOKResponse
// @Description 关联或解除硬件标记（action=add/remove）
// @Security BearerAuth
// @Router /api/items/{id}/tags [post]
func swaggerItemTagsMutate() {}

// swaggerItemActionsList documents GET /api/items/{id}/actions.
// @Summary 获取硬件操作记录
// @Tags 硬件操作记录
// @Produce json
// @Param id path int true "硬件资产 ID"
// @Success 200 {array} swaggerGenericObject
// @Description 获取硬件的维护操作记录
// @Security BearerAuth
// @Router /api/items/{id}/actions [get]
func swaggerItemActionsList() {}

// swaggerTagsNextID documents GET /api/tags/next-id.
// @Summary 获取下一个标记编号
// @Tags 标记
// @Produce json
// @Success 200 {object} swaggerGenericObject
// @Description 获取下一个可用的标记编号
// @Security BearerAuth
// @Router /api/tags/next-id [get]
func swaggerTagsNextID() {}

// swaggerSoftwareList documents GET /api/software.
// @Summary 获取软件许可列表
// @Tags 软件许可
// @Produce json
// @Param search query string false "搜索关键词"
// @Param limit query int false "返回数量，-1 表示全部"
// @Param offset query int false "偏移量"
// @Success 200 {array} swaggerGenericObject
// @Description 获取软件许可列表，支持关键字搜索
// @Security BearerAuth
// @Router /api/software [get]
func swaggerSoftwareList() {}

// swaggerSoftwareGet documents GET /api/software/{id}.
// @Summary 获取软件许可详情
// @Tags 软件许可
// @Produce json
// @Param id path int true "软件许可 ID"
// @Success 200 {object} swaggerGenericObject
// @Description 按编号获取软件许可详情及全部关联
// @Security BearerAuth
// @Router /api/software/{id} [get]
func swaggerSoftwareGet() {}

// swaggerSoftwareCreate documents POST /api/software.
// @Summary 创建软件许可
// @Tags 软件许可
// @Accept json
// @Produce json
// @Param body body softwarePayload true "软件许可"
// @Success 201 {object} swaggerIDResponse
// @Description 创建软件许可；标题加版本不允许重复，授权口径校验失败返回 409
// @Security BearerAuth
// @Router /api/software [post]
func swaggerSoftwareCreate() {}

// swaggerSoftwareUpdate documents PUT /api/software/{id}.
// @Summary 更新软件许可
// @Tags 软件许可
// @Accept json
// @Produce json
// @Param id path int true "软件许可 ID"
// @Param body body softwarePayload true "软件许可"
// @Success 200 {object} swaggerIDResponse
// @Description 更新软件许可；无任何修改时不写审计
// @Security BearerAuth
// @Router /api/software/{id} [put]
func swaggerSoftwareUpdate() {}

// swaggerSoftwareDelete documents DELETE /api/software/{id}.
// @Summary 删除软件许可
// @Tags 软件许可
// @Produce json
// @Param id path int true "软件许可 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除软件许可及其关联
// @Security BearerAuth
// @Router /api/software/{id} [delete]
func swaggerSoftwareDelete() {}

// swaggerSoftwareTagsMutate documents POST /api/software/{id}/tags.
// @Summary 关联或移除软件标签
// @Tags 软件许可
// @Accept json
// @Produce json
// @Param id path int true "软件许可 ID"
// @Param body body tagMutationPayload true "标签操作"
// @Success 200 {object} swaggerOKResponse
// @Description 关联或解除软件标记（action=add/remove）
// @Security BearerAuth
// @Router /api/software/{id}/tags [post]
func swaggerSoftwareTagsMutate() {}

// swaggerInvoicesList documents GET /api/invoices.
// @Summary 获取单据列表
// @Tags 单据
// @Produce json
// @Param search query string false "搜索关键词"
// @Param limit query int false "返回数量，-1 表示全部"
// @Param offset query int false "偏移量"
// @Success 200 {array} swaggerGenericObject
// @Description 获取单据（发票）列表，支持关键字搜索
// @Security BearerAuth
// @Router /api/invoices [get]
func swaggerInvoicesList() {}

// swaggerInvoicesGet documents GET /api/invoices/{id}.
// @Summary 获取单据详情
// @Tags 单据
// @Produce json
// @Param id path int true "发票 ID"
// @Success 200 {object} swaggerGenericObject
// @Description 按编号获取单据详情及全部关联
// @Security BearerAuth
// @Router /api/invoices/{id} [get]
func swaggerInvoicesGet() {}

// swaggerInvoicesCreate documents POST /api/invoices.
// @Summary 创建单据
// @Tags 单据
// @Accept json
// @Produce json
// @Param body body invoicePayload true "发票"
// @Success 201 {object} swaggerIDResponse
// @Description 创建单据
// @Security BearerAuth
// @Router /api/invoices [post]
func swaggerInvoicesCreate() {}

// swaggerInvoicesUpdate documents PUT /api/invoices/{id}.
// @Summary 更新单据
// @Tags 单据
// @Accept json
// @Produce json
// @Param id path int true "发票 ID"
// @Param body body invoicePayload true "发票"
// @Success 200 {object} swaggerIDResponse
// @Description 更新单据；无任何修改时不写审计
// @Security BearerAuth
// @Router /api/invoices/{id} [put]
func swaggerInvoicesUpdate() {}

// swaggerInvoicesDelete documents DELETE /api/invoices/{id}.
// @Summary 删除单据
// @Tags 单据
// @Produce json
// @Param id path int true "发票 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除单据及其关联
// @Security BearerAuth
// @Router /api/invoices/{id} [delete]
func swaggerInvoicesDelete() {}

// swaggerContractsList documents GET /api/contracts.
// @Summary 获取合同列表
// @Tags 合同
// @Produce json
// @Param search query string false "搜索关键词"
// @Param limit query int false "返回数量，-1 表示全部"
// @Param offset query int false "偏移量"
// @Success 200 {array} swaggerGenericObject
// @Description 获取合同列表，支持关键字搜索
// @Security BearerAuth
// @Router /api/contracts [get]
func swaggerContractsList() {}

// swaggerContractsGet documents GET /api/contracts/{id}.
// @Summary 获取合同详情
// @Tags 合同
// @Produce json
// @Param id path int true "合同 ID"
// @Success 200 {object} swaggerGenericObject
// @Description 按编号获取合同详情及全部关联
// @Security BearerAuth
// @Router /api/contracts/{id} [get]
func swaggerContractsGet() {}

// swaggerContractsCreate documents POST /api/contracts.
// @Summary 创建合同
// @Tags 合同
// @Accept json
// @Produce json
// @Param body body contractPayload true "合同"
// @Success 201 {object} swaggerIDResponse
// @Description 创建合同；合同标题不允许重复
// @Security BearerAuth
// @Router /api/contracts [post]
func swaggerContractsCreate() {}

// swaggerContractsUpdate documents PUT /api/contracts/{id}.
// @Summary 更新合同
// @Tags 合同
// @Accept json
// @Produce json
// @Param id path int true "合同 ID"
// @Param body body contractPayload true "合同"
// @Success 200 {object} swaggerIDResponse
// @Description 更新合同；备件、事件历史随表单一并保存，重名标题返回 409
// @Security BearerAuth
// @Router /api/contracts/{id} [put]
func swaggerContractsUpdate() {}

// swaggerContractsDelete documents DELETE /api/contracts/{id}.
// @Summary 删除合同
// @Tags 合同
// @Produce json
// @Param id path int true "合同 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除合同；被下级合同引用时返回 409
// @Security BearerAuth
// @Router /api/contracts/{id} [delete]
func swaggerContractsDelete() {}

// swaggerContractNextEventID documents GET /api/contracts/next-event-id.
// @Summary 获取下一个合同事件编号
// @Tags 合同事件
// @Produce json
// @Success 200 {object} swaggerGenericObject
// @Description 获取下一个可用的合同事件编号
// @Security BearerAuth
// @Router /api/contracts/next-event-id [get]
func swaggerContractNextEventID() {}

// swaggerContractEventsList documents GET /api/contracts/{id}/events.
// @Summary 获取合同事件
// @Tags 合同事件
// @Produce json
// @Param id path int true "合同 ID"
// @Success 200 {array} swaggerGenericObject
// @Description 获取指定合同的事件历史列表
// @Security BearerAuth
// @Router /api/contracts/{id}/events [get]
func swaggerContractEventsList() {}

// swaggerContractEventsCreate documents POST /api/contracts/{id}/events.
// @Summary 创建合同事件
// @Tags 合同事件
// @Accept json
// @Produce json
// @Param id path int true "合同 ID"
// @Param body body contractEventPayload true "合同事件"
// @Success 201 {object} swaggerIDResponse
// @Description 为指定合同新增一条事件历史
// @Security BearerAuth
// @Router /api/contracts/{id}/events [post]
func swaggerContractEventsCreate() {}

// swaggerContractEventsUpdate documents PUT /api/contracts/{id}/events/{eventId}.
// @Summary 更新合同事件
// @Tags 合同事件
// @Accept json
// @Produce json
// @Param id path int true "合同 ID"
// @Param eventId path int true "事件 ID"
// @Param body body contractEventPayload true "合同事件"
// @Success 200 {object} swaggerIDResponse
// @Description 更新指定的事件历史
// @Security BearerAuth
// @Router /api/contracts/{id}/events/{eventId} [put]
func swaggerContractEventsUpdate() {}

// swaggerContractEventsDelete documents DELETE /api/contracts/{id}/events/{eventId}.
// @Summary 删除合同事件
// @Tags 合同事件
// @Produce json
// @Param id path int true "合同 ID"
// @Param eventId path int true "事件 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除指定的事件历史
// @Security BearerAuth
// @Router /api/contracts/{id}/events/{eventId} [delete]
func swaggerContractEventsDelete() {}

// swaggerFilesList documents GET /api/files.
// @Summary 获取文件列表
// @Tags 文件
// @Produce json
// @Param search query string false "搜索关键词"
// @Success 200 {array} swaggerGenericObject
// @Description 获取文件列表，支持关键字搜索
// @Security BearerAuth
// @Router /api/files [get]
func swaggerFilesList() {}

// swaggerFilesGet documents GET /api/files/{id}.
// @Summary 获取文件详情
// @Tags 文件
// @Produce json
// @Param id path int true "文件 ID"
// @Success 200 {object} swaggerGenericObject
// @Description 按编号获取文件详情及全部关联
// @Security BearerAuth
// @Router /api/files/{id} [get]
func swaggerFilesGet() {}

// swaggerFilesDownload documents GET /api/files/{id}/download.
// @Summary 下载文件
// @Tags 文件
// @Produce application/octet-stream
// @Param id path int true "文件 ID"
// @Success 200 {file} file
// @Description 下载文件内容
// @Security BearerAuth
// @Router /api/files/{id}/download [get]
func swaggerFilesDownload() {}

// swaggerFilesCreate documents POST /api/files.
// @Summary 上传文件
// @Tags 文件
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "标题"
// @Param typeId formData int true "文件类型 ID"
// @Param date formData string false "签署日期"
// @Param file formData file true "文件"
// @Param itemLinks formData string false "关联硬件 ID，逗号分隔"
// @Param softwareLinks formData string false "关联软件 ID，逗号分隔"
// @Param invoiceLinks formData string false "关联发票 ID，逗号分隔"
// @Param contractLinks formData string false "关联合同 ID，逗号分隔"
// @Success 201 {object} swaggerIDResponse
// @Description 上传文件（multipart），支持关联到硬件、软件、单据、合同
// @Security BearerAuth
// @Router /api/files [post]
func swaggerFilesCreate() {}

// swaggerFilesUpdate documents PUT /api/files/{id}.
// @Summary 更新文件
// @Tags 文件
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "文件 ID"
// @Param title formData string true "标题"
// @Param typeId formData int true "文件类型 ID"
// @Param date formData string false "签署日期"
// @Param file formData file false "替换文件"
// @Param itemLinks formData string false "关联硬件 ID，逗号分隔"
// @Param softwareLinks formData string false "关联软件 ID，逗号分隔"
// @Param invoiceLinks formData string false "关联发票 ID，逗号分隔"
// @Param contractLinks formData string false "关联合同 ID，逗号分隔"
// @Success 200 {object} swaggerIDResponse
// @Description 更新文件信息，可选择替换文件内容；无变化时不写审计
// @Security BearerAuth
// @Router /api/files/{id} [put]
func swaggerFilesUpdate() {}

// swaggerFilesDelete documents DELETE /api/files/{id}.
// @Summary 删除文件
// @Tags 文件
// @Produce json
// @Param id path int true "文件 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除文件；仍被关联引用时返回 409
// @Security BearerAuth
// @Router /api/files/{id} [delete]
func swaggerFilesDelete() {}

// swaggerAgentsList documents GET /api/agents.
// @Summary 获取厂商/代理商列表
// @Tags 代理
// @Produce json
// @Param search query string false "搜索关键词"
// @Param limit query int false "返回数量，-1 表示全部"
// @Param offset query int false "偏移量"
// @Success 200 {array} swaggerGenericObject
// @Description 获取代理（硬件厂商、软件厂商、供应商、采购方、承包方）列表，支持关键字搜索
// @Security BearerAuth
// @Router /api/agents [get]
func swaggerAgentsList() {}

// swaggerAgentsGet documents GET /api/agents/{id}.
// @Summary 获取厂商/代理商详情
// @Tags 代理
// @Produce json
// @Param id path int true "厂商/代理商 ID"
// @Success 200 {object} swaggerGenericObject
// @Description 按编号获取代理详情
// @Security BearerAuth
// @Router /api/agents/{id} [get]
func swaggerAgentsGet() {}

// swaggerAgentsCreate documents POST /api/agents.
// @Summary 创建厂商/代理商
// @Tags 代理
// @Accept json
// @Produce json
// @Param body body agentPayload true "厂商/代理商"
// @Success 201 {object} swaggerIDResponse
// @Description 创建代理；同一名称（忽略大小写与首尾空格）不允许重复
// @Security BearerAuth
// @Router /api/agents [post]
func swaggerAgentsCreate() {}

// swaggerAgentsUpdate documents PUT /api/agents/{id}.
// @Summary 更新厂商/代理商
// @Tags 代理
// @Accept json
// @Produce json
// @Param id path int true "厂商/代理商 ID"
// @Param body body agentPayload true "厂商/代理商"
// @Success 200 {object} swaggerIDResponse
// @Description 更新代理；取消仍被引用的类型会被整体拒绝
// @Security BearerAuth
// @Router /api/agents/{id} [put]
func swaggerAgentsUpdate() {}

// swaggerAgentsDelete documents DELETE /api/agents/{id}.
// @Summary 删除厂商/代理商
// @Tags 代理
// @Produce json
// @Param id path int true "厂商/代理商 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除代理并置空硬件、软件、单据中的引用；仍被引用时返回 409
// @Security BearerAuth
// @Router /api/agents/{id} [delete]
func swaggerAgentsDelete() {}

// swaggerLocationsList documents GET /api/locations.
// @Summary 获取地点列表
// @Tags 地点
// @Produce json
// @Param search query string false "搜索关键词"
// @Success 200 {array} swaggerGenericObject
// @Description 获取地点列表，支持关键字搜索
// @Security BearerAuth
// @Router /api/locations [get]
func swaggerLocationsList() {}

// swaggerLocationsGet documents GET /api/locations/{id}.
// @Summary 获取地点详情
// @Tags 地点
// @Produce json
// @Param id path int true "位置 ID"
// @Success 200 {object} swaggerGenericObject
// @Description 按编号获取地点详情
// @Security BearerAuth
// @Router /api/locations/{id} [get]
func swaggerLocationsGet() {}

// swaggerLocationsFloorplan documents GET /api/locations/{id}/floorplan.
// @Summary 查看地点平面图
// @Tags 地点
// @Produce application/octet-stream
// @Param id path int true "位置 ID"
// @Success 200 {file} file
// @Description 查看地点平面图图片
// @Security BearerAuth
// @Router /api/locations/{id}/floorplan [get]
func swaggerLocationsFloorplan() {}

// swaggerLocationsCreate documents POST /api/locations.
// @Summary 创建地点
// @Tags 地点
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "位置名称"
// @Param floor formData string true "楼层"
// @Param file formData file false "平面图"
// @Success 201 {object} swaggerIDResponse
// @Description 创建地点；审计详情包含保存时提交的区域清单
// @Security BearerAuth
// @Router /api/locations [post]
func swaggerLocationsCreate() {}

// swaggerLocationsUpdate documents PUT /api/locations/{id}.
// @Summary 更新地点
// @Tags 地点
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "位置 ID"
// @Param name formData string true "位置名称"
// @Param floor formData string true "楼层"
// @Param file formData file false "平面图"
// @Success 200 {object} swaggerIDResponse
// @Description 更新地点；区域页签的最终清单随表单提交，无任何修改时不写审计
// @Security BearerAuth
// @Router /api/locations/{id} [put]
func swaggerLocationsUpdate() {}

// swaggerLocationsDelete documents DELETE /api/locations/{id}.
// @Summary 删除地点
// @Tags 地点
// @Produce json
// @Param id path int true "位置 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除地点；仍被机架或硬件引用时返回 409
// @Security BearerAuth
// @Router /api/locations/{id} [delete]
func swaggerLocationsDelete() {}

// swaggerLocationNextAreaID documents GET /api/locations/next-area-id.
// @Summary 获取下一个区域编号
// @Tags 地点区域
// @Produce json
// @Success 200 {object} swaggerGenericObject
// @Description 获取下一个可用的区域编号
// @Security BearerAuth
// @Router /api/locations/next-area-id [get]
func swaggerLocationNextAreaID() {}

// swaggerLocationAreasList documents GET /api/locations/{id}/areas.
// @Summary 获取地点区域
// @Tags 地点区域
// @Produce json
// @Param id path int true "位置 ID"
// @Success 200 {array} swaggerGenericObject
// @Description 获取地点下的全部区域
// @Security BearerAuth
// @Router /api/locations/{id}/areas [get]
func swaggerLocationAreasList() {}

// swaggerLocationAreasCreate documents POST /api/locations/{id}/areas.
// @Summary 创建地点区域
// @Tags 地点区域
// @Accept json
// @Produce json
// @Param id path int true "位置 ID"
// @Param body body locAreaPayload true "区域"
// @Success 201 {object} swaggerIDResponse
// @Description 为地点新增区域
// @Security BearerAuth
// @Router /api/locations/{id}/areas [post]
func swaggerLocationAreasCreate() {}

// swaggerLocationAreasUpdate documents PUT /api/locations/{id}/areas/{areaId}.
// @Summary 更新地点区域
// @Tags 地点区域
// @Accept json
// @Produce json
// @Param id path int true "位置 ID"
// @Param areaId path int true "区域 ID"
// @Param body body locAreaPayload true "区域"
// @Success 200 {object} swaggerIDResponse
// @Description 更新地点区域
// @Security BearerAuth
// @Router /api/locations/{id}/areas/{areaId} [put]
func swaggerLocationAreasUpdate() {}

// swaggerLocationAreasDelete documents DELETE /api/locations/{id}/areas/{areaId}.
// @Summary 删除地点区域
// @Tags 地点区域
// @Produce json
// @Param id path int true "位置 ID"
// @Param areaId path int true "区域 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除地点区域；仍被引用时返回 409
// @Security BearerAuth
// @Router /api/locations/{id}/areas/{areaId} [delete]
func swaggerLocationAreasDelete() {}

// swaggerRacksList documents GET /api/racks.
// @Summary 获取机架列表
// @Tags 机架
// @Produce json
// @Param search query string false "搜索关键词"
// @Success 200 {array} swaggerGenericObject
// @Description 获取机架列表，支持关键字搜索
// @Security BearerAuth
// @Router /api/racks [get]
func swaggerRacksList() {}

// swaggerRacksGet documents GET /api/racks/{id}.
// @Summary 获取机架详情
// @Tags 机架
// @Produce json
// @Param id path int true "机架 ID"
// @Success 200 {object} swaggerGenericObject
// @Description 按编号获取机架详情
// @Security BearerAuth
// @Router /api/racks/{id} [get]
func swaggerRacksGet() {}

// swaggerRacksCreate documents POST /api/racks.
// @Summary 创建机架
// @Tags 机架
// @Accept json
// @Produce json
// @Param body body rackPayload true "机架"
// @Success 201 {object} swaggerIDResponse
// @Description 创建机架
// @Security BearerAuth
// @Router /api/racks [post]
func swaggerRacksCreate() {}

// swaggerRacksUpdate documents PUT /api/racks/{id}.
// @Summary 更新机架
// @Tags 机架
// @Accept json
// @Produce json
// @Param id path int true "机架 ID"
// @Param body body rackPayload true "机架"
// @Success 200 {object} swaggerIDResponse
// @Description 更新机架；无任何修改时不写审计
// @Security BearerAuth
// @Router /api/racks/{id} [put]
func swaggerRacksUpdate() {}

// swaggerRacksDelete documents DELETE /api/racks/{id}.
// @Summary 删除机架
// @Tags 机架
// @Produce json
// @Param id path int true "机架 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除机架；仍被硬件引用时返回 409
// @Security BearerAuth
// @Router /api/racks/{id} [delete]
func swaggerRacksDelete() {}

// swaggerDictionariesList documents GET /api/dictionaries.
// @Summary 获取所有字典数据
// @Tags 字典
// @Produce json
// @Success 200 {object} swaggerGenericObject
// @Description 一次性获取硬件类型、部门、状态、标记等全部字典数据（按权限过滤可见分类）
// @Security BearerAuth
// @Router /api/dictionaries [get]
func swaggerDictionariesList() {}

// swaggerDictionariesCreate documents POST /api/dictionaries/{name}.
// @Summary 创建字典行
// @Tags 字典
// @Accept json
// @Produce json
// @Param name path string true "字典名称"
// @Param body body swaggerDictionaryPayload true "字典行"
// @Success 201 {object} swaggerIDResponse
// @Description 创建字典行；重名返回 409；携带 audit=0 时跳过逐行审计（批量导入专用）
// @Security BearerAuth
// @Router /api/dictionaries/{name} [post]
func swaggerDictionariesCreate() {}

// swaggerDictionariesUpdate documents PUT /api/dictionaries/{name}/{id}.
// @Summary 更新字典行
// @Tags 字典
// @Accept json
// @Produce json
// @Param name path string true "字典名称"
// @Param id path int true "字典行 ID"
// @Param body body swaggerDictionaryPayload true "字典行"
// @Success 200 {object} swaggerIDResponse
// @Description 更新字典行；内容无任何变化时不写审计
// @Security BearerAuth
// @Router /api/dictionaries/{name}/{id} [put]
func swaggerDictionariesUpdate() {}

// swaggerDictionariesDelete documents DELETE /api/dictionaries/{name}/{id}.
// @Summary 删除字典行
// @Tags 字典
// @Produce json
// @Param name path string true "字典名称"
// @Param id path int true "字典行 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除字典行；仍被引用时返回 409 且不写审计
// @Security BearerAuth
// @Router /api/dictionaries/{name}/{id} [delete]
func swaggerDictionariesDelete() {}

// swaggerTagsItems documents GET /api/tags/{id}/items.
// @Summary 获取标签关联硬件
// @Tags 标记
// @Produce json
// @Param id path int true "标签 ID"
// @Success 200 {array} swaggerGenericObject
// @Description 获取标记关联的硬件清单（仅现存硬件）
// @Security BearerAuth
// @Router /api/tags/{id}/items [get]
func swaggerTagsItems() {}

// swaggerTagsSoftware documents GET /api/tags/{id}/software.
// @Summary 获取标签关联软件
// @Tags 标记
// @Produce json
// @Param id path int true "标签 ID"
// @Success 200 {array} swaggerGenericObject
// @Description 获取标记关联的软件清单（仅现存软件）
// @Security BearerAuth
// @Router /api/tags/{id}/software [get]
func swaggerTagsSoftware() {}

// swaggerReportsList documents GET /api/reports.
// @Summary 获取报表列表
// @Tags 报表
// @Produce json
// @Success 200 {array} reportDefinition
// @Description 获取可用的统计报表清单
// @Security BearerAuth
// @Router /api/reports [get]
func swaggerReportsList() {}

// swaggerReportsRun documents GET /api/reports/{name}.
// @Summary 执行报表
// @Tags 报表
// @Produce json
// @Param name path string true "报表名称"
// @Param limit query int false "返回数量"
// @Success 200 {object} swaggerReportRunResponse
// @Description 执行指定报表并返回数据与图表配置
// @Security BearerAuth
// @Router /api/reports/{name} [get]
func swaggerReportsRun() {}

// swaggerImportDatabase documents POST /api/import/database.
// @Summary 导入 SQLite 数据库
// @Tags 数据库导入
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "SQLite .db 文件"
// @Success 200 {object} swaggerImportDatabaseResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 409 {object} swaggerErrorResponse
// @Description 上传 .db 或 zip 替换当前数据库，压缩包内的上传文件恢复到上传目录，导入前自动预备份（当前库引用的上传文件一并打包为 zip，未引用文件时仅备份数据库），导入成功后清理已打包的上传文件；旧版平台数据库自动转换，仅迁移资产管理与资料管理数据，用户按用户名与当前库合并（同名用户保留现状），新导入用户按旧库用户类型自动关联 admin 或 viewer 内置角色，系统配置、用户角色档案与审计历史从当前库恢复，硬件维护日志不迁移，标签预设保持当前默认；与当前项目结构一致的数据库则完整替换；当前数据库文件被外部工具占用时返回 409
// @Security BearerAuth
// @Router /api/import/database [post]
func swaggerImportDatabase() {}

// swaggerBrowseTree documents GET /api/browse/tree.
// @Summary 获取浏览树
// @Tags 浏览树
// @Produce json
// @Param id query string false "节点 ID"
// @Success 200 {array} browseNode
// @Description 按类型、部门、用户、代理等维度获取资产导航树
// @Security BearerAuth
// @Router /api/browse/tree [get]
func swaggerBrowseTree() {}

// swaggerLabelsItems documents GET /api/labels/items.
// @Summary 获取标签打印资产
// @Tags 标签打印
// @Produce json
// @Param search query string false "搜索关键词"
// @Param orderBy query string false "排序字段"
// @Param limit query int false "返回数量"
// @Param offset query int false "偏移量"
// @Success 200 {array} swaggerGenericObject
// @Description 获取可打印标签的硬件清单，支持关键字搜索与排序
// @Security BearerAuth
// @Router /api/labels/items [get]
func swaggerLabelsItems() {}

// swaggerLabelsPresets documents GET /api/labels/presets.
// @Summary 获取标签纸预设
// @Tags 标签打印
// @Produce json
// @Success 200 {array} swaggerGenericObject
// @Description 获取标签纸预设列表
// @Security BearerAuth
// @Router /api/labels/presets [get]
func swaggerLabelsPresets() {}

// swaggerLabelsPreview documents POST /api/labels/preview.
// @Summary 预览标签打印
// @Tags 标签打印
// @Accept json
// @Produce json
// @Param body body swaggerLabelPreviewRequest true "标签预览参数"
// @Success 200 {object} swaggerGenericObject
// @Description 按选中硬件生成标签预览数据，并写入打印标签审计
// @Security BearerAuth
// @Router /api/labels/preview [post]
func swaggerLabelsPreview() {}

// swaggerLabelsPresetsCreate documents POST /api/labels/presets.
// @Summary 保存标签纸预设（同名时更新）
// @Tags 标签打印
// @Accept json
// @Produce json
// @Param body body swaggerDictionaryPayload true "标签纸预设"
// @Success 200 {object} swaggerIDResponse
// @Success 201 {object} swaggerIDResponse
// @Description 保存标签预设；同名覆盖更新，无任何修改时不写审计
// @Security BearerAuth
// @Router /api/labels/presets [post]
func swaggerLabelsPresetsCreate() {}

// swaggerLabelsPresetsDelete documents DELETE /api/labels/presets/{id}.
// @Summary 删除标签纸预设
// @Tags 标签打印
// @Produce json
// @Param id path int true "标签纸预设 ID"
// @Success 200 {object} swaggerOKResponse
// @Description 删除标签纸预设
// @Security BearerAuth
// @Router /api/labels/presets/{id} [delete]
func swaggerLabelsPresetsDelete() {}

// swaggerSettingsUsersList documents GET /api/settings/users.
// @Summary 获取设置用户列表
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerSettingsUserListResponse
// @Failure 401 {object} swaggerErrorResponse
// @Description 获取用户列表及有效角色与权限
// @Security BearerAuth
// @Router /api/settings/users [get]
func swaggerSettingsUsersList() {}

// swaggerSettingsRoles documents GET /api/settings/roles.
// @Summary 获取设置角色列表
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerSettingsRoleListResponse
// @Description 获取角色列表及权限集合
// @Security BearerAuth
// @Router /api/settings/roles [get]
func swaggerSettingsRoles() {}

// swaggerSettingsGroups documents GET /api/settings/user-groups.
// @Summary 获取设置用户组列表
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerSettingsGroupListResponse
// @Description 获取用户群组列表及成员与角色
// @Security BearerAuth
// @Router /api/settings/user-groups [get]
func swaggerSettingsGroups() {}

// swaggerSettingsAuthProvider documents GET /api/settings/auth-provider.
// @Summary 获取认证配置
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerSettingsAuthProviderResponse
// @Description 获取 AD/LDAP 认证配置（敏感字段脱敏）
// @Security BearerAuth
// @Router /api/settings/auth-provider [get]
func swaggerSettingsAuthProvider() {}

// swaggerSettingsEmail documents GET /api/settings/email.
// @Summary 获取邮件配置
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerSettingsEmailResponse
// @Description 获取邮件配置（敏感字段脱敏）
// @Security BearerAuth
// @Router /api/settings/email [get]
func swaggerSettingsEmail() {}

// swaggerChangePassword documents POST /api/auth/change-password.
// @Summary 修改当前用户密码
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body changePasswordBody true "密码修改请求"
// @Success 200 {object} swaggerMessageResponse
// @Failure 400 {object} swaggerErrorResponse
// @Failure 401 {object} swaggerErrorResponse
// @Description 修改当前登录用户的密码，旧密码错误返回 400，成功后需重新登录
// @Security BearerAuth
// @Router /api/auth/change-password [post]
func swaggerChangePassword() {}

// swaggerPasswordResetCaptcha documents password reset captcha.
// @Summary 获取找回密码图形验证码
// @Tags 认证
// @Produce json
// @Success 200 {object} swaggerPasswordResetCaptchaResponse
// @Description 获取找回密码图形验证码，无需认证
// @Router /api/auth/password-reset/captcha [get]
func swaggerPasswordResetCaptcha() {}

// swaggerPasswordResetVerify documents password reset identity verification.
// @Summary 校验找回密码身份
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body swaggerPasswordResetVerifyRequest true "找回密码身份校验"
// @Success 200 {object} swaggerPasswordResetVerifyResponse
// @Description 校验用户名与图形验证码，换取找回密码流程令牌
// @Router /api/auth/password-reset/verify [post]
func swaggerPasswordResetVerify() {}

// swaggerPasswordResetSend documents password reset code delivery.
// @Summary 发送找回密码验证码
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body swaggerPasswordResetSendRequest true "找回密码验证码发送"
// @Success 200 {object} swaggerPasswordResetSendResponse
// @Description 向校验通过的邮箱发送找回密码验证码，受发送冷却与限流窗口约束
// @Router /api/auth/password-reset/send [post]
func swaggerPasswordResetSend() {}

// swaggerPasswordResetConfirm documents password reset confirmation.
// @Summary 确认找回密码
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body swaggerPasswordResetConfirmRequest true "找回密码确认"
// @Success 200 {object} swaggerMessageResponse
// @Description 凭邮箱验证码完成找回密码，重置账号密码并写入审计
// @Router /api/auth/password-reset/confirm [post]
func swaggerPasswordResetConfirm() {}

// swaggerSettingsUsersPost documents POST /api/settings/users.
// @Summary 创建用户
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 201 {object} swaggerSettingsUser
// @Failure 400 {object} swaggerErrorResponse
// @Failure 409 {object} swaggerErrorResponse
// @Security BearerAuth
// @Router /api/settings/users [post]
func swaggerSettingsUsersPost() {}

// swaggerSettingsUsersPut documents PUT /api/settings/users/{id}.
// @Summary 更新用户
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 200 {object} swaggerSettingsUser
// @Failure 400 {object} swaggerErrorResponse
// @Security BearerAuth
// @Router /api/settings/users/{id} [put]
func swaggerSettingsUsersPut() {}

// swaggerSettingsUsersDelete documents DELETE /api/settings/users/{id}.
// @Summary 删除用户
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerOKResponse
// @Failure 409 {object} swaggerErrorResponse
// @Security BearerAuth
// @Router /api/settings/users/{id} [delete]
func swaggerSettingsUsersDelete() {}

// swaggerSettingsUsersDisabled documents POST /api/settings/users/{id}/disabled.
// @Summary 启用或禁用用户
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 200 {object} swaggerSettingsUser
// @Security BearerAuth
// @Router /api/settings/users/{id}/disabled [post]
func swaggerSettingsUsersDisabled() {}

// swaggerSettingsRolesPost documents POST /api/settings/roles.
// @Summary 创建角色
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 201 {object} swaggerSettingsRole
// @Failure 409 {object} swaggerErrorResponse
// @Security BearerAuth
// @Router /api/settings/roles [post]
func swaggerSettingsRolesPost() {}

// swaggerSettingsRolesPut documents PUT /api/settings/roles/{id}.
// @Summary 更新角色
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 200 {object} swaggerSettingsRole
// @Security BearerAuth
// @Router /api/settings/roles/{id} [put]
func swaggerSettingsRolesPut() {}

// swaggerSettingsRolesDelete documents DELETE /api/settings/roles/{id}.
// @Summary 删除角色
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerOKResponse
// @Security BearerAuth
// @Router /api/settings/roles/{id} [delete]
func swaggerSettingsRolesDelete() {}

// swaggerSettingsRolesDisabled documents POST /api/settings/roles/{id}/disabled.
// @Summary 启用或禁用角色
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 200 {object} swaggerSettingsRole
// @Security BearerAuth
// @Router /api/settings/roles/{id}/disabled [post]
func swaggerSettingsRolesDisabled() {}

// swaggerSettingsGroupsPost documents POST /api/settings/user-groups.
// @Summary 创建用户群组
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 200 {object} swaggerSettingsGroup
// @Security BearerAuth
// @Router /api/settings/user-groups [post]
func swaggerSettingsGroupsPost() {}

// swaggerSettingsGroupsPut documents PUT /api/settings/user-groups/{id}.
// @Summary 更新用户群组
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 200 {object} swaggerSettingsGroup
// @Security BearerAuth
// @Router /api/settings/user-groups/{id} [put]
func swaggerSettingsGroupsPut() {}

// swaggerSettingsGroupsDelete documents DELETE /api/settings/user-groups/{id}.
// @Summary 删除用户群组
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerOKResponse
// @Security BearerAuth
// @Router /api/settings/user-groups/{id} [delete]
func swaggerSettingsGroupsDelete() {}

// swaggerSettingsAuthProviderPut documents PUT /api/settings/auth-provider.
// @Summary 保存 AD/LDAP 认证配置
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 200 {object} swaggerSettingsAuthProviderResponse
// @Security BearerAuth
// @Router /api/settings/auth-provider [put]
func swaggerSettingsAuthProviderPut() {}

// swaggerSettingsWecomProvider documents GET /api/settings/auth/wecom.
// @Summary 获取企业微信认证配置
// @Tags 系统设置
// @Produce json
// @Success 200 {object} swaggerWecomProvider
// @Description 获取企业微信认证配置（Secret 不回显，仅返回是否已配置）
// @Security BearerAuth
// @Router /api/settings/auth/wecom [get]
func swaggerSettingsWecomProvider() {}

// swaggerSettingsWecomProviderPut documents PUT /api/settings/auth/wecom.
// @Summary 保存企业微信认证配置
// @Tags 系统设置
// @Accept json
// @Produce json
// @Param body body swaggerWecomProviderRequest true "企业微信认证配置"
// @Success 200 {object} swaggerWecomProvider
// @Failure 400 {object} swaggerErrorResponse
// @Description 保存企业微信认证配置；认证方式 direct（直连）启用时要求企业 ID、AgentID 与 Secret 完整，sso（统一认证中心）启用时要求认证中心地址、应用标识与应用密钥完整，密钥加密存储不回显
// @Security BearerAuth
// @Router /api/settings/auth/wecom [put]
func swaggerSettingsWecomProviderPut() {}

// swaggerSettingsEmailPut documents PUT /api/settings/email.
// @Summary 保存邮件配置
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 200 {object} swaggerSettingsEmailResponse
// @Security BearerAuth
// @Router /api/settings/email [put]
func swaggerSettingsEmailPut() {}

// swaggerSettingsEmailTest documents POST /api/settings/email.
// @Summary 发送测试邮件
// @Tags 系统设置
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 502 {object} swaggerErrorResponse
// @Security BearerAuth
// @Router /api/settings/email [post]
func swaggerSettingsEmailTest() {}
