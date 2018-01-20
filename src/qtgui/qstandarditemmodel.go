//  header block begin
// /usr/include/qt/QtGui/qstandarditemmodel.h
// #include <qstandarditemmodel.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 84
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStandardItemModel struct {
	*qtcore.QAbstractItemModel
}

func (this *QStandardItemModel) GetCthis() unsafe.Pointer {
	return this.QAbstractItemModel.GetCthis()
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:326
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStandardItemModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:330
// index:0
// void QStandardItemModel(class QObject *)
func NewQStandardItemModel(parent unsafe.Pointer) *QStandardItemModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(cthis)
	return gothis
}
func NewQStandardItemModelFromPointer(cthis unsafe.Pointer) *QStandardItemModel {
	bcthis0 := qtcore.NewQAbstractItemModelFromPointer(cthis)
	return &QStandardItemModel{bcthis0}
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:331
// index:1
// void QStandardItemModel(int, int, class QObject *)
func NewQStandardItemModel_1(rows int, columns int, parent unsafe.Pointer) *QStandardItemModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModelC2EiiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &rows, &columns, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:425
// index:2
// void QStandardItemModel(class QStandardItemModelPrivate &, class QObject *)
func NewQStandardItemModel_2(dd unsafe.Pointer, parent unsafe.Pointer) *QStandardItemModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModelC2ER25QStandardItemModelPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:332
// index:0
// virtual
// void ~QStandardItemModel()
func DeleteQStandardItemModel(*QStandardItemModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:336
// index:0
// virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QStandardItemModel) Index(row int, column int, parent unsafe.Pointer) {
	// 0: (, row int, column int, parent const QModelIndex &), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:337
// index:0
// virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QStandardItemModel) Parent(child unsafe.Pointer) {
	// 0: (, child const QModelIndex &), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:339
// index:0
// virtual
// int rowCount(const class QModelIndex &)
func (this *QStandardItemModel) RowCount(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:340
// index:0
// virtual
// int columnCount(const class QModelIndex &)
func (this *QStandardItemModel) ColumnCount(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:341
// index:0
// virtual
// bool hasChildren(const class QModelIndex &)
func (this *QStandardItemModel) HasChildren(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:343
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QStandardItemModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, row int, column int, idx const QModelIndex &), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:345
// index:0
// virtual
// QVariant data(const class QModelIndex &, int)
func (this *QStandardItemModel) Data(index unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, role int), (index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:346
// index:0
// virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QStandardItemModel) SetData(index unsafe.Pointer, value unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, value const QVariant &, role int), (index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:348
// index:0
// virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QStandardItemModel) HeaderData(section int, orientation int, role int) {
	// 0: (, section int, orientation Qt::Orientation, role int), (&section, &orientation, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:350
// index:0
// virtual
// bool setHeaderData(int, Qt::Orientation, const class QVariant &, int)
func (this *QStandardItemModel) SetHeaderData(section int, orientation int, value unsafe.Pointer, role int) {
	// 0: (, section int, orientation Qt::Orientation, value const QVariant &, role int), (&section, &orientation, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section, &orientation, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:353
// index:0
// virtual
// bool insertRows(int, int, const class QModelIndex &)
func (this *QStandardItemModel) InsertRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, row int, count int, parent const QModelIndex &), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:354
// index:0
// virtual
// bool insertColumns(int, int, const class QModelIndex &)
func (this *QStandardItemModel) InsertColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, column int, count int, parent const QModelIndex &), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:355
// index:0
// virtual
// bool removeRows(int, int, const class QModelIndex &)
func (this *QStandardItemModel) RemoveRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, row int, count int, parent const QModelIndex &), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:356
// index:0
// virtual
// bool removeColumns(int, int, const class QModelIndex &)
func (this *QStandardItemModel) RemoveColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, column int, count int, parent const QModelIndex &), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:358
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QStandardItemModel) Flags(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:359
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QStandardItemModel) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:361
// index:0
// virtual
// QMap<int, QVariant> itemData(const class QModelIndex &)
func (this *QStandardItemModel) ItemData(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel8itemDataERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:364
// index:0
// void clear()
func (this *QStandardItemModel) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:368
// index:0
// virtual
// void sort(int, Qt::SortOrder)
func (this *QStandardItemModel) Sort(column int, order int) {
	// 0: (, column int, order Qt::SortOrder), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:370
// index:0
// QStandardItem * itemFromIndex(const class QModelIndex &)
func (this *QStandardItemModel) ItemFromIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:371
// index:0
// QModelIndex indexFromItem(const class QStandardItem *)
func (this *QStandardItemModel) IndexFromItem(item unsafe.Pointer) {
	// 0: (, item const QStandardItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel13indexFromItemEPK13QStandardItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:373
// index:0
// QStandardItem * item(int, int)
func (this *QStandardItemModel) Item(row int, column int) {
	// 0: (, row int, column int), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel4itemEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:374
// index:0
// void setItem(int, int, class QStandardItem *)
func (this *QStandardItemModel) SetItem(row int, column int, item unsafe.Pointer) {
	// 0: (, row int, column int, item QStandardItem *), (&row, &column, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel7setItemEiiP13QStandardItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:375
// index:1
// inline
// void setItem(int, class QStandardItem *)
func (this *QStandardItemModel) SetItem_1(row int, item unsafe.Pointer) {
	// 1: (, row int, item QStandardItem *), (&row, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel7setItemEiP13QStandardItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:376
// index:0
// QStandardItem * invisibleRootItem()
func (this *QStandardItemModel) InvisibleRootItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel17invisibleRootItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:378
// index:0
// QStandardItem * horizontalHeaderItem(int)
func (this *QStandardItemModel) HorizontalHeaderItem(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel20horizontalHeaderItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:379
// index:0
// void setHorizontalHeaderItem(int, class QStandardItem *)
func (this *QStandardItemModel) SetHorizontalHeaderItem(column int, item unsafe.Pointer) {
	// 0: (, column int, item QStandardItem *), (&column, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel23setHorizontalHeaderItemEiP13QStandardItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:380
// index:0
// QStandardItem * verticalHeaderItem(int)
func (this *QStandardItemModel) VerticalHeaderItem(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel18verticalHeaderItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:381
// index:0
// void setVerticalHeaderItem(int, class QStandardItem *)
func (this *QStandardItemModel) SetVerticalHeaderItem(row int, item unsafe.Pointer) {
	// 0: (, row int, item QStandardItem *), (&row, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel21setVerticalHeaderItemEiP13QStandardItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:383
// index:0
// void setHorizontalHeaderLabels(const class QStringList &)
func (this *QStandardItemModel) SetHorizontalHeaderLabels(labels unsafe.Pointer) {
	// 0: (, labels const QStringList &), (labels)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), labels)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:384
// index:0
// void setVerticalHeaderLabels(const class QStringList &)
func (this *QStandardItemModel) SetVerticalHeaderLabels(labels unsafe.Pointer) {
	// 0: (, labels const QStringList &), (labels)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), labels)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:386
// index:0
// void setRowCount(int)
func (this *QStandardItemModel) SetRowCount(rows int) {
	// 0: (, rows int), (&rows)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel11setRowCountEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:387
// index:0
// void setColumnCount(int)
func (this *QStandardItemModel) SetColumnCount(columns int) {
	// 0: (, columns int), (&columns)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel14setColumnCountEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:391
// index:0
// inline
// void appendRow(class QStandardItem *)
func (this *QStandardItemModel) AppendRow(item unsafe.Pointer) {
	// 0: (, item QStandardItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel9appendRowEP13QStandardItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:395
// index:0
// inline
// void insertRow(int, class QStandardItem *)
func (this *QStandardItemModel) InsertRow(row int, item unsafe.Pointer) {
	// 0: (, row int, item QStandardItem *), (&row, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel9insertRowEiP13QStandardItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:397
// index:1
// inline
// bool insertRow(int, const class QModelIndex &)
func (this *QStandardItemModel) InsertRow_1(row int, parent unsafe.Pointer) {
	// 1: (, row int, parent const QModelIndex &), (&row, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel9insertRowEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:398
// index:0
// inline
// bool insertColumn(int, const class QModelIndex &)
func (this *QStandardItemModel) InsertColumn(column int, parent unsafe.Pointer) {
	// 0: (, column int, parent const QModelIndex &), (&column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel12insertColumnEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:400
// index:0
// QStandardItem * takeItem(int, int)
func (this *QStandardItemModel) TakeItem(row int, column int) {
	// 0: (, row int, column int), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel8takeItemEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:401
// index:0
// QList<QStandardItem *> takeRow(int)
func (this *QStandardItemModel) TakeRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel7takeRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:402
// index:0
// QList<QStandardItem *> takeColumn(int)
func (this *QStandardItemModel) TakeColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel10takeColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:404
// index:0
// QStandardItem * takeHorizontalHeaderItem(int)
func (this *QStandardItemModel) TakeHorizontalHeaderItem(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel24takeHorizontalHeaderItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:405
// index:0
// QStandardItem * takeVerticalHeaderItem(int)
func (this *QStandardItemModel) TakeVerticalHeaderItem(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel22takeVerticalHeaderItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:407
// index:0
// const QStandardItem * itemPrototype()
func (this *QStandardItemModel) ItemPrototype() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel13itemPrototypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:408
// index:0
// void setItemPrototype(const class QStandardItem *)
func (this *QStandardItemModel) SetItemPrototype(item unsafe.Pointer) {
	// 0: (, item const QStandardItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel16setItemPrototypeEPK13QStandardItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:410
// index:0
// QList<QStandardItem *> findItems(const class QString &, Qt::MatchFlags, int)
func (this *QStandardItemModel) FindItems(text unsafe.Pointer, flags int, column int) {
	// 0: (, text const QString &, QFlags<Qt::MatchFlag> flags, column int), (text, &flags, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel9findItemsERK7QString6QFlagsIN2Qt9MatchFlagEEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text, &flags, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:414
// index:0
// int sortRole()
func (this *QStandardItemModel) SortRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel8sortRoleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:415
// index:0
// void setSortRole(int)
func (this *QStandardItemModel) SetSortRole(role int) {
	// 0: (, role int), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel11setSortRoleEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:417
// index:0
// virtual
// QStringList mimeTypes()
func (this *QStandardItemModel) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:419
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QStandardItemModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, data const QMimeData *, action Qt::DropAction, row int, column int, parent const QModelIndex &), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:422
// index:0
// void itemChanged(class QStandardItem *)
func (this *QStandardItemModel) ItemChanged(item unsafe.Pointer) {
	// 0: (, item QStandardItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel11itemChangedEP13QStandardItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

//  body block end