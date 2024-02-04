// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"sms/app/gateway/api/dao/model"
)

func newMsgProvider(db *gorm.DB, opts ...gen.DOOption) msgProvider {
	_msgProvider := msgProvider{}

	_msgProvider.msgProviderDo.UseDB(db, opts...)
	_msgProvider.msgProviderDo.UseModel(&model.MsgProvider{})

	tableName := _msgProvider.msgProviderDo.TableName()
	_msgProvider.ALL = field.NewAsterisk(tableName)
	_msgProvider.ID = field.NewInt64(tableName, "id")
	_msgProvider.ProviderCode = field.NewString(tableName, "provider_code")
	_msgProvider.ProviderName = field.NewString(tableName, "provider_name")
	_msgProvider.Description = field.NewString(tableName, "description")
	_msgProvider.CreateTime = field.NewTime(tableName, "create_time")
	_msgProvider.CreateBy = field.NewString(tableName, "create_by")
	_msgProvider.ModifyTime = field.NewTime(tableName, "modify_time")
	_msgProvider.ModifyBy = field.NewString(tableName, "modify_by")
	_msgProvider.Deleted = field.NewField(tableName, "deleted")

	_msgProvider.fillFieldMap()

	return _msgProvider
}

type msgProvider struct {
	msgProviderDo msgProviderDo

	ALL          field.Asterisk
	ID           field.Int64
	ProviderCode field.String
	ProviderName field.String
	Description  field.String
	CreateTime   field.Time
	CreateBy     field.String
	ModifyTime   field.Time
	ModifyBy     field.String
	Deleted      field.Field

	fieldMap map[string]field.Expr
}

func (m msgProvider) Table(newTableName string) *msgProvider {
	m.msgProviderDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m msgProvider) As(alias string) *msgProvider {
	m.msgProviderDo.DO = *(m.msgProviderDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *msgProvider) updateTableName(table string) *msgProvider {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.ProviderCode = field.NewString(table, "provider_code")
	m.ProviderName = field.NewString(table, "provider_name")
	m.Description = field.NewString(table, "description")
	m.CreateTime = field.NewTime(table, "create_time")
	m.CreateBy = field.NewString(table, "create_by")
	m.ModifyTime = field.NewTime(table, "modify_time")
	m.ModifyBy = field.NewString(table, "modify_by")
	m.Deleted = field.NewField(table, "deleted")

	m.fillFieldMap()

	return m
}

func (m *msgProvider) WithContext(ctx context.Context) *msgProviderDo {
	return m.msgProviderDo.WithContext(ctx)
}

func (m msgProvider) TableName() string { return m.msgProviderDo.TableName() }

func (m msgProvider) Alias() string { return m.msgProviderDo.Alias() }

func (m msgProvider) Columns(cols ...field.Expr) gen.Columns { return m.msgProviderDo.Columns(cols...) }

func (m *msgProvider) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *msgProvider) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 9)
	m.fieldMap["id"] = m.ID
	m.fieldMap["provider_code"] = m.ProviderCode
	m.fieldMap["provider_name"] = m.ProviderName
	m.fieldMap["description"] = m.Description
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["create_by"] = m.CreateBy
	m.fieldMap["modify_time"] = m.ModifyTime
	m.fieldMap["modify_by"] = m.ModifyBy
	m.fieldMap["deleted"] = m.Deleted
}

func (m msgProvider) clone(db *gorm.DB) msgProvider {
	m.msgProviderDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m msgProvider) replaceDB(db *gorm.DB) msgProvider {
	m.msgProviderDo.ReplaceDB(db)
	return m
}

type msgProviderDo struct{ gen.DO }

func (m msgProviderDo) Debug() *msgProviderDo {
	return m.withDO(m.DO.Debug())
}

func (m msgProviderDo) WithContext(ctx context.Context) *msgProviderDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m msgProviderDo) ReadDB() *msgProviderDo {
	return m.Clauses(dbresolver.Read)
}

func (m msgProviderDo) WriteDB() *msgProviderDo {
	return m.Clauses(dbresolver.Write)
}

func (m msgProviderDo) Session(config *gorm.Session) *msgProviderDo {
	return m.withDO(m.DO.Session(config))
}

func (m msgProviderDo) Clauses(conds ...clause.Expression) *msgProviderDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m msgProviderDo) Returning(value interface{}, columns ...string) *msgProviderDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m msgProviderDo) Not(conds ...gen.Condition) *msgProviderDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m msgProviderDo) Or(conds ...gen.Condition) *msgProviderDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m msgProviderDo) Select(conds ...field.Expr) *msgProviderDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m msgProviderDo) Where(conds ...gen.Condition) *msgProviderDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m msgProviderDo) Order(conds ...field.Expr) *msgProviderDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m msgProviderDo) Distinct(cols ...field.Expr) *msgProviderDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m msgProviderDo) Omit(cols ...field.Expr) *msgProviderDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m msgProviderDo) Join(table schema.Tabler, on ...field.Expr) *msgProviderDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m msgProviderDo) LeftJoin(table schema.Tabler, on ...field.Expr) *msgProviderDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m msgProviderDo) RightJoin(table schema.Tabler, on ...field.Expr) *msgProviderDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m msgProviderDo) Group(cols ...field.Expr) *msgProviderDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m msgProviderDo) Having(conds ...gen.Condition) *msgProviderDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m msgProviderDo) Limit(limit int) *msgProviderDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m msgProviderDo) Offset(offset int) *msgProviderDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m msgProviderDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *msgProviderDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m msgProviderDo) Unscoped() *msgProviderDo {
	return m.withDO(m.DO.Unscoped())
}

func (m msgProviderDo) Create(values ...*model.MsgProvider) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m msgProviderDo) CreateInBatches(values []*model.MsgProvider, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m msgProviderDo) Save(values ...*model.MsgProvider) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m msgProviderDo) First() (*model.MsgProvider, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgProvider), nil
	}
}

func (m msgProviderDo) Take() (*model.MsgProvider, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgProvider), nil
	}
}

func (m msgProviderDo) Last() (*model.MsgProvider, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgProvider), nil
	}
}

func (m msgProviderDo) Find() ([]*model.MsgProvider, error) {
	result, err := m.DO.Find()
	return result.([]*model.MsgProvider), err
}

func (m msgProviderDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MsgProvider, err error) {
	buf := make([]*model.MsgProvider, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m msgProviderDo) FindInBatches(result *[]*model.MsgProvider, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m msgProviderDo) Attrs(attrs ...field.AssignExpr) *msgProviderDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m msgProviderDo) Assign(attrs ...field.AssignExpr) *msgProviderDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m msgProviderDo) Joins(fields ...field.RelationField) *msgProviderDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m msgProviderDo) Preload(fields ...field.RelationField) *msgProviderDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m msgProviderDo) FirstOrInit() (*model.MsgProvider, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgProvider), nil
	}
}

func (m msgProviderDo) FirstOrCreate() (*model.MsgProvider, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MsgProvider), nil
	}
}

func (m msgProviderDo) FindByPage(offset int, limit int) (result []*model.MsgProvider, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m msgProviderDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m msgProviderDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m msgProviderDo) Delete(models ...*model.MsgProvider) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *msgProviderDo) withDO(do gen.Dao) *msgProviderDo {
	m.DO = *do.(*gen.DO)
	return m
}
