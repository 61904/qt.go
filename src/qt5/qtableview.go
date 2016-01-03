package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qtableview.h
// dst-file: /src/widgets/qtableview.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTableView::resizeRowsToContents();
extern void _ZN10QTableView20resizeRowsToContentsEv(void* qthis);
  // proto:  void QTableView::setRowHeight(int row, int height);
extern void _ZN10QTableView12setRowHeightEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QHeaderView * QTableView::verticalHeader();
extern void _ZNK10QTableView14verticalHeaderEv(void* qthis);
  // proto:  void QTableView::setSpan(int row, int column, int rowSpan, int columnSpan);
extern void _ZN10QTableView7setSpanEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QTableView::setSortingEnabled(bool enable);
extern void _ZN10QTableView17setSortingEnabledEb(void* qthis, bool arg0);
  // proto:  void QTableView::setColumnWidth(int column, int width);
extern void _ZN10QTableView14setColumnWidthEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QTableView::setWordWrap(bool on);
extern void _ZN10QTableView11setWordWrapEb(void* qthis, bool arg0);
  // proto:  void QTableView::doItemsLayout();
extern void _ZN10QTableView13doItemsLayoutEv(void* qthis);
  // proto:  void QTableView::setSelectionModel(QItemSelectionModel * selectionModel);
extern void _ZN10QTableView17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0);
  // proto:  void QTableView::setHorizontalHeader(QHeaderView * header);
extern void _ZN10QTableView19setHorizontalHeaderEP11QHeaderView(void* qthis, void* arg0);
  // proto:  void QTableView::setRowHidden(int row, bool hide);
extern void _ZN10QTableView12setRowHiddenEib(void* qthis, int32_t arg0, bool arg1);
  // proto:  int QTableView::rowViewportPosition(int row);
extern void _ZNK10QTableView19rowViewportPositionEi(void* qthis, int32_t arg0);
  // proto:  int QTableView::columnAt(int x);
extern void _ZNK10QTableView8columnAtEi(void* qthis, int32_t arg0);
  // proto:  bool QTableView::isRowHidden(int row);
extern void _ZNK10QTableView11isRowHiddenEi(void* qthis, int32_t arg0);
  // proto:  void QTableView::showColumn(int column);
extern void _ZN10QTableView10showColumnEi(void* qthis, int32_t arg0);
  // proto:  void QTableView::resizeRowToContents(int row);
extern void _ZN10QTableView19resizeRowToContentsEi(void* qthis, int32_t arg0);
  // proto:  void QTableView::setRootIndex(const QModelIndex & index);
extern void _ZN10QTableView12setRootIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QTableView::setColumnHidden(int column, bool hide);
extern void _ZN10QTableView15setColumnHiddenEib(void* qthis, int32_t arg0, bool arg1);
  // proto:  void QTableView::hideRow(int row);
extern void _ZN10QTableView7hideRowEi(void* qthis, int32_t arg0);
  // proto:  void QTableView::resizeColumnsToContents();
extern void _ZN10QTableView23resizeColumnsToContentsEv(void* qthis);
  // proto:  bool QTableView::wordWrap();
extern void _ZNK10QTableView8wordWrapEv(void* qthis);
  // proto:  void QTableView::setShowGrid(bool show);
extern void _ZN10QTableView11setShowGridEb(void* qthis, bool arg0);
  // proto:  bool QTableView::isColumnHidden(int column);
extern void _ZNK10QTableView14isColumnHiddenEi(void* qthis, int32_t arg0);
  // proto:  void QTableView::selectRow(int row);
extern void _ZN10QTableView9selectRowEi(void* qthis, int32_t arg0);
  // proto:  const QMetaObject * QTableView::metaObject();
extern void _ZNK10QTableView10metaObjectEv(void* qthis);
  // proto:  bool QTableView::isCornerButtonEnabled();
extern void _ZNK10QTableView21isCornerButtonEnabledEv(void* qthis);
  // proto:  void QTableView::selectColumn(int column);
extern void _ZN10QTableView12selectColumnEi(void* qthis, int32_t arg0);
  // proto:  void QTableView::~QTableView();
extern void _ZN10QTableViewD0Ev(void* qthis);
  // proto:  void QTableView::resizeColumnToContents(int column);
extern void _ZN10QTableView22resizeColumnToContentsEi(void* qthis, int32_t arg0);
  // proto:  void QTableView::QTableView(QWidget * parent);
extern void* dector_ZN10QTableViewC1EP7QWidget(void* arg0);
extern void _ZN10QTableViewC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QTableView::sortByColumn(int column);
extern void _ZN10QTableView12sortByColumnEi(void* qthis, int32_t arg0);
  // proto:  int QTableView::columnSpan(int row, int column);
extern void _ZNK10QTableView10columnSpanEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  int QTableView::columnWidth(int column);
extern void _ZNK10QTableView11columnWidthEi(void* qthis, int32_t arg0);
  // proto:  int QTableView::columnViewportPosition(int column);
extern void _ZNK10QTableView22columnViewportPositionEi(void* qthis, int32_t arg0);
  // proto:  int QTableView::rowHeight(int row);
extern void _ZNK10QTableView9rowHeightEi(void* qthis, int32_t arg0);
  // proto:  void QTableView::QTableView(const QTableView & );
extern void* dector_ZN10QTableViewC1ERKS_(void* arg0);
extern void _ZN10QTableViewC1ERKS_(void* qthis, void* arg0);
  // proto:  int QTableView::rowAt(int y);
extern void _ZNK10QTableView5rowAtEi(void* qthis, int32_t arg0);
  // proto:  int QTableView::rowSpan(int row, int column);
extern void _ZNK10QTableView7rowSpanEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QTableView::setCornerButtonEnabled(bool enable);
extern void _ZN10QTableView22setCornerButtonEnabledEb(void* qthis, bool arg0);
  // proto:  QRect QTableView::visualRect(const QModelIndex & index);
extern void _ZNK10QTableView10visualRectERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QTableView::indexAt(const QPoint & p);
extern void _ZNK10QTableView7indexAtERK6QPoint(void* qthis, void* arg0);
  // proto:  QHeaderView * QTableView::horizontalHeader();
extern void _ZNK10QTableView16horizontalHeaderEv(void* qthis);
  // proto:  void QTableView::setModel(QAbstractItemModel * model);
extern void _ZN10QTableView8setModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  bool QTableView::isSortingEnabled();
extern void _ZNK10QTableView16isSortingEnabledEv(void* qthis);
  // proto:  void QTableView::clearSpans();
extern void _ZN10QTableView10clearSpansEv(void* qthis);
  // proto:  void QTableView::setVerticalHeader(QHeaderView * header);
extern void _ZN10QTableView17setVerticalHeaderEP11QHeaderView(void* qthis, void* arg0);
  // proto:  bool QTableView::showGrid();
extern void _ZNK10QTableView8showGridEv(void* qthis);
  // proto:  void QTableView::showRow(int row);
extern void _ZN10QTableView7showRowEi(void* qthis, int32_t arg0);
  // proto:  void QTableView::hideColumn(int column);
extern void _ZN10QTableView10hideColumnEi(void* qthis, int32_t arg0);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTableView)=1
type QTableView struct {
  /*qbase*/ QAbstractItemView;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QTableView::resizeRowsToContents();
func (this *QTableView) resizeRowsToContents(args ...interface{}) () {
  // resizeRowsToContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView20resizeRowsToContentsEv
    // invoke: void resizeRowsToContents()
    C._ZN10QTableView20resizeRowsToContentsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "resizeRowsToContents", args)
  }

}

  // proto:  void QTableView::setRowHeight(int row, int height);
func (this *QTableView) setRowHeight(args ...interface{}) () {
  // setRowHeight(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12setRowHeightEii
    // invoke: void setRowHeight(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QTableView12setRowHeightEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "setRowHeight", args)
  }

}

  // proto:  QHeaderView * QTableView::verticalHeader();
func (this *QTableView) verticalHeader(args ...interface{}) () {
  // verticalHeader()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView14verticalHeaderEv
    // invoke: QHeaderView * verticalHeader()
    C._ZNK10QTableView14verticalHeaderEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "verticalHeader", args)
  }

}

  // proto:  void QTableView::setSpan(int row, int column, int rowSpan, int columnSpan);
func (this *QTableView) setSpan(args ...interface{}) () {
  // setSpan(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView7setSpanEiiii
    // invoke: void setSpan(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN10QTableView7setSpanEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTableView", "setSpan", args)
  }

}

  // proto:  void QTableView::setSortingEnabled(bool enable);
func (this *QTableView) setSortingEnabled(args ...interface{}) () {
  // setSortingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView17setSortingEnabledEb
    // invoke: void setSortingEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView17setSortingEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setSortingEnabled", args)
  }

}

  // proto:  void QTableView::setColumnWidth(int column, int width);
func (this *QTableView) setColumnWidth(args ...interface{}) () {
  // setColumnWidth(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView14setColumnWidthEii
    // invoke: void setColumnWidth(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QTableView14setColumnWidthEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "setColumnWidth", args)
  }

}

  // proto:  void QTableView::setWordWrap(bool on);
func (this *QTableView) setWordWrap(args ...interface{}) () {
  // setWordWrap(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView11setWordWrapEb
    // invoke: void setWordWrap(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView11setWordWrapEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setWordWrap", args)
  }

}

  // proto:  void QTableView::doItemsLayout();
func (this *QTableView) doItemsLayout(args ...interface{}) () {
  // doItemsLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView13doItemsLayoutEv
    // invoke: void doItemsLayout()
    C._ZN10QTableView13doItemsLayoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "doItemsLayout", args)
  }

}

  // proto:  void QTableView::setSelectionModel(QItemSelectionModel * selectionModel);
func (this *QTableView) setSelectionModel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView17setSelectionModelEP19QItemSelectionModel
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(QItemSelectionModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTableView17setSelectionModelEP19QItemSelectionModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setSelectionModel", args)
  }

}

  // proto:  void QTableView::setHorizontalHeader(QHeaderView * header);
func (this *QTableView) setHorizontalHeader(args ...interface{}) () {
  // setHorizontalHeader(class QHeaderView *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QHeaderView{}) // "QHeaderView *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView19setHorizontalHeaderEP11QHeaderView
    // invoke: void setHorizontalHeader(class QHeaderView *)
    var arg0 = args[0].(QHeaderView).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTableView19setHorizontalHeaderEP11QHeaderView(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setHorizontalHeader", args)
  }

}

  // proto:  void QTableView::setRowHidden(int row, bool hide);
func (this *QTableView) setRowHidden(args ...interface{}) () {
  // setRowHidden(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12setRowHiddenEib
    // invoke: void setRowHidden(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN10QTableView12setRowHiddenEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "setRowHidden", args)
  }

}

  // proto:  int QTableView::rowViewportPosition(int row);
func (this *QTableView) rowViewportPosition(args ...interface{}) () {
  // rowViewportPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView19rowViewportPositionEi
    // invoke: int rowViewportPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView19rowViewportPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "rowViewportPosition", args)
  }

}

  // proto:  int QTableView::columnAt(int x);
func (this *QTableView) columnAt(args ...interface{}) () {
  // columnAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView8columnAtEi
    // invoke: int columnAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView8columnAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "columnAt", args)
  }

}

  // proto:  bool QTableView::isRowHidden(int row);
func (this *QTableView) isRowHidden(args ...interface{}) () {
  // isRowHidden(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView11isRowHiddenEi
    // invoke: bool isRowHidden(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView11isRowHiddenEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "isRowHidden", args)
  }

}

  // proto:  void QTableView::showColumn(int column);
func (this *QTableView) showColumn(args ...interface{}) () {
  // showColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView10showColumnEi
    // invoke: void showColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView10showColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "showColumn", args)
  }

}

  // proto:  void QTableView::resizeRowToContents(int row);
func (this *QTableView) resizeRowToContents(args ...interface{}) () {
  // resizeRowToContents(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView19resizeRowToContentsEi
    // invoke: void resizeRowToContents(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView19resizeRowToContentsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "resizeRowToContents", args)
  }

}

  // proto:  void QTableView::setRootIndex(const QModelIndex & index);
func (this *QTableView) setRootIndex(args ...interface{}) () {
  // setRootIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12setRootIndexERK11QModelIndex
    // invoke: void setRootIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTableView12setRootIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setRootIndex", args)
  }

}

  // proto:  void QTableView::setColumnHidden(int column, bool hide);
func (this *QTableView) setColumnHidden(args ...interface{}) () {
  // setColumnHidden(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView15setColumnHiddenEib
    // invoke: void setColumnHidden(int, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN10QTableView15setColumnHiddenEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "setColumnHidden", args)
  }

}

  // proto:  void QTableView::hideRow(int row);
func (this *QTableView) hideRow(args ...interface{}) () {
  // hideRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView7hideRowEi
    // invoke: void hideRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView7hideRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "hideRow", args)
  }

}

  // proto:  void QTableView::resizeColumnsToContents();
func (this *QTableView) resizeColumnsToContents(args ...interface{}) () {
  // resizeColumnsToContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView23resizeColumnsToContentsEv
    // invoke: void resizeColumnsToContents()
    C._ZN10QTableView23resizeColumnsToContentsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "resizeColumnsToContents", args)
  }

}

  // proto:  bool QTableView::wordWrap();
func (this *QTableView) wordWrap(args ...interface{}) () {
  // wordWrap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView8wordWrapEv
    // invoke: bool wordWrap()
    C._ZNK10QTableView8wordWrapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "wordWrap", args)
  }

}

  // proto:  void QTableView::setShowGrid(bool show);
func (this *QTableView) setShowGrid(args ...interface{}) () {
  // setShowGrid(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView11setShowGridEb
    // invoke: void setShowGrid(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView11setShowGridEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setShowGrid", args)
  }

}

  // proto:  bool QTableView::isColumnHidden(int column);
func (this *QTableView) isColumnHidden(args ...interface{}) () {
  // isColumnHidden(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView14isColumnHiddenEi
    // invoke: bool isColumnHidden(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView14isColumnHiddenEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "isColumnHidden", args)
  }

}

  // proto:  void QTableView::selectRow(int row);
func (this *QTableView) selectRow(args ...interface{}) () {
  // selectRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView9selectRowEi
    // invoke: void selectRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView9selectRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "selectRow", args)
  }

}

  // proto:  const QMetaObject * QTableView::metaObject();
func (this *QTableView) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK10QTableView10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "metaObject", args)
  }

}

  // proto:  bool QTableView::isCornerButtonEnabled();
func (this *QTableView) isCornerButtonEnabled(args ...interface{}) () {
  // isCornerButtonEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView21isCornerButtonEnabledEv
    // invoke: bool isCornerButtonEnabled()
    C._ZNK10QTableView21isCornerButtonEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "isCornerButtonEnabled", args)
  }

}

  // proto:  void QTableView::selectColumn(int column);
func (this *QTableView) selectColumn(args ...interface{}) () {
  // selectColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12selectColumnEi
    // invoke: void selectColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView12selectColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "selectColumn", args)
  }

}

  // proto:  void QTableView::~QTableView();
func (this *QTableView) FreeQTableView(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTableView", "~QTableView", args)
  }

}

  // proto:  void QTableView::resizeColumnToContents(int column);
func (this *QTableView) resizeColumnToContents(args ...interface{}) () {
  // resizeColumnToContents(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView22resizeColumnToContentsEi
    // invoke: void resizeColumnToContents(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView22resizeColumnToContentsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "resizeColumnToContents", args)
  }

}

  // proto:  void QTableView::QTableView(QWidget * parent);
func NewQTableView(args ...interface{}) QTableView {
  return QTableView{}
}

  // proto:  void QTableView::sortByColumn(int column);
func (this *QTableView) sortByColumn(args ...interface{}) () {
  // sortByColumn(int, Qt::SortOrder)
  // sortByColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::SortOrder"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView12sortByColumnEi
    // invoke: void sortByColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView12sortByColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "sortByColumn", args)
  }

}

  // proto:  int QTableView::columnSpan(int row, int column);
func (this *QTableView) columnSpan(args ...interface{}) () {
  // columnSpan(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView10columnSpanEii
    // invoke: int columnSpan(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QTableView10columnSpanEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "columnSpan", args)
  }

}

  // proto:  int QTableView::columnWidth(int column);
func (this *QTableView) columnWidth(args ...interface{}) () {
  // columnWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView11columnWidthEi
    // invoke: int columnWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView11columnWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "columnWidth", args)
  }

}

  // proto:  int QTableView::columnViewportPosition(int column);
func (this *QTableView) columnViewportPosition(args ...interface{}) () {
  // columnViewportPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView22columnViewportPositionEi
    // invoke: int columnViewportPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView22columnViewportPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "columnViewportPosition", args)
  }

}

  // proto:  int QTableView::rowHeight(int row);
func (this *QTableView) rowHeight(args ...interface{}) () {
  // rowHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView9rowHeightEi
    // invoke: int rowHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView9rowHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "rowHeight", args)
  }

}

  // proto:  int QTableView::rowAt(int y);
func (this *QTableView) rowAt(args ...interface{}) () {
  // rowAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView5rowAtEi
    // invoke: int rowAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView5rowAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "rowAt", args)
  }

}

  // proto:  int QTableView::rowSpan(int row, int column);
func (this *QTableView) rowSpan(args ...interface{}) () {
  // rowSpan(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView7rowSpanEii
    // invoke: int rowSpan(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QTableView7rowSpanEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableView", "rowSpan", args)
  }

}

  // proto:  void QTableView::setCornerButtonEnabled(bool enable);
func (this *QTableView) setCornerButtonEnabled(args ...interface{}) () {
  // setCornerButtonEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView22setCornerButtonEnabledEb
    // invoke: void setCornerButtonEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView22setCornerButtonEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setCornerButtonEnabled", args)
  }

}

  // proto:  QRect QTableView::visualRect(const QModelIndex & index);
func (this *QTableView) visualRect(args ...interface{}) () {
  // visualRect(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView10visualRectERK11QModelIndex
    // invoke: QRect visualRect(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView10visualRectERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "visualRect", args)
  }

}

  // proto:  QModelIndex QTableView::indexAt(const QPoint & p);
func (this *QTableView) indexAt(args ...interface{}) () {
  // indexAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView7indexAtERK6QPoint
    // invoke: QModelIndex indexAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTableView7indexAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "indexAt", args)
  }

}

  // proto:  QHeaderView * QTableView::horizontalHeader();
func (this *QTableView) horizontalHeader(args ...interface{}) () {
  // horizontalHeader()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView16horizontalHeaderEv
    // invoke: QHeaderView * horizontalHeader()
    C._ZNK10QTableView16horizontalHeaderEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "horizontalHeader", args)
  }

}

  // proto:  void QTableView::setModel(QAbstractItemModel * model);
func (this *QTableView) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTableView8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setModel", args)
  }

}

  // proto:  bool QTableView::isSortingEnabled();
func (this *QTableView) isSortingEnabled(args ...interface{}) () {
  // isSortingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView16isSortingEnabledEv
    // invoke: bool isSortingEnabled()
    C._ZNK10QTableView16isSortingEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "isSortingEnabled", args)
  }

}

  // proto:  void QTableView::clearSpans();
func (this *QTableView) clearSpans(args ...interface{}) () {
  // clearSpans()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView10clearSpansEv
    // invoke: void clearSpans()
    C._ZN10QTableView10clearSpansEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "clearSpans", args)
  }

}

  // proto:  void QTableView::setVerticalHeader(QHeaderView * header);
func (this *QTableView) setVerticalHeader(args ...interface{}) () {
  // setVerticalHeader(class QHeaderView *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QHeaderView{}) // "QHeaderView *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView17setVerticalHeaderEP11QHeaderView
    // invoke: void setVerticalHeader(class QHeaderView *)
    var arg0 = args[0].(QHeaderView).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTableView17setVerticalHeaderEP11QHeaderView(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "setVerticalHeader", args)
  }

}

  // proto:  bool QTableView::showGrid();
func (this *QTableView) showGrid(args ...interface{}) () {
  // showGrid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTableView8showGridEv
    // invoke: bool showGrid()
    C._ZNK10QTableView8showGridEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableView", "showGrid", args)
  }

}

  // proto:  void QTableView::showRow(int row);
func (this *QTableView) showRow(args ...interface{}) () {
  // showRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView7showRowEi
    // invoke: void showRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView7showRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "showRow", args)
  }

}

  // proto:  void QTableView::hideColumn(int column);
func (this *QTableView) hideColumn(args ...interface{}) () {
  // hideColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTableView10hideColumnEi
    // invoke: void hideColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTableView10hideColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableView", "hideColumn", args)
  }

}

// <= body block end
