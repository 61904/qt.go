package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qabstractitemdelegate.h
// dst-file: /src/widgets/qabstractitemdelegate.go
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
  // proto:  QVector<int> QAbstractItemDelegate::paintingRoles();
extern void _ZNK21QAbstractItemDelegate13paintingRolesEv(void* qthis);
  // proto:  const QMetaObject * QAbstractItemDelegate::metaObject();
extern void _ZNK21QAbstractItemDelegate10metaObjectEv(void* qthis);
  // proto:  void QAbstractItemDelegate::updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK21QAbstractItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QAbstractItemDelegate::setEditorData(QWidget * editor, const QModelIndex & index);
extern void _ZNK21QAbstractItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex(void* qthis, void* arg0, void* arg1);
  // proto:  void QAbstractItemDelegate::QAbstractItemDelegate(QObject * parent);
extern void* dector_ZN21QAbstractItemDelegateC1EP7QObject(void* arg0);
extern void _ZN21QAbstractItemDelegateC1EP7QObject(void* qthis, void* arg0);
  // proto:  QWidget * QAbstractItemDelegate::createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZNK21QAbstractItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QAbstractItemDelegate::QAbstractItemDelegate(const QAbstractItemDelegate & );
extern void* dector_ZN21QAbstractItemDelegateC1ERKS_(void* arg0);
extern void _ZN21QAbstractItemDelegateC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QAbstractItemDelegate::editorEvent(QEvent * event, QAbstractItemModel * model, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZN21QAbstractItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3);
  // proto:  bool QAbstractItemDelegate::helpEvent(QHelpEvent * event, QAbstractItemView * view, const QStyleOptionViewItem & option, const QModelIndex & index);
extern void _ZN21QAbstractItemDelegate9helpEventEP10QHelpEventP17QAbstractItemViewRK20QStyleOptionViewItemRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3);
  // proto:  void QAbstractItemDelegate::~QAbstractItemDelegate();
extern void _ZN21QAbstractItemDelegateD0Ev(void* qthis);
  // proto:  void QAbstractItemDelegate::setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index);
extern void _ZNK21QAbstractItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QAbstractItemDelegate::destroyEditor(QWidget * editor, const QModelIndex & index);
extern void _ZNK21QAbstractItemDelegate13destroyEditorEP7QWidgetRK11QModelIndex(void* qthis, void* arg0, void* arg1);
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

// class sizeof(QAbstractItemDelegate)=1
type QAbstractItemDelegate struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _closeEditor QAbstractItemDelegate_closeEditor_signal;
//  _commitData QAbstractItemDelegate_commitData_signal;
//  _sizeHintChanged QAbstractItemDelegate_sizeHintChanged_signal;
}

  // proto:  QVector<int> QAbstractItemDelegate::paintingRoles();
func (this *QAbstractItemDelegate) paintingRoles(args ...interface{}) () {
  // paintingRoles()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QAbstractItemDelegate13paintingRolesEv
    // invoke: QVector<int> paintingRoles()
    C._ZNK21QAbstractItemDelegate13paintingRolesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "paintingRoles", args)
  }

}

  // proto:  const QMetaObject * QAbstractItemDelegate::metaObject();
func (this *QAbstractItemDelegate) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QAbstractItemDelegate10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK21QAbstractItemDelegate10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "metaObject", args)
  }

}

  // proto:  void QAbstractItemDelegate::updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index);
func (this *QAbstractItemDelegate) updateEditorGeometry(args ...interface{}) () {
  // updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QAbstractItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: void updateEditorGeometry(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionViewItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK21QAbstractItemDelegate20updateEditorGeometryEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "updateEditorGeometry", args)
  }

}

  // proto:  void QAbstractItemDelegate::setEditorData(QWidget * editor, const QModelIndex & index);
func (this *QAbstractItemDelegate) setEditorData(args ...interface{}) () {
  // setEditorData(class QWidget *, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QAbstractItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex
    // invoke: void setEditorData(class QWidget *, const class QModelIndex &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK21QAbstractItemDelegate13setEditorDataEP7QWidgetRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "setEditorData", args)
  }

}

  // proto:  void QAbstractItemDelegate::QAbstractItemDelegate(QObject * parent);
func NewQAbstractItemDelegate(args ...interface{}) QAbstractItemDelegate {
  return QAbstractItemDelegate{}
}

  // proto:  QWidget * QAbstractItemDelegate::createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index);
func (this *QAbstractItemDelegate) createEditor(args ...interface{}) () {
  // createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QAbstractItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: QWidget * createEditor(class QWidget *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStyleOptionViewItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK21QAbstractItemDelegate12createEditorEP7QWidgetRK20QStyleOptionViewItemRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "createEditor", args)
  }

}

  // proto:  bool QAbstractItemDelegate::editorEvent(QEvent * event, QAbstractItemModel * model, const QStyleOptionViewItem & option, const QModelIndex & index);
func (this *QAbstractItemDelegate) editorEvent(args ...interface{}) () {
  // editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"
  vtys[0][1] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"
  vtys[0][2] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][3] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QAbstractItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QStyleOptionViewItem).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QModelIndex).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN21QAbstractItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "editorEvent", args)
  }

}

  // proto:  bool QAbstractItemDelegate::helpEvent(QHelpEvent * event, QAbstractItemView * view, const QStyleOptionViewItem & option, const QModelIndex & index);
func (this *QAbstractItemDelegate) helpEvent(args ...interface{}) () {
  // helpEvent(class QHelpEvent *, class QAbstractItemView *, const class QStyleOptionViewItem &, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QHelpEvent{}) // "QHelpEvent *"
  vtys[0][1] = reflect.TypeOf(QAbstractItemView{}) // "QAbstractItemView *"
  vtys[0][2] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[0][3] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QAbstractItemDelegate9helpEventEP10QHelpEventP17QAbstractItemViewRK20QStyleOptionViewItemRK11QModelIndex
    // invoke: bool helpEvent(class QHelpEvent *, class QAbstractItemView *, const class QStyleOptionViewItem &, const class QModelIndex &)
    var arg0 = args[0].(QHelpEvent).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAbstractItemView).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QStyleOptionViewItem).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QModelIndex).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN21QAbstractItemDelegate9helpEventEP10QHelpEventP17QAbstractItemViewRK20QStyleOptionViewItemRK11QModelIndex(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "helpEvent", args)
  }

}

  // proto:  void QAbstractItemDelegate::~QAbstractItemDelegate();
func (this *QAbstractItemDelegate) FreeQAbstractItemDelegate(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "~QAbstractItemDelegate", args)
  }

}

  // proto:  void QAbstractItemDelegate::setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index);
func (this *QAbstractItemDelegate) setModelData(args ...interface{}) () {
  // setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QAbstractItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex
    // invoke: void setModelData(class QWidget *, class QAbstractItemModel *, const class QModelIndex &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK21QAbstractItemDelegate12setModelDataEP7QWidgetP18QAbstractItemModelRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "setModelData", args)
  }

}

  // proto:  void QAbstractItemDelegate::destroyEditor(QWidget * editor, const QModelIndex & index);
func (this *QAbstractItemDelegate) destroyEditor(args ...interface{}) () {
  // destroyEditor(class QWidget *, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QAbstractItemDelegate13destroyEditorEP7QWidgetRK11QModelIndex
    // invoke: void destroyEditor(class QWidget *, const class QModelIndex &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK21QAbstractItemDelegate13destroyEditorEP7QWidgetRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractItemDelegate", "destroyEditor", args)
  }

}

// <= body block end
