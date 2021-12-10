package accesscontroller // import "github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-ipfs-log/accesscontroller"

import (
	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-ipfs-log/identityprovider"
)

type LogEntry interface {
	GetPayload() []byte
	GetIdentity() *identityprovider.Identity
}

type CanAppendAdditionalContext interface {
	GetLogEntries() []LogEntry
}

type Interface interface {
	CanAppend(LogEntry, identityprovider.Interface, CanAppendAdditionalContext) error
}
