// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package hcl

import "github.com/coveo/gotemplate/v3/collections"

// List implementation of IGenericList for hclList
type List = hclList
type hclIList = collections.IGenericList
type hclList []interface{}

func (l hclList) AsArray() []interface{} { return []interface{}(l) }
func (l hclList) Cap() int               { return cap(l) }
func (l hclList) Capacity() int          { return cap(l) }
func (l hclList) Clone() hclIList        { return hclListHelper.Clone(l) }
func (l hclList) Contains(values ...interface{}) bool {
	return hclListHelper.Contains(l, values...)
}
func (l hclList) Count() int                  { return len(l) }
func (l hclList) Create(args ...int) hclIList { return hclListHelper.CreateList(args...) }
func (l hclList) CreateDict(args ...int) hclIDict {
	return hclListHelper.CreateDictionary(args...)
}
func (l hclList) First() interface{} { return hclListHelper.GetIndexes(l, 0) }
func (l hclList) Get(indexes ...int) interface{} {
	return hclListHelper.GetIndexes(l, indexes...)
}
func (l hclList) Has(values ...interface{}) bool   { return l.Contains(values...) }
func (l hclList) IsArrayOfSingleMap() bool         { return hclListHelper.IsArrayOfSingleMap(l) }
func (l hclList) Join(sep interface{}) str         { return l.StringArray().Join(sep) }
func (l hclList) Last() interface{}                { return hclListHelper.GetIndexes(l, len(l)-1) }
func (l hclList) Len() int                         { return len(l) }
func (l hclList) New(args ...interface{}) hclIList { return hclListHelper.NewList(args...) }
func (l hclList) Reverse() hclIList                { return hclListHelper.Reverse(l) }
func (l hclList) StringArray() strArray            { return hclListHelper.GetStringArray(l) }
func (l hclList) Strings() []string                { return hclListHelper.GetStrings(l) }
func (l hclList) TypeName() str                    { return "Hcl" }
func (l hclList) Unique() hclIList                 { return hclListHelper.Unique(l) }

func (l hclList) GetHelpers() (collections.IDictionaryHelper, collections.IListHelper) {
	return hclDictHelper, hclListHelper
}

func (l hclList) Append(values ...interface{}) hclIList {
	return hclListHelper.Add(l, false, values...)
}

func (l hclList) Intersect(values ...interface{}) hclIList {
	return hclListHelper.Intersect(l, values...)
}

func (l hclList) Pop(indexes ...int) (interface{}, hclIList) {
	if len(indexes) == 0 {
		indexes = []int{len(l) - 1}
	}
	return l.Get(indexes...), l.Remove(indexes...)
}

func (l hclList) Prepend(values ...interface{}) hclIList {
	return hclListHelper.Add(l, true, values...)
}

func (l hclList) Remove(indexes ...int) hclIList {
	return hclListHelper.Remove(l, indexes...)
}

func (l hclList) Set(i int, v interface{}) (hclIList, error) {
	return hclListHelper.SetIndex(l, i, v)
}

func (l hclList) Union(values ...interface{}) hclIList {
	return hclListHelper.Add(l, false, values...).Unique()
}

func (l hclList) Without(values ...interface{}) hclIList {
	return hclListHelper.Without(l, values...)
}

// Dictionary implementation of IDictionary for hclDict
type Dictionary = hclDict
type hclIDict = collections.IDictionary
type hclDict map[string]interface{}

func (d hclDict) Add(key, v interface{}) hclIDict     { return hclDictHelper.Add(d, key, v) }
func (d hclDict) AsMap() map[string]interface{}       { return (map[string]interface{})(d) }
func (d hclDict) Clone(keys ...interface{}) hclIDict  { return hclDictHelper.Clone(d, keys) }
func (d hclDict) Count() int                          { return len(d) }
func (d hclDict) Create(args ...int) hclIDict         { return hclListHelper.CreateDictionary(args...) }
func (d hclDict) CreateList(args ...int) hclIList     { return hclHelper.CreateList(args...) }
func (d hclDict) Flush(keys ...interface{}) hclIDict  { return hclDictHelper.Flush(d, keys) }
func (d hclDict) Get(keys ...interface{}) interface{} { return hclDictHelper.Get(d, keys) }
func (d hclDict) GetKeys() hclIList                   { return hclDictHelper.GetKeys(d, false) }
func (d hclDict) GetAllKeys() hclIList                { return hclDictHelper.GetKeys(d, true) }
func (d hclDict) GetValues() hclIList                 { return hclDictHelper.GetValues(d) }
func (d hclDict) Has(keys ...interface{}) bool        { return hclDictHelper.Has(d, keys) }
func (d hclDict) KeysAsString() strArray              { return hclDictHelper.KeysAsString(d, false) }
func (d hclDict) Len() int                            { return len(d) }
func (d hclDict) Native() interface{}                 { return collections.ToNativeRepresentation(d) }
func (d hclDict) Pop(keys ...interface{}) interface{} { return hclDictHelper.Pop(d, keys) }
func (d hclDict) Set(key, v interface{}) hclIDict     { return hclDictHelper.Set(d, key, v) }
func (d hclDict) SingleKey() string                   { return hclDictHelper.SingleKey(d) }
func (d hclDict) Transpose() hclIDict                 { return hclDictHelper.Transpose(d) }
func (d hclDict) TypeName() str                       { return "Hcl" }

func (d hclDict) GetHelpers() (collections.IDictionaryHelper, collections.IListHelper) {
	return hclDictHelper, hclListHelper
}

func (d hclDict) Default(key, defVal interface{}) interface{} {
	return hclDictHelper.Default(d, key, defVal)
}

func (d hclDict) Delete(key interface{}, otherKeys ...interface{}) (hclIDict, error) {
	return hclDictHelper.Delete(d, append([]interface{}{key}, otherKeys...))
}

func (d hclDict) Diff(dict hclIDict) hclIDict {
	return hclDictHelper.Diff(d, dict)
}

func (d hclDict) Merge(dict hclIDict, otherDicts ...hclIDict) hclIDict {
	return hclDictHelper.Merge(d, append([]hclIDict{dict}, otherDicts...))
}

func (d hclDict) Overwrite(dict hclIDict, otherDicts ...hclIDict) hclIDict {
	return hclDictHelper.MergeOverwrite(d, append([]hclIDict{dict}, otherDicts...))
}

func (d hclDict) Omit(key interface{}, otherKeys ...interface{}) hclIDict {
	return hclDictHelper.Omit(d, append([]interface{}{key}, otherKeys...))
}

// Generic helpers to simplify physical implementation
func hclListConvert(list hclIList) hclIList { return hclList(list.AsArray()) }
func hclDictConvert(dict hclIDict) hclIDict { return hclDict(dict.AsMap()) }
func needConversion(object interface{}, strict bool) bool {
	return needConversionImpl(object, strict, "Hcl")
}

var hclHelper = helperBase{ConvertList: hclListConvert, ConvertDict: hclDictConvert, NeedConversion: needConversion}
var hclListHelper = helperList{BaseHelper: hclHelper}
var hclDictHelper = helperDict{BaseHelper: hclHelper}

// DictionaryHelper gives public access to the basic dictionary functions
var DictionaryHelper collections.IDictionaryHelper = hclDictHelper

// GenericListHelper gives public access to the basic list functions
var GenericListHelper collections.IListHelper = hclListHelper

type (
	str      = collections.String
	strArray = collections.StringArray
)

var iif = collections.IIf
