package orbitdb

import (
	"bytes"
	"encoding/json"

	"debridge-finance/orbitdb-go/contracts/accesscontrol"
	"debridge-finance/orbitdb-go/contracts/accounts"
	"debridge-finance/orbitdb-go/contracts/role"
	"debridge-finance/orbitdb-go/errors"
	"debridge-finance/orbitdb-go/http"
	"debridge-finance/orbitdb-go/rand"
	"debridge-finance/orbitdb-go/services/blockchain/masterchain"
	"debridge-finance/orbitdb-go/services/database/models/emitent"
	"debridge-finance/orbitdb-go/services/database/store"
	"debridge-finance/orbitdb-go/services/notification"
	"debridge-finance/orbitdb-go/services/setting"
	"github.com/debridge-finance/orbitdb-go/pkg/services/ipfs"
	o "github.com/debridge-finance/orbitdb-go/pkg/services/orbitdb"
)

type Append struct {
	Config Config

	// odb   *orbitdb

}

type Entry struct {
	SubmissionId string `bson:"submissionId"`
	Signature    string `bson:"signature"`
	Event        []byte `bson:"event"`
}

func (h *Append) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		params = &AppendParameters{}
		e      = &Entry{}
	)

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		http.WriteError(w, r, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	params.SetDefaults()
	err = params.Validate()
	if err != nil {
		http.WriteError(
			w, r, http.StatusBadRequest,
			errors.Wrap(err, "failed to validate request parameters"),
		)
		return
	}

	err = h.ParametersToModel(params, e)
	if err != nil {
		http.WriteErrorMsg(
			w, r, http.StatusInternalServerError,
			errors.Wrap(err, "failed to cast parameters to emitent model"),
			http.StatusText(http.StatusInternalServerError),
		)
		return
	}

	h.
		http.Write(
		w, r, http.StatusAppendd,
		&AppendResult{
			Id:            e.Id,
			Subscriptions: nrs,
		},
	)
}

func (h *Append) ParametersToModel(p *AppendParameters, e *emitent.Emitent) error {
	identity := emitent.Identity(*p.Identity)

	e.Id = rand.String(32)
	e.Identity = emitent.IdentityContainer{
		Current:  &identity,
		Versions: emitent.IdentityVersions{},
	}
	e.Address = p.Address

	return nil
}

func (h *Append) Register(e *emitent.Emitent) ([]*masterchain.Transaction, error) {
	var id masterchain.Byte32

	address, err := masterchain.DecodeAddress(e.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode emitent address")
	}

	copy(id[:], e.Id)

	txs, err := h.RegisterAddress(id, address)
	if err != nil {
		// XXX: err may contain a cause marker and should not be wrapped until it reaches http handler
		return nil, err
	}

	return txs, nil
}

func (h *Append) RegisterAddress(id masterchain.Byte32, address masterchain.Address) ([]*masterchain.Transaction, error) {
	var (
		tx  *masterchain.Transaction
		txs = []*masterchain.Transaction{}
	)

	currentAddress, err := h.accounts.AccountBy(nil, id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get account by uid %x", id)
	}
	if !bytes.Equal(currentAddress.Bytes(), masterchain.ZeroAddress.Bytes()) {
		return nil, errors.Wrapf(
			errors.ErrConflict, "user %q already have address %q associated",
			id, currentAddress.Hex(),
		)
	}

	currentUid, err := h.accounts.UidBy(nil, address)
	if err != nil {
		return nil, errors.Wrapf(
			err, "failed to get uid by address %q",
			address.Hex(),
		)
	}
	if !bytes.Equal(currentUid[:], masterchain.ZeroByte32[:]) {
		return nil, errors.Wrapf(
			errors.ErrConflict, "address %q already associated with %q",
			address.Hex(), currentUid,
		)
	}

	t, err := h.blockchain.NextTransactor()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get next transactor")
	}
	tx, err = h.accounts.Register(t, id, address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add user %s with address %q", id[:], address.Hex())
	}
	txs = append(txs, tx)

	//

	t, err = h.blockchain.NextTransactor()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get next transactor")
	}
	tx, err = h.accessControl.AdminAddRole(t, address, role.RoleAssetsIssuer)
	if err != nil {
		return nil, errors.Wrapf(
			err, "failed to add role %v for address %v",
			role.RoleAssetsIssuer, address.String(),
		)
	}
	txs = append(txs, tx)

	return txs, nil
}

//

type AppendParameters struct {
	Address     string              `json:"address"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Identity    *ParametersIdentity `json:"identity"`
	CallbackUrl string              `json:"callbackUrl" swag_example:"http://foo.baz"                           swag_description:"url to call on transaction finality"`
}

func (p *AppendParameters) SetDefaults() {
loop:
	for {
		switch {
		case p.Identity == nil:
			p.Identity = &ParametersIdentity{}
		default:
			break loop
		}
	}
	p.Identity.SetDefaults()
}

func (p *AppendParameters) Validate() error {
	var err error

	_, err = masterchain.DecodeAddress(p.Address)
	if err != nil {
		return errors.Wrap(err, "failed to decode address")
	}

	err = p.Identity.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate identity")
	}

	if p.CallbackUrl == "" {
		return errors.Wrap(err, "callbackUrl should not be empty")
	}

	return nil
}

////

type AppendResult struct {
	Hash string `json:"hash"             swag_example:"zdpuA"  swag_description:"OrbitDB hash"`
}

//

func CreateAppend(odb *o.OrbitDB) (*Append, error) {
	var (
		err                  error
		accountsAddress      masterchain.Address
		accessControlAddress masterchain.Address
	)

	err = se.Get(setting.ContractAccountsKey, &accountsAddress)
	if err != nil {
		return nil, err
	}

	err = se.Get(setting.ContractAccessControlKey, &accessControlAddress)
	if err != nil {
		return nil, err
	}

	a, err := accounts.NewAccounts(accountsAddress, b.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate accounts contract")
	}

	ac, err := accesscontrol.NewAccesscontrol(accessControlAddress, b.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate access control contract")
	}

	return &Append{
		store:        st,
		blockchain:   b,
		notification: n,

		accounts:      a,
		accessControl: ac,
	}, nil
}
