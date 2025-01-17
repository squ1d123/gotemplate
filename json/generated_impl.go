// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package json

import (
	"github.com/coveooss/gotemplate/v3/collections"
	"github.com/coveooss/gotemplate/v3/errors"
)

// List implementation of IGenericList for jsonList
type List = jsonList
type jsonIList = collections.IGenericList
type jsonList []interface{}

func (l jsonList) AsArray() []interface{} { return []interface{}(l) }
func (l jsonList) Cap() int               { return cap(l) }
func (l jsonList) Capacity() int          { return cap(l) }
func (l jsonList) Clone() jsonIList       { return jsonListHelper.Clone(l) }
func (l jsonList) Contains(values ...interface{}) bool {
	return jsonListHelper.Contains(l, values...)
}
func (l jsonList) Count() int                   { return len(l) }
func (l jsonList) Create(args ...int) jsonIList { return jsonListHelper.CreateList(args...) }
func (l jsonList) CreateDict(args ...int) jsonIDict {
	return jsonListHelper.CreateDictionary(args...)
}
func (l jsonList) First() interface{} { return jsonListHelper.GetIndexes(l, 0) }
func (l jsonList) Get(indexes ...int) interface{} {
	return jsonListHelper.GetIndexes(l, indexes...)
}
func (l jsonList) Has(values ...interface{}) bool    { return l.Contains(values...) }
func (l jsonList) Join(sep interface{}) str          { return l.StringArray().Join(sep) }
func (l jsonList) Last() interface{}                 { return jsonListHelper.GetIndexes(l, len(l)-1) }
func (l jsonList) Len() int                          { return len(l) }
func (l jsonList) New(args ...interface{}) jsonIList { return jsonListHelper.NewList(args...) }
func (l jsonList) Reverse() jsonIList                { return jsonListHelper.Reverse(l) }
func (l jsonList) StringArray() strArray             { return jsonListHelper.GetStringArray(l) }
func (l jsonList) Strings() []string                 { return jsonListHelper.GetStrings(l) }
func (l jsonList) TypeName() str                     { return "Json" }
func (l jsonList) Unique() jsonIList                 { return jsonListHelper.Unique(l) }

func (l jsonList) GetHelpers() (collections.IDictionaryHelper, collections.IListHelper) {
	return jsonDictHelper, jsonListHelper
}

func (l jsonList) Append(values ...interface{}) jsonIList {
	return jsonListHelper.Add(l, false, values...)
}

func (l jsonList) Intersect(values ...interface{}) jsonIList {
	return jsonListHelper.Intersect(l, values...)
}

func (l jsonList) Pop(indexes ...int) (interface{}, jsonIList) {
	if len(indexes) == 0 {
		indexes = []int{len(l) - 1}
	}
	return l.Get(indexes...), l.Remove(indexes...)
}

func (l jsonList) Prepend(values ...interface{}) jsonIList {
	return jsonListHelper.Add(l, true, values...)
}

func (l jsonList) Remove(indexes ...int) jsonIList {
	return jsonListHelper.Remove(l, indexes...)
}

func (l jsonList) Set(i int, v interface{}) (jsonIList, error) {
	return jsonListHelper.SetIndex(l, i, v)
}

func (l jsonList) Union(values ...interface{}) jsonIList {
	return jsonListHelper.Add(l, false, values...).Unique()
}

func (l jsonList) Without(values ...interface{}) jsonIList {
	return jsonListHelper.Without(l, values...)
}

// Dictionary implementation of IDictionary for jsonDict
type Dictionary = jsonDict
type jsonIDict = collections.IDictionary
type jsonDict map[string]interface{}

func (d jsonDict) Add(key, v interface{}) jsonIDict    { return jsonDictHelper.Add(d, key, v) }
func (d jsonDict) AsMap() map[string]interface{}       { return (map[string]interface{})(d) }
func (d jsonDict) Clone(keys ...interface{}) jsonIDict { return jsonDictHelper.Clone(d, keys) }
func (d jsonDict) Count() int                          { return len(d) }
func (d jsonDict) Create(args ...int) jsonIDict        { return jsonListHelper.CreateDictionary(args...) }
func (d jsonDict) CreateList(args ...int) jsonIList    { return jsonHelper.CreateList(args...) }
func (d jsonDict) Flush(keys ...interface{}) jsonIDict { return jsonDictHelper.Flush(d, keys) }
func (d jsonDict) Get(keys ...interface{}) interface{} { return jsonDictHelper.Get(d, keys) }
func (d jsonDict) GetKeys() jsonIList                  { return jsonDictHelper.GetKeys(d) }
func (d jsonDict) GetValues() jsonIList                { return jsonDictHelper.GetValues(d) }
func (d jsonDict) Has(keys ...interface{}) bool        { return jsonDictHelper.Has(d, keys) }
func (d jsonDict) KeysAsString() strArray              { return jsonDictHelper.KeysAsString(d) }
func (d jsonDict) Len() int                            { return len(d) }
func (d jsonDict) Native() interface{}                 { return must(collections.MarshalGo(d)) }
func (d jsonDict) Pop(keys ...interface{}) interface{} { return jsonDictHelper.Pop(d, keys) }
func (d jsonDict) Set(key, v interface{}) jsonIDict    { return jsonDictHelper.Set(d, key, v) }
func (d jsonDict) Transpose() jsonIDict                { return jsonDictHelper.Transpose(d) }
func (d jsonDict) TypeName() str                       { return "Json" }

func (d jsonDict) GetHelpers() (collections.IDictionaryHelper, collections.IListHelper) {
	return jsonDictHelper, jsonListHelper
}

func (d jsonDict) Default(key, defVal interface{}) interface{} {
	return jsonDictHelper.Default(d, key, defVal)
}

func (d jsonDict) Delete(key interface{}, otherKeys ...interface{}) (jsonIDict, error) {
	return jsonDictHelper.Delete(d, append([]interface{}{key}, otherKeys...))
}

func (d jsonDict) Merge(dict jsonIDict, otherDicts ...jsonIDict) jsonIDict {
	return jsonDictHelper.Merge(d, append([]jsonIDict{dict}, otherDicts...))
}

func (d jsonDict) Omit(key interface{}, otherKeys ...interface{}) jsonIDict {
	return jsonDictHelper.Omit(d, append([]interface{}{key}, otherKeys...))
}

// Generic helpers to simplify physical implementation
func jsonListConvert(list jsonIList) jsonIList { return jsonList(list.AsArray()) }
func jsonDictConvert(dict jsonIDict) jsonIDict { return jsonDict(dict.AsMap()) }
func needConversion(object interface{}, strict bool) bool {
	return needConversionImpl(object, strict, "Json")
}

var jsonHelper = helperBase{ConvertList: jsonListConvert, ConvertDict: jsonDictConvert, NeedConversion: needConversion}
var jsonListHelper = helperList{BaseHelper: jsonHelper}
var jsonDictHelper = helperDict{BaseHelper: jsonHelper}

// DictionaryHelper gives public access to the basic dictionary functions
var DictionaryHelper collections.IDictionaryHelper = jsonDictHelper

// GenericListHelper gives public access to the basic list functions
var GenericListHelper collections.IListHelper = jsonListHelper

type (
	str      = collections.String
	strArray = collections.StringArray
)

var (
	iif  = collections.IIf
	must = errors.Must
)
