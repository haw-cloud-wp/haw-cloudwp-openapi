package commons

import (
	"errors"
	"github.com/scrapes/haw-cloudwp-openapi/src/middleware"
	"io"
	"os"
	"time"
)

type IPermission interface {
	Init(claims middleware.CustomClaims) IPermission
	HasPermission(permission string) bool
}

type AllowAllPermission struct {
}

func (a *AllowAllPermission) Init(claims middleware.CustomClaims) IPermission {
	return a
}

func (a *AllowAllPermission) HasPermission(permission string) bool {
	return true
}

type IObjectInfo interface {
	Init(name string, size int64, lastMod time.Time) IObjectInfo
	GetName() string
	GetSize() int64
	GetLastMod() time.Time
}

type IObject interface {
	Init(bucket IBucket, name string) IObject
	GetName() string
	Delete() error

	Set(file *os.File) error
	Get() (error, *os.File)
	Stream() (error, *io.Reader)

	GetSize() (error, int64)
	GetLastMod() (error, time.Time)
}

type IBucket interface {
	Init(store IStorage, name string) IBucket
	GetName() string
	Delete() error
	GetStorage() IStorage

	GetObjects() (error, []IObjectInfo)

	DeleteObject(object IObject) error
	GetObjectStream(object IObject) (error, *io.Reader)
	GetObject(object IObject) (error, *os.File)
	SetObject(object IObject, data *os.File) error

	GetObjectSize(object IObject) (error, int64)
	GetObjectLastMod(object IObject) (error, time.Time)
}

type IStorage interface {
	io.Closer
	Init(permissions IPermission) IStorage
	DeleteBucket(bucket IBucket) error
	CreateBucket(name string) (error, IBucket)

	GetObjects(bucket IBucket) (error, []IObjectInfo)
	DeleteObject(bucket IBucket, object IObject) error
	GetObjectStream(bucket IBucket, object IObject) (error, *io.Reader)
	GetObject(bucket IBucket, object IObject) (error, *os.File)
	SetObject(bucket IBucket, object IObject, data *os.File) error
	GetPermission(permission string) bool

	GetObjectSize(bucket IBucket, object IObject) (error, int64)
	GetObjectLastMod(bucket IBucket, object IObject) (error, time.Time)
}

type Bucket struct {
	storage IStorage
	name    string
}

func (b *Bucket) GetObjectSize(object IObject) (error, int64) {
	return b.storage.GetObjectSize(b, object)
}

func (b *Bucket) GetObjectLastMod(object IObject) (error, time.Time) {
	return b.storage.GetObjectLastMod(b, object)
}

func (b *Bucket) GetStorage() IStorage {
	return b.storage
}

func (b *Bucket) Init(store IStorage, name string) IBucket {
	b.storage = store
	b.name = name
	return b
}

var errorPermissionDenied = errors.New("permission denied")

func (b *Bucket) GetName() string {
	return b.name
}

func (b *Bucket) Delete() error {
	if b.storage.GetPermission("own:bucket_" + b.GetName()) {
		return b.storage.DeleteBucket(b)
	} else {
		return errorPermissionDenied
	}
}

func (b *Bucket) GetObjects() (error, []IObjectInfo) {
	if b.storage.GetPermission("read:bucket_" + b.GetName()) {
		return b.storage.GetObjects(b)
	} else {
		return errorPermissionDenied, nil
	}
}

func (b *Bucket) DeleteObject(object IObject) error {
	if b.storage.GetPermission("delete:bucket_" + b.GetName()) {
		return b.storage.DeleteObject(b, object)
	}
	return errorPermissionDenied
}

func (b *Bucket) GetObjectStream(object IObject) (error, *io.Reader) {
	if b.storage.GetPermission("get:bucket_" + b.GetName()) {
		return b.storage.GetObjectStream(b, object)
	}
	return errorPermissionDenied, nil
}

func (b *Bucket) GetObject(object IObject) (error, *os.File) {
	if b.storage.GetPermission("get:bucket_" + b.GetName()) {
		return b.storage.GetObject(b, object)
	}
	return errorPermissionDenied, nil
}

func (b *Bucket) SetObject(object IObject, data *os.File) error {
	if b.storage.GetPermission("set:bucket_" + b.GetName()) {
		return b.storage.SetObject(b, object, data)
	}
	return errorPermissionDenied
}

type Object struct {
	bucket IBucket
	name   string
}

func (o *Object) GetSize() (error, int64) {
	return o.bucket.GetObjectSize(o)
}

func (o *Object) GetLastMod() (error, time.Time) {
	return o.bucket.GetObjectLastMod(o)
}

func (o *Object) Init(bucket IBucket, name string) IObject {
	o.bucket = bucket
	o.name = name
	return o
}

func (o *Object) GetName() string {
	return o.name
}

func (o *Object) Delete() error {
	return o.bucket.DeleteObject(o)
}

func (o *Object) Set(file *os.File) error {
	return o.bucket.SetObject(o, file)
}

func (o *Object) Get() (error, *os.File) {
	return o.bucket.GetObject(o)
}

func (o *Object) Stream() (error, *io.Reader) {
	return o.bucket.GetObjectStream(o)
}

type ObjectInfo struct {
	name    string
	size    int64
	lastMod time.Time
}

func (o *ObjectInfo) Init(name string, size int64, lastMod time.Time) IObjectInfo {
	o.name = name
	o.size = size
	o.lastMod = lastMod
	return o
}

func (o *ObjectInfo) GetName() string {
	return o.name
}

func (o *ObjectInfo) GetSize() int64 {
	return o.size
}

func (o *ObjectInfo) GetLastMod() time.Time {
	return o.lastMod
}
