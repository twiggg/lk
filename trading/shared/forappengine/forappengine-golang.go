package forappengine

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	//"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	//"cloud.google.com/go/datastore"
	//ds "cloud.google.com/go/datastore"
	logger "twiggg.com/lk/logger/v1.0.0"
	"twiggg.com/lk/trading/shared/errors"
	langs "twiggg.com/lk/trading/shared/langs/v1.0.0"
)

//NewMemcacheKey returns a KeyString with enforced format
func NewMemcacheKey(namespace string, service string, id string) string {
	//return fmt.Sprintf("%s_%s_%s", namespace, service, id)
	return fmt.Sprintf("%s_%s", service, id)
}

//NewCtxWithNamespace returns a new appengine context from src context and a namespace
func NewCtxWithNamespace(ctx context.Context, r *http.Request, namespace string) (context.Context, error) {
	var actx context.Context
	/*if ctx != nil {
		actx = appengine.WithContext(ctx, r)
	} else {
		//actx = appengine.NewContext(r)
		actx = r.Context()
	}
	*/
	actx = r.Context()
	return appengine.Namespace(actx, namespace)
}

//NewClient returns a new appengine Urlfetch client
func NewClient(ctx context.Context, r *http.Request, namespace string) (*http.Client, error) {
	if ctx == nil {
		ctx = appengine.NewContext(r)
	}
	actx, err := NewCtxWithNamespace(ctx, r, namespace)
	if err != nil {
		return nil, err
	}
	return urlfetch.Client(actx), nil
}

//NewKey generates a key
func NewKey(actx context.Context, namespace string, entityType string, id int64) *datastore.Key {

	/*k := datastore.NewKey(actx, entityType, "", id, nil)
	if namespace != "" {
		k.Namespace = namespace
	}
	return k*/
	return datastore.NewKey(actx, entityType, "", id, nil)
}

//NewStringKey generates a key
func NewStringKey(actx context.Context, namespace string, entityType string, stringID string) *datastore.Key {
	/*k := datastore.NameKey(entityType, stringID, nil)
	if namespace != "" {
		k.Namespace = namespace
	}
	return k*/
	return datastore.NewKey(actx, entityType, stringID, 0, nil)
}

//NewIncompleteKey generates a new unused int key
func NewIncompleteKey(actx context.Context, namespace string, entityType string) *datastore.Key {
	/*k := datastore.IncompleteKey(entityType, nil)
	if namespace != "" {
		k.Namespace = namespace
	}
	return k*/
	return datastore.NewIncompleteKey(actx, entityType, nil)
}

//KeysFromIds returns keys from a []int64
func KeysFromIds(actx context.Context, namespace string, ids []int64, entityType string) []*datastore.Key {
	keys := make([]*datastore.Key, len(ids))
	id := int64(0)
	for i := range ids {
		id = ids[i]
		//accs[i].KeyID = fmt.Sprintf("%d", id)
		keys[i] = NewKey(actx, namespace, entityType, id) //datastore.NewKey(actx, entityType, "", id, nil)
	}
	return keys
}

//DecodeKeys returns keys from a []string
func DecodeKeys(actx context.Context, lang string, keyStrings []string) ([]*datastore.Key, error) {
	keys := make([]*datastore.Key, len(keyStrings))
	keyString := ""
	for i := range keyStrings {
		keyString = keyStrings[i]
		key, err := datastore.DecodeKey(keyString)
		if err != nil {
			return nil, errors.NewErrorInvalidKeyID(lang, keyString)
		}
		keys[i] = key
	}
	return keys, nil
}

//EncodeKeys returns keys from a []*datastore.Key
func EncodeKeys(keys []*datastore.Key) []string {
	keyStrings := make([]string, len(keys))
	var key *datastore.Key
	keyString := ""
	for i := range keys {
		key = keys[i]
		keyString = key.Encode()
		keyStrings[i] = keyString
	}
	return keyStrings
}

//Key holds a key
type Key struct {
	IntID    int64
	StringID string
}

//KeyStringFromURLIntID generate the encoded key
func KeyStringFromURLIntID(actx context.Context, lang string, entityType string, namespace string, intID string) (string, error) {

	i, err := strconv.ParseInt(intID, 10, 64)
	if err != nil {
		return "", newErrorInvalidIntID(lang, intID)
	}
	key := NewKey(actx, namespace, entityType, i)
	return key.Encode(), nil
}

func newErrorInvalidIntID(lang string, id string) error {
	switch lang {
	case langs.FR:
		return fmt.Errorf("identifiant '%s' invalide: doit être un nombre entier", id)
	case langs.ENG:
		return fmt.Errorf("invalid id '%s': must be an integer", id)
	default:
		return fmt.Errorf("invalid id '%s': must be an integer", id)
	}
}

//Delete deletes an entity by Key
func Delete(actx context.Context, projectID string /* r *http.Request,*/, lang string, namespace string, keyString string) (int, error) {
	//e := &ErrorResponse{}
	/*actx, err := NewCtxWithNamespace(actx, r, namespace)
	if err != nil {
		//msg := errors.NewErrorGcdContext(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorGcdContext(lang, err)
	}*/

	key, err := datastore.DecodeKey(keyString)
	if err != nil {
		//msg := err.Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	/*if namespace != "" {
		key.Namespace = namespace
	}
	client, err := datastore.NewClient(actx, projectID)
	if err != nil {
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	err = client.Delete(actx, key)*/
	err = datastore.Delete(actx, key)
	if err != nil {
		if err.Error() == datastore.ErrNoSuchEntity.Error() {
			//msg := errors.NewErrorNoSuchEntity(lang, keyString).Error() // errors.NewErrorNoSuchEntity(lang, keyString).Error()
			//e.Key = &msg
			return http.StatusNotFound, errors.NewErrorNoSuchEntity(lang, keyString)
		}
		//msg := errors.NewErrorInternal(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	return http.StatusOK, nil
}

//DeleteMulti deletes multiples entities by Keys
func DeleteMulti(actx context.Context, projectID string /*r *http.Request,*/, lang string, namespace string, keyStrings []string) (int, error) {
	//e := &ErrorResponse{}
	/*actx, err := NewCtxWithNamespace(ctx, r, namespace)
	if err != nil {
		//msg := errors.NewErrorGcdContext(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorGcdContext(lang, err)
	}*/
	keys, err := DecodeKeys(actx, lang, keyStrings)
	if err != nil {
		//msg := err.Error()
		//e.Key = &msg
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	/*if namespace != "" {
		for i := range keys {
			keys[i].Namespace = namespace
		}
	}
	//err = datastore.DeleteMulti(actx, keys)
	client, err := datastore.NewClient(actx, projectID)
	if err != nil {
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	err = client.DeleteMulti(actx, keys)*/
	err = datastore.DeleteMulti(actx, keys)
	if err != nil {
		//msg := errors.NewErrorInternal(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	return http.StatusOK, nil
}

//GetIntID extract the IntID from a keystring after decoding it
func GetIntID(lang string, keystring string) (int64, error) {
	key, err := datastore.DecodeKey(keystring)
	if err != nil {
		return 0, errors.NewErrorInvalidKeyID(lang, keystring)
	}
	//if key.StringID() != "" {
	/*
		if key.Name != "" {
			return 0, errors.NewErrorNotIntID(lang)
		}
	*/
	//return key.ID, nil
	return key.IntID(), nil
}

//SetLogger sets the logger
func SetLogger(l logger.Logger) {
	if l != nil {
		log = l
	}
}

//GetByKey fetches an entity by Key
func GetByKey(actx context.Context, projectID string /* r *http.Request,*/, lang string, namespace string, keyString string, dstPointer interface{}) (int, error) {
	//e := &ErrorResponse{}
	/*actx, err := NewCtxWithNamespace(ctx, r, namespace)
	if err != nil {
		//msg := errors.NewErrorGcdContext(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorGcdContext(lang, err)
	}*/
	//res := &Customer{}
	//key := newDatastoreKey(actx, keyString)
	key, err := datastore.DecodeKey(keyString)
	if err != nil {
		log.Infof(actx, "DecodeKey: err: '%s'\n", err.Error())
		//msg := errors.NewErrorInvalidKeyID(lang, keyString).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorInvalidKeyID(lang, keyString)
	}
	/*if namespace != "" {
		key.Namespace = namespace
	}
	//err = datastore.Get(actx, key, dstPointer)
	client, err := datastore.NewClient(actx, projectID)
	if err != nil {
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}

	err = client.Get(actx, key, dstPointer)*/
	err = datastore.Get(actx, key, dstPointer)
	if err != nil {
		if err.Error() == datastore.ErrNoSuchEntity.Error() {
			//msg := errors.NewErrorNoSuchEntity(lang, keyString).Error() // errors.NewErrorNoSuchEntity(lang, keyString).Error()
			//e.Key = &msg
			id := key.StringID()
			//id := key.Name
			if id == "" {
				id = fmt.Sprintf("%d", key.IntID())
				//id = fmt.Sprintf("%d", key.ID)
			}
			return http.StatusNotFound, errors.NewErrorNoSuchEntity(lang, id)
		}
		//msg := errors.NewErrorInternal(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	//res.KeyID = key.Encode()
	//res.ID = key.IntID()
	return http.StatusOK, nil
}

//GetMulti fetches entities by Keys
func GetMulti(actx context.Context, projectID string /*r *http.Request, */, lang string, namespace string, keyStrings []string, slice interface{}) (int, error) {
	//e := &ErrorResponse{}
	/*actx, err := NewCtxWithNamespace(ctx, r, namespace)
	if err != nil {
		//msg := errors.NewErrorGcdContext(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorGcdContext(lang, err)
	}*/
	//res := &Customer{}
	//key := newDatastoreKey(actx, keyString)
	keys, err := DecodeKeys(actx, lang, keyStrings)
	//key, err := datastore.DecodeKey(keyString)
	if err != nil {
		//logger.Infof(r, true, "DecodeKey: err: '%s'\n", err.Error())
		//msg := errors.NewErrorInvalidKeyID(lang, keyString).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, err
	}
	/*if namespace != "" {
		for i := range keys {
			keys[i].Namespace = namespace
		}
	}*/
	log.Infof(actx, "len(keys): %d", len(keys))
	//logger.Infof(r, true, "len(dst): %d", len(slicePointer))
	//err = datastore.GetMulti(actx, keys, slice)
	/*client, err := datastore.NewClient(actx, projectID)
	if err != nil {
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}*/
	//err = client.GetMulti(actx, keys, slice)
	err = datastore.GetMulti(actx, keys, slice)
	if err != nil {
		if err.Error() == datastore.ErrNoSuchEntity.Error() {
			//msg := errors.NewErrorNoSuchEntity(lang, keyString).Error() // errors.NewErrorNoSuchEntity(lang, keyString).Error()
			//e.Key = &msg
			return http.StatusNotFound, errors.NewErrorInternal(lang, err)
		}
		//msg := errors.NewErrorInternal(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	//res.KeyID = key.Encode()
	//res.ID = key.IntID()
	return http.StatusOK, nil
}

//PutNew generates a newKey and puts and stores a new entity
func PutNew(actx context.Context, projectID string /*r *http.Request,*/, lang string, namespace string, entityType string, payloadPointer interface{}) (string, int, error) {
	//e := &ErrorResponse{}
	/*actx, err := NewCtxWithNamespace(actx, r, namespace)
	if err != nil {
		//msg := errors.NewErrorGcdContext(lang, err).Error()
		//e.Internal = &msg
		return "", http.StatusInternalServerError, errors.NewErrorGcdContext(lang, err)
	}*/
	//key := newDatastoreKey(actx, payload.Name)
	//key:=datastore.NewIncompleteKey(actx, entityType, nil)
	key := NewIncompleteKey(actx, namespace, entityType)
	/*if namespace != "" {
		key.Namespace = namespace
	}
	//key, err := datastore.Put(actx, key, payloadPointer)
	client, err := datastore.NewClient(actx, projectID)
	if err != nil {
		return "", http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}

	key, err = client.Put(actx, key, payloadPointer)
	*/
	key, err := datastore.Put(actx, key, payloadPointer)
	if err != nil {
		//msg := errors.NewErrorInternal(lang, err).Error()
		//e.Internal = &msg
		return "", http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	keyString := key.Encode()
	return keyString, http.StatusOK, nil
}

//Put store an entity using keyString (after decoding it)
func Put(actx context.Context, projectID string /*r *http.Request, */, lang string, namespace string, keyString string, payloadPointer interface{}) (int, error) {
	//e := &ErrorResponse{}
	/*actx, err := NewCtxWithNamespace(actx, r, namespace)
	if err != nil {
		//msg := errors.NewErrorGcdContext(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorGcdContext(lang, err)
	}*/
	key, err := datastore.DecodeKey(keyString)
	if err != nil {
		log.Infof(actx, "DecodeKey: err: '%s'\n", err.Error())
		//msg := errors.NewErrorInvalidKeyID(lang, keyString).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorInvalidKeyID(lang, keyString)
	}
	/*if namespace != "" {
		key.Namespace = namespace
	}
	//key, err = datastore.Put(actx, key, payloadPointer)
	client, err := datastore.NewClient(actx, projectID)
	if err != nil {
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	key, err = client.Put(actx, key, payloadPointer)*/
	key, err = datastore.Put(actx, key, payloadPointer)
	if err != nil {
		//msg := errors.NewErrorInternal(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	//keyString := key.Encode()
	return http.StatusOK, nil
}

//PutMulti stores multiple entities with specified keys (after decoding them)
func PutMulti(actx context.Context, projectID string /*r *http.Request, */, lang string, namespace string, keyStrings []string, slice interface{}) (int, error) {
	//e := &ErrorResponse{}
	/*actx, err := NewCtxWithNamespace(ctx, r, namespace)
	if err != nil {
		//msg := errors.NewErrorGcdContext(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorGcdContext(lang, err)
	}*/
	keys, err := DecodeKeys(actx, lang, keyStrings)
	if err != nil {
		//msg := err.Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, err
	}
	//keys, err = datastore.PutMulti(actx, keys, slice)
	/*client, err := datastore.NewClient(actx, projectID)
	if err != nil {
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	_, err = client.PutMulti(actx, keys, slice)*/
	_, err = datastore.PutMulti(actx, keys, slice)
	if err != nil {
		//msg := errors.NewErrorInternal(lang, err).Error()
		//e.Internal = &msg
		return http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	return http.StatusOK, nil
}

//LookUp searches for entities matching the query's criteria
func LookUp(actx context.Context, projectID string /*r *http.Request,*/, lang string, namespace string, q Query, dstSlicePointer interface{}) ([]string, int, error) {
	//ctx0 := appengine.NewContext(r)
	//log.Infof(ctx0, "userid: '%s'\n", userid)
	//e := &ErrorResponse{}
	/*actx, err := NewCtxWithNamespace(ctx, r, namespace)
	if err != nil {
		//msg := errors.NewErrorGcdContext(lang, err).Error()
		//e.Internal = &msg
		return nil, http.StatusInternalServerError, errors.NewErrorGcdContext(lang, err)
	}*/
	dq := q.DatastoreQuery()
	/*client, err := datastore.NewClient(actx, projectID)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	keys, err := client.GetAll(actx, dq, dstSlicePointer)
	*/
	keys, err := dq.GetAll(actx, dstSlicePointer)
	if err != nil {
		if err.Error() == datastore.ErrNoSuchEntity.Error() {
			//msg := newErrorNotFoundByUsername(lang, userid).Error()
			//e.Key = &msg
			return nil, http.StatusNotFound, errors.NewErrorInternal(lang, err)
		}
		//msg := errors.NewErrorInternal(lang, err).Error()
		//e.Internal = &msg
		return nil, http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	/*
		if len(keys) == 0 {
			msg := newErrorNotFoundByUsername(lang, userid).Error()
			e.Key = &msg
			return nil, nil, http.StatusNotFound, e
		}
	*/

	keyStrings := EncodeKeys(keys)
	return keyStrings, http.StatusOK, nil
}

//Count searches for entities matching the query's criteria
func Count(actx context.Context, projectID string /*r *http.Request,*/, lang string, namespace string, q Query) (int, int, error) {
	//ctx0 := appengine.NewContext(r)
	//log.Infof(ctx0, "userid: '%s'\n", userid)
	//e := &ErrorResponse{}
	/*actx, err := NewCtxWithNamespace(ctx, r, namespace)
	if err != nil {
		//msg := errors.NewErrorGcdContext(lang, err).Error()
		//e.Internal = &msg
		return 0, http.StatusInternalServerError, errors.NewErrorGcdContext(lang, err)
	}*/
	dq := q.DatastoreQuery()
	//log.Infof(actx, "query: %+v", *dq)
	/*client, err := datastore.NewClient(actx, projectID)
	if err != nil {
		return 0, http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}


	n, err := client.Count(actx, dq)*/
	n, err := dq.Count(actx)
	if err != nil {
		if err.Error() == datastore.ErrNoSuchEntity.Error() {
			//msg := newErrorNotFoundByUsername(lang, userid).Error()
			//e.Key = &msg
			return 0, http.StatusNotFound, errors.NewErrorInternal(lang, err)
		}
		//msg := errors.NewErrorInternal(lang, err).Error()
		//e.Internal = &msg
		return 0, http.StatusInternalServerError, errors.NewErrorInternal(lang, err)
	}
	return n, http.StatusOK, nil
}

//NewQuery returns a new standard Query object able to generate a datastore.Query from its inner data
func NewQuery(entityType string, filters map[string]interface{}, limit int, offset int, orderBy string, keysOnly bool) *Query {
	if filters == nil {
		filters = map[string]interface{}{}
	}
	return &Query{
		entityType: entityType,
		filters:    filters,
		limit:      limit,
		offset:     offset,
		order:      orderBy,
		keysOnly:   keysOnly,
	}
}

//NewComplexQuery returns a new standard Query object able to generate a datastore.Query from its inner data
func NewComplexQuery(entityType string, filters map[string]interface{}, limit int, offset int, orderBy string, keysOnly bool, distinct bool) *Query {
	if filters == nil {
		filters = map[string]interface{}{}
	}
	return &Query{
		entityType: entityType,
		filters:    filters,
		limit:      limit,
		offset:     offset,
		order:      orderBy,
		keysOnly:   keysOnly,
		distinct:   distinct,
	}

}

//Query object
type Query struct {
	entityType string
	filters    map[string]interface{}
	limit      int
	offset     int
	order      string
	keysOnly   bool
	distinct   bool
}

//DatastoreQuery builds an appengine query
func (q *Query) DatastoreQuery() *datastore.Query {

	dq := datastore.NewQuery(q.entityType)
	for key, val := range q.filters {
		//if val != nil {
		dq = dq.Filter(key, val)
		//}
	}
	if q.limit > 0 {
		dq = dq.Limit(q.limit)
	}
	if q.offset > 0 {
		dq = dq.Offset(q.offset)
	}
	if q.order != "" {
		dq = dq.Order(q.order)
	}
	if q.distinct {
		dq = dq.Distinct()
	}
	return dq
}
