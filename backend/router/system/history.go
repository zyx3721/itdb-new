package system

import (
	"itdb-backend/internal/common/primitives"
	"itdb-backend/router/common"
	"net/http"
	"strings"

	"itdb-backend/internal/service"
)

type auditHistoryItem struct {
	ID          int64  `json:"id"`
	Timestamp   int64  `json:"timestamp"`
	Username    string `json:"username"`
	Module      string `json:"module"`
	Action      string `json:"action"`
	Target      string `json:"target"`
	TargetTitle string `json:"targetTitle,omitempty"`
	IP          string `json:"ip"`
	Result      string `json:"result"`
	Detail      string `json:"detail"`
}

// handleHistory 返回结构化审计日志；导入类事件的目标为全部名称清单，
// 响应中 target 仅保留前两个名称，完整内容随 targetTitle 返回供悬浮提示
func (a *Router) handleHistory(w http.ResponseWriter, r *http.Request) {
	limit := primitives.ListLimitParam(r.URL.Query().Get("limit"), -1)

	var total int64
	if err := a.domains.System.QueryRow(`SELECT COUNT(*) FROM history`).Scan(&total); err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	query := `SELECT id, date, authuser, ip, sql, module, action, target, detail, result FROM history ORDER BY id DESC`
	args := []interface{}{}
	if limit >= 0 {
		query += ` LIMIT ?`
		args = append(args, limit)
	}
	rows, err := a.fetchRows(query, args...)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	items := make([]auditHistoryItem, 0, len(rows))
	for _, row := range rows {
		username := strings.TrimSpace(primitives.AsString(row["authuser"]))
		sqlText := primitives.AsString(row["sql"])
		module := strings.TrimSpace(primitives.AsString(row["module"]))
		action := strings.TrimSpace(primitives.AsString(row["action"]))
		if module == "" || action == "" {
			legacyModule, legacyAction := service.ClassifyAuditLabels(username, sqlText)
			if module == "" {
				module = legacyModule
			}
			if action == "" {
				action = legacyAction
			}
		}
		detail := strings.TrimSpace(primitives.AsString(row["detail"]))
		if detail == "" {
			detail = sqlText
		}
		result := strings.TrimSpace(primitives.AsString(row["result"]))
		if result == "" {
			result = service.AuditResultSuccess
		}
		target := strings.TrimSpace(primitives.AsString(row["target"]))
		targetTitle := ""
		display, truncated := service.TruncateAuditTarget(target, 2)
		if truncated || strings.HasPrefix(action, "导入") {
			targetTitle = target
			target = display
		}
		items = append(items, auditHistoryItem{
			ID:          primitives.AsInt64(row["id"]),
			Timestamp:   primitives.AsInt64(row["date"]),
			Username:    username,
			Module:      module,
			Action:      action,
			Target:      target,
			TargetTitle: targetTitle,
			IP:          strings.TrimSpace(primitives.AsString(row["ip"])),
			Result:      result,
			Detail:      detail,
		})
	}

	common.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"items": items,
		"total": total,
		"limit": limit,
	})
}
