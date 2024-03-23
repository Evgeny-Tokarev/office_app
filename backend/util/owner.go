package util

import (
	"net/http"
)

func HasAccessRights(w http.ResponseWriter, r *http.Request, requiredOwner string) bool {
	ownerValue := r.Context().Value("owner")
	owner, ok := ownerValue.(string)
	if !ok {
		SendTranscribedError(w, "Forbidden: Request owner not specified", http.StatusBadRequest)
		return false
	}

	requiredLevel, ok := ReqOwners[requiredOwner]
	if !ok {
		SendTranscribedError(w, "Invalid required owner", http.StatusBadRequest)
		return false
	}

	ownerLevel, exists := ReqOwners[owner]
	if !exists {
		SendTranscribedError(w, "Forbidden: Insufficient privileges", http.StatusForbidden)
		return false
	}

	if ownerLevel < requiredLevel {
		SendTranscribedError(w, "Forbidden: Insufficient privileges", http.StatusForbidden)
		return false
	}

	return true
}
