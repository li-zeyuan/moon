package meta

import (
	"fmt"
	"github.com/fatih/structtag"
	"log"
	"reflect"
	"runtime/debug"
	"strings"
)

type Meta struct {
	info  map[uintptr]Key
	sf    map[uintptr]Key // 只装exported struct fields
	name  Key
	alias Key
}

func (m *Meta) recursive(i interface{}) {
	rv := reflect.ValueOf(i).Elem()
	rt := reflect.TypeOf(i).Elem()

	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		field := rt.Field(i)

		// exported field
		if field.PkgPath == "" {
			_, ok := m.sf[fv.UnsafeAddr()]
			// 这里必须这么做, 不然会被顶掉
			if !ok {
				m.sf[fv.UnsafeAddr()] = Key(field.Name)
			}
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			if fv.Addr().CanInterface() {
				m.recursive(fv.Addr().Interface())
			}
		default:
			break
		}

		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			continue
		}

		tag, err := tags.Get("gorm")
		if err != nil {
			continue
		}

		m.info[rv.Field(i).UnsafeAddr()] = Key(tag.Name)
	}
}

func (m *Meta) Init(i interface{}) {
	m.info = make(map[uintptr]Key)
	m.sf = make(map[uintptr]Key)

	if reflect.TypeOf(i).Elem().Kind() != reflect.Struct &&
		reflect.TypeOf(i).Kind() != reflect.Ptr {
		log.Fatalln("NEED A PTR TO STRUCT")
	}

	name, ok := reflect.TypeOf(i).Elem().FieldByName("tableName")
	if !ok {
		log.Fatalln("NO TABLENAME FIELD")
	}

	tags, err := structtag.Parse(string(name.Tag))
	if err != nil {
		log.Fatalln("parse tag err: ", err)
	}

	tag, err := tags.Get("gorm")
	if err != nil {
		log.Fatalln("get tag err: ", err)
	}

	m.name = Key(tag.Name)

	for _, option := range tag.Options {
		if strings.Contains(option, "alias") {
			strs := strings.Split(option, ":")
			m.alias = Key(strs[1])
			break
		}
	}

	if len(m.alias) == 0 {
		log.Fatalln("NO ALIAS PROVIDED")
	}

	m.recursive(i)
}

func (m *Meta) TableName() string {
	return string(m.name)
}

func (m *Meta) AliasPk() string {
	return m.alias.V() + "." + "id"
}

func (m *Meta) AliasIsDelete() Key {
	return Key(m.alias.V() + ".is_delete")
}

func (m *Meta) Alias() string {
	return string(m.alias)
}

func (m *Meta) AliasCus(cus string) string {
	return fmt.Sprintf("%s_%s", string(m.alias), cus)
}

func (m *Meta) AliasAny() string {
	return fmt.Sprintf("%s.*", string(m.alias))
}

func (m *Meta) AliasTag(i interface{}) Key {
	k := m.Tag(i)
	return Key(fmt.Sprintf("%s.%s", m.alias, k))
}

func (m *Meta) AliasCusTag(cus string, i interface{}) Key {
	k := m.Tag(i)
	return Key(fmt.Sprintf("%s.%s", m.AliasCus(cus), k))
}

func (m *Meta) AliasTagEscape(i interface{}) Key {
	k := m.Tag(i)
	return Key(fmt.Sprintf(`%s."%s"`, m.alias, k))
}

func (m *Meta) Tag(i interface{}) Key {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		debug.PrintStack()
		log.Fatalln("NEED PTR")
	}
	addr := reflect.ValueOf(i).Elem().UnsafeAddr()
	if m.info[addr] == "" {
		debug.PrintStack()
		log.Fatal("NO FIELD")
	}
	return Key(m.info[addr])
}

func (m *Meta) Field(i interface{}) Key {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		debug.PrintStack()
		log.Fatalln("NEED PTR")
	}
	addr := reflect.ValueOf(i).Elem().UnsafeAddr()
	if m.sf[addr] == "" {
		debug.PrintStack()
		log.Fatalln("NO FIELD")
	}
	return Key(m.sf[addr])
}
