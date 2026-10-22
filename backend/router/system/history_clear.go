package system

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"itdb-backend/internal/service"
	"itdb-backend/router/common"
)

type historyClearRequest struct {
	IDs []int64 `json:"ids"`
}

// handleHistoryClear 按前端筛选结果批量清空审计日志，删除成功后补记一条清空操作审计
func (a *Router) handleHistoryClear(w http.ResponseWriter, r *http.Request) {
	var req historyClearRequest
	if !common.DecodeJSON(w, r, &req) {
		return
	}
	ids := dedupeHistoryIDs(req.IDs)
	if len(ids) == 0 {
		common.WriteError(w, http.StatusBadRequest, "no audit log ids provided")
		return
	}
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}
	result, err := a.db.Exec(
		fmt.Sprintf(`DELETE FROM history WHERE id IN (%s)`, strings.Join(placeholders, ",")),
		args...,
	)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	deleted, err := result.RowsAffected()
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.recordHistoryClearAudit(r, deleted)
	common.WriteJSON(w, http.StatusOK, map[string]interface{}{"deleted": deleted})
}

// dedupeHistoryIDs 过滤非法编号并去重
func dedupeHistoryIDs(raw []int64) []int64 {
	ids := make([]int64, 0, len(raw))
	seen := make(map[int64]bool, len(raw))
	for _, id := range raw {
		if id <= 0 || seen[id] {
			continue
		}
		seen[id] = true
		ids = append(ids, id)
	}
	return ids
}

// recordHistoryClearAudit 在删除完成后写入清空操作审计，保证该记录自身不被删除
func (a *Router) recordHistoryClearAudit(r *http.Request, deleted int64) {
	if a.audit == nil {
		return
	}
	operator, _ := common.CurrentUser(r.Context())
	event := service.AuditEvent{
		Module: service.AuditModuleAudit,
		Action: "清空审计日志",
		Target: fmt.Sprintf("%d 条", deleted),
		Detail: fmt.Sprintf("清空 %d 条审计日志记录", deleted),
		Result: service.AuditResultSuccess,
	}
	if err := a.audit.RecordEvent(r.Context(), operator.Username, common.ClientIP(r), event); err != nil {
		log.Printf("Record audit clear event failed: %s", err)
	}
}
