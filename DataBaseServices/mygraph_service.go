// file: services/movie_service.go
/*eslint-disable */
package DataBaseServices

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"

	grpc "google.golang.org/grpc"
)

type MygraphService interface {
	MutateRDF([]byte, string) interface{}
	Mutate([]byte) (interface{}, map[string]string, error)
	Query(string) ([]byte, error)
	VQuery(string, map[string]string) ([]byte, error)
}

func NewDgraphTrasn() *mygraphService {
	return &mygraphService{
		key: "dgraph",
		//dg:  newClient(),
		txn: newClient().NewTxn(),
	}
}

//newClient :To create a client, dial a connection to Dgraph’s external gRPC port (typically 9080). The following code snippet shows just one connection.
func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

type mygraphService struct {
	key string
	//dg  *dgo.Dgraph
	txn *dgo.Txn
}

/*
all mutation api type
	SetJson             []byte   `protobuf:"bytes,1,opt,name=set_json,json=setJson,proto3" json:"set_json,omitempty"`
	DeleteJson          []byte   `protobuf:"bytes,2,opt,name=delete_json,json=deleteJson,proto3" json:"delete_json,omitempty"`
	SetNquads           []byte   `protobuf:"bytes,3,opt,name=set_nquads,json=setNquads,proto3" json:"set_nquads,omitempty"`
	DelNquads           []byte   `protobuf:"bytes,4,opt,name=del_nquads,json=delNquads,proto3" json:"del_nquads,omitempty"`
*/
//MutateRDF : mutate database with  JSON and RDF N-Quad.
//allowed type here is
// Ttype : transaction type
//	* "set" : insert or update data
//	* "delete" : delete data
func (s *mygraphService) MutateRDF(qry []byte, Ttype string) interface{} {
	defer s.txn.Discard(context.Background())
	switch Ttype {
	case "set":
		res, err := s.txn.Mutate(context.Background(), &api.Mutation{
			CommitNow: true,
			SetNquads: qry,
		})
		if err != nil {
			log.Fatal(err)
		}
		return res
	case "delete":
		res, err := s.txn.Mutate(context.Background(), &api.Mutation{
			CommitNow: true,
			DelNquads: qry,
		})
		if err != nil {
			log.Fatal(err)
		}
		return res
	}
	return "null"
}

//Mutate : mutate the data base with json qry
// UIDs are map of string in the patern of blank-0 , blank-1 , ...
func (s *mygraphService) Mutate(qry []byte) (FULL interface{}, UIDs map[string]string, err error) {
	defer s.txn.Discard(context.Background())
	res, err := s.txn.Mutate(context.Background(), &api.Mutation{
		CommitNow: true,
		SetJson:   qry,
	})
	return res, res.Uids, err
}

//Query : query , format like
/*
	q=`{
		data(func:uid(0xXXX)){
			uid
			expand(_all_)
			...
		}
	}
		`

return json formated stirng and err
*/
func (s *mygraphService) Query(qry string) (jsonStr []byte, err error) {
	res, err := s.txn.Query(context.Background(), qry)

	return res.Json, err
}

//VQuery : query with variables , format like
/*
	q=`{
		data(func:uid($uid)){
			uid
			expand(_all_)
			...
		}
	}
		`
	data := map[string]string{"$id": "0xXXX"}

return json formated stirng and err
*/
func (s *mygraphService) VQuery(qry string, data map[string]string) (jsonStr []byte, err error) {
	res, err := s.txn.QueryWithVars(context.Background(), qry, data)
	return res.Json, err
}
